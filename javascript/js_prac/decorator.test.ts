describe('decorator_test', () => {
  it('descriptor practice', () => {
    // https://wonism.github.io/what-is-decorator/
    const example = {
      name: 'qwerty',
      age: 0,
      birthDate: new Date('1945-08-15'),
    };

    Object.defineProperty(example, 'name', { writable: false });

    Object.defineProperty(example, 'age', {
      get() {
        return new Date().getFullYear() - this.birthDate.getFullYear();
      },
      set(desc) {
        this.birthDate = new Date(desc);
      },
    });
    console.log(Object.getOwnPropertyDescriptor(example, 'name'));
    console.log('age', example.age);
    example.age = 1402021540000;
    console.log('age', example.age);
    /*
    {
      value: 'qwerty',
      writable: false,
      enumerable: true,
      configurable: true
    }
    */
    // writable: writable은 객체의 프로퍼티가 쓰기 가능한 지의 여부. false일 경우 값 쓰기 불가.
    // enumerable: enumerable은 객체의 프로퍼티가 열거 가능한 지의 여부. false일 경우 Object.keys에서 해당 프로퍼티를 볼 수 없음.(Object.values, Object.entries도 마찬가지.)
    // configurable: configurable은 객체의 프로퍼티가 defineProperty를 통해 설정 될 수 있는 지의 여부. false일 경우 Object.defineProperty로 해당 프로퍼티를 수정할 수 없음.
  });

  it('method decorator', () => {
    // 데코레이터는
    // target은 데코레이터를 사용할 class의 instance.
    const readOnly = (target: Person, property: string, descriptor: PropertyDescriptor) => {
      descriptor.writable = false;
      console.log('target', target);
      console.log('property', property);
      console.log('descriptor', descriptor);
      return descriptor; // 데코레이터 함수는 반드시 descriptor를 return 해야 함.
    };

    class Person {
      firstName: string;
      lastName: string;

      constructor(firstName: string, lastName: string) {
        this.firstName = firstName;
        this.lastName = lastName;
      }

      @readOnly
      getFullName() {
        return `${this.firstName} ${this.lastName}`;
      }
    }

    const p = new Person('John', 'Doe');
    console.log(p.getFullName());
    // Person.prototype.getFullName = () => 'CRACKED'; //오류 발생!
  });

  it('using arg in decorator by closer', () => {
    const readOnly = (arg: any) => (target: Person, property: string, descriptor: PropertyDescriptor) => {
      const originMethod = descriptor.value;

      descriptor.value = function (...args: any[]) {
        console.log(arg);
        return originMethod.apply(this, args);
      };
      return descriptor;
    };

    class Person {
      firstName: string;
      lastName: string;

      constructor(firstName: string, lastName: string) {
        this.firstName = firstName;
        this.lastName = lastName;
      }

      @readOnly('TEST')
      getFullName() {
        return `${this.firstName} ${this.lastName}`;
      }
    }

    const p = new Person('John', 'Doe');
    console.log(p.getFullName());
  });

  it('class decorator', () => {
    // class에 사용하는 데코레이터의 인자는 하나
    const ClassDeco = (constructor: typeof Person) => {
      const originMethod = constructor.prototype.getFullName;
      constructor.prototype.getFullName = function () {
        return originMethod.apply(this) + ' TESTING';
      };
    };

    function ClassDeco2(classes: typeof Person) {
      const newclass = class extends classes {
        dob: string;
        constructor(...args) {
          super(args[0], args[1]);
          this.dob = new Date().toString();
        }
        setDob(dob: string) {
          this.dob = dob;
        }
      };
      return newclass;
    }

    @ClassDeco2
    @ClassDeco
    class Person {
      firstName: string;
      lastName: string;

      constructor(firstName: string, lastName: string) {
        this.firstName = firstName;
        this.lastName = lastName;
      }

      getFullName() {
        return `${this.firstName} ${this.lastName}`;
      }
    }

    const p = new Person('John', 'Doe');
    console.log(p.getFullName());
    console.log(p);
  });
});

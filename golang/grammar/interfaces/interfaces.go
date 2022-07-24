package interfaces

import (
	"fmt"
)

// 1. 반드시 메서드명 포함헤야.
// 2. 매개변수와 반환이 다르더라도 이름이 같은 메서드는 불가능
// 3. 인터페이스에는 메서드 구현 포함 X.
type InterfaceName interface {
	Test()
	TestResult(value string) bool
}

// 해당하는 인터페이스의 모든 메서드를 구현해야 인터페이스에 속함.

// 1. 인터페이스를 포함하는 인터페이스
type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
	//이 경우 원래라면 Close 메서드가 겹치지만, 매개변수, 리턴 타입이 동일한 같은 메서드 형식이므로 하나로 합쳐지게 된다.
}

// 2. 빈 인터페이스
//TS의 any와 비슷? 모든 값을 받을 수 있는 함수, 메서드, 변숫값을 만들 때 사용
type Any interface{}

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %.3f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

// 3. 인터페이스의 기본값은 nil

//인터페이스 변환하기
//본래의 구체화된 타입으로 변환할 때 주로 사용
type Student struct {
	Age  int
	Name string
}

type Hellower interface {
	Hello() string
}

func (s *Student) Hello() string {
	return fmt.Sprintf("Hello Age:%d", s.Age)
}

func PrintAge(hellower Hellower) {
	a := hellower.(*Student) //Student와  같은 ConcreteType만 가능. 해당 ConcreteType이 인터페이스를 구현해야 가능.
	fmt.Printf("Age: %d\n", a.Age)
	//ConcreteType이 두 가지 interface(InterfaceA, InterfaceB)를 구현하고 있다면 다른 인터페이스로도 변경 가능.
}

func ConvertInterface() {
	s := &Student{20, "alice"}

	PrintAge(s)

}

package method

import "fmt"

// 기본 메서드 선언 형식
// func (r Reciever) mothodName() type {
// 	...
// 	return typesValue
// }
// 리시버는 모든 로컬 타입들(패키지 내에서 type으로 선언된 타입들. 구조체, 별칭 타입 등)이 가능

type Account struct {
	AccountNo int
	balance   int
	name      string
}

func (a *Account) Withdraw(amount int) bool {
	if a.balance < amount {
		fmt.Println("잔고가 부족합니다.")
		return false
	}
	a.balance -= amount
	return true
}

func Method() {
	a := &Account{AccountNo: 33331502, balance: 10000, name: "민혁"}
	fmt.Println(a)
	a.Withdraw(5000)
	fmt.Println(a)
	a.Withdraw(6000)
	fmt.Println(a)
}

func PointerTypeOrValueTypeMethod() {
	var mainA *Account = &Account{111111, 100000, "민혁"}
	mainA.withdrawPointer(3000)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(2000)
	fmt.Println(mainA.balance)

	mainB := mainA.withdrawReturnValue(1000)
	fmt.Println(mainB.balance)
}

func (a1 *Account) withdrawPointer(amount int) {
	a1.balance -= amount
}
func (a2 Account) withdrawValue(amount int) {
	a2.balance -= amount
}
func (a3 Account) withdrawReturnValue(amount int) Account {
	a3.balance -= amount
	return a3
}

//포인터 메서드: 내부에서 리시버의 값을 변경 가능. 인스턴스 중심.
//값 타입 메서드: 내부에서 리시버의 값을 변경 불가능. 값 중심.

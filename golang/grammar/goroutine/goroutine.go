package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func GoroutineBasicTest() {
	go printKorean()
	go printNumber()

	time.Sleep(3 * time.Second) //main 함수에서 이렇게 기다려주지 않으면 아무리 고루틴이 많아도 다 종료되고 프로그램 끝남.
	fmt.Println("done")
}

func printKorean() {
	koreans := []rune{'가', '나', '다', '라', '마'}
	for _, v := range koreans {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", v)
	}
}
func printNumber() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Println(i)
	}
}

var wg sync.WaitGroup

func GoroutineWaitGroupTest() {
	wg.Add(10) //서브 고루틴 개수 세팅
	for i := 0; i < 10; i++ {
		go sumAtoB(1, 100*i)
	}
	wg.Wait()
}

func sumAtoB(a, b int) {
	sum := 0
	for i := a; i < b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지의 합계는 %d입니다.\n", a, b, sum)
	wg.Done() //호출할 때 마다 wg의 남은 작업 개수 1씩 감소
}

// 동시성 프로그래밍의 문제점 //
// 여러 고루틴에서 동일한 메모리 공간에 접근해 값을 변경하면 예기치 못한 문제가 발생할 수 있다.
// 이러한 문제점을 막기 위한 방법으로 뮤텍스가 있다.

type account struct {
	balance int
}

var mutex sync.Mutex //문제 해결을 위해 추가 되는 코드

func MutexTest() {
	var wg sync.WaitGroup

	ac := &account{0}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(ac)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(ac *account) {
	//문제 해결을 위해 추가 되는 코드.
	mutex.Lock() // 뮤텍스 획득. 다른 고루틴이 뮤텍스를 획득했다면 그 고루틴이 뮤텍스를 놓을 때까지 대기함.
	//문제 해결을 위해 추가 되는 코드.
	defer mutex.Unlock() // defer를 활용해 함수 종료 시에 뮤텍스 반납. 뮤텍스를 획득했다면 반드시 반납해줘야 함.

	if ac.balance < 0 {
		panic("잔액 부족")
	}
	ac.balance += 10000
	time.Sleep(time.Microsecond * 2)
	ac.balance -= 10000
}

// 뮤텍스 사용의 문제점 //
// 1. 고루틴에 의한 성능 개선 효과를 누리지 못함.
// 2. 데드락 발생. 어떠한 고루틴도 뮤텍스를 획득하지 못하게 되어 프로그램이 완전히 멈춰버리는 것.

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
// 2. 데드락 발생 가능. 어떠한 고루틴도 뮤텍스를 획득하지 못하게 되어 프로그램이 완전히 멈춰버리는 것.
// 해결책 2가지
// 1. 작업 영역의 분리
// 2. 역할의 분리

//---------------------------

//채널
func ChannelExample() {
	//채널 인스턴스 생성. 일반적으로 생성하면 크기가 0인 채널이 생성됨.
	var messages chan string = make(chan string)

	//채널에 값 입력
	messages <- "this is message"

	//채널에서 값 빼기
	//뺄 때 채널에 데이터가 없으면 데이터가 생길 때까지 대기
	//채널의 값이 비워지지 않으면 고루틴이 종료되지 않음. => 데드락
	var msg1 string = <-messages
	msg2 := <-messages
	
	//버퍼를 가진 채널. 크기가 3짜리
	bufferedMessage := make(chan string, 3)
	fmt.Println(msg1, msg2, bufferedMessage)
}

//---------------------------

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { // ❷ 데이터를 계속 기다린다. for n:= range chan 구문으로 채널의 데이터를 계속 기다릴 수 있음.
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done() // ❹ for n:= range ch에서 채널을 계속 기다리기 때문에 실행되지 않는다. => 데드락!
}

func ChanWaitingTest() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2 // ❶ 데이터를 넣는다.
	}
	close(ch) //채널을 모두 사용했다면 닫아서 끝내줘야 함. 이걸 추가하면 4번에서의 데드락 문제 해결! 채널이 닫히고 데이터가 모두 처리되면 2번도 끝나게 됨.
	wg.Wait() // ❸ 작업 완료를 기다린다.
}


// select문 활용
// select문을 활용하면 여러 채널을 기다릴 수 있음.
func SelectTest() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool) // ❶ 종료 채널

	wg.Add(1)
	go square1(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}

func square1(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select { // ❷ ch와 quit 양쪽을 모두 기다린다.
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}
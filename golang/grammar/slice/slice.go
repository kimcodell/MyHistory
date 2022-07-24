package slice

import (
	"fmt"
	"sort"
)

//슬라이스 기본
func SliceStudy() {
	slice1 := make([]int, 3, 5) //len = 3 ,cap = 5. 값들은 0(int의 기본값)으로 초기화
	slice2 := []int{0, 1, 2, 3, 4}
	// arr1 := [...]int{0, 1, 2, 3, 4} //이건 배열을 만드므로 주의

	slice1 = append(slice1, 1)
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6) //여러 개도 쌉가능
	fmt.Println("slice1 :", slice1, "slice2 :", slice2)
}

//슬라이스 동작원리
type SliceHeader struct { //전체 24바이트
	Date uintptr //실제 배열을 가리키는 포인터
	Len  int     //배열의 길이(요소 개수)
	Cap  int     //배열의 용량(실제 배열의 전체 길이)
}

func DifferenceArrayAndSlice() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("before change", "arr", arr, "slice", slice)

	changeArray(arr)
	changeSlice(slice)

	fmt.Println("after change", "arr", arr, "slice", slice) //결과: arr는 변경 X, slice는 변경 O
}

func changeArray(arr [5]int) {
	arr[2] = 200
}

func changeSlice(slice []int) {
	slice[2] = 200
}

//append 사용 시 예기치 못하게 발생할 수 있는 문제 1
func AppendProblem1() {
	//남은 빈 공간 = cap - len
	//남은 빈 공간의 수가 추가하려는 수보다 부족하면 cap을 늘리고 len을 늘려서 값을 추가.
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)
	//이제 slice1은 cap 5, len 3 이고 slice2는 cap 5, len 5 인 상태로 같은 슬라이스(배열)을 가리키고 있음.

	slice1[1] = 1000 //slice2도 두번째 값이 변경됨.

	slice1 = append(slice1, 500)
	//이렇게 되면 slice1[len] 위치에 값을 넣고 len을 1 증가시킴. 근데 이렇게 되면 slice2의 네번째 값이 500으로 바뀜.
	fmt.Println("slice1", slice1, "slice2", slice2)
}

//append 사용 시 예기치 못하게 발생할 수 있는 문제 2
func AppendProblem2() {
	//cap만큼 꽉 찬 상태에서 값을 추가하면 기존 크기의 2배 크기 배열을 새로 만들어 값을 복사하고 return 함.

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5) //slice1 - cap: 3, len: 3, slice2 - cap: 6, len: 5.
	//지금 slice1과 slice2의 포인터는 다른 곳을 가리킴.

	slice1[1] = 500 //이렇게 해도 slice2는 변경 X
	slice1 = append(slice1, 600)

	fmt.Println("slice1", slice1, "slice2", slice2)
}

//슬라이싱
func Slicing() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := slice1[1:3] //슬라이싱은 새로운 배열을 만드는게 아니라 해당 부분을 가리키는 포인터를 만드는 것.
	// 슬라이싱을 하면 cap은 여유롭게 지정하기 위해 대상 배열의 총 길이에서 시작 인덱스를 뺸만큼 가지게 됨.

	slice1[1] = 100 //이렇게 수정하면 slice2도 해당 인덱스를 포함하므로 같이 변경됨

	slice2 = append(slice2, 500)
	//이렇게 하면 slice2의 cap은 충분하므로 500이 들어가고, slice1 입장에서는 네번째 항목이 변경됨.

	slice3 := slice1[1:3:4] //시작 인덱스, 끝 인덱스, 최대 인덱스. cap = 최대 인덱스 - 시작 인덱스

	fmt.Println("slice1", slice1, "slice2", slice2, "slice", "slice3", slice3, "slice3.cap", cap(slice3))
}

//각종 팁
func SliceTip() {
	copySlice()
	popAtMiddle()
	appendAtMiddle()
	sortSlice()
}

func copySlice() {
	//복제
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	slice2 := append([]int{}, slice1...)
	slice3 := make([]int, len(slice1))
	copyAmount := copy(slice3, slice1) //만약 첫 인자의 슬라이스의 cap이 부족하면 첫 인자의 cap만큼만 복제됨.
	// 그래서 복사될 슬라이스는 []type{}으로 만들면 안 됨. cap이 0이니까

	fmt.Println(slice1, slice2, slice3, copyAmount)
}

func popAtMiddle() {
	//중간 삭제
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice1)

	index := 2 //index 2인 3을 제거
	slice1 = append(slice1[:index], slice1[index+1:]...)
	fmt.Println(slice1)
}

func appendAtMiddle() {
	//중간 추가
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice1)

	index := 2
	inputValue := 1000
	slice1 = append(slice1[:index], append([]int{inputValue}, slice1[index+1:]...)...) //방법 1. 그러나 임시 슬라이스를 생성하므로 불필요하게 메모리를 사용하게 됨.

	//방법 2.
	slice1 = append(slice1, 0)             //1 요소(공간) 추가
	copy(slice1[index+1:], slice1[index:]) //2 값 복사
	slice1[index] = 2000                   //3 값 변경

	fmt.Println(slice1)
}

type student struct {
	name string
	age  int
}

type students []student

func (s students) Len() int           { return len(s) }
func (s students) Less(i, j int) bool { return s[i].age < s[j].age }
func (s students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//정렬
func sortSlice() {
	//int 슬라이스 정렬
	s := []int{5, 3, 1, 6, 4, 2}
	fmt.Println("before sorting s", s)
	sort.Ints(s)
	fmt.Println("after sorting s", s)

	//구조체 정렬
	stds := []student{
		{"네임2", 21},
		{"민혁", 24},
		{"서현", 21},
		{"네임1", 19},
	}
	fmt.Println("before sorting stds", stds)
	sort.Sort(students(stds))
	fmt.Println("after sorting stds", stds)
}

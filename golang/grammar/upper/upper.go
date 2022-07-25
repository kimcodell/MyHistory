package upper

import (
	"fmt"
	"os"
)

func UpperSkill() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))
	deferTest()
	mapTest()
	errorHandling()
}

// 1. 가변 인수
func sum(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

// 2. defer
func deferTest() {
	f, e := os.Create("test.txt")
	if e != nil {
		fmt.Println("fail creating file!")
		return
	}

	defer fmt.Println("반드시 호출됨.")
	defer f.Close()
	defer fmt.Println("파일 닫음.")
	// 지연 호출은 역순. 파일 닫음 출력 => 실제로 닫음 => 반드시 호출됨 출력

	fmt.Println("파일에 글 작성")
	fmt.Fprintln(f, "Hello World")
}

// 3. 맵
func mapTest() {
	m := make(map[string]string)
	m["a"] = "seoul"
	m["b"] = "busan"
	m["c"] = "gyeonggi"

	fmt.Println(m)

	delete(m, "c") //요소 키로 삭제. 키가 존재하지 않으면 아무 동작도 X.
	fmt.Println(m)
}

// 4. 에러 처리
func errorHandling() {

}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World")
	})

	fmt.Println("3000번 포트에서 서버 실행 중...")
	http.ListenAndServe(":3000", nil)
}

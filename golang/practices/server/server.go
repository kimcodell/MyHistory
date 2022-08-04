package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type student struct {
	Age int
	Name string
	Score int
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World")
	})

	http.HandleFunc("/bar", barHandler)

	fs := http.FileServer(http.Dir("./statics"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/json", jsonHandler)

	fmt.Println("3000번 포트에서 서버 실행 중...")
	http.ListenAndServe(":3000", nil)

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprint(res, "Hello World")
	// })
	// http.ListenAndServe(":3001", mux)
}

//쿼리 파라미터
func barHandler(res http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()  // ❶ 쿼리 인수 가져오기
	name := values.Get("name") // ❷ 특정 키값이 있는지 확인
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id")) // ❸ id값을 가져와서 int타입 변환
	fmt.Fprintf(res, "Hello %s! id:%d", name, id)
}

func jsonHandler(res http.ResponseWriter, req *http.Request) {
	student := student{20, "J", 99}
	data, _ := json.Marshal(student);
	res.Header().Add("content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, string(data))
}
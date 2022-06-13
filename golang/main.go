package main

import (
	"fmt"
	"reflect"
)

type datas struct {
	hash int;
	data []int;
	address int;
}

func main() {
	a := datas{1, []int{1,2,3, 4}, 4};
	p := &a;
	(*p).hash = 2;
	p.hash = 23;
	strgf := "sdfsdfdsf";
	result := reflect.TypeOf(strgf);
	fmt.Print(a, *p, p, p.hash, p.data, p.address, 123213, strgf, result, "\n");
}
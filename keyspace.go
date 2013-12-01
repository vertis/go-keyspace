package main

import (
		"fmt"
)

var keyspace = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")


func toKeyspaceString(arr []int) {
	x := make([]byte, len(arr))
    for idx, value := range arr {
      x[idx]=keyspace[value]
      //fmt.Println("str: ", value)
    }
    fmt.Println("str: ", string(x[:]))
}

func incrementString(x * []int, offset int) {
	arr := *x
	if(arr[offset]==61) {
		arr[offset]=0
		incrementString(x, offset-1)
	} else {
		arr[offset]++
	}
	*x = arr
}

func main() {
	start := []int{ 0, 0, 0, 0, 0, 0, 0, 0}
	//toKeyspaceString(start)  
    // 218340105584895
    for i := 0; i < 1000000000; i++ {
      if(i % 1000000000 == 0) {
        fmt.Println(i)
      }
      incrementString(&start, len(start)-1)
    }
	toKeyspaceString(start)
	//x := [...]byte { 48 }
	//fmt.Println("Key: ", string(x[:])) //, reflect.TypeOf(dk))
}
package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
)

func main() {
	//b()
	myMap:=map[string]string{}
	myMap["key"] = "value"
	for k, v := range myMap {
		fmt.Println("key:%s,value:%s",k,v)
	}
	slice:=[]int{}
	slice = append(slice, 3)
	for e := range slice {
		fmt.Println("条目:%s",slice[e])
	}
}

func b() (sum int) {
	//var client http.Client

	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://wwww.baidu.com", strings.NewReader(""))
	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("jsonStr", string(body))
	response.Body.Close()
	return 1
}

type firstInterface interface {
	call() int
}

type myStruct struct {
	a int
}

func (b myStruct) call() int {
	println("h")
	return 32
}

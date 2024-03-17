package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("에러")
		panic(err)
	}
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello World!")
}

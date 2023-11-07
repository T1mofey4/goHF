package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(url string) int {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
	return len(body)
}

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}
func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func greeting() string {
	return "Hi!"
}

func main() {
	// 	go responseSize("http://example.com")
	// 	go responseSize("http://golang.com")
	// 	go responseSize("http://golang.com/doc")
	// 	time.Sleep(time.Second * 3)

	// go a()
	// go b()
	// fmt.Println("end main")
	myChannel := make(chan float64)
	fmt.Println(go greeting())
}

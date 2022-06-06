package main

import (
	"fmt"
	"protoc-go-hello-world-plugin/example"
)

func main() {
	example := example.Example{}
	fmt.Println(example.Hello())
}

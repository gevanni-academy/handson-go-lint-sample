package main

import "fmt"

func unusedFunc() {} // ← 使用されていない公開関数（警告）

func main() {
	fmt.Println("Hello World")
}

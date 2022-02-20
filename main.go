package main

import "fmt"

const MAGIC_NUMBER int = 16

func main() {
	items := GetVotiLive(26, Codice["Lazio"], MAGIC_NUMBER)
	fmt.Println(items)
}

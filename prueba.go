package main

import (
	"dmnuw"
	"fmt"
)

func main() {
	texto := dmnuw.Parametrs("nombre", "edad", "altura")
	fmt.Println(len(texto))
	texto = dmnuw.Onlystring(texto)
	fmt.Println(len(texto))

}

package main

import (
	_ "dmnuw"
	"fmt"
	"strings"
	"time"
)

func main() {
	var sl []string
	r := time.Now()
	l := r.String()
	array := strings.Fields(l)
	fmt.Println(l)
	for i, v := range array {
		if i < 2 {
			sl = append(sl, v[i])
		}
	}
	fmt.Println("\n", len(array), "\n\t el nuevo slice es", sl, " y su longitud es",
		len(sl))

}

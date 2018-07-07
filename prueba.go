package main

import (
	_ "dmnuw"
	"fmt"
	_ "io/ioutil"
	_ "os"
	_ "strconv"
	_ "strings"
	"time"
)

/*---type fecha struct {
	dia int
	mes time.Month
	año int
}*/

func main() {
	/*---	var actual fecha
	r := time.Now()
	actual.año = r.Year()
	actual.dia = r.Day()
	actual.mes = r.Month()
	d, s := r.Zone()
	//m["mes"] = int(actual.mes)
	fmt.Println(actual, d, s)*/

	//otrofile, _ := os.Create("destiny.txt")

	/*
		para añadir strings a un archivo

		file, _ := os.Create("source.txt")
		for i := 0; i < 1001; i++ {
			file.WriteString(strconv.Itoa(i) + "\n")
		}
		sl, _ := ioutil.ReadFile("source.txt")
		//contenido de source en la varible d
		d := string(sl)
		d = d + "\n"
		file.Close()
		n, _ := os.OpenFile("source.txt", os.O_APPEND, 0777)
		var st string
		for i := 1000; i >= 0; i-- {
			st = st + strconv.Itoa(i) + "\n"
		}
		//nw, _ := os.OpenFile("destiny.txt", os.O_APPEND, 0777)
		d = d + st
		d = d + "\n"
		//escribimos el archivo ya con los datos anteriores y con los actuales
		n.WriteString(d)
		fmt.Println("success!!")*/

	start := time.Now()
	s := start.String()

	run := []rune(s)
	fmt.Println("before", len(run))
	run = append(run[:0], run[:20]...)
	fmt.Println("after", len(run))
	for i, v := range run {

		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%c", v)
	}

}

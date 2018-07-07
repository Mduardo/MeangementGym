package main

import (
	"bufio"
	"dmnuw"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type fecha struct {
	dia int
	mes time.Month
	año int
}

type cliente struct {
	name            string
	age             string
	dateofpay       string
	mounthprecius   int //precio del mes
	cellphone       string
	dateinscription fecha
}

func limpiar() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func dormir() {
	time.Sleep(2 * (1000 * time.Millisecond))
}

func leer() string {
	bufer := bufio.NewReader(os.Stdin)
	texto, _ := bufer.ReadString('\n')

	texto = dmnuw.Onlystring(texto)
	return texto
}

func generarfechadentrode(d *fecha) {
	s := time.Now()
	d.año, d.mes, d.dia = s.Date() //generamos la fecha

}

/*func convertcliente(c cliente) (r []byte) {
	var b []byte
	for i, v := range c {
		b = append(b, v)
	}
	b = append(b, '\n')
	return b
}*/

/*func backup(r *os.File) {
	f, _ := os.Create("tmp.txt")
	sl := ioutil.ReadFile("no tocar/data.txt")
	//convertimos a string todo lo que tenemos aqui en r
	l := string(sl)
	//escribimos todo en temp.txt
	f.WriteString(l)

}*/

func createfiletemp() {
	f, _ := os.Create("no tocar/tmp.txt")
	sl, _ := ioutil.ReadFile("no tocar/data.txt")
	//convertimos a string todo lo que tenemos aqui en r
	l := string(sl)
	//añadimos lo que ya tenemos en el documento
	//escribimos todo en temp.txt
	f.WriteString(l)
}

func añadirAarchivo(f *os.File, datoscliente cliente) {
	j, _ := f.Stat()
	n := j.Name()
	//leemos todos los datos del archivo en un string
	//firstslice, _ := ioutil.ReadFile(n)
	//contentbeforechanges := string(sl)
	//cerramos ese flujo acia el file
	f.Close()
	//abrimos un nuvo flujo acia el archivo en modo añadir
	l, _ := os.OpenFile(n, os.O_APPEND, 0777)
	//leemos el archivo cual es proporcionado pr
	//sl := ioutil.ReadFile(n)

	datoscliente.dateofpay = strconv.Itoa(datoscliente.dateinscription.dia)
	l.WriteString("nombre : ")
	l.WriteString(datoscliente.name)
	l.WriteString(" edad: ")
	l.WriteString(datoscliente.age)
	l.WriteString(" cellular : ")
	l.WriteString(datoscliente.cellphone)
	l.WriteString(" dia de pago: ")
	l.WriteString(datoscliente.dateofpay)
	l.WriteString(" fecha inscripcion: ")
	tmp := time.Now().String() + "\n"
	l.WriteString(tmp)

}

func main() {
	//creamos un flujo de lectura
	//entradabufer := bufio.NewReader(os.Stdin)

	var contador, l int = 0, 0
	var c1 cliente
	//var date fecha
	_, err := os.Stat("no tocar/data.txt")
	if err != nil {
		if os.IsNotExist(err) {
			//os.Create("clientes.txt")
			os.Mkdir("no tocar", 0777)
			os.Create("no tocar/data.txt")
			//os.Create("no tocar/fechasinscripciones.txt")

		}
	}

	for contador == 0 {
		respuesta := dmnuw.Parametrs("Administrar Clientes : 1",
			"Consultar Clientes: 2",
			"Mensualidades de Hoy : 3",
			"Crear Respaldo de datos: 4",
			"Salir de programa : 5")
		r, _ := strconv.Atoi(respuesta)
		fmt.Println(r)
		switch r {
		case 1: //case principal

			limpiar()
			for l == 0 { //for clientes
				fmt.Println("Clientes:\n\n")
				respuesta = dmnuw.Parametrs("Ingresar Cliente : 1",
					"Modificar Cliente : 2",
					"Borrar Cliente : 3",
					"Buscar Cliente : 4",
					"Regresar al menu principal : 5")
				//convertimoa a entero la respuesta ingresdad en forma de string
				r, _ = strconv.Atoi(respuesta)
				fmt.Println(r)
				switch r {
				case 1:
					l = 0
					for l == 0 { //for case 1
						createfiletemp()
						fmt.Println("ingrese nombre del cliente")
						c1.name = leer()
						fmt.Println("ingrese edad del cliente")
						c1.age = leer()
						fmt.Println("ingrese celular del cliente")
						c1.cellphone = leer()
						fmt.Println("el nombre del cliente es \t\n ", c1.name,
							"\n\tsi desea Guardar este nombre ingrese 1")
						respuesta = leer()
						r, _ = strconv.Atoi(respuesta)
						if r == 1 {
							generarfechadentrode(&c1.dateinscription)
							file, _ := os.OpenFile("no tocar/data.txt", os.O_APPEND|os.O_RDWR, 0777)
							///añadirAarchivo(file, c1)
							//inscrip, err := os.OpenFile("no tocar/fechasinscripciones.txt", os.O_RDWR, 0777)
							/**/
							c1.dateofpay = strconv.Itoa(c1.dateinscription.dia)
							//Guardamos cliente pero antes guardamos toda la informacione en c1
							file.WriteString("nombre : ")
							file.WriteString(c1.name)
							file.WriteString(" edad: ")
							file.WriteString(c1.age)
							file.WriteString(" cellular : ")
							file.WriteString(c1.cellphone)
							file.WriteString(" dia de pago: ")
							file.WriteString(c1.dateofpay)
							file.WriteString(" fecha inscripcion: ")
							tmp := time.Now().String()
							run := []rune(tmp)
							//reducimos el string con append
							run = append(run[:0], run[:20]...)
							tmp = string(run)
							tmp = tmp + "\n"
							//tmp := time.Now().String() + "\n"
							file.WriteString(tmp) /**/
							file.Close()
							limpiar()
							l++ //para salir del for
							r = 0
						} else {
							continue
						}

					} //finaliza for case 1
					l = 0 //para que permita ingresar otra vez al menu de clientes
				case 2:
				case 3:
				case 4:
				case 5:
					fmt.Println("saliendo del menu Clientes...")
					dormir()
					l++
					r = 0
					limpiar()
				default:
					fmt.Println("opcion no valida")
					dormir()
					limpiar()
				}

			} //termina for clientes
			l = 0 //para que permita regresar al for principal

		case 2:
			l = 0
		case 3:
			fmt.Println("algo.--")
		case 4:
			fmt.Println("respaldo??")
		case 5:
			fmt.Println("la opcion fue 5")
			fmt.Println("Gracias por utilizar Strong sistem cerrando....")
			dormir()
			limpiar()
			contador++

		default:

			fmt.Println("opcion no valida")
			dormir()
			limpiar()
		} //termina swith principal

	}

}

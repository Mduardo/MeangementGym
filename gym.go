package main

import (
	"bufio"
	"dmnuw"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type cliente struct {
	name          string
	age           string
	dateofpay     string
	mounthprecius int //precio del mes
	cellphone     string
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
	bufer, _ := bufio.NewReader(os.Stdin)
	texto := bufer.ReadString('\n')

	texto = dmnuw.Onlystring(texto)
	return texto
}
func fechaActual() string {
	r := time.Now()
	tiempo := r.Date()
}

func main() {
	//creamos un flujo de lectura
	//entradabufer := bufio.NewReader(os.Stdin)
	var contador, l int = 0, 0
	var c1 cliente
	_, err := os.Stat("clientes.txt")
	if err != nil {
		if os.IsNotExist(err) {
			//os.Create("clientes.txt")
			os.Mkdir("no tocar", 0777)
			os.Create("no tocar/clientes.txt")
			os.Create("no tocar/mensualidades.txt")

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

				respuesta = dmnuw.Parametrs("Ingresar Cliente : 1",
					"Modificar Cliente : 2",
					"Borrar Cliente : 3",
					"Buscar Cliente : 4",
					"Regresar al menu principal : 5")
				//convertimoa a entero la respuesta ingresdad en forma de string
				r, _ = strconv.Atoi(respuesta)
				switch r {
				case 1:
					l = 0
					for l == 0 {

						fmt.Println("ingrese nombre del cliente")
						c1.name = leer()
						fmt.Println("ingrese edad del cliente")
						c1.age = leer()
						fmt.Println("ingrese celular del cliente")
						c1.cellphone = leer()
						fmt.Println("el nombre del cliente es \t ", c1.name,
							"si desea Guardar este nombre ingrese 1")
						respuesta = leer()
						r, _ = strconv.Atoi(respuesta)
						if r == 1 {

							l++
						} else {
							continue
						}

					}

				case 2:
				case 3:
				case 4:
				case 5:
					fmt.Println("saliendo del menu Clientes...")
					dormir()
					l++
					limpiar()
				default:
					fmt.Println("opcion no valida")
					dormir()
					limpiar()
				}

			} //termina for clientes

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
		}

	}

}

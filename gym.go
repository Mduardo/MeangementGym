package main

import (
	"dmnuw"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type cliente struct {
	name          string
	age           int
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

func main() {
	//creamos un flujo de lectura
	//entradabufer := bufio.NewReader(os.Stdin)
	var contador int = 0
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
		/*fmt.Println("Binvenido a Strong System \n" +
		"Que desea hacer\n\tAdministrar clientes : 1\n\t" +
		"Administrar Consultas: 2\n\tAdministrar mensualidades: 3\n\t" +
		"Crear respaldo: 5\n\t" +
		"Salir de sistema :6\n\t")*/

		respuesta := dmnuw.Parametrs("Administrar Clientes : 1",
			"Consultar Clientes: 2",
			"Mensualidades de Hoy : 3",
			"Crear Respaldo de datos: 4",
			"Salir de programa : 5")
		//respuesta = dmnuw.Onlystring(respuesta)
		//fmt.Println(len(respuesta))
		r, _ := strconv.Atoi(respuesta)
		fmt.Println(r)
		switch r {
		case 1:
			fmt.Println("la opcion fue 1")
		case 2:
			fmt.Println("la opcion fue 2")
		case 3:
			fmt.Println("la opcion fue 3")
		case 4:
			fmt.Println("la opcion fue 4")
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

		/*if respuesta == "5" {
			break
		}*/

	}

}

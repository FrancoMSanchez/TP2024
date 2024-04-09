package main

import (
	"fmt"
)

func main() {
	on := true
	for on {
		fmt.Println("Bienvenido al sistema de ejercicios")
		fmt.Println("1. Crear rutina")
		fmt.Println("2. Mostrar rutina")
		fmt.Println("3. Modificar rutina")
		fmt.Println("4. Eliminar rutina")
		fmt.Println("5. Salir")
		fmt.Print("Ingrese una opcion: ")
		var opcion int
		fmt.Scanf("%d ", &opcion)
		switch opcion {
		case 1:
			//createRoutine()
		case 2:
			//showRoutine()
		case 3:
			//modifyRoutine()
		case 4:
			//deleteRoutine()
		case 5:
			on = false
		default:
			fmt.Printf("Error: Opcion %d no reconocida\n", opcion)
		}
	}
}

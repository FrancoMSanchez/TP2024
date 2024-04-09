package main

import (
	"errors"
	"fmt"
)

// Iniciar e importar variables y listas necesarias para el funcionamiento del sistema
func iniciar() error {
	flag1 := ObtenerEjercicios()
	if flag1 == nil {
		e := errors.New("error al obtener los datos del archivo de ejercicios")
		return e
	}
	flag2 := ObtenerEtiquetas()
	if flag2 == nil {
		e := errors.New("error al obtener los datos del archivo de etiquetas")
		return e
	}
	flag3 := ObtenerRutinas()
	if flag3 == nil {
		e := errors.New("error al obtener los datos del archivo de rutinas")
		return e
	}
	flag4 := ObtenerUsuario()
	if flag4 == nil {
		e := errors.New("error al obtener los datos del archivo de usuarios")
		return e
	}
	return nil
}

// Obtencion de listados guardados en csv
func ObtenerEjercicios() []Ejercicio {
	return nil
}
func ObtenerEtiquetas() []Etiqueta {
	return nil
}
func ObtenerRutinas() []Rutina {
	return nil
}
func ObtenerUsuario() []Usuario {
	return nil
}

//Creacion de nuevas variables

func newPuntos() *Puntos {
	return &Puntos{0, 0, 0, 0}
}
func newEtiqueta(nombre string) *Etiqueta {
	return &Etiqueta{nombre}
}
func newEjercicio() *Ejercicio {
	return &Ejercicio{}
}
func newUsuario() *Usuario {
	return &Usuario{}
}

//Miscellaneous:

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

package Structs

type Rutina struct {
	Titulo           string
	Descripcion      string
	Duracion         int
	Dificultad       int
	CaloriasQuemadas int
	Etiquetas        *[]Etiqueta
	Ejercicios       *[]Ejercicio
}

package main

type Puntos struct {
	Cardio       int
	Flexibilidad int
	Resistencia  int
	Fuerza       int
}
type Etiqueta struct {
	Nombre string
}

type Ejercicio []struct {
	Titulo           string
	Descripcion      string
	Duracion         int
	Puntos           *Puntos
	Dificultad       int
	CaloriasQuemadas int
	TipoDeEjercicio  string
	Etiquetas        *[]Etiqueta
}
type Rutina struct {
	Titulo           string
	Descripcion      string
	Duracion         int
	Dificultad       int
	CaloriasQuemadas int
	Etiquetas        *[]Etiqueta
	Ejercicios       *[]Ejercicio
}
type Usuario struct {
	Nombre string
	Rutina *Rutina
	Peso   float64
	Altura float64
	Puntos *Puntos
}

package Structs

type Ejercicio struct {
	Titulo           string
	Descripcion      string
	Duracion         int
	Puntos           *Puntos
	Dificultad       string
	CaloriasQuemadas int
	Etiquetas        *[]Etiqueta
}

func (e *Ejercicio) NewEjercicio(string titulo)

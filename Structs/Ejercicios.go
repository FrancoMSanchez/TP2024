package Structs

import (
	"errors"
)

type Ejercicio struct {
	Titulo           string
	Descripcion      string
	Duracion         int //en minutos
	Puntos           *Puntos
	Dificultad       string //media,baja,alta
	CaloriasQuemadas int    // en cal
	Etiquetas        *[]Etiqueta
}

func NewEjercicio() *Ejercicio {
	return &Ejercicio{}
}
func (e *Ejercicio) Set_titulo(titulo string) error {
	if titulo != "" {
		e.Titulo = titulo
		return nil
	}
	return errors.New("el título ingresado no es válido")
}
func (e *Ejercicio) Set_descripcion(descripcion string) error {
	if descripcion != "" {
		e.Descripcion = descripcion
		return nil
	}
	return errors.New("la descripcion ingresada no es válida")
}
func (e *Ejercicio) Set_duracion(duracion int) error {
	if duracion != 0 {
		e.Duracion = duracion
		return nil
	}
	return errors.New("la duración ingresada no es válida")
}
func (e *Ejercicio) Set_Puntos(puntos *Puntos) error {
	if puntos != nil {
		e.Puntos = puntos
		return nil
	}
	return errors.New("los puntos ingresados no son válidos")
}
func (e *Ejercicio) Set_dificultad(dificultad string) error {
	if dificultad == "baja" || dificultad == "media" || dificultad == "alta" {
		e.Dificultad = dificultad
		return nil
	}
	return errors.New("la descripcion ingresada no es válida")
}

func (e *Ejercicio) Set_calorias_quemadas(calorias_quemadas int) error {
	if calorias_quemadas != 0 {
		e.CaloriasQuemadas = calorias_quemadas
		return nil
	}
	return errors.New("las calorias ingresadas no son válidas")
}

//Etiquetas podria ser de tipo nodo y podemos implementar
//una lista de etiquetas dentro de cada ejercicio/rutina
/*
func (e *Ejercicio) Set_etiquetas(etiquetas *Etiqueta) error {
	if etiquetas != 0 {
		e.Etiquetas = etiquetas
		return nil
	}
	return errors.New("las calorias ingresadas no son válidas")
}*/

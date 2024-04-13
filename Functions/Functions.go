package Functions

import (
	"errors"
	"fmt"

	st "github.com/untref-ayp2/TP-2024-el_barto/Structs"
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
func ObtenerEjercicios() []st.Ejercicio {
	return nil
}
func ObtenerEtiquetas() []st.Etiqueta {
	return nil
}
func ObtenerRutinas() []st.Rutina {
	return nil
}
func ObtenerUsuario() []st.Usuario {
	return nil
}

//Creacion de nuevas variables

/*func newPuntos() *st.Puntos {
	return &st.Puntos{0, 0, 0, 0}
}
func newEtiqueta(nombre string) *Etiqueta {
	return &st.Etiqueta{nombre}
}
func newEjercicio() *Ejercicio {
	return &Ejercicio{}
}
func newUsuario() *Usuario {
	return &Usuario{}
}*/

//Miscellaneous:

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

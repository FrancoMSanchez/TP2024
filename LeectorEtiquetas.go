package Almacenamiento

import (
	"encoding/csv"
	"fmt"
	"os"
)

func LeerEtiquetas() ([][]string, error) {
	f, err := os.Open("Almacenamiento/ArchivosCSV/rutinas.csv")
	if err != nil {
		fmt.Println("error abriendo el archivo: ", err)
	}
	defer f.Close() //cierro el archivo
	r := csv.NewReader(f)
	r.Comma = ','         //indico que esta separado por comas
	r.Comment = '#'       //indico los comentarios
	r.FieldsPerRecord = 7 //indico las columnas

	//el read devuelve un ->[][]string

	rawData, _ := r.ReadAll()

	fmt.Println(rawData)

	//var rawData [][]string
	return rawData, nil
}

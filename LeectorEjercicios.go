package Almacenamiento

import (
	"encoding/csv"
	"fmt"
	"os"
)

func LeerEjecicios() ([][]string, error) {
	f, err := os.Open("Almacenamiento/ArchivosCSV/ejercicios.csv")
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

func GuardarEjercicios(data [][]string) error {
	// Crea un nuevo archivo CSV
	file, err := os.Open("Almacenamiento/ArchivosCSV/ejercicios.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	// Crea un escritor CSV utilizando el archivo
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribe todos los datos en el archivo CSV
	err = writer.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}

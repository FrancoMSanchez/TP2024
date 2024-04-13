package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("error abriendo el archivo: ", err)
	}
	defer f.Close() //cierro el archivo
	r := csv.NewReader(f)
	r.Comma = ','         //indico que esta separado por comas
	r.Comment = '#'       //indico los comentarios
	r.FieldsPerRecord = 3 //indico las columnas

	//el read devuelve un ->[][]string

	rawData, err := r.ReadAll()
	if err != nil {
		fmt.Println("error leyendo la info del archivo: ", err)
	}
	fmt.Println(rawData)

	//var rawData [][]string

}

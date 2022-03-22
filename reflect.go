package main

import (
	"fmt"
	"reflect"
)

// func main(){
// 	var number = 23
// 	var reflectValue - reflect.Valueof(number)

// 	fmt.Println("type variabel :", reflectValue.Type())

// 	if reflectValue.Kind() == reflect.Int{
// 		fmt.Println("nilai variabel :", reflectValue.Int())
// 	}
// }

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Airell", 23}

	var sampleType = reflect.TypeOf(person)
	var valueType = reflect.ValueOf(person)

	fmt.Println(samlpeType.NumField())
	fmt.Println(samlpeType.Field(0).Name)
	fmt.Println(samlpeType)
	fmt.Println(valueType.Field(0).Interface())
}

package main

import (
	"encoding/json"
	"fmt"
	"intetech/interaction"
	"log"
)

func main() {
	//data input
	var n1, m1, n2, m2 int
	fmt.Scan(&n1, &m1)
	arr1 := make([][]int, n1)
	for i := 0; i < n1; i++ {
		arr1[i] = make([]int, m1)
		for j := 0; j < m1; j++ {
			fmt.Scan(&arr1[i][j])
		}
	}
	fmt.Scan(&n2, &m2)
	arr2 := make([][]int, n2)
	for i := 0; i < n2; i++ {
		arr2[i] = make([]int, m2)
		for j := 0; j < m2; j++ {
			fmt.Scan(&arr2[i][j])
		}
	}

	//work with python
	array1JSON, err := json.Marshal(arr1)
	if err != nil {
		log.Fatal(err)
	}
	array2JSON, err := json.Marshal(arr2)
	if err != nil {
		log.Fatal(err)
	}
	err = interaction.SetLibs()
	if err != nil {
		log.Fatal(err)
	}
	path, err := interaction.FindPath()
	if err != nil {
		log.Fatal(err)
	}
	output, err := interaction.PyRun(*path, array1JSON, array2JSON)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(output))
}

package main

import (
	"fmt"
	gofuncs "github.com/ervitis/go-funcs"
)

type MyStruct struct {
	Name    string
	Surname string
	Age     int8
	Country string
}

func main() {
	data := []MyStruct{
		{Name: "AAAAA", Surname: "AAAAA", Age: 20, Country: "Spain",},
		{Name: "BBBBB", Surname: "BBBBB", Age: 45, Country: "England",},
		{Name: "CCCCCC", Surname: "CCCCCC", Age: 20, Country: "Spain",},
		{Name: "DDDDD", Surname: "DDDDD", Age: 60, Country: "Spain",},
	}

	grouped := gofuncs.NewGroup(data)
	fmt.Println(grouped.GroupBy("Surname").Collect())
}

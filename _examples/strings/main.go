package main

import (
	"fmt"
	go_funcs "github.com/ervitis/go-funcs"
	"strings"
)

func main() {
	example := []string{"    monday    ", "tuesday", "    wednesday    ", "thursday", "friday", "saturday", "sunday"}

	r := go_funcs.NewStrings(example...).F().Map(strings.ToUpper).Map(strings.TrimSpace).Reduce(func(a string, b string) string {
		if len(a) < len(b) {
			return a
		}
		return b
	}).Filter(func(s string) bool {
		return s == "SUNDAY"
	})
	fmt.Println(r.Collect().DataResult())
}

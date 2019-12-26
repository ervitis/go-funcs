package go_funcs

import (
	. "github.com/ervitis/zen"
	"testing"
)

func TestGroup_GroupBy(t *testing.T) {
	type Test struct {
		Name    string
		Surname string
		Age     int8
		Country string
	}

	Desc(t, "Testing Group By", func(it It) {
		data := []Test{
			{Name: "AAAAA", Surname: "AAAAA", Age: 20, Country: "Spain",},
			{Name: "BBBBB", Surname: "BBBBB", Age: 45, Country: "England",},
			{Name: "CCCCCC", Surname: "CCCCCC", Age: 20, Country: "Spain",},
			{Name: "DDDDD", Surname: "DDDDD", Age: 60, Country: "Spain",},
			{Name: "EEEEE", Surname: "EEEEE", Age: 60, Country: "Spain",},
			{Name: "FFFFF", Surname: "FFFFF", Age: 60, Country: "Spain",},
			{Name: "GGGGG", Surname: "GGGGG", Age: 60, Country: "England",},
			{Name: "HHHHH", Surname: "HHHHH", Age: 60, Country: "France",},
		}
		group := NewGroup(data)

		it("when I group the data by country", func(expect Expect) {
			grouped := group.GroupBy("Country").Collect()
			expect(len(grouped["Spain"])).ToEqual(5)
			expect(len(grouped["England"])).ToEqual(2)
			expect(len(grouped["France"])).ToEqual(1)
		})

		it("when I group by an unsupported key type", func(expect Expect) {
			// TODO work in progress to support types

			defer func() {
				if r := recover(); r == nil {
					t.Errorf("the function should panicked")
				}
			}()

			group.GroupBy("Age")
		})
	})
}
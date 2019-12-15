package go_funcs

import (
	. "github.com/ervitis/zen"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	Desc(t, "Testing strings functional", func(it It) {
		data := []string{"amanda", "birges", "charles", "daniella", "eugene", "felix", "ginebra", "helen"}
		fstrings := NewStrings(data...)

		it("when I create NewStrings", func(expect Expect) {
			expect(fstrings).ToExist()
		})

		it("when I apply Map function to capitalize the first word", func(expect Expect) {
			data := fstrings.F().Map(strings.Title).Collect().DataResult()
			txt := strings.Join(data, ",")
			expect(txt).ToEqual("Amanda,Birges,Charles,Daniella,Eugene,Felix,Ginebra,Helen")
		})

		it("when I apply Reduce function to get the max length", func(expect Expect) {
			data := fstrings.F().Reduce(func(a string, b string) string {
				if len(a) > len(b) {
					return a
				}
				return b
			}).Collect().DataResult()
			expect(data[0]).ToEqual("daniella")
		})

		it("when I apply Reduce in only two elements or less", func(expect Expect) {
			fs := NewStrings([]string{"hello", "bye"}...)
			data := fs.F().Reduce(func(a string, b string) string {
				if string(a[0]) == "h" {
					return a
				}
				return b
			}).Collect().DataResult()
			expect(data[0]).ToEqual("hello")
		})

		it("when I apply filter it returns an array", func(expect Expect) {
			data := strings.Join(fstrings.F().Filter(func(s string) bool {
				return strings.Contains(s, "ge")
			}).Collect().DataResult(), ",")
			expect(data).ToEqual("birges,eugene")
		})

		it("when I apply EqualsTo it returns an array with the exactly element found or empty", func(expect Expect) {
			data := strings.Join(fstrings.F().EqualsTo("eugene").Collect().DataResult(), ",")
			expect(data).ToEqual("eugene")
			data = strings.Join(fstrings.F().EqualsTo("nothing").Collect().DataResult(), ",")
			expect(data).ToEqual("")
		})

		it("when I apply Include it returns if an element is included", func(expect Expect) {
			expect(fstrings.F().Include("eugene")).ToEqual(true)
			expect(fstrings.F().Include("nothere")).ToEqual(false)
		})

		it("when I apply Any it returns an array of boolean elements", func(expect Expect) {
			data := fstrings.F().Any(func(s string) bool {
				return len(s) > 2
			}).Collect().DataSatisfies()
			for _, v := range data {
				expect(v).ToEqual(true)
			}

			data = fstrings.F().Any(func(s string) bool {
				return len(s) > 30
			}).Collect().DataSatisfies()
			for _, v := range data {
				expect(v).ToEqual(false)
			}
		})

		it("when I apply All it returns an array of boolean elemens", func(expect Expect) {
			data := fstrings.F().All(func(s string) bool {
				return len(s) > 2
			}).Collect().DataSatisfies()
			for _, v := range data {
				expect(v).ToEqual(true)
			}

			data = fstrings.F().All(func(s string) bool {
				return len(s) > 30
			}).Collect().DataSatisfies()
			for _, v := range data {
				expect(v).ToEqual(false)
			}
		})
	})
}

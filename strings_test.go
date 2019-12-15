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
	})
}

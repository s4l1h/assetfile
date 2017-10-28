package assetfile

import "testing"

func Test(t *testing.T) {
	css := New()
	css.Add("a.css")
	css.Add("b.css")
	css.Add("a.css")
	css.Add("c.css")
	css.Add("b.css")
	css.List()
}
func BenchmarkCSS1000(b *testing.B) {

	for n := 0; n < b.N; n++ {
		css := New()
		css.Add("a.css")
		css.Add("b.css")
		css.Add("a.css")
		css.Add("c.css")
		css.Add("b.css")
		css.List()
	}
}

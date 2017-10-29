package assetfilepongo

import "testing"

func Test(t *testing.T) {
	css := New()
	css.Add("a.css")
	css.Add("b.css")
	css.Add("a.css")
	css.Add("c.css")
	css.Add("b.css")
	css.Add("foo_%s.css", "module")
	css.Add("foo_%s_%d.css", "module", 1)

	result := css.List()
	if len(result) != 5 {
		t.Error("list size")
	}
	if result[0] != "a.css" ||
		result[2] != "c.css" ||
		result[3] != "foo_module.css" ||
		result[4] != "foo_module_1.css" {
		t.Error(result)
	}
	css.AddIndex("bar.css", 0)
	css.AddIndex("bar%s.css", 1, "_mods")

	result = css.List()
	if result[0] != "bar.css" || result[1] != "bar_mods.css" {
		t.Error(result)
	}

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

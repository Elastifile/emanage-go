package size_test

import (
	"strings"
	"testing"

	"size"

	. "tester"

	. "github.com/onsi/gomega"
)

var testValues = []struct {
	sz  size.Size
	str string
}{
	{0, "0"},
	{1, "1 byte"},
	{512, "512 bytes"},
	{1024, "1 KiB"},
	{376542, "367.7 KiB"},
	{3 * 1024 * 1024, "3 MiB"},
	{2936013, "2.8 MiB"},
	{129394278, "123.4 MiB"},
	{5 * 1024 * 1024 * 1024, "5 GiB"},
	{5 * 1024 * 1024 * 1024, "5GiB"},
	{5 * 1024 * 1024 * 1024, "5 G"},
	{7 * 1024 * 1024 * 1024 * 1024, "7 TiB"},
	{6047313952768, "5.5 TiB"},
	{4 * 1024 * 1024 * 1024 * 1024 * 1024, "4 PiB"},
	{4 * 1024 * 1024 * 1024 * 1024 * 1024, "4 P"},
	{4 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, "4 E"},
}

func canonical(s string) string {
	s = strings.TrimSuffix(s, "iB")
	s = strings.Replace(s, " ", "", -1)
	return s
}

var _ = Describe("Size", func() {
	It("Convert size correctly", func() {
		for _, item := range testValues {
			Expect("<" + canonical(size.Size(item.sz).String()) + ">").To(Equal("<" + canonical(item.str) + ">"))
		}
	})
})

func TestSizeString(t *testing.T) {
	for _, item := range testValues {
		t.Logf("in=%#v, out=%s", item.sz, item.str)
		str := size.Size(item.sz).String()
		if canonical(str) != canonical(item.str) {
			t.Fatalf("Parsing error, got=%v, expected=%v\n", str, item.str)
		}
	}
}

func TestSizeParse(t *testing.T) {
	var epsilon = 0.05
	for _, item := range testValues {
		t.Logf("in=%#v, out=%s", item.str, item.sz)
		sz, err := size.Parse(item.str)
		check(t, err)
		if !size.Similar(sz, item.sz, epsilon) {
			t.Fatalf("Parsing error, got=%#v, expected=%#v", sz, item.sz)
		}
	}
}

///////////////////////////////////////////////////////////////////////////////
func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

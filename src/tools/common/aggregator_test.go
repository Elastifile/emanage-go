package common

import (
	"bufio"
	"strings"
	"testing"
)

var text = `line one
line two
line three
line four
line five
line six
line seven
line eight
line nine
line ten`

func TestHeadAggregator(t *testing.T) {
	expected := `line one
line two
line three
line four
line five`
	ha := NewHeadAggregator(5)
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		ha.ProcessLine(scanner.Text())
	}
	actual := ha.String()
	if actual != expected {
		t.Fatalf("Expected: %s, but got %s", expected, actual)
	}
}

func TestTailAggregator(t *testing.T) {
	expected := `line six
line seven
line eight
line nine
line ten`
	ha := NewTailAggregator(5)
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		ha.ProcessLine(scanner.Text())
	}
	actual := ha.String()
	if actual != expected {
		t.Fatalf("Expected: %s, but got %s", expected, actual)
	}
}

func TestCombinedAggregator(t *testing.T) {
	expected := `line one
line two
line three
line eight
line nine
line ten`
	ha := NewHeadAggregator(3)
	ta := NewTailAggregator(3)
	ca := CombinedAggregator{ha, ta}
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		ca.ProcessLine(scanner.Text())
	}
	actual := ca.String()
	if actual != expected {
		t.Fatalf("Expected: %s, but got %s", expected, actual)
	}
}

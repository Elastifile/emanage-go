package ioutil2

import (
	"bytes"
	"encoding/json"
	"testing"
	"unicode"
)

func TestFilterReader(t *testing.T) {
	data := "\x00\x01{\"foo\":4\x02\x032}\x04\x05"
	buf := bytes.NewBuffer([]byte(data))
	fr := NewFilterReader(buf, unicode.IsPrint)
	dec := json.NewDecoder(fr)
	m := map[string]int{}
	if e := dec.Decode(&m); e != nil {
		t.Fatalf("Reading failed because: %s", e)
	}
	if m["foo"] != 42 {
		t.Fatalf("Reading failed. Expected 42, got %d", m["foo"])
	}
}

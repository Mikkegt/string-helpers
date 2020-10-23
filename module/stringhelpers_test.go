package module

import "testing"

func TestJoin(t *testing.T) {
	v := Join([]string{"阿部", "ひろし"})
	if v != "阿部ひろし" {
		t.Error("Expected 阿部ひろし, got", v)
	}
}

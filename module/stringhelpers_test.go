package module

import "testing"

func TestJoin(t *testing.T) {
	v := Join([]string{"阿部", "ひろし"})
	if v != "阿部ひろし" {
		t.Error("Expected 阿部ひろし, got", v)
	}
}

func TestReverse(t *testing.T) {
	v := Reverse("テルマエ・ロマエ")
	if v != "エマロ・エマルテ" {
		t.Error("Expected エマロ・エマルテ, got", v)
	}
}

package simpleFactory

import "testing"

func TestGoodMorning_Say(t *testing.T) {
	gm := NewGretting("morning")
	content := gm.Say("yu")
	if content != "hi good morning yu" {
		t.Fatal("TestGoodMorning_Say test fail")
	}
}

func TestGoodEvening_Say(t *testing.T) {
	ge := NewGretting("evening")
	in := "jake"
	expected := "hi good evening jake"
	content := ge.Say("jake")
	if content != expected {
		t.Errorf("TestGoodEvening_Say test fail; Say(%v) expected %v", in, expected)
	}
}

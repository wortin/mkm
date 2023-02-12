package nchart

import (
	"testing"
)

func TestBz_String(t *testing.T) {
	bz := Bz{0, 0, 1, 1, 2, 2, 3, 3}
	bzS := bz.String()
	if bzS != "甲子乙丑丙寅丁卯" {
		t.Errorf("ba zi to string error, exp 甲子乙丑丙寅丁卯 but %s", bzS)
	}
}

func TestBzPermutations(t *testing.T) {
	bzs := BzPermutations()
	if len(bzs) != 518400 {
		t.Errorf("ba zi permutations length error, exp 518400 but %d", len(bzs))
	}
	ls := bzs[len(bzs)-1].String()
	if ls != "癸亥乙丑癸亥癸亥" {
		t.Errorf("ba zi permutations last one error, exp 癸亥乙丑癸亥癸亥 but %s", ls)
	}
}

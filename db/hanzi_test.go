package db

import (
	"testing"
)

func TestQueryHanZi(t *testing.T) {
	w, _ := QueryHanZi("我")
	if w.KxBiHua != 7 {
		t.Errorf("query hanzi error, exp kx bihua 7 but %d", w.KxBiHua)
	}
}

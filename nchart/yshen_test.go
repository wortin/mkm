package nchart

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBz_getRgScore(t *testing.T) {
	var scores []int
	bzs := BzPermutations()
	for _, bz := range bzs {
		c := bz.GetBzNChart()
		s := bz.getRgScore(c)
		scores = append(scores, s)
	}
	var min, max int
	for _, s := range scores {
		if s < min {
			min = s
		}
		if s > max {
			max = s
		}
	}
	gc := 20
	step := (max - min) / gc
	var tj []int
	for i := 0; i < gc; i++ {
		tj = append(tj, 0)
	}
	for _, s := range scores {
		for i := 0; i < gc; i++ {
			if s >= min+i*step && s < min+i*step+step {
				tj[i]++
				break
			}
		}
	}
	for i := 0; i < gc; i++ {
		fmt.Printf("%d-%d\n", min+i*step, min+i*step+step)
	}
	sb := strings.Builder{}
	for _, s := range scores {
		sb.WriteString(fmt.Sprintf("%d\n", s))
	}
	f, err := os.Create("/Users/wortin/Desktop/bzscore.txt")
	if err != nil {
		return
	}
	io.WriteString(f, sb.String())
	// 目测符合正态分布，拖尾偏向右边
	// 正态分布的峰值是多少呢？
}

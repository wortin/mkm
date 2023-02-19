package sql

import (
	"fmt"
	os "os"
	"strings"
	"testing"
)

func TestM(t *testing.T) {
	file, _ := os.ReadFile("hanzi.sql")
	fs := string(file)
	ss := strings.Split(fs, "\n")
	sm := make(map[string]bool)
	c := 0
	for _, s := range ss {
		if strings.HasPrefix(s, "INSERT INTO `hanzi` VALUES (") {
			c++
			sub := s[strings.Index(s, ",")+1:]
			sm[sub] = true
		}
	}
	fmt.Println(c)
	fmt.Println(len(sm))
	sb := strings.Builder{}
	id := 0
	for k, _ := range sm {
		sb.WriteString(fmt.Sprintf("%s%d,", "INSERT INTO `hanzi` VALUES (", id))
		sb.WriteString(k)
		sb.WriteString("\n")
		id++
	}
	os.WriteFile("hanzi.sql", []byte(sb.String()), 777)
}

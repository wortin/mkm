package name

import (
	"fmt"
	"mkm/db"
)

type Name struct {
	*Xing
	*Ming
}

type Xing struct {
	X1    *db.HanZi
	X2    *db.HanZi
	IsFuX bool
}

type Ming struct {
	M1    *db.HanZi
	M2    *db.HanZi
	IsFuM bool
}

func ParseName(xing string, ming string) (*Name, error) {
	x, err := ParseXing(xing)
	if err != nil {
		return nil, err
	}
	m, err := ParseMing(ming)
	if err != nil {
		return nil, err
	}
	return &Name{x, m}, nil
}

func ParseXing(xing string) (*Xing, error) {
	var xs []string
	for _, x := range xing {
		xs = append(xs, string(x))
	}
	if len(xs) == 0 || len(xs) > 2 {
		return nil, fmt.Errorf("only support 1 or 2 han zi in xing {%s}", xing)
	}
	xz1, err := db.QueryHanZi(xs[0])
	if err != nil {
		return nil, err
	}
	if len(xs) == 2 {
		xz2, err := db.QueryHanZi(xs[1])
		if err != nil {
			return nil, err
		}
		return &Xing{xz1, xz2, true}, nil
	} else {
		return &Xing{xz1, nil, false}, nil
	}
}

func ParseMing(ming string) (*Ming, error) {
	var ms []string
	for _, m := range ming {
		ms = append(ms, string(m))
	}
	if len(ms) == 0 || len(ms) > 2 {
		return nil, fmt.Errorf("only support 1 or 2 han zi in ming {%s}", ming)
	}
	mz1, err := db.QueryHanZi(ms[0])
	if err != nil {
		return nil, err
	}
	if len(ms) == 2 {
		mz2, err := db.QueryHanZi(ms[1])
		if err != nil {
			return nil, err
		}
		return &Ming{mz1, mz2, true}, nil
	} else {
		return &Ming{mz1, nil, false}, nil
	}
}

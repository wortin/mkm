package nchart

type Shen10 int

const (
	S10Bj Shen10 = iota
	S10Jc
	S10Ss
	S10Sg
	S10Pc
	S10Zc
	S10Qs
	S10Zg
	S10Py
	S10Zy
)

var Shen10s = [10]string{"比肩", "劫财", "食神", "伤官", "偏财", "正财", "七杀", "正官", "偏印", "正印"}

func (s Shen10) String() string {
	if s < 0 || int(s) >= len(Shen10s) {
		return ""
	}
	return Shen10s[s]
}

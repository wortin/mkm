package nchart

type Yy int

const (
	YyYin Yy = iota
	YyYang
)

var YyChars = [2]string{"阴", "阳"}

func (y Yy) String() string {
	if y < 0 || y > 2 {
		return ""
	}
	return YyChars[y]
}

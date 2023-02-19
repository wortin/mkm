package fyin

type Yum string

const (
	YumA   Yum = "a"
	YumO   Yum = "o"
	YumE   Yum = "e"
	YumI   Yum = "i"
	YumU   Yum = "u"
	YumV   Yum = "v"
	YumAi  Yum = "ai"
	YumEi  Yum = "ei"
	YumUi  Yum = "ui"
	YumAo  Yum = "ao"
	YumOu  Yum = "ou"
	YumIu  Yum = "iu"
	YumIe  Yum = "ie"
	YumVe  Yum = "ve"
	YumEr  Yum = "er"
	YumAn  Yum = "an"
	YumEn  Yum = "en"
	YumIn  Yum = "in"
	YumUn  Yum = "un"
	YumVn  Yum = "vn"
	YumAng Yum = "ang"
	YumEng Yum = "eng"
	YumIng Yum = "ing"
	YumOng Yum = "ong"
)

var YumChar = [24]Yum{YumA, YumO, YumE, YumI, YumU, YumV, YumAi, YumEi, YumUi, YumAo, YumOu, YumIu, YumIe, YumVe, YumEr, YumAn, YumEn, YumIn, YumUn, YumVn, YumAng, YumEng, YumIng, YumOng}

// 该韵母作为姓氏时，名的第一个字不宜用什么
func (s Yum) AsX2IsM1Ji(o string) bool {
	return false
}

package db

type HanZi struct {
	ID          int    `gorm:"primary_key;column:id"`
	HanZi       string `gorm:"column:hanzi"`
	Initial     string `gorm:"column:initial"`
	PinYin      string `gorm:"column:pinyin"`
	JtBuShou    string `gorm:"column:jt_bushou"`
	JtBiHua     int    `gorm:"column:jt_zongbihua"`
	FtBuShou    string `gorm:"column:ft_bushou"`
	KxHanZi     string `gorm:"column:kx_hanzi"`
	KxBiHua     int    `gorm:"column:kx_bihua"`
	WuXing      string `gorm:"column:wuxing"`
	JiXiong     string `gorm:"column:jixiong"`
	IsChangYong string `gorm:"column:is_changyong"`
}

func (HanZi) TableName() string {
	return "hanzi"
}

func QueryHanZi(zi string) (*HanZi, error) {
	hanzi := &HanZi{}
	DB.Table("hanzi").Where("hanzi=?", zi).Find(hanzi)
	return hanzi, nil
}

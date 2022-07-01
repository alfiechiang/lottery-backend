package model

type SixLotteryList struct {
	ID        uint
	Date      Date   `gorm:"type:date"`
	SixNum    string `gorm:"type:varchar(30)"`
	SpecilNum int
}

func (SixLotteryList) TableName() string {
	return "six_lottery_list"
}

package model

type Menu struct {
	ID       uint
	Name     string
	ParentId int
	Path     string
}

func (Menu) TableName() string {
	return "menu"
}

package menu

import (
	"lottery/model"
	"lottery/serializer"
	"github.com/gin-gonic/gin"
)

type MenuListRequest struct {
}

type MenuList struct {
	ID       uint
	Name     string
	ParentId int
	Path     string

	SubMenu []MenuList
}

func (request *MenuListRequest) MenuList(c *gin.Context) serializer.Response {

	var m1 []model.Menu
	model.DB.Table("menu").Where("parent_id", nil).Find(&m1)
	parent := make([]MenuList, 0)
	for _, v := range m1 {
		v1 := MenuList{
			ID:       v.ID,
			Name:     v.Name,
			ParentId: v.ParentId,
			Path:     v.Path,
		}
		parent = append(parent, v1)
	}

	var m2 []model.Menu
	model.DB.Table("menu").Where("parent_id IS NOT NULL").Find(&m2)
	sun := make([]MenuList, 0)

	for _, v := range m2 {
		v1 := MenuList{
			ID:       v.ID,
			Name:     v.Name,
			ParentId: v.ParentId,
			Path:     v.Path,
		}
		sun = append(sun, v1)
	}

	for i := 0; i < len(parent); i++ {
		parent[i].SubMenu = tree(parent[i], sun)
	}


	return serializer.Response{
		Code: 200,
		Data: parent,
	}

}


//遞迴
func tree(menu MenuList, sun []MenuList) []MenuList {

	pack := make([]MenuList, 0)
	menu.SubMenu = make([]MenuList, 0)
	for i := 0; i < len(sun); i++ {
		if menu.ID == uint(sun[i].ParentId) {
			pack = append(pack, sun[i])
			menu.SubMenu = pack
		}
	}

	for i := 0; i < len(menu.SubMenu); i++ {
		menu.SubMenu[i].SubMenu = tree(menu.SubMenu[i], sun)
	}

	return menu.SubMenu
}

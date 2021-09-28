package model

type Depart struct {
	ID   int64  `gorm:"primaryKey;unique;not null" json:"id"` // 部门id
	Name string `gorm:"not null" json:"name"`                 // 部门名字
	Desc string `gorm:"not null" json:"desc"`                 // 部门描述
}

type GetAllDepartsResp struct {
	Count   int       `json:"count" example:"4"`  // 部门个数
	Departs []*Depart `json:"departs" example:""` // 详细部门数组
}

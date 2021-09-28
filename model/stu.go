package model

type Student struct {
	ID        int64  `gorm:"primaryKey;unique;not null" json:"-" swaggerignore:"true" example:"11"` // id
	StaffID   int64  `gorm:"unique;not null" json:"staffID" example:"21051101"`                     // 学号
	StaffName string `gorm:"not null" json:"staffName" example:"小明"`                                // 姓名
	Show      int    `gorm:"not null" json:"show" example:"1"`                                      // 是否显示照片，1为显示，2为不显示
	Photo     int64  `gorm:"not null" json:"photo" example:"6"`                                     // 默认为-1
	Depart    int64  `gorm:"not null" json:"depart" example:"3"`                                    // 无部门为-1
	TimeHook
}

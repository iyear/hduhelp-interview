package model

type Photo struct {
	ID   int64  `gorm:"primaryKey;unique;not null" json:"-" example:"12" swaggerignore:"true"`          // 照片ID
	File string `gorm:"unique;not null" json:"file" example:"e744fd4e-313f-4977-9642-448927a283c1.jpg"` // 文件名
	Size int64  `gorm:"not null" json:"size" example:"1110524"`                                         // 照片大小
	TimeHook
}

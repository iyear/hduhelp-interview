package model

type Option struct {
	ID   int64  `json:"id"`   // 学生ID
	Name string `json:"name"` // 学生姓名
}
type GetOptionsReq struct {
	Num    int   `form:"num" example:"4" binding:"required,gte=1"`
	Depart int64 `form:"depart" example:"-1" binding:"required"`
}
type GetOptionsResp struct {
	Photo   string    `json:"photo" example:"be4aed10-aaf4-453a-8d4e-e0ace7bc56ea.jpeg"` // 照片
	Options []*Option `json:"options"`                                                   // 选项数组
}
type JudgeOptionReq struct {
	Photo string `form:"photo" binding:"required" example:"be4aed10-aaf4-453a-8d4e-e0ace7bc56ea.jpeg"`
	ID    int64  `form:"id" binding:"required" example:"4"`
}
type JudgeOptionResp struct {
	Right int      `json:"right" example:"2"`                 // 是否正确，1为正确，2为不正确
	Name  []string `json:"name,omitempty" example:"小明,小红,小黑"` // 返回所有使用了这张照片的人，大部分情况下只有一张
}

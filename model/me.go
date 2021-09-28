package model

type GetMeResp struct {
	StaffID   int64  `json:"staff_id" example:"21051101"`                              // 学号
	StaffName string `json:"staff_name" example:"小明"`                                  // 姓名
	Photo     string `json:"photo" example:"df93c5bb-b957-4e24-afa1-9cr2372ef8fe.jpg"` // 照片文件名
	Show      int    `json:"show" example:"1"`                                         // 是否显示照片，1为显示，2为不显示
	Depart    string `json:"depart" example:"未选择"`                                     // 部门名称
}

type UpdateMeReq struct {
	Photo  string `json:"photo" binding:"required" example:"df93c5bb-b957-4e24-afa1-9cr2372ef8fe.jpg"` // 照片文件名
	Show   int    `json:"show" binding:"required,gte=1,lte=2" example:"2"`                             // 是否显示照片，1为显示，2为不显示
	Depart int64  `json:"depart" binding:"required" example:"3"`                                       // 部门ID
}

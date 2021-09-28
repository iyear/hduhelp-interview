package e

var errMsg = map[int]string{
	SUCCESS:                    "ok",
	UNKNOWN_ERROR:              "未知错误",
	ERROR_EXIST_STU:            "该学生已存在",
	ERROR_EXIST_STU_FAIL:       "获取学生存在状态失败",
	ERROR_NOT_EXIST_STU:        "该学生不存在",
	ERROR_ADD_STU_FAIL:         "添加学生失败",
	ERROR_GET_STU_FAIL:         "获取学生失败",
	ERROR_GET_ME_FAIL:          "获取个人信息失败",
	ERROR_UPLOAD_PHOTO_FAIL:    "上传照片失败",
	ERROR_GET_STAFF_PHOTO_FAIL: "获取个人照片失败",
	ERROR_UPDATE_ME_FAIL:       "更新个人信息失败",
	ERROR_GET_ALL_DEPARTS_FAIL: "获取所有部门失败",
	ERROR_GET_OPTIONS_FAIL:     "获取选项失败",
	ERROR_JUDGE_OPTION_FAIL:    "判断选项失败",
	AUTH_FAIL:                  "鉴权失败",
}

func GetMsg(code int) string {
	return errMsg[code]
}

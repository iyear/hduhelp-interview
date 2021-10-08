package srv_depart

import (
	"github.com/iyear/hduhelp-interview/db"
	"github.com/iyear/hduhelp-interview/model"
)

func GetAllDeparts() ([]*model.Depart, error) {
	var departs []*model.Depart
	if err := db.Mysql.Find(&departs).Error; err != nil {
		return nil, err
	}
	return departs, nil
}

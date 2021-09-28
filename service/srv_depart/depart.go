package srv_depart

import (
	"github.com/iyear/hduhelp-interview/db"
	"github.com/iyear/hduhelp-interview/model"
)

func GetDepart(id int64) (*model.Depart, error) {
	var d *model.Depart
	if err := db.Mysql.Where("id = ?", id).Limit(1).First(&d).Error; err != nil {
		return nil, err
	}
	return d, nil
}
func GetAllDeparts() ([]*model.Depart, error) {
	var departs []*model.Depart
	if err := db.Mysql.Find(&departs).Error; err != nil {
		return nil, err
	}
	return departs, nil
}

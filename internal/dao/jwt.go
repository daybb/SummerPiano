package dao

import (
	model "helloworld/internal/model/jwt"
)

func (d *Dao) LoginCheck(req *model.LoginReq) (*model.User, error) {
	var user model.User
	db := d.db
	err := db.Table("user").Where("user_name = ?", req.UserName).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *Dao) CheckUser(req *model.RegisterReq) (cnt int64, err error) {
	db := d.db
	err = db.Table("user").Where("user_name = ? and password = ?", req.Name, req.Password).Count(&cnt).Error
	if err != nil {
		return
	}
	return cnt, nil
}

func (d *Dao) AddUser(req *model.UserCreate) (err error) {
	err = d.db.Table("user").Create(&req).Error
	if err != nil {
		return
	}
	return nil
}

package dao

import (
	"context"
	"helloworld/internal/model/piano"
)

func (d *Dao) CreateList(ctx context.Context, req *piano.Create) (err error) {
	db := d.db.WithContext(ctx)
	err = db.Table("piano_info").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

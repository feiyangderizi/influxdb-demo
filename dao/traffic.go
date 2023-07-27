package dao

import (
	"errors"
	"github.com/feiyangderizi/ginServer/global"
	"github.com/feiyangderizi/ginServer/model"
	"gorm.io/gorm"
)

type TrafficDao struct {
}

func (dao *TrafficDao) GetList(startId, limit int) (*[]model.Traffic, error) {
	var traffics = &[]model.Traffic{}
	err := global.DB.Where("id > ? and timestamp>'1683648000'", startId).Limit(limit).Find(&traffics).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return traffics, nil
}

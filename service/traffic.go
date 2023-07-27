package service

import (
	"github.com/feiyangderizi/ginServer/dao"
	"github.com/feiyangderizi/ginServer/model/result"
)

type TrafficService struct{}

var trafficDao dao.TrafficDao

func (service TrafficService) GetList(startId, limit int) result.Result {
	traffics, err := trafficDao.GetList(startId, limit)
	if err != nil {
		return result.FailWithMsg("流量信息查询失败：" + err.Error())
	}
	return result.SuccessWithData(traffics)
}

package controller

import (
	"context"
	"fmt"
	"github.com/feiyangderizi/ginServer/model"
	"github.com/feiyangderizi/ginServer/service"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/loebfly/eztools"
	"log"
	"time"
)

type ImportController struct{}

func (c *ImportController) Import() {
	trafficService := service.TrafficService{}

	var startId = 1
	var limit = 100000
	result := trafficService.GetList(startId, limit)
	if !result.IsOK() {
		return
	}

	import2influxdb(result.Data.(*[]model.Traffic))
}

func import2influxdb(traffics *[]model.Traffic) {
	token := "********************************"
	url := "https://us-east-1-1.aws.cloud2.influxdata.com"
	client := influxdb2.NewClient(url, token)

	org := "feiyang"
	bucket := "cdn"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	for i := 0; i < len(*traffics); i++ {
		item := (*traffics)[i]
		tags := map[string]string{
			"domain": item.Domain,
		}

		flowValue := eztools.Str(item.FlowValue).ToFloat64()
		timestamp := time.Unix(eztools.Str(item.Timestamp).ToInt64(), 0)
		fields := map[string]interface{}{
			"traffic": flowValue,
			//"traffic":     flowValue * 8 * 300,
		}
		point := write.NewPoint("traffic_"+item.FlowType, tags, fields, timestamp)

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Println(fmt.Sprintf("%d导入失败：%s。", item.Id, err.Error()))
		} else {
			log.Println(fmt.Sprintf("%d导入成功。", item.Id))
		}

	}
}

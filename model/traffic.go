package model

type Traffic struct {
	Id         int    `json:"id" gorm:"column:id;auto_increment;primary_key"` //用户编号
	Domain     string `json:"domain" gorm:"column:domain"`
	Timestamp  string `json:"timestamp" gorm:"column:timestamp"`
	FlowType   string `json:"flowType" gorm:"column:flow_type"`
	FlowValue  string `json:"flowValue" gorm:"column:flow_value"`
	UpdateTime string `json:"updateTime" gorm:"column:update_time"`
}

func (traffic *Traffic) TableName() string {
	return "resource_log_2305"
}

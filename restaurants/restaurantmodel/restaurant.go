package restaurantmodel

import (
	"encoding/json"
	"fooddlv/common"
)

const (
	EntityName = "Restaurant"
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string          `json:"name"         gorm:"column:name;"`
	Address         string          `json:"address"      gorm:"column:address;"`
	Lat             float64         `json:"lat"          gorm:"column:lat;"`
	Long            float64         `json:"long"         gorm:"column:long;"`
	Logo            json.RawMessage `json:"logo"         gorm:"column:logo;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

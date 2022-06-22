package model

import "gbsw-laundry-API/config"

type LaundryRoom struct {
	Washer [config.MaxWasherCount + 1]Machine `json:"washer"`
	Dryer  [config.MaxDryerCount + 1]Machine  `json:"dryer"`
}

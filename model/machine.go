package model

import "time"

type Machine struct {
	IsUsing    bool      `json:"isUsing"`
	ExpireTime time.Time `json:"expireTime"`
	User       User      `json:"user"`
	LastUser   User      `json:"lastUser"`
}

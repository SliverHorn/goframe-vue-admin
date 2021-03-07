package model

import "github.com/flipped-aurora/gf-vue-admin/server/library/global"

type JwtBlacklist struct {
	global.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (j *JwtBlacklist) TableName() string {
	return "jwt_blacklists"
}


package models

import "gorm.io/gorm"

type Cidade struct {
	gorm.Model
	Nome       string `json:"nome"`
	CodigoIBGE string `json:"codigoibge"`
	UFID       uint   `json:"uf_id"`
	UF         UF     `gorm:"foreignKey:UFID"`
}

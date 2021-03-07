package request

import system "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"

type ExcelInfo struct {
	FileName string        `json:"fileName"`
	InfoList []system.Menu `json:"infoList"`
}

package request

import system "gf-vue-admin/server/app/model/system"

type ExcelInfo struct {
	FileName string        `json:"fileName"`
	InfoList []system.Menu `json:"infoList"`
}

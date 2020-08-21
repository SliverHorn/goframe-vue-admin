package boot

import "server/app/model"

func InitializeData() (err error) {
	err = model.DataAdmins()
	err = model.DataApis()
	err = model.DataAuthorityMenus()
	err = model.DataCasbinRule()
	err = model.DataAuthorities()
	err = model.DataDictionaries()
	err = model.DataDictionaryDetails()
	err = model.DataMenus()
	return err
}

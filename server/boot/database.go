package boot

import "server/app/model"

func InitializeDatabase() (err error) {
	err = model.TableAdmins()
	err = model.TableApis()
	err = model.TableApis()
	err = model.TableAuthorities()
	err = model.TableAuthorityMenu()
	err = model.TableAuthorityResources()
	err = model.TableBreakpointChucks()
	err = model.TableBreakpointFiles()
	err = model.TableCasbinRule()
	err = model.TableCustomers()
	err = model.TableDictionaries()
	err = model.TableDictionaryDetails()
	err = model.TableFiles()
	err = model.TableJwts()
	err = model.TableMenus()
	err = model.TableOperations()
	return err
}

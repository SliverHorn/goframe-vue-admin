package service

import (
	"server/app/api/request"
	"server/app/model/customers"
	"server/library/global"

	"github.com/gogf/gf/frame/g"
)

// CreateCustomers create a Customers
// CreateCustomers 创建Customers
func CreateCustomers(create *request.CreateCustomer) (err error) {
	insert := customers.Entity{
		CustomerName:       create.CustomerName,
		CustomerPhoneData:  create.CustomerPhoneData,
		SysUserId:          create.SysUserId,
		SysUserAuthorityId: create.SysUserAuthorityId,
	}
	_, err = customers.Insert(&insert)
	return err
}

// DeleteCustomers delete Customers
// DeleteCustomers 删除 Customers
func DeleteCustomers(delete *request.DeleteById) (err error) {
	_, err = customers.Delete(g.Map{"id": delete.Id})
	return err
}

// DeleteCustomers batch deletion Customers
// DeleteCustomers 批量删除 Customers
func DeleteCustomersByIds(deletes *request.DeleteByIds) (err error) {
	_, err = customers.Delete(g.Map{"id IN(?)": deletes.Ids})
	return err
}

// UpdateCustomers update Customers
// UpdateCustomers 更新 Customers
func UpdateCustomers(update *request.UpdateCustomer) (err error) {
	condition := g.Map{"id": update.Id}
	updateData := g.Map{
		"customer_name":         update.CustomerName,
		"customer_phone_data":   update.CustomerPhoneData,
		"sys_user_id":           update.SysUserId,
		"sys_user_authority_id": update.SysUserAuthorityId,
	}
	_, err = customers.Update(updateData, condition)
	return err
}

// FindCustomers Gets a single Customers based on id
// FindCustomers 根据id获取单条Customers
func FindCustomers(find *request.FindById) (data *customers.Customers, err error) {
	data = (*customers.Customers)(nil)
	db := g.DB(global.Db).Table("customers").Safe()
	adminDb := g.DB(global.Db).Table("admins").Safe()
	err = db.Where(g.Map{"id": find.Id}).Struct(&data)
	err = adminDb.Where(g.Map{"authority_id": data.SysUserAuthorityId}).Struct(&data.SysUser)
	return
}

// GetCustomersList Page out the Customers list
// GetCustomersList 分页获取Customers列表
func GetCustomersList(info *request.GetCustomerList) (list interface{}, total int, err error) {
	datalist := ([]*customers.Customers)(nil)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.DB(global.Db).Table("customers").Safe()
	adminDb := g.DB(global.Db).Table("admins").Safe()
	total, err = db.Count()
	err = db.Limit(limit).Offset(offset).Structs(&datalist)
	for _, v := range datalist {
		err = adminDb.Where(g.Map{"authority_id": info.SysUserAuthorityId}).Struct(&v.SysUser)
	}
	return datalist, total, err
}

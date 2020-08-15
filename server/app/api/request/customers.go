package request

type CreateCustomer struct {
	CustomerName       string `p:"customerName" v:"required|length:1,1000#请输入客户名|客户名长度为:min到:max位"`
	CustomerPhoneData  string `p:"customerPhoneData" v:"required|length:1,1000#请输入客户电话|客户电话长度为:min到:max位"`
	SysUserId          uint   `p:"sysUserId"`
	SysUserAuthorityId string `p:"sysUserAuthorityId"`
}

type UpdateCustomer struct {
	Id                 int    `p:"id" v:"required|length:1,1000#请输入自增ID|自增ID长度为:min到:max位"`
	CustomerName       string `p:"customerName" v:"required|length:1,1000#请输入客户名|客户名长度为:min到:max位"`
	CustomerPhoneData  string `p:"customerPhoneData" v:"required|length:1,1000#请输入客户电话|客户电话长度为:min到:max位"`
	SysUserId          int    `p:"sysUserId"`
	SysUserAuthorityId string `p:"sysUserAuthorityId"`
}

type FindCustomer struct {
	Id                 int    `p:"id" v:"required|length:1,1000#请输入自增ID|自增ID长度为:min到:max位"`
	CustomerName       string `p:"customerName" v:"required|length:1,1000#请输入客户名|客户名长度为:min到:max位"`
	CustomerPhoneData  string `p:"customerPhoneData" v:"required|length:1,1000#请输入客户电话|客户电话长度为:min到:max位"`
	SysUserId          int    `p:"sysUserId" v:"required|length:1,1000#请输入负责员工id|负责员工id长度为:min到:max位"`
	SysUserAuthorityId string `p:"sysUserAuthorityId" v:"required|length:1,1000#请输入负责员工角色|负责员工角色长度为:min到:max位"`
}

type GetCustomerList struct {
	CustomerName       string `p:"customerName"`
	CustomerPhoneData  string `p:"customerPhoneData"`
	SysUserId          int    `p:"sysUserId"`
	SysUserAuthorityId string `p:"sysUserAuthorityId"`
	PageInfo
}

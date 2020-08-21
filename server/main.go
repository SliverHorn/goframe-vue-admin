package main

import (
	"server/boot"
)

func main() {
	// 如需自动导入初始数据表及数据,打开以下注释内容,初始化完成后请删掉或注释以下内容
	//if err := boot.InitializeDatabase(); err != nil { // 初始化表,如果出错则不会执行初始化数据
	//	fmt.Println(err)
	//	return
	//}
	//// 初始化数据,并且数据插入是10条10条这样插入的,每个表插入数据都有加事务
	//if err := boot.InitializeData(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	boot.InitializeI18n()      // 初始化I18N国际化
	boot.InitializeRedis()     // 初始化Redis连接, 如果use_multipoint为false,则不会初始化redis的服务
	boot.InitializeRunServer() // 初始化gf服务器
}

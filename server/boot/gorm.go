package boot

import (
	"database/sql"
	extra "github.com/flipped-aurora/gf-vue-admin/server/app/model/extra"
	system "github.com/flipped-aurora/gf-vue-admin/server/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/server/boot/internal"
	"github.com/flipped-aurora/gf-vue-admin/server/library/config"
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/gogf/gf/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) Initialize() {
	switch global.Config.System.DbType {
	case "mysql":
		Mysql.Initialize()
	}
}

func init() {
	link := g.Cfg().GetString("database.default.link")
	a := strings.Split(link, ":")
	if len(a) < 4 {
		g.Log().Error("获取失败!")
	}
	user := a[1]
	passAndHost := strings.Split(a[2], "@")
	portAndName := strings.Split(a[3], "/")
	global.Config.Mysql = config.Mysql{
		Path:         strings.Split(passAndHost[1], "(")[1] + ":" + strings.Split(strings.Split(portAndName[0], "/")[0], ")")[0],
		Config:        "charset=utf8mb4&parseTime=True&loc=Local",
		Dbname:        portAndName[1],
		Username:      user,
		Password:      passAndHost[0],
		MaxIdleConnes: 10,
		MaxOpenConnes: 10,
		LogMode:       false,
		LogZap:        "",
	}
}

var Mysql = new(_mysql)

type _mysql struct {
	db  *gorm.DB
	err error
	sql *sql.DB
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: gorm连接mysql数据库
func (m *_mysql) Initialize() {
	if m.db, m.err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       global.Config.Mysql.Dsn(), // DSN data source name
		DefaultStringSize:         191,                       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                     // 根据版本自动配置
	}), internal.GenerateConfig()); m.err != nil {
		g.Log().Error(`Gorm连接MySQL异常!`, g.Map{"err": m.err})
		os.Exit(0)
	} else {
		if m.sql, m.err = m.db.DB(); m.err != nil {
			g.Log().Error(`DatabaseSql对象获取异常!`, g.Map{"err": m.err})
		} else {
			global.Db = m.db
			m.AutoMigrateTables()
			m.sql.SetMaxIdleConns(global.Config.Mysql.GetMaxIdleConnes())
			m.sql.SetMaxOpenConns(global.Config.Mysql.GetMaxOpenConnes())
		}
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: gorm 同步模型 生成mysql表
func (m *_mysql) AutoMigrateTables() {
	m.err = m.db.AutoMigrate(
		new(system.Api),
		new(system.Admin),
		new(system.Menu),
		new(system.Authority),
		new(system.Dictionary),
		new(system.JwtBlacklist),
		new(system.MenuParameter),
		new(system.OperationRecord),
		new(system.DictionaryDetail),

		new(extra.File),
		new(extra.SimpleUploader),
		new(extra.BreakpointContinue),
		new(extra.BreakpointContinueChunk),
	)
	if m.err != nil {
		g.Log().Error(`注册表失败!`, g.Map{"err": m.err})
		os.Exit(0)
	}
	g.Log().Info(`注册表成功!`)
}

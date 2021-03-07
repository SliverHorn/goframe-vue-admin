/*
Copyright © 2020 SliverHorn

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gfva

import (
	"github.com/flipped-aurora/gf-vue-admin/server/boot"
	"github.com/flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/spf13/cobra"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "github.com/flipped-aurora/gf-vue-admin/server初始化数据",
	Long:  `gf-vue-admin初始化数据适配数据库情况: 1. mysql完美适配, 2. postgresql未适配, 3. sqlite未适配, 4. sqlserver未适配`,
	Run: func(cmd *cobra.Command, args []string) {
		frame, _ := cmd.Flags().GetString("frame")
		path, _ := cmd.Flags().GetString("path")
		switch frame {
		case "gin":
			boot.Viper.Initialize(path)
			Mysql.CheckDatabase()
			Mysql.CheckUtf8mb4()
			Mysql.Info()
			Mysql.Init()
			if global.Config.System.DbType == "mysql" {
				Mysql.AutoMigrateTables()
				Mysql.InitData()
			}
		case "gf":
			boot.Viper.Initialize(path)
			Mysql.CheckDatabase()
			Mysql.CheckUtf8mb4()
			Mysql.Info()
			Mysql.Init()
			if global.Config.System.DbType == "mysql" {
				Mysql.AutoMigrateTables()
				Mysql.InitData()
			}
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("path", "p", "./config/viper.yaml", "自定配置文件路径(绝对路径)")
	initdbCmd.Flags().StringP("frame", "f", "gf", "可选参数为gin,gf")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql,postgresql,sqlite,sqlserver")
}

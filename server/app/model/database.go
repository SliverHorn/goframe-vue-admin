package model

import (
	"server/library/global"

	"github.com/gogf/gf/frame/g"
)

func TableAdmins() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `admins`;\nCREATE TABLE `admins` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',\n  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',\n  `uuid` varchar(255) DEFAULT NULL COMMENT '用户唯一标识UUID',\n  `nickname` varchar(255) DEFAULT 'QMPlusUser' COMMENT '用户昵称',\n  `header_img` varchar(255) DEFAULT 'http://www.henrongyi.top/avatar/lufu.jpg' COMMENT '用户头像',\n  `authority_id` varchar(255) DEFAULT '888' COMMENT '用户角色ID',\n  `username` varchar(255) DEFAULT NULL COMMENT '用户名',\n  `password` varchar(255) DEFAULT NULL COMMENT '用户登录密码',\n  PRIMARY KEY (`id`),\n  KEY `idx_admins_deleted_at` (`delete_at`)\n) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;")
	return
}

func TableApis() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `apis`;\nCREATE TABLE `apis` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',\n  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',\n  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'api路径',\n  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'api中文描述',\n  `api_group` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'api组',\n  `method` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'POST' COMMENT '方法',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_apis_delete_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB AUTO_INCREMENT=148 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
	return
}

func TableAuthorities() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `authorities`;\nCREATE TABLE `authorities` (\n  `authority_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色ID',\n  `authority_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '角色名',\n  `parent_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '父角色ID',\n  `create_at` datetime DEFAULT NULL COMMENT '创建时间',\n  `update_at` datetime DEFAULT NULL COMMENT '更新时间',\n  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',\n  PRIMARY KEY (`authority_id`) USING BTREE,\n  UNIQUE KEY `authority_id` (`authority_id`) USING BTREE,\n  KEY `idx_sys_authorities_deleted_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
	return
}

func TableAuthorityMenu() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `authority_menu`;\nCREATE TABLE `authority_menu` (\n  `authority_id` varchar(255) NOT NULL COMMENT '权限id',\n  `menu_id` varchar(255) NOT NULL COMMENT '菜单id',\n  KEY `menu_id` (`menu_id`),\n  KEY `authority_id` (`authority_id`)\n) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
	return
}

func TableAuthorityResources() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `authority_resources`;\nCREATE TABLE `authority_resources` (\n  `authority_id` varchar(255) NOT NULL COMMENT '权限id',\n  `resources_id` varchar(255) DEFAULT NULL COMMENT '资源id',\n  KEY `resources_id` (`resources_id`),\n  KEY `authority_id` (`authority_id`)\n) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
	return
}

func TableBreakpointChucks() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `breakpoint_chucks`;\nCREATE TABLE `breakpoint_chucks` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',\n  `exa_file_id` int unsigned DEFAULT NULL COMMENT '文件id',\n  `file_chunk_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '切片路径',\n  `file_chunk_number` int DEFAULT NULL COMMENT '切片标号',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_exa_file_chunks_deleted_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
	return
}

func TableBreakpointFiles() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `breakpoint_files`;\nCREATE TABLE `breakpoint_files` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',\n  `file_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件名',\n  `file_md5` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件md5',\n  `file_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件路径',\n  `chunk_id` int DEFAULT NULL COMMENT '关联id',\n  `chunk_total` int DEFAULT NULL COMMENT '切片总数',\n  `is_finish` tinyint(1) DEFAULT NULL COMMENT '是否完整',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_breakpoint_files_deleted_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
	return
}

func TableCasbinRule() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `casbin_rule`;\nCREATE TABLE `casbin_rule` (\n  `ptype` varchar(10) DEFAULT NULL,\n  `v0` varchar(256) DEFAULT NULL,\n  `v1` varchar(256) DEFAULT NULL,\n  `v2` varchar(256) DEFAULT NULL,\n  `v3` varchar(256) DEFAULT NULL,\n  `v4` varchar(256) DEFAULT NULL,\n  `v5` varchar(256) DEFAULT NULL\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
	return
}

func TableCustomers() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `customers`;\nCREATE TABLE `customers` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',\n  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',\n  `customer_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户名',\n  `customer_phone_data` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '客户电话',\n  `sys_user_id` int unsigned DEFAULT NULL COMMENT '负责员工id',\n  `sys_user_authority_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '负责员工角色',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_exa_customers_deleted_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=COMPACT;")
	return
}

func TableDictionaries() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `dictionaries`;\nCREATE TABLE `dictionaries` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` datetime DEFAULT NULL COMMENT '创建时间',\n  `update_at` datetime DEFAULT NULL COMMENT '更新时间',\n  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',\n  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字典名（中）',\n  `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字典名（英）',\n  `status` tinyint(1) DEFAULT NULL COMMENT '状态',\n  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '描述',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_sys_dictionaries_deleted_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
	return
}

func TableDictionaryDetails() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `dictionary_details`;\nCREATE TABLE `dictionary_details` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` datetime DEFAULT NULL COMMENT '创建时间',\n  `update_at` datetime DEFAULT NULL COMMENT '更新时间',\n  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',\n  `label` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '展示值',\n  `value` int DEFAULT NULL COMMENT '字典值',\n  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',\n  `sort` int DEFAULT NULL COMMENT '排序标记',\n  `dictionary_id` int DEFAULT NULL COMMENT '关联标记',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_sys_dictionary_details_deleted_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
	return
}

func TableFiles() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `files`;\nCREATE TABLE `files` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',\n  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',\n  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件名',\n  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件地址',\n  `tag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件标签',\n  `key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '编号',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_files_deleted_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
	return
}

func TableJwts() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `jwts`;\nCREATE TABLE `jwts` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',\n  `jwt` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT 'jwt',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_jwts_deleted_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB AUTO_INCREMENT=145 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
	return
}

func TableMenus() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `menus`;\nCREATE TABLE `menus` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',\n  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',\n  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',\n  `menu_level` int unsigned DEFAULT NULL COMMENT '菜单等级(预留字段)',\n  `parent_id` varchar(255) DEFAULT NULL COMMENT '父菜单ID',\n  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '路由path',\n  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '路由name',\n  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',\n  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '前端文件路径',\n  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '菜单名',\n  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '菜单图标',\n  `sort` int DEFAULT NULL COMMENT '排序标记',\n  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '是否缓存',\n  `default_menu` tinyint(1) DEFAULT NULL COMMENT '是否是基础路由(开发中)',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_menus_deleted_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB AUTO_INCREMENT=513 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
	return
}

func TableOperations() (err error) {
	_, err = g.DB(global.Db).Exec("DROP TABLE IF EXISTS `operations`;\nCREATE TABLE `operations` (\n  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',\n  `create_at` datetime DEFAULT NULL COMMENT '创建时间',\n  `update_at` datetime DEFAULT NULL COMMENT '更新时间',\n  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',\n  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求ip',\n  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求方法',\n  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求路由',\n  `status` int DEFAULT NULL COMMENT '状态',\n  `latency` bigint DEFAULT NULL COMMENT '延迟',\n  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '代理',\n  `error_message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '报错信息',\n  `request` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '请求Body',\n  `user_id` int DEFAULT NULL COMMENT '用户id',\n  `response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '响应Body',\n  PRIMARY KEY (`id`) USING BTREE,\n  KEY `idx_operations_deleted_at` (`delete_at`) USING BTREE\n) ENGINE=InnoDB AUTO_INCREMENT=392 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=COMPACT;")
	return
}

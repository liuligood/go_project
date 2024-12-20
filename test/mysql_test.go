package test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMysql(t *testing.T) {
	// 构建 DSN
	dsn := "root:password@tcp(127.0.0.1:3306)/crmeb?charset=utf8mb4&parseTime=True&loc=Local"

	// 建立数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
	defer db.Close()

	// 测试连接是否正常
	if err := db.Ping(); err != nil {
		log.Fatalf("无法 Ping 数据库: %v", err)
	}

	// 查询所有包含 is_del 或 is_delete 字段的表和字段
	query := `
        SELECT TABLE_NAME, COLUMN_NAME
        FROM information_schema.COLUMNS
        WHERE TABLE_SCHEMA = ? AND COLUMN_NAME IN ('is_del', 'is_delte')
    `
	rows, err := db.Query(query, "crmeb")
	if err != nil {
		log.Fatalf("查询字段失败: %v", err)
	}
	defer rows.Close()

	// 收集需要删除的表字段信息
	tableFieldsMap := make(map[string][]string)
	for rows.Next() {
		var tableName, columnName string
		if err := rows.Scan(&tableName, &columnName); err != nil {
			log.Fatalf("扫描行失败: %v", err)
		}
		tableFieldsMap[tableName] = append(tableFieldsMap[tableName], columnName)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("读取行时发生错误: %v", err)
	}

	if len(tableFieldsMap) == 0 {
		fmt.Println("没有找到包含 is_del 或 is_delte 字段的表。")
		return
	}

	// 遍历表并删除对应的字段
	for table, columns := range tableFieldsMap {
		for _, column := range columns {
			alterStmt := fmt.Sprintf("ALTER TABLE `%s` DROP COLUMN `%s`;", table, column)
			if _, err := db.Exec(alterStmt); err != nil {
				log.Printf("删除表 %s 中的字段 %s 失败: %v", table, column, err)
			} else {
				fmt.Printf("成功删除表 %s 中的字段 %s。\n", table, column)
			}
		}
	}

	fmt.Println("字段删除操作完成。")
}

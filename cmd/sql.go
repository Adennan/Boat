package cmd

import (
	"Boat/internal/sql2struct"
	"log"

	"github.com/spf13/cobra"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

// sqlCmd represents the sql command
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql 转换和处理",
	Long:  `sql 转换和处理`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql 转换",
	Long:  `sql 转换`,
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			Username: username,
			Password: password,
			Charset:  charset,
		}

		dbModel := sql2struct.NewDBModel(dbInfo)
		if err := dbModel.Connect(); err != nil {
			log.Fatalf("dbModel.Connect error: %v", err)
		}

		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns error: %v", err)
		}

		temp := sql2struct.NewStructTeplate()
		templateColumns := temp.AssemblyColumns(columns)
		err = temp.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("templat.Generate error: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "u", "root", "请输入数据库的帐号")
	sql2structCmd.Flags().StringVarP(&password, "password", "p", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的 host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库的实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库的名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}

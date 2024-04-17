package sql_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "gitee.com/chunanyong/dm"          //达梦数据库 驱动
	_ "github.com/denisenkom/go-mssqldb" //sqlservice 驱动
	_ "github.com/go-sql-driver/mysql"   //mysql 驱动
	_ "github.com/godror/godror"         //oracle 驱动
	_ "github.com/lib/pq"                //postgres 驱动
	lssql "github.com/zmloong/lotus/sys/sql"
)

// /测试SqlServer 存储过程
func Test_SqlServer(t *testing.T) {
	err := lssql.OnInit(nil,
		lssql.SetSqlType(lssql.SqlServer),
		lssql.SetSqlUrl(fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s;encrypt=disable", "172.27.100.118", 1433, "test", "sa", "123456")),
	)
	if err != nil {
		fmt.Printf("初始化失败=%s", err.Error())
	} else {
		row, err := lssql.QueryContext("Login", sql.Named("@Account", "liwei1dao"), sql.Named("@Password", "li13451234"))
		if err != nil {
			fmt.Printf("执行存储工程失败:%s", err.Error())
		} else {
			fmt.Printf("执行存储工程成功:%v", row)
			//获取结果集
			for row.Next() {
				var uid int
				var a string
				var p string
				var n string
				row.Scan(&uid, &a, &p, &n)
				fmt.Printf("%d,%s,%s,%s", uid, a, p, n)
			}
		}
	}
}

func Test_MySql(t *testing.T) {
	err := lssql.OnInit(nil,
		lssql.SetSqlType(lssql.MySql),
		lssql.SetSqlUrl("root:Idss@sjzt2021@tcp(172.20.27.125:3306)/mysql"),
	)
	if err != nil {
		fmt.Printf("初始化失败=%v\n", err)
	} else {
		fmt.Printf("初始化成功")
		// if data, err := Query("select table_name from information_schema.tables where table_schema='mysql' and table_type='base table'"); err == nil {
		//SELECT * FROM (Select test.*,@rowno:=@rowno+1 as INCREMENTAL From test) AS T WHERE INCREMENTAL >= 0 and INCREMENTAL < 100 order by INCREMENTAL
		if data, err := lssql.Query("select test.*,(@rowno:=@rowno+1) as rownum from test, (select @rowno:=0) as init;"); err == nil {
			if coluns, err := data.Columns(); err == nil {
				fmt.Printf("coluns:%v\n", coluns)
			} else {
				fmt.Printf("coluns err:%v\n", err)
			}
			a1 := ""
			b2 := ""
			id := 0
			for data.Next() {
				if err := data.Scan(&a1, &b2, &id); err == nil {
					fmt.Printf("a1:%s b2:%s id:%v\n", a1, b2, id)
				} else {
					fmt.Printf("tablename err :%v\n", err)
				}
			}
		} else {
			fmt.Printf("Query err:%v\n", err)
		}
	}
}

func Test_Oracle(t *testing.T) {
	err := lssql.OnInit(nil,
		lssql.SetSqlType(lssql.Oracle),
		lssql.SetSqlUrl("idss_sjzt/idss1234@172.20.27.125:1521/nek"),
	)
	if err != nil {
		fmt.Printf("初始化失败=%v\n", err)
	} else {
		fmt.Printf("初始化成功\n")
	}
}

func Test_DM(t *testing.T) {
	err := lssql.OnInit(nil,
		lssql.SetSqlType(lssql.DM),
		lssql.SetSqlUrl("dm://Sysdba:Sysdba@172.20.27.145:5236"),
	)
	if err != nil {
		fmt.Printf("初始化失败=%v\n", err)
	} else {
		fmt.Printf("初始化成功\n")
	}
}

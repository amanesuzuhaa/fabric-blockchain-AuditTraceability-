package main

// 将其中的math/rand包注释为rand2 ，因为在所有包中存在相应的rand包重名
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	rand2 "math/rand"
	"time"
)



var (
	Db *sqlx.DB
	MysqlPath = "root:@tcp(127.0.0.1:3306)/blockchain"
	sql = "select * from logtest"
)

// mysqlInit 打开数据库的操作
func mysqlInit() {
	database, err := sqlx.Open("mysql", MysqlPath)
	// database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	fmt.Println("open mysql success")
	rand2.Seed(time.Now().Unix())
	// rand是一种假随机数，取决于seed，如果让seed随着时间进行相应的变化，从假随机数实现真随机数
	// test code
	 generateTestdata(300)
}

func getAssetFromMysql(sqlString string) ([]Asset, error) {

	tableData := make([]map[string]interface{}, 0)		// table数据
	assets := make([]Asset, 0)

	rows, err := Db.Query(sqlString)		// 进行返回行查询
	if err != nil {
		return assets,err
	}
	defer rows.Close()
	columns, err := rows.Columns()			// 返回行数据的column， columns 列
	if err != nil {
		return assets, err
	}

	count := len(columns)
	fmt.Print(columns)

	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		entry := make(map[string]interface{})	// 一个columns到values的映射

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	for _, v := range tableData{
		asset := Asset{
			ID:		v[columns[0]].(string),
			TIME: 	v[columns[1]].(string),
			USER: 	v[columns[2]].(string),
			INFO: 	v[columns[3]].(string),
		}
		assets = append(assets, asset)
	}
	return assets, nil
}



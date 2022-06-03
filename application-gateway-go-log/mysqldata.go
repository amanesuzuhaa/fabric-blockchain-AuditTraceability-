package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

type MysqlAsset struct {
	Id  	string `db:"id"`
	Time	string `db:"time"`
	User	string `db:"user"`
	Info    string `db:"info"`
}
// generateTestdata 生成制定行测试数据
func generateTestdata(Imax int) {
	conn, err := Db.Begin()
	defer 	conn.Commit()  // 进行事务操作，在函数结束时提交相应的事务操作
	if err != nil {
		fmt.Println("tx begin failed :", err)
		return
	}

	sql := "insert into logtest(id, time, user, info)values(?, ?,?, ?)"
	// 实现数据插入的sql语句 user_name (root user guest)

	rand2.Seed(time.Now().UnixNano())
	for i := 13; i < Imax; i++{
		_, err = conn.Exec(sql, i,time.RFC3339, randUser(), randInfo())
		if err != nil {
			fmt.Println("exec failed, ", err)
			conn.Rollback()
			return
		}
	}
}

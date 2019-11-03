package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/*
// 建表
CREATE TABLE userinfo
(
    uid serial NOT NULL,
    username character varying(100) NOT NULL,
    departname character varying(500) NOT NULL,
    Created date,
    CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);

CREATE TABLE userdeatail
(
    uid integer,
    intro character varying(100),
    profile character varying(100)
)
WITH(OIDS=FALSE);
*/

func main() {
	// 打开数据库
	db, err := sql.Open("postgres", "user@unix(/path/to/socket)/dbname?charset=utf8")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo (username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)
	res, err := stmt.Exec("aaa", "science", "2019-10-12")
	checkErr(err)

	// pg 没有类似 MySQL 的自增ID，所以不支持该函数
	//id, err := res.LastInsertId()
	//checkErr(err)
	//
	//fmt.Println(id)

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "aaa", "science", "2019-10-12").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("last inserted id = ", lastInsertId)

	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("aaaupdate", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// 删除数据
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=$1")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

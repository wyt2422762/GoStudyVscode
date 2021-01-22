package db

import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func connectMysql(){
	fmt.Println("连接mysql数据库")

	db, err :=sql.Open("mysql", "CVAS:P@$$W0RD@tcp(192.168.52.55:3106)/csdbdx?charset=utf8")
	checkErr(err)
	//查询数据
	rows, err := db.Query("SELECT id FROM employ")
	checkErr(err)

	for rows.Next(){
		var id int
		rows.Scan(&id)
		fmt.Println("id：", id)
	}

}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
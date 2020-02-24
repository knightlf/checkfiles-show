package lib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//数据库连接信息
const (
	USERNAME = "8lab"
	PASSWORD = "8lab"
	NETWORK = "tcp"
	SERVER = "192.168.1.193"
	PORT = 3306
	DATABASE = "redmine"
)

//user表结构体定义
type User struct {
	Id int `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Status int   `json:"status" form:"status"`      // 0 正常状态， 1删除
	Createtime int64 `json:"createtime" form:"createtime"`
}

//crshow 表结构体定义
type Crshow struct {
	Showstring string `json:"showstring" form:"showstring"`
	Ca string `json:"ca" form:"ca"`
	Status int   `json:"status" form:"status"`
	Createtime int `json:"createtime" form:"createtime"`
}

//func main() {
//	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
//	DB, err := sql.Open("mysql", conn)
//	if err != nil {
//		fmt.Println("connection to mysql failed:", err)
//		return
//	}
//
//	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超时的连接就close
//	DB.SetMaxOpenConns(100)                //设置最大连接数
//	CreateTable(DB)
//	InsertData(DB)
//	QueryOne(DB)
//	QueryMulti(DB)
//	UpdateData(DB)
//	DeleteData(DB)
//}

func CreateTable(DB *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS users(
	id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	username VARCHAR(64),
	password VARCHAR(64),
	status INT(4),
	createtime INT(10)
	); `

	if _, err := DB.Exec(sql); err != nil {
		fmt.Println("create table failed:", err)
		return
	}
	fmt.Println("create table successd")
}

//插入数据
func InsertData(DB *sql.DB) {
	result,err := DB.Exec("insert INTO users(username,password) values(?,?)","test","123456")
	if err != nil{
		fmt.Printf("Insert data failed,err:%v", err)
		return
	}
	lastInsertID,err := result.LastInsertId()    //获取插入数据的自增ID
	if err != nil {
		fmt.Printf("Get insert id failed,err:%v", err)
		return
	}
	fmt.Println("Insert data id:", lastInsertID)

	rowsaffected,err := result.RowsAffected()  //通过RowsAffected获取受影响的行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v",err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}

//查询单行
func QueryOne(DB *sql.DB) {
	cs := new(Crshow)   //用new()函数初始化一个结构体对象
	row := DB.QueryRow("SELECT * FROM crshow where ca=? ", "cc")
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&cs.Showstring,&cs.Ca,&cs.Status,&cs.Createtime); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println("Single row data:", *cs)
}

//查询多行
func QueryMulti(DB *sql.DB) {
	user := new(User)
	rows, err := DB.Query("select id,username,password from users where id = ?", 2)

	defer func() {
		if rows != nil {
			rows.Close()   //关闭掉未scan的sql连接
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password)  //不scan会导致连接不释放
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		fmt.Println("scan successd:", *user)
	}
}

//更新数据
func UpdateData(DB *sql.DB){
	result,err := DB.Exec("UPDATE users set password=? where id=?","111111",1)
	if err != nil{
		fmt.Printf("Insert failed,err:%v\n", err)
		return
	}
	fmt.Println("update data successd:", result)

	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}

//删除数据
func DeleteData(DB *sql.DB){
	result,err := DB.Exec("delete from users where id=?",1)
	if err != nil{
		fmt.Printf("Insert failed,err:%v\n",err)
		return
	}
	fmt.Println("delete data successd:", result)

	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}

//有返回值  两个表即不一样 "crshow","crshow1",返回被修改的行数
func CompareData(DB *sql.DB, tb1 string,tb2 string) (int64){
	//querysql:="SELECT * FROM " +
	//	"(SELECT *  FROM ? UNION ALL SELECT * FROM ?) tbl " +
	//	"GROUP BY showstring,`status`, ca,`createtime` " +
	//	"HAVING COUNT(*) = 1"
	querysql:="SELECT * FROM (SELECT * FROM ? UNION ALL SELECT * FROM ?) tbl GROUP BY showstring,`status`, ca,`createtime` HAVING COUNT(*) = 1"
	InfoHander(querysql)
	result,err := DB.Exec(querysql,tb1,tb2)
	//result,err := DB.Query(querysql,tb1,tb2)

	if err == nil{
		//fmt.Printf("CompareData failed,err:%v\n",err)
		LogHander("CompareData failed,err:%v\n",err)
		return -1
	}
	fmt.Println("compare data successd:", result)
	InfoHander("compare data successd. ")

	rowsaffected,err := result.RowsAffected()
	if err == nil {
		//fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		LogHander("Get RowsAffected failed,err:%v\n",err)
		return -1
	}
	//fmt.Println("Affected rows:", rowsaffected)
	InfoHander("Affected rows: "+string(rowsaffected))
	return rowsaffected
}

func CompareTables(DB *sql.DB, tb1 string,tb2 string) (int) {
	//cs:=new(Crshow)
	var cs Crshow
	count:=0

	//querysql:="SELECT * FROM " +
	//	"(SELECT *  FROM ? UNION ALL SELECT * FROM ?) tbl " +
	//	"GROUP BY showstring,`status`, ca,`createtime` " +
	//	"HAVING COUNT(*) = 1"
	querysql:=`SELECT * FROM
		(SELECT *  FROM ? UNION ALL SELECT * FROM ?) tbl
		GROUP BY showstring,status,ca,createtime
		HAVING COUNT(*) = 1`
	InfoHander(querysql)
	//querysql := "SELECT *  FROM ?"

	rows,err:=DB.Query(querysql,tb1,tb2)
	//rows,err:=DB.Query(querysql,tb1)

	//stmt,_:=DB.Prepare(querysql)
	//defer  stmt.Close()
	//rows,err:=stmt.Query(tb1,tb2)

	if err != nil{
		//fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		LogHander("exec sql failed,err:%v\n",err)
		//return -1
	}

	//defer rows.Close()
	for rows.Next() {
		//不scan会导致连接不释放
		err = rows.Scan(&cs.Showstring, &cs.Ca, &cs.Status, &cs.Createtime)
		if err != nil {
			//fmt.Printf("Scan failed,err:%v\n", err)
			LogHander("Scan failed,err:%v\n",err)
			return -1
		}
		count+=1
	}
	//rows,_:=stmt.Query(tb1,tb2)



	defer func() {
		if rows != nil {
			rows.Close()   //关闭掉未scan的sql连接
		}
	}()
	//if err != nil {
	//	fmt.Printf("Query failed,err:%v\n", err)
	//	return -1
	//}
	//for rows.Next() {
	//	err = rows.Scan(&user.Id, &user.Username, &user.Password)  //不scan会导致连接不释放
	//	if err != nil {
	//		fmt.Printf("Scan failed,err:%v\n", err)
	//		return -1
	//	}
	//	fmt.Println("scan successd:", *user)
	//}


	return count

}

//查询单行
func TableOne(DB *sql.DB, tb1 string,tb2 string)  (int) {
	querysql:=`SELECT * FROM (SELECT *  FROM ? UNION ALL SELECT * FROM ?) tbl GROUP BY showstring,STATUS,ca,createtime HAVING COUNT(*) = 1  LIMIT 1`
	cs := new(Crshow)   //用new()函数初始化一个结构体对象
	//row := DB.QueryRow("SELECT * FROM crshow where ca=? ", "cc")
	row := DB.QueryRow(querysql,tb1,tb2)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&cs.Showstring,&cs.Ca,&cs.Status,&cs.Createtime); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		//return -1
	}
	fmt.Println("Single row data:", *cs)
	return 1
}

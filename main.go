package main

import (
	"checkfiles-show/lib"
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

//const table1 = "crshow"
//const table2 = "crshow1"
const (
	USERNAME = "8lab"
	PASSWORD = "8lab"
	NETWORK = "tcp"
	SERVER = "192.168.1.193"
	PORT = 3306
	DATABASE = "redmine"
)


//help print
func helper() {
	//help show
	fmt.Printf("|%-6s|%-6s|\n", "rsql", "--show mysql table restore.")
	fmt.Printf("|%-6s|%-6s|\n", "rfile", "-- show files restore.")
	fmt.Printf("|%-6s|%-6s|\n", "smd5", "--test show mysql compare data Affected rows.")
	fmt.Printf("|%-6s|%-6s|\n", "ssql", "--test show files type md5.")

}

func rsql(){
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		lib.LogHander("connection to mysql failed:", err)
		return
	}

	lib.QueryOne(DB)
	fmt.Println(".........................db queryone finished! ")

}

func rfiles(){

}

func showmd5(){
	sourcemd5,err := lib.GetFileName("./source")
	if err !=nil {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: show md5 error ")
		fmt.Println(".........................show files md5 exception.")
	}
	fmt.Println(".........................show files md5 finished.")
	fmt.Println(sourcemd5)

}

var DB *sql.DB

func showsql(){
	tb1:="crshow"
	tb2:="crshow"
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		lib.LogHander("connection to mysql failed:", err)
		return
	}
	defer DB.Close()

	//srows := lib.CompareData(DB,table1,table2)
	//srows := lib.CompareTables(DB,table1,table2)
	srows := lib.CompareTables(DB,tb1,tb2)
	//srows := CompareTables(DB,table1,table2)
	//fmt.Println(srows)
	if srows ==-1 {
		//fmt.Println(err.Error())
		lib.InfoHander("db table have not edit! ")
		fmt.Println(".........................db table have not affected rows.")

	}
	fmt.Println(".........................db table have edit: ")
	fmt.Println(srows)

}

func showssql(){
	tb1:="crshow"
	tb2:="crshow"
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		lib.LogHander("connection to mysql failed:", err)
		return
	}
	defer DB.Close()

	//srows := lib.CompareData(DB,table1,table2)
	//srows := lib.CompareTables(DB,table1,table2)
	srows := lib.TableOne(DB,tb1,tb2)
	//srows := CompareTables(DB,table1,table2)
	//fmt.Println(srows)
	if srows ==-1 {
		//fmt.Println(err.Error())
		lib.InfoHander("mysql operate exciption! ")
		fmt.Println(".........................db table have excption.")

	}
	fmt.Println(".........................db table have edit: ")
	fmt.Println(srows)

}
////crshow 表结构体定义
//type Crshow struct {
//	Showstring string `json:"showstring" form:"showstring"`
//	Ca string `json:"ca" form:"ca"`
//	Status int   `json:"status" form:"status"`
//	Createtime int `json:"createtime" form:"createtime"`
//}
//
//func CompareTables(DB *sql.DB, tb1 string,tb2 string) (int) {
//	cs:=new(Crshow)
//	//querysql:="SELECT * FROM " +
//	//	"(SELECT *  FROM ? UNION ALL SELECT * FROM ?) tbl " +
//	//	"GROUP BY showstring,`status`, ca,`createtime` " +
//	//	"HAVING COUNT(*) = 1"
//	querysql:=`SELECT * FROM
//		(SELECT *  FROM ? UNION ALL SELECT * FROM ?) tbl
//		GROUP BY showstring,status,ca,createtime
//		HAVING COUNT(*) = 1`
//
//	rows,err:=DB.Query(querysql,tb1,tb2)
//	defer  DB.Close()
//
//	//stmt,_:=DB.Prepare(querysql)
//	//defer  stmt.Close()
//	//rows,err:=stmt.Query(tb1,tb2)
//
//	if err != nil{
//		//fmt.Printf("Get RowsAffected failed,err:%v\n",err)
//		lib.LogHander("exec sql failed,err:%v\n",err)
//		//return -1
//	}
//
//	count:=0
//	for rows.Next(){
//		err = rows.Scan(&cs.Showstring, &cs.Ca, &cs.Status, &cs.Createtime)  //不scan会导致连接不释放
//		if err != nil {
//			//fmt.Printf("Scan failed,err:%v\n", err)
//			lib.LogHander("Scan failed,err:%v\n",err)
//			return -1
//		}
//		count+=1
//		//fmt.Println("scan successd:", *user)
//	}
//	//rows,_:=stmt.Query(tb1,tb2)
//	rows.Close()
//	return count
//
//}

func main(){
	if len(os.Args)!=2{
		helper()
		//for idx, args := range os.Args {
		//	fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
		//}
	}
	if len(os.Args)==2{
		if string(os.Args[1])=="rsql"{
			rsql()
			fmt.Println("arg:", string(os.Args[1]))
			//fmt.Println("upgrade mysql table restore.")
		}
		if string(os.Args[1])=="rfile"{
			rfiles()
			fmt.Println("arg:", string(os.Args[1]))
			fmt.Println("upgrade show files restore.")
		}
		if string(os.Args[1])=="smd5"{
			showmd5()
			fmt.Println("arg:", string(os.Args[1]))
			//fmt.Println("show files type to md5.")
		}
		if string(os.Args[1])=="ssql"{
			//showsql()
			showssql()
			fmt.Println("arg:", string(os.Args[1]))
			//fmt.Println("show sql type to ssql.")
		}
	}
}

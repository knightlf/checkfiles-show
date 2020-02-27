package main

import (
	"cf-show/lib"
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

const tbSource = "crshow1"
const tbDist = "crshow"
const fmd="28cf7f56356350f48559ad848635c7e5"

const dirSource="./source"
const dirDist="./destin"

const (
	USERNAME = "8lab"
	PASSWORD = "8lab"
	NETWORK = "tcp"
	SERVER = "192.168.1.193"
	PORT = 3306
	DATABASE = "redmine"
)

var DB *sql.DB

//help print
func helper() {
	//help show
	fmt.Printf("|%-6s|%-6s|\n", "rsql", "--show mysql table restore.")
	fmt.Printf("|%-6s|%-6s|\n", "rfile", "-- show files restore.")
	fmt.Printf("|%-6s|%-6s|\n", "smd5", "--test show mysql compare data Affected rows.")
	fmt.Printf("|%-6s|%-6s|\n", "ssql", "--test show files type md5.")

}

func rsql(){
	//tb1:="crshow"
	//tb2:="crshow1"
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		lib.LogHander("connection to mysql failed:", err)
		return
	}

	srows:=lib.CompareTables(DB,tbSource,tbDist)
	fmt.Println(".........................db compare finished! is "+ string(srows))
	if srows !=0{
		resTable:=lib.RestoreData(DB,tbSource,tbDist)
		fmt.Println(".........................db restore finished! is "+string(resTable))
	}

}

func rfiles(){
	sourcemd5,serr := lib.GetFileName(dirSource)
	if serr !=nil {
		//fmt.Println(err.Error())
		//lib.InfoHander("exec faild: get dirSource md5 error ")
		lib.LogHander("exec faild: get dirSource md5 error ",serr)
		fmt.Println(".........................get dirSource md5 exception.")
	}
	destinmd5,derr := lib.GetFileName(dirDist)
	if derr !=nil {
		//fmt.Println(err.Error())
		lib.LogHander("exec faild: get dirDist md5 error ",derr)
		fmt.Println(".........................get dirDist md5 exception.")
	}
	if sourcemd5==destinmd5{
		lib.InfoHander("the file md5 exec has equal. ")
	}else{
		cpStr:=lib.CmdBash("cp -av "+dirSource+"/* "+dirDist)
		lib.InfoHander("exec cp: "+cpStr)
		//fmt.Println(err.Error())
		destinmd5,derr := lib.GetFileName(dirDist)
		if derr !=nil {
			//fmt.Println(err.Error())
			lib.LogHander("exec faild: get dirDist md5 error ",derr)
			fmt.Println(".........................get dirDist md5 exception.")
		}
		lib.LogHander("exec faild: get dirDist md5 error ",derr)
		fmt.Println("dirDist: "+fmd+"=="+"dirDist: "+destinmd5)
	}
	fmt.Println("dirDist: "+fmd+"=="+"dirDist: "+destinmd5)
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


func showssql(){
	tb1:="crshow"
	tb2:="crshow1"
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		lib.LogHander("connection to mysql failed:", err)
		return
	}
	defer DB.Close()

	//srows := lib.CompareData(DB,table1,table2)
	srows := lib.CompareTables(DB,tb1,tb2)
	//lib.TableOne(DB,tb1)

	//lib.QueryOne(DB)
	//srows := CompareTables(DB,table1,table2)
	//fmt.Println(srows)
	//if srows ==-1 {
		//fmt.Println(err.Error())
	//	lib.InfoHander("mysql operate exciption! ")
	//	fmt.Println(".........................db table have excption.")

	//}
	fmt.Println(".........................db table have edit: ")
	fmt.Println(srows)

}


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

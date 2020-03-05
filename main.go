package main

import (
	"cf-show/lib"
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const tbSource = "xcr_source"
const tbDist = "crshow"
const fmd="cd20059c676a1f35d4a34f897b736430"
const dtime=10

const DING_TOKEN = "dd84405981561e0f67af319ece4059f8d06fa56eb5f79d298443765c3024c95f"
const DING_SECRET = "SEC3b6b33bb310bfcaf200d99fc47ef72030437287435bdaf1741531433da21fb67"

const dirSource="./source"
const dirDist="/chongren"

const nodeIp="211.141.83.125"

const (
	USERNAME = "8lab"
	PASSWORD = "8lab"
	NETWORK = "tcp"
	//SERVER = "192.168.1.193"
	SERVER = "172.21.0.25"
	PORT = 3306
	//DATABASE = "redmine"
	DATABASE = "attack_defense_info"
	AlertDB = "attack_defense_info"
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
	defer DB.Close()

	//connalert := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, AlertDB)
	//DBalert, err := sql.Open("mysql", connalert)
	//if err != nil {
	//	fmt.Println("connection to mysql failed:", err)
	//	lib.LogHander("connection to mysql failed:", err)
	//	return
	//}
	//defer DBalert.Close()

	srows:=lib.CompareTables(DB,tbSource,tbDist)
	fmt.Println(".........................db compare finished! is "+ string(srows))
	if srows !=0{
		//update attack number attack_name: file=node_attack  sql=data_tampering
		upAttack:=lib.UpdateData(DB,"data_tampering")
		if upAttack==-1{
			lib.InfoHander("update attack number err! is -1")
			fmt.Println(".........................update attack number err!!")
		}
		lib.InfoHander("update attack number +1! is "+string(upAttack))
		fmt.Println(".........................update attack number +1! is "+string(upAttack))

		//update node_info
		//return 1 is sucessful.
		upNodeInfo:=lib.UpdateInfo(DB,nodeIp)
		if upNodeInfo!=1{
			lib.InfoHander("update node_info number err! is -1")
			fmt.Println(".........................update node_info err!!")
		}
		lib.InfoHander("update node_info +1! is "+string(upNodeInfo))
		fmt.Println(".........................update node_info +1! is "+string(upNodeInfo))


		//add attack_info
		//at_type is 2 or 3 , attack_name: 2 = node_attack;  3 = data_tampering
		//return 1 is sucessful.
		addInfo:=lib.InsertInfo(DB,"hacker",3)
		if addInfo==-1{
			lib.InfoHander("add attack_info err! is -1")
			fmt.Println(".........................add attack_info err!!")
		}
		lib.InfoHander("add attack_info is "+string(addInfo))
		fmt.Println(".........................add attack_info is "+string(addInfo))


		//add attack_log
		//at_type is 2 or 3 , attack_name: 2 = node_attack;  3 = data_tampering
		//return 1 is sucessful.
		addLog:=lib.InsertLog(DB,3)
		if addLog==-1{
			lib.InfoHander("add attack_log is err!")
			fmt.Println(".........................add attack_log is err!")
		}
		lib.InfoHander("add attack_info is "+string(addLog))
		fmt.Println(".........................add attack_info is "+string(addLog))


		//restore the table data
		resTable:=lib.RestoreData(DB,tbSource,tbDist)
		if resTable!=1{
			lib.InfoHander("restore db table err!")
			fmt.Println(".........................restore db table err!!")
		}
		fmt.Println(".........................db restore finished! is "+string(resTable))

		defer DB.Close()

		//send dingding
		msg:="崇仁存证节点受到数据篡改攻击，系统已阻断。"

		var sendd=lib.DingTalk{}
		sendd.AccessToken=DING_TOKEN
		sendd.Secret=DING_SECRET
		//defaults timeout is 30
		dt:=10
		dresult,err:=sendd.SendDingMsg(msg,dt)
		if err!=nil{
			lib.LogHander("send dinging faild: ",err)
			fmt.Println(".........................send dinging faild!")
		}
		fmt.Println(".........................send dinging success! http code "+string(dresult.ErrCode))

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
	//if sourcemd5==destinmd5{
	if string(sourcemd5)==string(destinmd5) {
		lib.InfoHander("the file md5 exec has equal. ")
		fmt.Println(".........................the file md5 exec has equal.")
	} else {
		//update attack number
		conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
		DB, err := sql.Open("mysql", conn)
		if err != nil {
			fmt.Println("connection to mysql failed:", err)
			lib.LogHander("connection to mysql failed:", err)
			return
		}

		//update attack number attack_name: file=node_attack  sql=data_tampering
		upAttack := lib.UpdateData(DB, "node_attack")
		if upAttack == -1 {
			lib.InfoHander("update attack number err! is -1")
			fmt.Println(".........................update attack number err!!")
		}
		lib.InfoHander("update attack file number +1! is " + string(upAttack))
		fmt.Println(".........................update attack number +1! is " + string(upAttack))
		defer DB.Close()

		//update node_info
		//return 1 is sucessful.
		upNodeInfo := lib.UpdateInfo(DB,nodeIp)
		if upNodeInfo != 1 {
			lib.InfoHander("update node_info number err! is -1")
			fmt.Println(".........................update node_info err!!")
		}
		lib.InfoHander("update node_info +1! is " + string(upNodeInfo))
		fmt.Println(".........................update node_info +1! is " + string(upNodeInfo))

		//add attack_info
		//at_type is 2 or 3 , attack_name: 2 = node_attack;  3 = data_tampering
		//return 1 is sucessful.
		addInfo := lib.InsertInfo(DB, "hacker", 2)
		if addInfo == -1 {
			lib.InfoHander("add attack_info err! is -1")
			fmt.Println(".........................add attack_info err!!")
		}
		lib.InfoHander("add attack_info is " + string(addInfo))
		fmt.Println(".........................add attack_info is " + string(addInfo))

		//add attack_log
		//at_type is 2 or 3 , attack_name: 2 = node_attack;  3 = data_tampering
		//return 1 is sucessful.
		addLog := lib.InsertLog(DB, 2)
		if addLog == -1 {
			lib.InfoHander("add attack_log is err!")
			fmt.Println(".........................add attack_log is err!")
		}
		lib.InfoHander("add attack_info is " + string(addLog))
		fmt.Println(".........................add attack_info is " + string(addLog))

		//exec restrofile to destetion
		rmStr := lib.CmdBash("rm -rf " + dirDist + "/* ")
		lib.InfoHander("exec rm: " + rmStr)
		cpStr := lib.CmdBash("cp -av " + dirSource + "/* " + dirDist)
		lib.InfoHander("exec cp: " + cpStr)

		//fmt.Println(err.Error())
		destinmd5, derr := lib.GetFileName(dirDist)
		if derr != nil {
			//fmt.Println(err.Error())
			lib.LogHander("exec faild: get dirDist md5 error ", derr)
			fmt.Println(".........................get dirDist md5 exception.")
		}
		lib.LogHander("exec faild: get dirDist md5 error ", derr)
		fmt.Println("dirDist: " + fmd + "==" + "dirDist: " + destinmd5)


		msg:="崇仁验证节点受到文件篡改攻击，系统已阻断。"
		var sendd=lib.DingTalk{}
		sendd.AccessToken=DING_TOKEN
		sendd.Secret=DING_SECRET
		//defaults timeout is 30
		dt:=10
		dresult,err:=sendd.SendDingMsg(msg,dt)
		if err!=nil{
			lib.LogHander("send dinging faild: ",err)
			fmt.Println(".........................send dinging faild!")
		}
		fmt.Println(".........................send dinging success! http code "+string(dresult.ErrCode))

	}
	fmt.Println("dirDist: "+fmd+"=="+"dirDist: "+destinmd5)
}

func showmd5(){
	//sourcemd5,err := lib.GetFileName("./source")
	//destmd5,err := lib.GetFileName("./destin")
	sourcemd5,err := lib.GetFileName(dirSource)
	destmd5,err := lib.GetFileName(dirDist)
	if err !=nil {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: show md5 error ")
		fmt.Println(".........................show files md5 exception.")
	}
	fmt.Println(".........................show source files md5: ")
	fmt.Println(sourcemd5)
	fmt.Println(".........................show destetion files md5: ")
	fmt.Println(destmd5)

}


func showssql(){
	//tb1:="crshow"
	//tb2:="crshow1"
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		lib.LogHander("connection to mysql failed:", err)
		return
	}
	defer DB.Close()

	//srows := lib.CompareData(DB,table1,table2)
	srows := lib.CompareTables(DB,tbSource,tbDist)
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

func hackersql(){
	//tb1:="crshow"
	//tb2:="crshow1"
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		lib.LogHander("connection to mysql failed:", err)
		return
	}
	defer DB.Close()

	upNodeInfo := lib.UpdateData(DB,"hacker")
	if upNodeInfo == -1 {
		lib.InfoHander("hacker sql number err! is -1")
		fmt.Println(".........................hacker sql number err!!")
	}
	lib.InfoHander("hacker sql number success! is " + string(upNodeInfo))
	fmt.Println(".........................hacker sql number success! is " + string(upNodeInfo))

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
			for {
				time.Sleep(3 * time.Second)
				rsql()
				fmt.Println("arg:", string(os.Args[1]))
				//fmt.Println("upgrade mysql table restore.")
			}

		}
		if string(os.Args[1])=="rfile"{
			for {
				time.Sleep(3 * time.Second)
				rfiles()
				fmt.Println("arg:", string(os.Args[1]))
				fmt.Println("upgrade show files restore.")
			}

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
		if string(os.Args[1])=="hsql"{
			//showsql()
			hackersql()
			fmt.Println("arg:", string(os.Args[1]))
			//fmt.Println("show sql type to ssql.")
		}
	}
}

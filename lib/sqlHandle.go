package lib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"upgrader/lib"
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

//插入数据 attack_info,
// at_type is 2 or 3 , attack_name: 2 = node_attack;  3 = data_tampering
func InsertInfo(DB *sql.DB,defender string,at_type int) int64 {
	//INSERT INTO `attack_defense_info`.`attack_info` (`defender`,`status`,`attack_type`,`createtime`) VALUES ('defender', 0, 2, NOW());
	//result,err := DB.Exec("INSERT INTO `attack_info`(`defender`,`status`,`attack_type`,`createtime`) VALUES ('defender', 0, 2, NOW()) ;","test","123456")
	sql:="INSERT INTO `attack_info`(`defender`,`status`,`attack_type`,`createtime`) VALUES (?, 0, ?, NOW())"
	if at_type ==2 {
		//esql="UPDATE all_attack SET 'count'='count'+1 WHERE attack_name=?"
		result,err := DB.Exec(sql,defender,at_type)
		if err != nil{
			//fmt.Printf("Insert attack_info failed,err:%v\n", err)
			LogHander("Insert attack_info failed,err:%v\n", err)
			return -1
		}
		//fmt.Println("Insert attack_info successd:", result)
		InfoHander("Insert attack_info successd.")

		rowsaffected,err := result.RowsAffected()
		if err != nil {
			//fmt.Printf("Get Insert attack_info RowsAffected failed,err:%v\n",err)
			LogHander("Get Insert attack_info RowsAffected failed,err:%v\n",err)
			return -1
		}
		//fmt.Println("Affected rows:", rowsaffected)
		InfoHander("Affected rows: "+string(rowsaffected))
		return rowsaffected
	}
	if at_type ==3 {
		//esql="UPDATE all_attack SET 'count'='count'+1 WHERE attack_name=?"
		result,err := DB.Exec(sql,defender,at_type)
		if err != nil{
			//fmt.Printf("Insert attack_info failed,err:%v\n", err)
			LogHander("Insert attack_info failed,err:%v\n", err)
			return -1
		}
		//fmt.Println("Insert attack_info successd:", result)
		InfoHander("Insert attack_info successd.")

		rowsaffected,err := result.RowsAffected()
		if err != nil {
			//fmt.Printf("Get Insert attack_info RowsAffected failed,err:%v\n",err)
			LogHander("Get Insert attack_info RowsAffected failed,err:%v\n",err)
			return -1
		}
		//fmt.Println("Affected rows:", rowsaffected)
		InfoHander("Affected rows: "+string(rowsaffected))
		return rowsaffected
	}
	return -1
}

//插入数据 attack_log
func InsertLog(DB *sql.DB,at_type int) int64 {
	//INSERT INTO `attack_defense_info`.`attack_log` (`attack_type`, `attack_node`,`node_ip`,`create_time`,`is_block`,`is_alarm`) VALUES(2,'attack_node','node_ip',NOW(),1,1) ;
	//result,err := DB.Exec("insert INTO users(username,password) values(?,?)","test","123456")
	sql:="INSERT INTO `attack_log`(`attack_type`, `attack_node`,`node_ip`,`create_time`,`is_block`,`is_alarm`) VALUES(?,'崇仁','192.168.1.3',NOW(),1,1)"
	if at_type ==2 {
		//esql="UPDATE all_attack SET 'count'='count'+1 WHERE attack_name=?"
		result,err := DB.Exec(sql,at_type)
		if err != nil{
			//fmt.Printf("Insert attack_info failed,err:%v\n", err)
			LogHander("Insert attack_info failed,err:%v\n", err)
			return -1
		}
		//fmt.Println("Insert attack_info successd:", result)
		InfoHander("Insert attack_info successd.")

		rowsaffected,err := result.RowsAffected()
		if err != nil {
			//fmt.Printf("Get Insert attack_info RowsAffected failed,err:%v\n",err)
			LogHander("Get Insert attack_info RowsAffected failed,err:%v\n",err)
			return -1
		}
		//fmt.Println("Affected rows:", rowsaffected)
		InfoHander("Affected rows: "+string(rowsaffected))
		return rowsaffected
	}
	if at_type ==3 {
		//esql="UPDATE all_attack SET 'count'='count'+1 WHERE attack_name=?"
		result,err := DB.Exec(sql,at_type)
		if err != nil{
			//fmt.Printf("Insert attack_info failed,err:%v\n", err)
			LogHander("Insert attack_info failed,err:%v\n", err)
			return -1
		}
		//fmt.Println("Insert attack_info successd:", result)
		InfoHander("Insert attack_info successd.")

		rowsaffected,err := result.RowsAffected()
		if err != nil {
			//fmt.Printf("Get Insert attack_info RowsAffected failed,err:%v\n",err)
			LogHander("Get Insert attack_info RowsAffected failed,err:%v\n",err)
			return -1
		}
		//fmt.Println("Affected rows:", rowsaffected)
		InfoHander("Affected rows: "+string(rowsaffected))
		return rowsaffected
	}
	return -1
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

//查询单行 , tb1 string,tb2 string   (int)
func TableOne(DB *sql.DB, tb1 string) {
	//querysql:=`SELECT * FROM (SELECT *  FROM ? UNION ALL SELECT * FROM ?) tbl GROUP BY showstring,STATUS,ca,createtime HAVING COUNT(*) = 1  LIMIT 1`
	querysql:="SELECT * FROM "+tb1+" limit 1"

	cs := new(Crshow)   //用new()函数初始化一个结构体对象
	//row := DB.QueryRow("SELECT * FROM crshow where ca=? ", tb1)
	//row := DB.QueryRow("SELECT * FROM ? limit 1 ", tb1)
	row := DB.QueryRow(querysql)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&cs.Showstring,&cs.Ca,&cs.Status,&cs.Createtime); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		//return -1
	}
	fmt.Println("Single row data:", *cs)
	//return 1
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

//更新attack数据 get all_attack table. attack_name: node_attack  data_tampering
func UpdateData(DB *sql.DB,who string) (int64){
	//result,err := DB.Exec("UPDATE xx set password=? where id=?","111111",1)
	//rowsaffected:=0
	esql:="UPDATE all_attack SET `count`=`count`+1 WHERE attack_name=?"
	//if who =="file" {
	if who =="node_attack" {
		//esql="UPDATE all_attack SET 'count'='count'+1 WHERE attack_name=?"
		result,err := DB.Exec(esql,who)
		if err != nil{
			//fmt.Printf("Insert failed,err:%v\n", err)
			LogHander("Update all_attack failed,err:%v\n", err)
			return -1
		}
		//fmt.Println("update data successd:", result)
		InfoHander("update data all_attack successd. ")

		rowsaffected,err := result.RowsAffected()
		if err != nil {
			//fmt.Printf("Get RowsAffected failed,err:%v\n",err)
			LogHander("Get RowsAffected failed,err:%v\n",err)
			return -1
		}
		//fmt.Println("Affected rows:", rowsaffected)
		InfoHander("Affected rows: "+string(rowsaffected))
		return rowsaffected
	}
	//if who == "sql"{
	if who == "data_tampering"{
		//esql="UPDATE all_attack SET 'count'='count'+1 WHERE attack_name=?"
		result,err := DB.Exec(esql,who)
		if err != nil{
			//fmt.Printf("Insert failed,err:%v\n", err)
			LogHander("Update all_attack failed,err:%v\n", err)
			return -1
		}
		//fmt.Println("update data successd:", result)
		InfoHander("update data all_attack successd. ")

		rowsaffected,err := result.RowsAffected()
		if err != nil {
			//fmt.Printf("Get RowsAffected failed,err:%v\n",err)
			LogHander("Get RowsAffected failed,err:%v\n",err)
			return -1
		}
		//fmt.Println("Affected rows:", rowsaffected)
		InfoHander("Affected rows: "+string(rowsaffected))
		return rowsaffected
	}
	return -1
}

//更新数据 node_info  waiting get attack table. attack_name: node_attack  data_tampering
func UpdateInfo (DB *sql.DB,nodeIp string) (int64){
	//result,err := DB.Exec("UPDATE xx set password=? where id=?","111111",1)
	//rowsaffected:=0
	//UPDATE node_info SET 'block_height'='block_height'+1
	//UPDATE node_info SET 'attack_count'='attack_count'+1 WHERE ip='192.168.1.3'
	asql:="UPDATE node_info SET `block_height`=`block_height`+1 "
	csql:="UPDATE node_info SET `attack_count`=`attack_count`+1 WHERE ip=?"

	aresult,aerr := DB.Exec(asql)
	//cresult,cerr := DB.Exec(csql,"192.168.1.3")
	cresult,cerr := DB.Exec(csql,nodeIp)
	if aerr != nil && cerr!=nil{
		//fmt.Printf("Insert failed,err:%v\n", err)
		LogHander("Update node_info failed,err:%v\n", aerr)
		LogHander("Update node_info failed,err:%v\n", cerr)
		return -1
	}
	//fmt.Println("update data successd:", result)
	InfoHander("update data node_info successd. ")

	arowsaffected,aerr := aresult.RowsAffected()
	crowsaffected,cerr := cresult.RowsAffected()
	if aerr != nil && cerr!=nil {
		//fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		LogHander("Get RowsAffected failed,err:%v\n",aerr)
		LogHander("Get RowsAffected failed,err:%v\n",cerr)
		return -1
	}
	//fmt.Println("Affected rows:", rowsaffected)
	InfoHander("Affected rows: "+string(arowsaffected))
	InfoHander("Affected rows: "+string(crowsaffected))
	return 1
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
func CompareTables(DB *sql.DB, tb1 string,tb2 string) (int) {
	//cs:=new(Crshow)
	var cs Crshow
	count:=0

	//querysql:="SELECT * FROM " +
	//	"(SELECT *  FROM ? UNION ALL SELECT * FROM ?) tbl " +
	//	"GROUP BY showstring,`status`, ca,`createtime` " +
	//	"HAVING COUNT(*) = 1"
	querysql:=`SELECT * FROM
		(SELECT *  FROM ` +tb1+` UNION ALL SELECT * FROM `+tb2+ `) tbl
		GROUP BY showstring,status,ca,createtime
		HAVING COUNT(*) = 1`
	//InfoHander(querysql)
	//querysql := "SELECT *  FROM ?"

	//rows,err:=DB.Query(querysql,tb1,tb2)
	rows,err:=DB.Query(querysql)

	//stmt,_:=DB.Prepare(querysql)
	//defer  stmt.Close()
	//rows,err:=stmt.Query(tb1,tb2)

	if err != nil{
		//fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		LogHander("exec sql failed,err:%v\n",err)
		return -1
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

	defer func() {
		if rows != nil {
			rows.Close()   //关闭掉未scan的sql连接
		}
	}()

	return count

}


//恢复数据 TRUNCATE TABLE crshow 
//INSERT INTO crshow SELECT * FROM crshow1;
func RestoreData(DB *sql.DB,tbsource string,tbdist string) (int){
	cleansql:="TRUNCATE TABLE "+tbdist
	writesql:="INSERT INTO "+tbdist+" SELECT * FROM "+tbsource+";"

	_,cerr := DB.Exec(cleansql)
	if cerr != nil{
		//fmt.Printf("Clean table crshow failed,err:%v\n", err)
		lib.LogHander("Clean table crshow failed,err:%v\n",cerr)
		return -1
	}
	_,werr := DB.Exec(writesql)
	if werr != nil{
		//fmt.Printf("Insert failed,err:%v\n", err)
		lib.LogHander("Insert table crshow failed,err:%v\n",werr)
		return -1
	}
	fmt.Println("update data successd:")
	lib.InfoHander("Restore data table crshow successd!")
	return 1

}


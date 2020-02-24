package lib

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CmdBash(cstr string) string{
	//strs:=""

	out,err:=exec.Command("/bin/bash", "-c", cstr).Output()

	if err != nil {
		LogHander(cstr+"cmd exec failed!", err)
		return "failed"
	}
	return string(out)
}

func CmdStr(cstr ...string) string{
	strs:=""
	for _, arge:=range cstr  {
		strs+=arge+" "
	}

	cmd:=exec.Command(strs)

	err :=cmd.Start()
	if err != nil {
		LogHander(strs+"cmd exec failed!", err)
		return "failed"
	}
	errc:=cmd.Wait()
	if errc!=nil{
		LogHander(strs+"cmd exec failed!", errc)
		return "failed"
	}
	return cmd.ProcessState.String()
}

func TarZxvf(str string) string {
	cmd:=exec.Command("tar","zxvf", str,"-C","/")
	err :=cmd.Run()
	if err != nil {
		LogHander("cmd exec failed!", err)
		return "failed"
	}
	return "sucesss"
}

func Chmod(str string) string {
	cmd:=exec.Command("tar","zxvf", str,"-C","/")
	err :=cmd.Run()
	if err != nil {
		LogHander("cmd exec failed!", err)
		return "failed"
	}
	return "sucesss"
}

func AddSource(apts string){
	df:=os.Remove("/etc/apt/sources.list")
	if df != nil {
		InfoHander("del apt file faild")
	}
	//f,err := os.Create("/etc/apt/sources.list")
	f,err:=os.OpenFile("/etc/apt/sources.list",os.O_RDWR|os.O_CREATE,0666)
	defer f.Close()

	if err !=nil {
		//fmt.Println(err.Error())
		LogHander("Create apt file faild",err)
	}

	conts:="deb [trusted=yes] http://"+apts+"/ octa18 test "
	contb:=[]byte(conts)
	_,err=f.Write(contb)
	if err!=nil {
		LogHander("Write apt file faild",err)
	}
}


//changes blackbox config
func SetBlackbboxConf(){
	//get the conncet ip
	hn:=GetPulicIP()
	//hn:= strings.Replace(CmdBash("hostname"),"\n","",-1)
	input,err:=ioutil.ReadFile("/8lab/conf/configure.json")
	if err!=nil{
		LogHander("read blackbox config err: ",err)
	}

	//content:=strings.Replace(string(input),"\n","",-1)
	content:=string(input)
	newcontent:=strings.Replace(content,"192.168.1.141",hn,-1)

	errw:=ioutil.WriteFile("/8lab/conf/configure.json",[]byte(newcontent),0)
	if errw!=nil{
		LogHander("wirte blackbox config err: ",errw)
	}
}

type ReplaceHandle struct {
	Root    string //根目录
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHandle) DoWrok() error {
	return filepath.Walk(h.Root, h.walkCallback)
}

func (h ReplaceHandle) walkCallback(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		//fmt.Pringln("DIR:",path)
		return nil
	}
	//文件类型需要进行过滤
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)
	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)
	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)
	return err
}

//get request public ip
func GetExternal() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	//LogHander(string(content),err)
	return string(content)
}

//get all local listen ip
func GetIntranetIp() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println("ip:", ipnet.IP.String())
				//LogHander(ipnet.IP.String(),err)
			}
		}
	}
}

//in order of dns to ensure public ip
func GetPulicIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx:= strings.LastIndex(localAddr, ":")
	//LogHander(localAddr[0:idx],err)
	return localAddr[0:idx]
}

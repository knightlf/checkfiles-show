package lib

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const ddir  = "/data/tmp"

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		//os.Chdir("/data/tmp")
		return false, nil
	}
	return false, err
}


func GetBin(wfile string,url string){
	//exist, err := PathExists(ddir)
	_, err := PathExists(ddir)
	if err!=nil{
		LogHander("get dir err!",err)
		return
	} else {
		InfoHander("create dir default is /data/tmp")
		err := os.Mkdir(ddir, os.ModePerm)
		if err !=nil{
			LogHander("create dir err!",err)
		}
	}

	fout,err:=os.Create(wfile)
	defer fout.Close()

	if err!=nil{
		LogHander("Create download file failed!",err)
	}

	res,err:=http.Get(url)
	if err!=nil{
		LogHander("http download file failed!",err)
	}

	buf:=make([]byte,1024)
	for{
		size,_:=res.Body.Read(buf)
		if size==0{
			break
		}else {
			fout.Write(buf[:size])
		}
	}
}

const bufferSize = 65536

// MD5sum returns MD5 checksum of one filename
func MD5sum(file string) (string) {
	//if info, err := os.Stat(filename); err != nil {
	//	return "", err
	//} else if info.IsDir() {
	//	return "", nil
	//}
	//
	//file, err := os.Open(filename)
	//if err != nil {
	//	return "", err
	//}
	//defer file.Close()

	hash := md5.New()
	bfile:=[]byte(file)
	hash.Write(bfile)

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum

	//for buf, reader := make([]byte, bufferSize), bufio.NewReader(file); ; {
	//	n, err := reader.Read(buf)
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//		return "", err
	//	}
	//
	//	hash.Write(buf[:n])
	//}
	//
	//checksum := fmt.Sprintf("%x", hash.Sum(nil))
	//return checksum, nil
}

func CreateMd5(filename string, md5str *string) {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open", err)
		return
	}

	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		fmt.Println("Copy", err)
		return
	}

	//*md5str=md5hash.Sum(nil)
	md5hash.Sum(nil)
	*md5str = fmt.Sprintf("%x", md5hash.Sum(nil))

	//InfoHander("this md5 create: \n"+string(*md5str))
	//fmt.Println("this md5 create: \n"+string(*md5str))

}

// MD5sum returns MD5 checksum of many files
func GetFileName(dir string) (string, error) {
	var md5str string
	//var md5str []byte

	//获取指定文件下的所有文件
	//var md5total []string
	var strmd5 string

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//去除子文件夹
			if info.IsDir() == false {
				//调用上面CreateMd5函数，为每个文件创建MD5，这里的path就是给定目录下的文件的绝对路径
				CreateMd5(path, &md5str)
				//md5total=append(md5total, md5str)
				strmd5=strmd5+md5str
				//把MD5sr存入redis
				//_, err = redisClient.Do("HSET", "XzWxClientMd5Sign", path, md5str)
				if err != nil {
					//log.Println("Set key err: ", err)
					LogHander("Set key err: ", err)
				}
			}
			return nil
		})
	if err != nil {
		//log.Println(err)
		LogHander("get dir err!",err)
	}

	//MD5 after sorting
	return MD5sum(strmd5),nil
	//sort.Strings(md5total)
	//for md5s := range md5total{
	//	strmd5=strmd5+string(md5s)
	//	InfoHander("this md5 sort: \n"+string(md5s))
	//	fmt.Println("this md5 sort: \n"+string(md5s))
	//}
	////MD5sum(strmd5)
	//return MD5sum(strmd5),nil

}

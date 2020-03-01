package lib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)


// DingTalk dingtalk client
type DingTalk struct {
	AccessToken string
	Secret      string
}

// Response response struct
type Response struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int64  `json:"errcode"`
}

const httpTimoutSecond = time.Duration(30) * time.Second
// https://oapi.dingtalk.com/robot/send?access_token=xxx
const dingTalkOAPI = "oapi.dingtalk.com"

var dingTalkURL url.URL = url.URL{
	Scheme: "https",
	Host:   dingTalkOAPI,
	Path:   "robot/send",
}

var timestamp = strconv.FormatInt(time.Now().Unix()*1000, 10)

// GetDingTalkURL get DingTalk URL with accessToken & secret
// If no signature is set, the secret is set to ""
// 如果没有加签，secret 设置为 "" 即可
func GetDingTalkURL(accessToken string, secret string) (string, error) {
	dtu := dingTalkURL
	value := url.Values{}
	value.Set("access_token", accessToken)

	if secret == "" {
		dtu.RawQuery = value.Encode()
		return dtu.String(), nil
	}

	sign, err := sign(timestamp, secret)
	if err != nil {
		dtu.RawQuery = value.Encode()
		return dtu.String(), err
	}

	value.Set("timestamp", timestamp)
	value.Set("sign", sign)
	dtu.RawQuery = value.Encode()
	return dtu.String(), nil
}

func sign(timestamp string, secret string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}


func (d *DingTalk) SendDingMsg(msg string) (Response, error) {
	//请求地址模板
	//safe tag: SEC3b6b33bb310bfcaf200d99fc47ef72030437287435bdaf1741531433da21fb67

	res := Response{}
	pushURL, err := GetDingTalkURL(d.AccessToken, d.Secret)
	if err != nil {
		LogHander("get ding sign err: ",err)
		return res, err
	}

	content := `{"msgtype": "text",
		"text": {"content": "`+ msg + `"}
	}`
	req, err := http.NewRequest("POST", pushURL, strings.NewReader(content))
	if err != nil {
		LogHander("get http request err: ",err)
		return res, err
	}
	req.Header.Add("Accept-Charset", "utf8")
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	client.Timeout = httpTimoutSecond
	resp, err := client.Do(req)
	if err != nil {
		LogHander("set http client err: ",err)
		return res, err
	}

	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogHander("get result byte err: ",err)
		return res, err
	}

	err = json.Unmarshal(resultByte, &res)
	if err != nil {
		//LogHander("byte json unmarshal err: ",err)
		return res, fmt.Errorf("unmarshal http response body from json error = %v", err)
	}

	if res.ErrCode != 0 {
		return res, fmt.Errorf("send message to dingtalk error = %s", res.ErrMsg)
	}

	return res, nil

	//sign=
	//webHook := `https://oapi.dingtalk.com/robot/send?access_token=dd84405981561e0f67af319ece4059f8d06fa56eb5f79d298443765c3024c95f`
	//content := `{"msgtype": "text",
	//	"text": {"content": "`+ msg + `"}
	//}`
	////创建一个请求
	//req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	//InfoHander("send msg is: \n"+content)
	//if err != nil {
	//	// handle error
	//	LogHander("get dingding url err: ",err)
	//}
	//
	//client := &http.Client{}
	////设置请求头
	//req.Header.Set("Content-Type", "application/json; charset=utf-8")
	////发送请求
	//resp, err := client.Do(req)
	////关闭请求
	//defer resp.Body.Close()
	//InfoHander("send msg is: \n"+resp.Status)
	//
	//if err != nil {
	//	// handle error
	//	LogHander("send alert message faild: ",err)
	//}
}


// Send message
//func (d *DingTalk) Send(message message.Message) (Response, error) {
//	res := Response{}
//
//	reqBytes, err := message.ToByte()
//	if err != nil {
//		return res, err
//	}
//
//	//pushURL, err := security.GetDingTalkURL(d.AccessToken, d.Secret)
//	pushURL, err := GetDingTalkURL(d.AccessToken, d.Secret)
//	if err != nil {
//		return res, err
//	}
//
//	req, err := http.NewRequest("POST", pushURL, bytes.NewReader(reqBytes))
//	if err != nil {
//		return res, err
//	}
//	req.Header.Add("Accept-Charset", "utf8")
//	req.Header.Add("Content-Type", "application/json")
//
//	client := new(http.Client)
//	client.Timeout = httpTimoutSecond
//	resp, err := client.Do(req)
//	if err != nil {
//		return res, err
//	}
//
//	resultByte, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return res, err
//	}
//
//	err = json.Unmarshal(resultByte, &res)
//	if err != nil {
//		return res, fmt.Errorf("unmarshal http response body from json error = %v", err)
//	}
//
//	if res.ErrCode != 0 {
//		return res, fmt.Errorf("send message to dingtalk error = %s", res.ErrMsg)
//	}
//
//	return res, nil
//}
//



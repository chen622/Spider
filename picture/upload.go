package picture

//sm.ms api for golang
//athor: https://github.com/sndnvaps/sm_ms_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	//"sync"
)

func main() {
	fmt.Printf("let's upload the file to https://sm.ms\n")
	msg, err := Upload("http://i2.hdslb.com/bfs/space/c919a9818172a8297f8b0597722f96504a1e1d88.jpg")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Upload %s\n", msg.Code)
		if msg.Msg != "" {
			fmt.Println("Msg = %s\n", msg.Msg)
		} else {
			fmt.Printf("Filename: %s\n", msg.Data.FileName)
			fmt.Printf("FileInfo: %d x %d\n", msg.Data.Width, msg.Data.Height)
			fmt.Printf("StoreName: %s\n", msg.Data.StoreName)
			fmt.Printf("Size: %d\n", msg.Data.Size)
			fmt.Printf("Path: %s\n", msg.Data.Path)
			fmt.Printf("Hash: %s\n", msg.Data.Hash)
			fmt.Printf("TimeStamp: %d\n", msg.Data.TimeStamp)
			fmt.Printf("Url: %s\n", msg.Data.Url)
			fmt.Printf("Delete: %s\n", msg.Data.Delete)
		}

	}
	//if err := os.Remove("picture/img/c919a9818172a8297f8b0597722f96504a1e1d88.png"); err != nil {
	//	panic(err)
	//}
}

//锁，用于 Upload() func
//var mutex = &sync.Mutex{}

//用于 Upload() func
type MsgBody struct {
	Code string   `json:"code"`
	Data DataInfo `json:"data,omitempty"`
	Msg  string   `json:"msg,omitempty"` //用于接收错误信息
}

//用于 ListUploadHistory() func
type HistoryMsgBody struct {
	Code string     `json:"code"`
	Data []DataInfo `json:"data"`
}

//用于获取上传图片的信息
type DataInfo struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	FileName  string `json:"filename"`
	StoreName string `json:"storename"`
	Size      int    `json:"size"`
	Path      string `json:"path"`
	Hash      string `json:"hash"`
	TimeStamp int64  `json:"timestamp"`
	Ip        string `json:"ip"`
	Url       string `json:"url"`
	Delete    string `json:"delete"`
}

//Check file suffix , only support jpeg,jpg,png,gif,bmp
func CheckFileSuffix(filename string) (bool, string) {
	suffix := path.Ext(filename)
	if (suffix == ".jpeg") || (suffix == ".jpg") ||
		(suffix == ".png") || (suffix == ".gif") ||
		(suffix == ".bmp") {
		return true, suffix
	}

	return false, suffix
}

//doc link  https://sm.ms/doc/
func Upload(filename string) (MsgBody, error) {
	//mutex.Lock()
	url := "https://sm.ms/api/upload"
	var msg MsgBody
	msg = MsgBody{
		Code: "Error",
		Msg:  "Internal function error",
	}
	reg, _ := regexp.Compile(`(\w|\d|_)*.(png|jpg|jpeg)`) //感谢ydy同学
	name := reg.FindStringSubmatch(filename)[0]
	fmt.Print(name)

	resp, _ := http.Get(filename)
	body, _ := ioutil.ReadAll(resp.Body)

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	fn, _ := filepath.Abs("picture/img/" + name)
	formFile, err := writer.CreateFormFile("smfile", fn)

	if err != nil {
		return msg, err
	}

	_, err = io.Copy(formFile, bytes.NewReader(body))
	if err != nil {
		return msg, err
	}

	writer.WriteField("ssl", "0")

	//发送表单
	contentType := writer.FormDataContentType()
	writer.Close() //发送之前必须调用Close()以写入结尾行
	resp, err = http.Post(url, contentType, buf)
	//mutex.Unlock()
	if err != nil {
		return msg, err
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return msg, err
		}

	}
	msg.Msg = "" //最后返回值，设置为 空
	return msg, nil
}

////用指定的 删除地址来 删除已经上传的图片
//func Delete(delUrlLink string) string {
//
//	url := delUrlLink
//	doc, _ := htmlquery.LoadURL(url)
//	resp_msg := htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='container']"))
//	fmt.Printf("Delete [%s] from  https://sm.ms\n", url)
//	return resp_msg
//}

//获得过去一小时内上传的文件列表
func ListUploadHistory() (HistoryMsgBody, error) {
	var msg HistoryMsgBody
	resp, err := http.Get("https://sm.ms/api/list")
	if err != nil {
		return msg, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &msg); err == nil {
		return msg, nil
	} else {
		return msg, err
	}

	return msg, nil
}

func Clear() (MsgBody, error) {
	var msg MsgBody
	resp, err := http.Get("https://sm.ms/api/clear")
	if err != nil {
		return msg, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &msg); err == nil {
		return msg, nil
	} else {
		return msg, err
	}

	return msg, nil
}

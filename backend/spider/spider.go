package spider

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"io/ioutil"
	"net/http"
	"time"
)

type Video struct {
	Author  string `jpath:"author"`
	Mid     int    `jpath:"mid"`
	Aid     int    `jpath:"aid"`
	Title   string `jpath:"title"`
	Created int64 `jpath:"created"`
}

func (v Video) GetTime() time.Time {
	unix := time.Unix(v.Created, 0)
	return unix
}

type Bilibili struct {
	Status bool    `jpath:"status"`
	Vlist  []Video `jpath:"data.vlist"`
}

const BILIBILI = "https://space.bilibili.com/ajax/member/getSubmitVideos?mid="

func GetVideoList(mid int) (list []Video, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprint(BILIBILI, mid), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36=anny")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(body), &dat); err != nil {
		panic(err)
	}

	if dat["status"].(bool) {
		var bilibili Bilibili
		mapstructure.DecodePath(dat, &bilibili)
		return bilibili.Vlist, nil
	} else {
		return nil, SpiderError{fmt.Sprint(BILIBILI, mid), dat["msg"].(string)}
	}
}

type SpiderError struct {
	Url string
	Msg string
}

func (err SpiderError) Error() string {
	return fmt.Sprintf("Error from: '%s', msg: '%s'", err.Url, err.Msg)
}

func main() {

}

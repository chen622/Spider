package spider

import (
	"Spider/database"
	"Spider/model"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
)

type Bilibili struct {
	Status bool                  `jpath:"status"`
	Vlist  []model.BilibiliVideo `jpath:"data.vlist"`
}

const bilibili_up_info = "https://api.bilibili.com/x/space/acc/info?jsonp=jsonp&mid="
const bilibili_video = "https://space.bilibili.com/ajax/member/getSubmitVideos?mid="

func GetUpInfo(mid int64) (*model.BilibiliUp, error) {
	up := model.BilibiliUp{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprint(bilibili_up_info, mid), nil)
	if err != nil {
		return &up, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36=anny")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &up, err
	}
	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		return &up, err
	}
	data = data["data"].(map[string]interface{})
	if err = mapstructure.Decode(data, &up); err != nil {
		return &up, err
	}
	database.DB.Create(&up)
	return &up, nil
}

/**
爬取up主的视频列表
*/
func GetVideoList(mid int64) (list []model.BilibiliVideo, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprint(bilibili_video, mid, "&pagesize=5&page=1"), nil)
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
		mapstructure.Decode(dat, &bilibili)
		for index, video := range bilibili.Vlist {
			bilibili.Vlist[index].Pic = fmt.Sprint("http:", video.Pic)
		}
		return bilibili.Vlist, nil
	} else {
		return nil, SpiderError{fmt.Sprint(bilibili_video, mid), dat["msg"].(string)}
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
	GetUpInfo(11336264)
}

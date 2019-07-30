package spider

import (
	"Spider/database"
	"Spider/model"
	"Spider/picture"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
)

type Bilibili struct {
	Status bool
	Data   struct {
		Vlist []model.BilibiliVideo
	}
	Msg string
}

const bilibili_up_info = "https://api.bilibili.com/x/space/acc/info?jsonp=jsonp&mid="
const bilibili_video = "https://space.bilibili.com/ajax/member/getSubmitVideos?mid="

func GetUpInfo(mid int64) (*model.BilibiliUp, error) {
	up := model.BilibiliUp{}

	//创建请求
	client := &http.Client{}

	if req, err := http.NewRequest("GET", fmt.Sprint(bilibili_up_info, mid), nil); err != nil {
		return nil, err
	} else {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36=anny")
		req.Header.Set("Accept", "application/json, text/plain, */*")
		if resp, err := client.Do(req); err != nil {
			return nil, err
		} else {
			defer resp.Body.Close()
			if body, err := ioutil.ReadAll(resp.Body); err != nil {
				return nil, err
			} else {

				//解析响应数据
				var data map[string]interface{}
				if err = json.Unmarshal(body, &data); err != nil {
					return nil, err
				}
				data = data["data"].(map[string]interface{})
				if err = mapstructure.Decode(data, &up); err != nil {
					return nil, err
				}

				//替换图片url
				if msg, err := picture.Upload(up.TopPhoto); err != nil {
					return nil, err
				} else {
					up.TopPhoto = msg.Data.Url
				}
				if msg, err := picture.Upload(up.Face); err != nil {
					return nil, err
				} else {
					up.Face = msg.Data.Url
					database.DB.Create(&up)
					return &up, nil
				}
			}
		}
	}
}

/**
爬取up主的视频列表
*/
func GetVideoList(mid uint64) (list []model.BilibiliVideo, err error) {

	//创建请求
	client := &http.Client{}
	if req, err := http.NewRequest("GET", fmt.Sprint(bilibili_video, mid, "&pagesize=5&page=1"), nil); err != nil {
		return nil, err
	} else {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36=anny")
		req.Header.Set("Accept", "application/json, text/plain, */*")
		if resp, err := client.Do(req); err != nil {
			return nil, err
		} else {
			defer resp.Body.Close()
			if body, err := ioutil.ReadAll(resp.Body); err != nil {
				return nil, err
			} else {
				//解析数据
				var bilibili Bilibili
				if err := json.Unmarshal([]byte(body), &bilibili); err != nil {
					return nil, err
				}
				if bilibili.Status {
					for index, video := range bilibili.Data.Vlist {
						bilibili.Data.Vlist[index].Pic = fmt.Sprint("http:", video.Pic)
					}
					return bilibili.Data.Vlist, nil
				} else {
					return nil, Error{fmt.Sprint(bilibili_video, mid), bilibili.Msg}
				}
			}
		}
	}
}

type Error struct {
	Url string
	Msg string
}

func (err Error) Error() string {
	return fmt.Sprintf("Error from: '%s', msg: '%s'", err.Url, err.Msg)
}

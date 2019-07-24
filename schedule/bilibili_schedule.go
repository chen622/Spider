package schedule

import (
	"Spider/database"
	"Spider/model"
	"Spider/spider"
	"Spider/utils"
	"fmt"
	"github.com/robfig/cron"
)

func New() {
	temp := cron.New()
	if err := temp.AddFunc("@every 1m", BilibiliSchedule); err != nil {
		fmt.Println(err)
	}
	temp.Start()
	//return temp
}

func BilibiliSchedule() {
	utils.Logger.Errorf("RUN SCHEDULE")
	biliUps := []model.BilibiliUp{}
	database.DB.Preload("Users").Find(&biliUps)
	for _, up := range biliUps {
		go refreshUp(&up)
	}
}

func refreshUp(up *model.BilibiliUp) error {
	hasNew := false
	if videos, err := spider.GetVideoList(up.ID); err != nil {
		return err
	} else {
		for _, video := range videos {
			if database.DB.First(&model.BilibiliVideo{}, video.ID).RecordNotFound() {
				database.DB.Model(&up).Association("BilibiliVideo").Append(video)
				hasNew = true
			}
		}
		if hasNew {

		}
		return nil
	}
}

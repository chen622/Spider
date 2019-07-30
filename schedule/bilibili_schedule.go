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
	utils.Logger.Info("RUN SCHEDULE")
	fmt.Println("RUN SCHEDULE")
	biliUps := []model.BilibiliUp{}
	database.DB.Preload("Users").Find(&biliUps)
	for _, up := range biliUps {
		go refreshUp(&up)
	}
}

func refreshUp(up *model.BilibiliUp) {
	hasNew := false
	if videos, err := spider.GetVideoList(up.ID); err != nil {
		utils.Logger.Error(err.Error())
	} else {
		for _, video := range videos {
			if database.DB.First(&model.BilibiliVideo{}, video.ID).RecordNotFound() {
				if err := database.DB.Model(&up).Association("BilibiliVideos").Append(video).Error; err != nil {
					utils.Logger.Error(err.Error())
				} else {
					hasNew = true
				}
			}
		}
		if hasNew {
			var users []model.User
			if err := database.DB.Model(&up).Association("Users").Find(&users).Error; err != nil {
				utils.Logger.Error(err.Error())
			} else {

			}
		}
	}
}

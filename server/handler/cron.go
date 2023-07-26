package handler

import (
	"fmt"
	"main/server/services/arena"
	"main/server/utils"

	"github.com/robfig/cron/v3"
)

func StartCron() {
	fmt.Println("Cron is running")
	c := cron.New()
	c.AddFunc(utils.EASY_PERK, func() {
		arena.ArenaPerk("easy")
	})
	c.AddFunc(utils.MEDIUM_PERK, func() {
		arena.ArenaPerk("medium")
	})
	c.AddFunc(utils.HARD_PERK, func() {
		arena.ArenaPerk("hard")
	})

	c.Start()
}

// @Desc: \\todo
// @Author: QianQingnian 2021/3/26 16:11

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	initUpCfg()
	rand.Seed(time.Now().Unix())
	upgrade()
}

type config struct {
	Level    int32
	Count    int32
	BaseRate int32
	AddRate  int32
}

var upCfg = make(map[int32]config, 16)

func initUpCfg() {
	upCfg[0] = config{0, 5, 50, 6}
	upCfg[1] = config{1, 10, 40, 5}
	upCfg[2] = config{2, 15, 30, 4}
	upCfg[3] = config{3, 20, 20, 3}
	upCfg[4] = config{4, 25, 10, 2}
	upCfg[5] = config{5, 30, 5, 1}
	//upCfg[6] = config{6, 60, 40, 3}
	//upCfg[7] = config{7, 70, 36, 3}
	//upCfg[8] = config{8, 80, 32, 3}
	//upCfg[9] = config{9, 90, 28, 3}
	//upCfg[10] = config{10, 100, 24, 2}
	//upCfg[11] = config{11, 110, 20, 2}
	//upCfg[12] = config{12, 120, 16, 2}
	//upCfg[13] = config{13, 130, 12, 2}
	//upCfg[14] = config{14, 140, 8, 1}
	//upCfg[15] = config{15, 150, 4, 1}
}

func upgrade() {
	var level, targetLevel int32 = 0, 1
	var count int32 = 0

	for level < 5 {
		log := ""

		cfg := upCfg[level]
		rate := cfg.BaseRate
		targetCfg := upCfg[targetLevel]
		log += fmt.Sprintf("当前等级:%d 最高等级:%d 总次数:%d", level, targetLevel, count)
		log += fmt.Sprintf(" 最高等级次数:%d 概率:%d", targetCfg.Count, cfg.BaseRate)
		if count > targetCfg.Count {
			rate += cfg.AddRate * (count - targetCfg.Count)
			log += fmt.Sprintf(" 总次数超过最高等级次数 总概率:%d 额外加的概率:%d*%d=%d", rate, cfg.AddRate, count-targetCfg.Count, cfg.AddRate*(count-targetCfg.Count))
		}
		random := rand.Int31n(100) + 1
		log += fmt.Sprintf(" 随机数:%d", random)
		count += 1
		if random <= rate {
			level += 1
			log += fmt.Sprintf(" 升级成功,新等级:%d", level)
			if targetLevel <= level {
				targetLevel = level + 1
				log += fmt.Sprintf(" 最高等级升一级")
			}
			log += fmt.Sprintf(" 最高等级:%d", targetLevel)
			fmt.Println(log)
		} else {
			var loss int32
			if level > 0 {
				loss = rand.Int31n(level) + 1
				level -= loss
			}
			log += fmt.Sprintf(" 升级失败掉%d级 新等级:%d 最高等级:%d", loss, level, targetLevel)
			fmt.Println(log)
			continue
		}
	}
}

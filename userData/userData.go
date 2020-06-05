package userData

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iotqq-plugins-demo/Go/achievement"
	"iotqq-plugins-demo/Go/building"
	"iotqq-plugins-demo/Go/cards"
	"iotqq-plugins-demo/Go/common"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var UserMap sync.Map
var MaxCollectionNum = 0

type User struct {
	Udid                int64
	SummonCardNum       int
	Water               int
	UnHitNumber         int
	CardIndex           []int
	BuildIndex          []common.BuildRecord
	AchievementList     []common.AchievementRecord
	LastVolunterGetTime time.Time
	Static              Static
}

type Static struct {
	VolunterReiceiveTime int
	VolunterReiceiveMax  int
}

var userinfoPath = "d:\\userinfo"

func GetUser(udid int64) *User {
	user, _ := UserMap.LoadOrStore(udid, &User{
		Udid:          udid,
		SummonCardNum: 500,
	})
	return user.(*User)
}

func UserRange(f func(key, value interface{}) bool) {
	UserMap.Range(f)
}

func UserDataSave() {
	//GetUser(10000)
	fmt.Println("enter UserDataSave")
	//str, _ := os.Getwd()
	//fmt.Println(str)
	s, err := os.Stat(userinfoPath) //os.Stat获取文件信息
	if err != nil || !s.IsDir() {
		fmt.Println("enter mkdir")
		os.Remove(userinfoPath)
		os.Mkdir(userinfoPath, 0)
	}
	UserMap.Range(func(key, value interface{}) bool {
		//fmt.Println("enter UserMap")
		f, err := os.Create(userinfoPath + "\\" + strconv.Itoa(int(key.(int64))) + ".data")
		defer f.Close()

		if err != nil {
			fmt.Println(err.Error())
		} else {
			b, err := json.Marshal(value)
			_, err = f.Write(b)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		return true
	})
}

func UserDataLoad() {
	s, err := os.Stat(userinfoPath)
	if err != nil {
		fmt.Println("could not find userinfo", err.Error())
		return
	}

	if !s.IsDir() {
		fmt.Println("userinfo is not a dir")
	}

	filepath.Walk(userinfoPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println("userdata path is " + path)
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("could not open file", err)
		}
		//fmt.Println("filecontent is ", string(b))
		var user User
		err = json.Unmarshal(b, &user)
		if err != nil {
			fmt.Println("unmarshal faild", err)
			return nil
		}
		UserMap.Store(user.Udid, &user)
		//fmt.Println(path,info)
		return nil
	})
}

func (u *User) GetAccountInfo() string {
	res := ""
	res += fmt.Sprintf("资产一览 召唤卷:%d🎟,水滴:%d💧", u.SummonCardNum, u.Water)
	return res
}

func (u *User) GetCollection() string {
	res := ""
	//c := cards.GetCardsAnalysis(u.CardIndex)
	//res += fmt.Sprintf("图鉴一览:五星角色%d/%d,四星角色%d/%d,三星角色%d/%d\n",
	//	c[0], common.FiveStarCharacterNum, c[1], common.FourStarCharacterNum, c[2], common.ThreeStarCharacterNum)
	//res += fmt.Sprintf("五星龙%d/%d,四星龙%d/%d,三星龙%d/%d",
	//	c[3], common.FiveStarDragonNum, c[4], common.FourStarDragonNum, c[5], common.ThreeStarDragonNum)
	res += fmt.Sprintf("图鉴完成度:%d/%d", len(u.CardIndex), len(cards.Cards))
	return res
}

func (u *User) GetBuildInfo() string {
	if len(u.BuildIndex) <= 0 {
		return "建筑无"
	}
	var res string
	var item []string
	for _, b := range u.BuildIndex {
		item = append(item, fmt.Sprintf("%slv%d", building.BuildList[b.Index].Title, b.Level))
	}
	res += fmt.Sprintf("拥有的建筑:%s;", strings.Join(item, ","))
	eff := building.GetBuildEffect(u.BuildIndex)
	if eff.VolunterMineProduct != 0 {
		ft := u.LastVolunterGetTime.Add(common.VolunterMineProductPeriod).Sub(time.Now())
		res += fmt.Sprintf("金币矿山将在%d小时%d分钟之后产出%d🎟", int(ft.Minutes())/60, int(ft.Minutes())%60, eff.VolunterMineProduct)
	}
	return res
}

func (u *User) GetMyHitRate(nickName string) string {
	return fmt.Sprintf("%s殿下的概率:%.1f%%,继续%d次召唤提高概率", nickName, float32(common.BaseSSRProbality+u.UnHitNumber/10*5)/10, 10-u.UnHitNumber%10)
}

func (u *User) GetHitRate() int {
	return u.UnHitNumber
}

func (u *User) GetStatic() string {
	return fmt.Sprintf("被赠送%d次,最高被赠送%d张", u.Static.VolunterReiceiveTime, u.Static.VolunterReiceiveMax)
}

func (u *User) GetAchievement() string {
	m := map[int]time.Time{}
	for i := range u.AchievementList {
		m[u.AchievementList[i].Index] = u.AchievementList[i].AchievementTime
	}
	outStr := ""
	outStr += fmt.Sprintf("%-24s%10s\n", "成就名称", "达成时间")
	for i, value := range achievement.AchievementList {
		time := "未完成"
		if value, ok := m[i]; ok {
			time = value.Format("2006-01-02 15:04:05")
		}
		outStr += fmt.Sprintf("%-24s%10s\n", value.Title, time)
	}
	return outStr
}

func (u *User) Achieve(id int) bool {
	for _, record := range u.AchievementList {
		if record.Index == id {
			return false
		}
	}
	u.AchievementList = append(u.AchievementList, common.AchievementRecord{
		Index:           id,
		AchievementTime: time.Now(),
	})
	return true
}

package wordTriggerBot

import (
	"math"
	"regexp"
	"time"
)

type wordTriggerRule struct {
	regex       *regexp.Regexp
	probability int
	times       int
	coolDown    time.Duration
	response    string
}

type wordTriggerState struct {
	wordTriggerRule
	lastTriggerTime time.Time
	triggerTimes    int
}

var stateList []*wordTriggerState

func init() {
	NewRule("天堂", "天堂武藏抽了么?😊", 20, math.MaxInt64, time.Minute*1)
	NewRule("复读", "复读我很专业呢", 30, math.MaxInt64, time.Minute*5)
	NewRule("我想要(.*?)召唤券", "抽卡要氪金的啊", 30, math.MaxInt64, time.Hour*1)
	NewRule("机器人", "虽然我是机器人,但是希望叫我修玛吉亚", 10, math.MaxInt64, time.Hour*1)
	NewRule("Du娘", "叫我修玛吉亚", 2, math.MaxInt64, time.Hour*24)
	NewRule("[Dd]ulang", "Dulang?把我弄的满身bug的大叔,一定在摸鱼呢", 70, math.MaxInt64, time.Hour*10)
	NewRule("(.*?)有(.*?)妹妹", "反正我没有妹妹...想要一个妹妹..", 100, math.MaxInt64, time.Hour*12)
	NewRule("有(.*?)车[吗嘛]", "我帮你招募一下如何?", 100, math.MaxInt64, time.Minute*5)
	NewRule("^\\?$", "为什么要单打一个问号呢？你有遇上什么烦恼吗，或许我可以帮你...", 80, math.MaxInt64, time.Minute*5)
}

func NewRule(regex, resp string, probability, times int, coolDown time.Duration) {
	r := regexp.MustCompile(regex)
	stateList = append(stateList, &wordTriggerState{
		wordTriggerRule: wordTriggerRule{r, probability, times, coolDown, resp},
		lastTriggerTime: time.Time{},
		triggerTimes:    0,
	})
}

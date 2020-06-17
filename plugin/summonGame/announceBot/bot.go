package announceBot

import (
	"iotqq-plugins-demo/Go/plugin"
)

func init() {
	plugin.FactoryInstance.RegisterPlugin(&announceBot{3})
}

type announceBot struct {
	priority int //[0~1000)
}

func (a *announceBot) IsTrigger(req *plugin.Request) (res bool, vNext bool) {
	//fmt.Println(req.Content)
	if req.IsAtMe && req.Content == "公告" {
		return true, false
	}
	return false, true
}

func (a *announceBot) Process(req *plugin.Request) []*plugin.Result {
	res := &plugin.Result{}
	//out := "现在的卡池:fes卡池,利弗,三藏,八戒 0.5% 铁扇公主 0.8% 其他fes小概率"
	out := "建筑功能 觉醒之岚树,水祭坛,金币矿山 \n@修玛吉亚-Du 建造[建筑名称] 触发\n觉醒之岚树影响赠送召唤卷的数量,每一级提高赠送量\n建造费用等于等级*10w💧\n"
	out += "水祭坛影响被复读赠送召唤卷的概率,每一级微小提高概率,累计赠送次数会极微小的减少概率\n建造费用等于等级*20w💧\n"
	out += "金币矿山每六小时赠送50张🎟,如果有任何事件触发🎟赠送逻辑,金币矿山的计时将重新从六小时开始计算,每升一级提高赠送🎟的数量,\n建造费用等于等级*20w💧\n"
	res.Content = out
	return []*plugin.Result{res}
}

func (a *announceBot) Priority() int {
	return a.priority
}

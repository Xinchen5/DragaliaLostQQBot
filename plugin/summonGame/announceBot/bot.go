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
	out := "卡池变化:大幅增加了fes卡池出现的概率,但是总体概率都是跟随着现在游戏卡池概率,也就是fes卡池的虹率也是4%\n"
	res.Content = out

	res2 := &plugin.Result{}
	out2 := "建筑功能 觉醒之岚树,水祭坛,金币矿山 \n@修玛吉亚-Du 建造[建筑名称] 触发\n觉醒之岚树影响赠送召唤卷的数量,每一级提高赠送量\n建造费用等于等级*10w💧\n"
	out2 += "水祭坛影响被复读赠送召唤卷的概率,每一级微小提高概率,累计赠送次数会极微小的减少概率,在一个金币矿山周期内被赠送的次数微小减少赠券几率,金币矿山出现产出后刷新该次数\n建造费用等于等级*20w💧\n"
	out2 += "金币矿山每六小时赠送50张🎟,如果有任何事件触发🎟赠送逻辑,金币矿山的计时将重新从六小时开始计算,每升一级提高赠送🎟的数量,\n建造费用等于等级*20w💧\n"
	res2.Content = out2
	return []*plugin.Result{res, res2}
}

func (a *announceBot) Priority() int {
	return a.priority
}

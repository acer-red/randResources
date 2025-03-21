package randresources

import (
	"math/rand"
	"time"
)

var humorousAdjectives = []string{
	"会飞的", "忧郁的", "吵闹的", "迷路的", "聪明的", "傻傻的", "奇怪的", "贪吃的", "懒惰的", "爱放屁的",
	"倒霉的", "戏精的", "傲娇的", "呆萌的", "脑洞大的", "戏很多的", "认真的", "离谱的", "魔性的",
	"土味的", "憨憨的", "社恐的", "社牛的", "摸鱼的", "暴躁的", "佛系的", "咸鱼的",
	"嘴强的", "手残的", "戏多的", "脑回路清奇的", "眼神犀利的",
}

var humorousNouns = []string{
	"香蕉", "幽灵", "企鹅", "袜子", "拖鞋", "键盘", "鼠标", "猫", "熊", "猩猩", "猴子", "狮子", "老虎",
	"吃货", "程序猿", "铲屎官", "网红", "干饭人", "打工人",
	"瘦猴", "咕咕叽", "吃瓜群众", "气氛组", "显眼包", "小透明", "冤种", "大冤种", "社畜",
}

var humorousAdverbs = []string{
	"疯狂地", "悄悄地", "大声地", "笨拙地", "优雅地", "得意地", "伤心地", "兴奋地", "好奇地", "漫无目的地", "无聊地", "热情地", "冷静地", "狡猾地",
	"突然地", "认真地", "默默地", "夸张地", "一本正经地", "猝不及防地", "明目张胆地", "理直气壮地", "偷偷摸摸地", "心安理得地", "不屑一顾地", "不知所措地",
	"兴高采烈地", "垂头丧气地", "若无其事地", "装模作样地", "煞有介事地", "鬼鬼祟祟地", "慢悠悠地", "狠狠地", "优雅地", "激情澎湃地", "敷衍地", "懒散地", "疲惫地", "梦想着", "假装",
}

var humorousVerbs = []string{
	"跳舞", "唱歌", "思考", "睡觉", "奔跑", "飞翔", "爬行", "尖叫",
	"喜欢", "讨厌", "寻找", "摸鱼", "躺平", "内卷", "吐槽",
	"点赞", "转发", "围观", "裂开", "emo了", "蚌埠住了", "栓Q",
}

var humorousSuffixes = []string{
	"侠", "大人", "怪", "兽", "精", "王", "子",
	"狗子", "铁子", "老铁", "大佬", "小可爱", "憨憨",
}

func getRandomElement(slice []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(slice))
	return slice[index]
}

func Text() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var nickname string
	pattern := r.Intn(7)

	switch pattern {
	case 0: // 形容词 + 名词
		nickname = getRandomElement(humorousAdjectives) + getRandomElement(humorousNouns)
	case 1: // 名词 + 动词
		nickname = getRandomElement(humorousNouns) + getRandomElement(humorousVerbs)
	case 2: // 形容词 + 动词 + 名词
		nickname = getRandomElement(humorousAdjectives) + getRandomElement(humorousVerbs) + getRandomElement(humorousNouns)
	case 3: // 名词 + 形容词
		nickname = getRandomElement(humorousAdjectives) + getRandomElement(humorousNouns)
	case 4: // 名词 + 名词
		nickname = getRandomElement(humorousNouns) + getRandomElement(humorousNouns)
	case 5: // 名词 + 动词 + 名词
		nickname = getRandomElement(humorousNouns) + getRandomElement(humorousVerbs) + getRandomElement(humorousNouns)
	case 6: // 副词 + 动词
		nickname = getRandomElement(humorousAdverbs) + getRandomElement(humorousVerbs)
	default: // 形容词 + 名词 + 状态
		nickname = getRandomElement(humorousAdjectives) + getRandomElement(humorousNouns)
	}

	// 随机添加后缀 (降低概率)
	if rand.Float64() < 0.15 {
		nickname += getRandomElement(humorousSuffixes)
	}

	return nickname
}

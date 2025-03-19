# RandResources

RandResources 是一个随机昵称和随机头像的golang工具库。

RandResources A golang tool library for random nicknames and random avatars.

## quick start

```bash
go get github.com/acer-red/randResources
```

## 用法

### 生成随机图片

使用 `image.go` 文件中的函数生成随机图片。以下是一个简单的示例：

```go
package main

import (
    "fmt"
    rs "github.com/acer-red/randResources"
)

func main() {
    img,err := rs.NewIBuildImagemage("your_email_or_nickanme")
    if err !=nil{
        panic(err)
    }

    // 输出string的BASE64
    _ := img.Base64()

    // 输出为[]byte
    _ := img.Bytes()

    // 输出到本地文件
    if err := img.Save("your_file_path"); err !=nil{
        panic(err)
    }
}
```

### 生成随机文本

使用 `text.go` 文件中的函数生成随机文本。以下是一个简单的示例：

```go
package main

import (
    "fmt"
    rs "github.com/acer-red/randResources"

)

func main() {
    nickname := rs.Text()
    fmt.Println(nickname)
}
```

## 输出示例

### 随机图片

生成的随机图片将以图像对象的形式输出，可以保存为文件或直接在应用中使用。

![随机图片示例](example/random.png)

### 随机文本

```text
得意地内卷
猩猩思考
鼠标飞翔气氛组怪
佛系的狮子
倒霉的拖鞋
戏很多的梦想着非酋
优雅地放屁
瘦猴耳语
冤种尖叫
非酋键盘喷子
显眼包耳语
脑洞大的显眼包
傲娇的跳舞冤种
社畜熊猫
键盘侠内卷
戏多的摸鱼狮子
香蕉干饭人
熊猫老虎
聪明的幽灵
伤心地蚌埠住了
社畜讨厌
垂头丧气地躺平
小透明摸鱼
熬夜冠军裂开
离谱的吐槽柠檬精
佛系的思考气氛组
社畜狮子
认真的企鹅
优雅地讨厌
迷路的社畜狗子
迷路的香蕉
认真的非酋
理直气壮地喜欢
憨憨的干饭人
幽灵鼠标
气氛组耳语铁子
咕咕机假装社畜
魔性的鼠标
会飞的狮子
无聊地蚌埠住了子
大声地裂开老铁
手残的假装冤种
程序猿键盘侠
企鹅吃瓜群众
嘴强的喜欢干饭人
狗程序猿
默默地点赞
咸鱼的香蕉
脑回路清奇的emo了老虎
鼠标耳语
社牛的睡觉杠精
认真的吐槽猪
暴躁的秃头怪
兴奋地裂开
疯狂地尖叫
暴躁的香蕉
懒惰的显眼包
铲屎官耳语
戏多的瘦猴
拖鞋气氛组狗子
小透明转发秃头怪
戏精的转发秃头怪兽
肥宅拖鞋
嘴强的打工人
嘴强的气氛组
猴子喜欢猴子
迷路的放屁干饭人
香蕉讨厌狮子
魔性的键盘
鼠标思考
瘦猴飞翔
小透明幽灵憨憨
咸鱼的非酋
猪吐槽狮子
鼠标放屁熬夜冠军
猴子点赞
呆萌的拖鞋
瘦猴非酋
戏很多的鼠标
傻傻的干饭人
兴高采烈地睡觉
眼神犀利的秃头怪
傻傻的熊
眼神犀利的显眼包怪
吃瓜群众猩猩
不屑一顾地奔跑
傻傻的冤种
懒散地耳语
吃瓜群众睡觉
小透明放屁
打工人思考社畜
激情澎湃地耳语精
离谱的猫
袜子尖叫
佛系的秃头怪
认真的点赞猪
戏多的假装非酋
干饭人摸鱼
认真的栓Q香蕉
魔性的小透明
```

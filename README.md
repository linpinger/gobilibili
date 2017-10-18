# gobilibili fork by fox

## 缘起

- 最近决定将Linux作为主系统，大部分需求都解决了，就是看B站直播没有弾幕比较苦恼

- 在github上找到了这个项目，正好懂一点golang，修改一下很好用

- 后期，准备将字幕加上颜色来区分各种弾幕，这个有现成的库

## 安装

    go get github.com/linpinger/gobilibili

## ChangeLog

- 2017-10-13: 修改 example/main.go 使用命令行 main roomID 来显示弾幕

*******************************

# gobilibili

B 站直播弹幕 Go 版。

## 安装

    go get github.com/lyyyuna/gobilibili

## 简单说明

参考 example/main.go，打印实时直播弹幕。

程序逻辑来自我的 [B 站直播弹幕姬 Python 版](https://github.com/lyyyuna/bilibili_danmu)。逻辑基本和 Python 版本保持一致，可以对着理解。

正如 [B 站直播弹幕姬 Python 版](https://github.com/lyyyuna/bilibili_danmu) 指出的，B 站直播协议会变，不保证向后兼容性。

在 example 的 main.go 中，roomid 没有自动转换，对于部分 UP 主，不要使用短 ID，请使用他的原始 ID。（Python 的版本做了处理，Go 我偷懒了。）

## 写后感

写并发，Go 比 Python 的 asyncio 舒服。但由于我是 Go 新手，总体上，写起来不舒服，Json 的处理也稍显麻烦。

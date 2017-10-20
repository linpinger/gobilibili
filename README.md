# gobilibili fork by fox

## 缘起

- 最近决定将Linux作为主系统，大部分需求都解决了，就是看B站直播没有弾幕比较苦恼

- 在github上找到了这个项目，正好懂一点golang，修改一下很好用

- **修改:** 字幕加上颜色来区分弾幕与系统消息，这个调用了这个 https://github.com/daviddengcn/go-colortext

## 下载并编译

    go get   github.com/linpinger/gobilibili
	./bin/gobilibili 97202

- 最小编译: go build -ldflags "-s -w" github.com/linpinger/gobilibili

**预编译版的下载地址:**
- 目前包含 linux x86/x64位, win 32位
- 七牛(速度快，就是每月有流量限制): http://linpinger.qiniudn.com/prj/gobilibili-bin.zip
- SF(慢，好像没有限制): http://master.dl.sourceforge.net/project/foxtestphp/prj/gobilibili-bin.zip

**用法:**
- gobilibili 97202

- 这个 97202 是真实的房间号，有时候这个跟网址后面的数字是不一样的

- 例如 http://live.bilibili.com/225  这里的 225 不是真实房间号，在浏览器上点右键，查看源代码，找到接近头部的

```javascript
    <script>
        document.domain = 'bilibili.com';

        var ROOMID = 97202;
        var DANMU_RND = 1508483734;
        var NEED_VIDEO = 1;
        var ROOMURL = 225;
        var INITTIME = Date.now();
        var p2p = 0;
    </script>
```

- 这里的ROOMID 就是 97202，而不是 ROOMURL 225

## ChangeLog

- 2017-10-20: 修正只能编译为64位的问题，原因是uid数字比int大，修改为int64，这样就ok了

- 2017-10-18: 修改弾幕颜色

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

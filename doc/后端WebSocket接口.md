* [流程](#流程)
* [WebSocket数据格式](#websocket数据格式)
  * [请求](#请求)
  * [响应](#响应)
    * [请求成功时的响应](#请求成功时的响应)
    * [出现错误时的响应](#出现错误时的响应)
  * [弹幕和信号](#弹幕和信号)
* [心跳包](#心跳包)
* [登陆](#登陆)
* [命令类型](#命令类型)
  * [获取弹幕](#获取弹幕)
  * [停止获取弹幕](#停止获取弹幕)
  * [直播间观众列表](#直播间观众列表)
  * [礼物贡献榜](#礼物贡献榜)
  * [直播总结信息](#直播总结信息)
  * [抢红包结果（未完成）](#抢红包结果未完成)
  * [直播回放](#直播回放)
  * [全部礼物列表](#全部礼物列表)
  * [账户钱包](#账户钱包)
  * [登陆用户的房管列表](#登陆用户的房管列表)
  * [添加房管](#添加房管)
  * [删除房管](#删除房管)
  * [踢人的历史记录（未完成）](#踢人的历史记录未完成)
  * [房管踢人](#房管踢人)
  * [主播踢人](#主播踢人)
* [弹幕和信号类型](#弹幕和信号类型)
  * [弹幕类型](#弹幕类型)
    * [弹幕](#弹幕)
    * [点赞](#点赞)
    * [进入直播间](#进入直播间)
    * [关注主播](#关注主播)
    * [投蕉](#投蕉)
    * [礼物](#礼物)
    * [富文本](#富文本)
    * [加入守护团](#加入守护团)
  * [状态信号类型](#状态信号类型)
    * [获取弹幕结束](#获取弹幕结束)
    * [直播间收到香蕉总数](#直播间收到香蕉总数)
    * [在线观众和点赞数量](#在线观众和点赞数量)
    * [在线观众前三名](#在线观众前三名)
    * [进直播间时显示的最近弹幕](#进直播间时显示的最近弹幕)
    * [红包列表](#红包列表)
    * [Chat Call](#chat-call)
    * [Chat Accept](#chat-accept)
    * [Chat Ready](#chat-ready)
    * [Chat End](#chat-end)
  * [通知信号类型](#通知信号类型)
    * [被踢出直播间](#被踢出直播间)
    * [ViolationAlert](#violationalert)
    * [登陆帐号的房管状态](#登陆帐号的房管状态)
* [错误码](#错误码)

### 流程
1. 客户端连接服务端，默认端口为15368
2. 客户端每5秒发送一次[心跳包](#心跳包)给服务端，接收服务端的心跳包
3. 客户端请求[登陆](#登陆)AcFun帐号，一个连接只能同时登陆一个帐号
4. 客户端发送[命令](#命令类型)请求，接受服务端的响应

客户端请求[获取直播间弹幕](#获取弹幕)后，服务端会不断发送[弹幕和信号数据](#弹幕和信号类型)给客户端。可同时请求多个直播间的弹幕。

### WebSocket数据格式
#### 请求
```json
{
    "type": 110,
    "requestID": "abc",
    "data": {}
}
```

`type`：请求类型

`requestID`：请求ID，用于区分同一类型的请求，具体内容由客户端决定

`data`：请求的数据，可选，一些请求不需要

#### 响应
##### 请求成功时的响应
```json
{
    "type": 110,
    "requestID": "abc",
    "result": 1,
    "data": {}
}
```

`type`：响应类型，和请求类型一样

`requestID`：请求ID，用于区分同一类型的请求，内容和对应的请求一样

`result`：请求结果，请求成功时为1

`data`：响应的数据，可选，绝大部分响应都会有

##### 出现错误时的响应
```json
{
    "type": 110,
    "requestID": "abc",
    "result": 11,
    "error": "error message"
}
```

`type`：响应类型，和请求类型一样，不需要时可为0

`requestID`：请求ID，用于区分同一类型的请求，内容和对应的请求一样

`result`：请求结果，出现错误时（不为1）为[错误码](#错误码)

`error`：错误信息

#### 弹幕和信号
```json
{
    "liverUID": 12345,
    "type": 1100,
    "data": {}
}
```

`liverUID`：主播uid

`type`：弹幕和信号类型

`data`：弹幕和信号的数据

[弹幕和信号数据](#弹幕和信号类型)由服务端发送到客户端，在客户端请求[获取弹幕](#获取弹幕)后才会发送

### 心跳包
```json
{
    "type": 1
}
```

`type`：心跳包类型为1

客户端每5秒发送一次心跳包给服务端，服务端收到心跳包后会发送心跳包给客户端

### 登陆
#### 请求
```json
{
    "type": 2,
    "requestID": "abc",
    "data": {
        "account": "account",
        "password": "password"
    }
}
```

`account`：AcFun帐号邮箱或手机号

`password`：AcFun帐号密码

`account`或`password`为空时以匿名游客的身份登陆

#### 响应
```json
{
    "type": 2,
    "requestID": "abc",
    "result": 1,
    "data": {
        "tokenInfo": {
            "userID": 1000000083957782, // AcFun帐号或匿名用户的uid
            "securityKey": "1oVtckMbS958PQwD9oYA==", // 密钥
            "serviceToken": "ChRhY2Z1bi5hcGkudmlzaXRvci5zdBJwcn4Q5oc4RhQVng1kCvHAwrY4_Atih1mCLjV4Hf5O7HrdQkFHwjxQZNv0pvtU0cZhhsW1BfCPtYLvVc2DomsyqZuPkTp_AGzij4d5bnpHDlnSWAbqFmR4V09QeY_ACmrtq0VSz_eN1RV9Il7XvvUgKxoSyrmOnZEFeoExMdFq6-X8nnzoIiAMQakYBJwiJRPaToN7BdKVd33_gZ5y7Kfm1wm9PcAK7ig", // token
            "deviceID": "web_7919352416EF8C", // 设备id
            "cookies": [] // AcFun帐号的cookie，匿名用户为空
        }
    }
}
```

### 命令类型
#### 获取弹幕
##### 请求
```json
{
    "type": 100,
    "requestID": "abc",
    "data": {
        "liverUID": 123456
    }
}
```

`liverUID`：主播的uid

客户端可以同时请求获取多个不同主播的弹幕

##### 响应
```json
{
    "type": 100,
    "result": 1,
    "requestID": "abc",
    "data": {
        "StreamInfo": {
            "liveID": "-ZwJdCYApS4", // 直播ID
            "title": "不瘦十斤不改标题", // 直播间标题
            "liveStartTime": 1608433802998, // 直播开始的时间，是以毫秒为单位的Unix time
            "panoramic": false, // 是否全景直播
            "streamList": [ // 直播源列表
                {
                    "url": "https://ali-acfun-adaptive.pull.etoote.com/livecloud/kszt_z9EyK1pmwg4_sd1000.flv?auth_key=1611034859-0-0-1de78274cd73a99d9ce9c3b080431855\u0026oidc=alihb\u0026tsc=origin", // 直播源链接
                    "bitrate": 1000, // 直播源码率，不一定是实际码率
                    "qualityType": "STANDARD", // 直播源类型，一般是"STANDARD"、"HIGH"、"SUPER"、"BLUE_RAY"
                    "qualityName": "高清" // 直播源类型的中文名字，一般是"高清"、"超清"、"蓝光 4M"、"蓝光 5M"、"蓝光 6M"、"蓝光 7M"、"蓝光 8M"
                },
                {
                    "url": "https://ali-acfun-adaptive.pull.etoote.com/livecloud/kszt_z9EyK1pmwg4_hd2000.flv?auth_key=1611034859-0-0-03a64a59accf8ad1d1dcc52d46002211\u0026oidc=alihb\u0026tsc=origin",
                    "bitrate": 2000,
                    "qualityType": "HIGH",
                    "qualityName": "超清"
                },
                {
                    "url": "https://ali-acfun-adaptive.pull.etoote.com/livecloud/kszt_z9EyK1pmwg4_hd4000.flv?auth_key=1611034859-0-0-9bcb3496372141c5d354a454c9f33266\u0026oidc=alihb\u0026tsc=origin",
                    "bitrate": 4000,
                    "qualityType": "SUPER",
                    "qualityName": "蓝光 4M"
                },
                {
                    "url": "https://ali-acfun-adaptive.pull.etoote.com/livecloud/kszt_z9EyK1pmwg4.flv?auth_key=1611034859-0-0-4c5e8a6375ca0bb2dbc85501a09c7510\u0026oidc=alihb",
                    "bitrate": 5000,
                    "qualityType": "BLUE_RAY",
                    "qualityName": "蓝光 5M"
                }
            ],
            "streamName": "kszt_z9EyK1pmwg4" // 直播源名字（ID）
        }
    }
}
```

请求成功后服务端会不断发送[弹幕和信号数据](#弹幕和信号类型)给客户端，直到直播结束或请求[停止获取弹幕](#停止获取弹幕)为止

#### 停止获取弹幕
##### 请求
```json
{
    "type": 101,
    "requestID": "abc",
    "data": {
        "liverUID": 123456
    }
}
```

##### 响应
```json
{
    "type": 101,
    "requestID": "abc",
    "result": 1
}
```

#### 直播间观众列表
##### 请求
```json
{
    "type": 102,
    "requestID": "abc",
    "data": {
        "liveID": "cgbKNA8R5nY"
    }
}
```

`liveID`：直播ID

##### 响应
```json
{
    "type": 102,
    "requestID": "abc",
    "result": 1,
    "data": [
        {
            "userInfo": {
                "userID": 541323, // 用户uid
                "nickname": "天然猪肉丸", // 用户昵称
                "avatar": "https://imgs.aixifan.com/content/2020_11_5/1.604508397967681E9.png", // 用户头像
                "medal": { // 没有守护徽章
                    "uperID": 0,
                    "userID": 0,
                    "clubName": "",
                    "level": 0
                },
                "managerType": 0 // 没有房管类型
            },
            "anonymousUser": false, // 是否匿名用户
            "displaySendAmount": "1000", // 赠送的全部礼物的价值，单位是AC币，注意不一定是纯数字的字符串
            "customData": "{\"userInfo\":{\"verified\":0,\"verifiedTypes\":[5],\"joinUpCollege\":true},\"countInfo\":{\"fansCount\":1398}}" // 用户的一些额外信息，格式为json
        }
    ]
}
```

#### 礼物贡献榜
##### 请求
```json
{
    "type": 103,
    "requestID": "abc",
    "data": {
        "liverUID": 12345
    }
}
```

##### 响应
```json
{
    "type": 103,
    "requestID": "abc",
    "result": 1,
    "data": [
        {
            "userInfo": {
                "userID": 13614296,
                "nickname": "某个帕克",
                "avatar": "https://imgs.aixifan.com/FqArNeselDeOoPTXd-xqga9TDs4z",
                "medal": { // 没有守护徽章
                    "uperID": 0,
                    "userID": 0,
                    "clubName": "",
                    "level": 0
                },
                "managerType": 0 // 没有房管类型
            },
            "anonymousUser": false,
            "displaySendAmount": "1.2万",
            "customData": "{\"userInfo\":{\"verified\":2,\"verifiedTypes\":[2,5],\"joinUpCollege\":true},\"countInfo\":{\"fansCount\":351}}"
        }
    ]
}
```

#### 直播总结信息
##### 请求
```json
{
    "type": 104,
    "requestID": "abc",
    "data": {
        "liveID": "cgbKNA8R5nY"
    }
}
```

##### 响应
```json
{
    "type": 104,
    "requestID": "abc",
    "result": 1,
    "data": {
        "liveDuration": 13506514, // 直播时长，单位为毫秒
        "likeCount": "39380", // 点赞总数
        "watchCount": "13475" // 观看过直播的人数总数
    }
}
```

#### 抢红包结果（未完成）
##### 请求
```json
{
    "type": 105,
    "requestID": "abc",
    "data": {
        "liveID": "cgbKNA8R5nY",
        "redpackID": "PSCfFH10v-w"
    }
}
```

##### 响应
```json
```

#### 直播回放
##### 请求
```json
{
    "type": 106,
    "requestID": "abc",
    "data": {
        "liveID": "cgbKNA8R5nY"
    }
}
```

##### 响应
```json
{
    "type": 106,
    "requestID": "abc",
    "result": 1,
    "data": {
        "duration": 13508162, // 录播视频时长，单位是毫秒
        "url": "http://alivod.a.yximgs.com/livedvr/flv2ts/livecloud/kszt_o0SzIE0GizQ.1608550343521-13497998.0.m3u8?auth_key=1608995849-1372638012-0-0802fab3d4f8ce29986a07fe7cdcee3c", // 录播源链接，目前分为阿里云和腾讯云两种
        "backupURL": "http://txvod.a.yximgs.com/livedvr/flv2ts/livecloud/kszt_o0SzIE0GizQ.1608550343521-13497998.0.m3u8?sign=1608995849-2424432917894601301-0-e4f023847bebae3fc7889686e6d87ebd", // 备份录播源链接
        "m3u8Slice": "#EXTM3U\r\n#EXT-X-PLAYLIST-TYPE:VOD\r\n#EXT-X-VERSION:4\r\n#EXT-X-MEDIA-SEQUENCE:0\r\n#EXT-X-TARGETDURATION:13509\r\n#EXTINF:11.984,\r\nhttp://alivod.a.yximgs.com/livedvr/flv2ts/livecloud/kszt_o0SzIE0GizQ.1608550343521.alihbs1.ts?auth_key=1608995849-1372638012-0-0802fab3d4f8ce29986a07fe7cdcee3c",
        "width": 1920, // 录播视频宽度
        "height": 1080 // 录播视频高度
    }
}
```

#### 全部礼物列表
##### 请求
```json
{
    "type": 107,
    "requestID": "abc",
}
```

##### 响应
```json
{
    "type": 107,
    "requestID": "abc",
    "result": 1,
    "data": [
        {
            "giftID": 1, // 礼物ID
            "giftName": "香蕉", // 礼物名字
            "arLiveName": "", // 不为空时礼物属于虚拟偶像区的特殊礼物
            "payWalletType": 2, // 1为非免费礼物，2为免费礼物
            "price": 1, // 礼物价格，非免费礼物时单位为AC币，免费礼物（香蕉）时为1
            "webpPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200316101317UbXssBoH.webp", // 礼物的webp格式图片（动图）
            "pngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200812141711JRxMyUWH.png", // 礼物的png格式图片（大）
            "smallPngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200316101519KncIIcdd.png", // 礼物的png格式图片（小）
            "allowBatchSendSizeList": [ // 网页或APP单次能够赠送的礼物数量列表
                1,
                5
            ],
            "canCombo": false, // 是否能连击，一般免费礼物（香蕉）不能连击，其余能连击
            "canDraw": false, // 是否能涂鸦？
            "magicFaceID": 0,
            "description": "达成蕉易，投蕉鼓励！", // 礼物的描述
            "redpackPrice": 0 // 礼物红包价格总额，单位为AC币
        },
        {
            "giftID": 2,
            "giftName": "吃瓜",
            "arLiveName": "",
            "payWalletType": 1,
            "price": 1,
            "webpPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200326112056bpqDFUpE.webp",
            "pngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200812141654NvIlrtUX.png",
            "smallPngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200316101616sGxRzHkg.png",
            "allowBatchSendSizeList": [
                1,
                6,
                10,
                66,
                233
            ],
            "canCombo": true,
            "canDraw": true,
            "magicFaceID": 0,
            "description": "哎，我跟你讲，这瓜超甜的！",
            "redpackPrice": 0
        }
    ]
}
```

#### 账户钱包
##### 请求
```json
{
    "type": 108,
    "requestID": "abc",
}
```

##### 响应
```json
{
    "type": 108,
    "requestID": "abc",
    "result": 1,
    "data": {
        "acCoin": 57, // 账户AC币数量
        "banana": 289 // 账户香蕉数量
    }
}
```

#### 登陆用户的房管列表
##### 请求
```json
{
    "type": 200,
    "requestID": "abc",
}
```

##### 响应
```json
{
    "type": 200,
    "requestID": "abc",
    "result": 1,
    "data": [
        {
            "userInfo": {
                "userID": 23682490,
                "nickname": "AC娘本体",
                "avatar": "https://imgs.aixifan.com/FnlcvTfQHideC2bGgfRO2u9gfig_",
                "medal": {
                    "uperID": 0,
                    "userID": 0,
                    "clubName": "",
                    "level": 0
                },
                "managerType": 0
            },
            "customData": "{\"userInfo\":{\"verified\":1,\"verifiedTypes\":[1,5,3],\"joinUpCollege\":true},\"countInfo\":{\"fansCount\":365877}}",
            "online": false // 是否直播间在线？（可能不准确）
        }
    ]
}
```

#### 添加房管
##### 请求
```json
{
    "type": 201,
    "requestID": "abc",
    "data": {
        "managerUID": 12345
    }
}
```

`managerUID`：房管的uid

##### 响应
```json
{
    "type": 201,
    "requestID": "abc",
    "result": 1
}

```

#### 删除房管
##### 请求
```json
{
    "type": 202,
    "requestID": "abc",
    "data": {
        "managerUID": 12345
    }
}
```

##### 响应
```json
{
    "type": 202,
    "requestID": "abc",
    "result": 1
}
```

#### 踢人的历史记录（未完成）
##### 请求
```json
{
    "type": 203,
    "requestID": "abc",
}
```

##### 响应
```json
```

#### 房管踢人
##### 请求
```json
{
    "type": 204,
    "requestID": "abc",
    "data": {
        "kickedUID": 12345
    }
}
```

`kickedUID`：被踢的用户的uid

##### 响应
```json
{
    "type": 204,
    "requestID": "abc",
    "result": 1
}
```

#### 主播踢人
##### 请求
```json
{
    "type": 205,
    "requestID": "abc",
    "data": {
        "kickedUID": 12345
    }
}
```

##### 响应
```json
{
    "type": 205,
    "requestID": "abc",
    "result": 1
}
```

### 弹幕和信号类型
弹幕和信号数据在客户端请求[获取弹幕](#获取弹幕)后由服务端发送给客户端

#### 弹幕类型
##### 弹幕
```json
{
    "liverUID": 4537972,
    "type": 1000,
    "data": {
        "danmuInfo": {
            "sendTime": 1608379795363, // 弹幕发送时间，是以毫秒为单位的Unix时间
            "userInfo": { // 发送弹幕的用户的信息
                "userID": 666609, // 用户uid
                "nickname": "泼墨造一匹快马追回十年前姑娘", // 用户名字
                "avatar": "https://imgs.aixifan.com/content/2020_09_20/1600575703124.JPG", // 用户头像
                "medal": { // 用户正在佩戴的守护徽章
                    "uperID": 4537972, // UP主的uid
                    "userID": 666609, // 用户的uid
                    "clubName": "有猫饼", // 守护徽章名字
                    "level": 8 // 守护徽章等级
                },
                "managerType": 0 // 用户是否房管，0是房管，1不是房管
            }
        },
        "content": "哈哈哈哈.." // 弹幕文字
    }
}
```

##### 点赞
```json
{
    "liverUID": 4537972,
    "type": 1001,
    "data": {
        "sendTime": 1608379805737,
        "userInfo": {
            "userID": 35929956,
            "nickname": "甜不辣椒酱",
            "avatar": "https://imgs.aixifan.com/content/2020_10_29/1.6039781036291416E9.gif",
            "medal": {
                "uperID": 36782183,
                "userID": 35929956,
                "clubName": "个正",
                "level": 8
            },
            "managerType": 0
        }
    }
}
```

##### 进入直播间
```json
{
    "liverUID": 2889712,
    "type": 1002,
    "data": {
        "sendTime": 1608390440400,
        "userInfo": {
            "userID": 37976639,
            "nickname": "言晶",
            "avatar": "https://imgs.aixifan.com/style/image/202007/Uzm4NfSfo8mjNp7r6BJEQVR9nMjhDO2L.jpg",
            "medal": {
                "uperID": 23512715,
                "userID": 37976639,
                "clubName": "吳彦祖",
                "level": 9
            },
            "managerType": 0
        }
    }
}
```

##### 关注主播
```json
{
    "liverUID": 2889712,
    "type": 1003,
    "data": {
        "sendTime": 1608390202644,
        "userInfo": {
            "userID": 609092,
            "nickname": "非天然猪肉丸",
            "avatar": "https://imgs.aixifan.com/style/image/defaultAvatar.jpg",
            "medal": {
                "uperID": 20680343,
                "userID": 609092,
                "clubName": "酱紫鸭",
                "level": 7
            },
            "managerType": 0
        }
    }
}
```

##### 投蕉
```json
{
    "liverUID": 4537972,
    "type": 1004,
    "data": {
        "danmuInfo": {
            "sendTime": 1608379795363,
            "userInfo": {
                "userID": 666609,
                "nickname": "泼墨造一匹快马追回十年前姑娘",
                "avatar": "", // 没有用户头像
                "medal": { // 没有守护徽章信息
                    "uperID": 0,
                    "userID": 0,
                    "clubName": "",
                    "level": 0
                },
                "managerType": 0 // 没有房管类型
            }
        },
        "bananaCount": 5 // 投蕉数量
    }
}
```

现在基本不用这个，而是用[礼物](#礼物)代替

##### 礼物
```json
{
    "liverUID": 4537972,
    "type": 1005,
    "data": {
        "danmuInfo": {
            "sendTime": 1608379838216,
            "userInfo": {
                "userID": 532848,
                "nickname": "D.H.T",
                "avatar": "https://imgs.aixifan.com/content/2019_09_14/1568459103154.JPG",
                "medal": {
                    "uperID": 34743261,
                    "userID": 532848,
                    "clubName": "扶她人",
                    "level": 7
                },
                "managerType": 0
            }
        },
        "giftDetail": { // 礼物详细信息
            "giftID": 12, // 礼物ID
            "giftName": "打Call", // 礼物名字
            "arLiveName": "", // 不为空时礼物属于虚拟偶像区的特殊礼物
            "payWalletType": 1, // 1为非免费礼物，2为免费礼物
            "price": 10, // 礼物价格，非免费礼物时单位为AC币，免费礼物（香蕉）时为1
            "webpPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200326112443kfWQhpaG.webp", // 礼物的webp格式图片（动图）
            "pngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200812141131ukNHkGeU.png", // 礼物的png格式图片（大）
            "smallPngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200316110407BdolKFLb.png", // 礼物的png格式图片（小）
            "allowBatchSendSizeList": [ // 网页或APP单次能够赠送的礼物数量列表
                1,
                6,
                10,
                66,
                233
            ],
            "canCombo": true, // 是否能连击，一般免费礼物（香蕉）不能连击，其余能连击
            "canDraw": true, // 是否能涂鸦？
            "magicFaceID": 0,
            "description": "为你转身为你爆灯，为你打CALL日夜不分", // 礼物的描述
            "redpackPrice": 0 // 礼物红包价格总额，单位为AC币
        },
        "count": 5, // 礼物单次赠送的数量，礼物总数是Count * Combo
        "combo": 1, // 礼物连击数量，礼物总数是Count * Combo
        "value": 50000, // 礼物价值，非免费礼物时单位为AC币*1000，免费礼物（香蕉）时单位为礼物数量
        "comboID": "FD7E07B5-DF69-4257-84BC-7FEA377E9C85", // 礼物连击ID
        "slotDisplayDuration": 3000, // 应该是礼物动画持续的时间，单位为毫秒，送礼物后在该时间内再送一次可以实现礼物连击
        "ExpireDuration": 300000,
        "drawGiftInfo": { // 礼物涂鸦
            "screenWidth": 1440, // 手机屏幕宽度
            "screenHeight": 2560, // 手机屏幕高度
            "drawPoint": [ // 涂鸦里各个礼物的位置
                {
                    "marginLeft": 393, // 到手机屏幕左边的距离
                    "marginTop": 263, // 到手机屏幕顶部的距离
                    "scaleRatio": 1, // 放大倍数？
                    "handup": false
                }
            ]
        }
    }
}
```

##### 富文本
```json
{
    "liverUID": 2889712,
    "type": 1006,
    "data": {
        "sendTime": 0, // 富文本的发送时间，是以毫秒为单位的Unix时间，可能为0
        "segments": [ // 富文本各部分，类型分别是RichTextUserInfo、RichTextPlain或RichTextImage
            {
                "type": 1900, // RichTextUserInfo
                "segment": {
                    "userInfo": {
                        "userID": 2334509,
                        "nickname": "aqilili",
                        "avatar": "https://imgs.aixifan.com/style/image/defaultAvatar.jpg",
                        "medal": {
                            "uperID": 0,
                            "userID": 0,
                            "clubName": "",
                            "level": 0
                        },
                        "managerType": 0
                    },
                    "color": "#409BEF" // 一般是用户昵称的颜色
                }
            },
            {
                "type": 1901, // RichTextPlain
                "segment": {
                    "text": " 领取了 ", // 文字
                    "color": "" // 文字的颜色
                }
            },
            {
                "type": 1900,
                "segment": {
                    "userInfo": {
                        "userID": 702914,
                        "nickname": "昊东",
                        "avatar": "https://imgs.aixifan.com/content/2020_12_16/1.6080943970212917E9.png",
                        "medal": {
                            "uperID": 11461714,
                            "userID": 702914,
                            "clubName": "级变态",
                            "level": 6
                        },
                        "managerType": 0
                    },
                    "color": "#409BEF"
                }
            },
            {
                "type": 1901,
                "segment": {
                    "text": " 发的5AC币",
                    "color": ""
                }
            },
            {
                "type": 1902, // RichTextImage
                "segment": {
                    "pictures": [ // 图片
                        "http://abc.jpg"
                    ],
                    "alternativeText": "abcd", // 可选的文本
                    "alternativeColor": "#409BEF" // 可选的文本的颜色
                }
            }
        ]
    }
}
```

富文本有三种：`RichTextUserInfo`、`RichTextPlain`、`RichTextImage`，对应的`type`分别为1900、1901、1902

##### 加入守护团
```json
{
    "liverUID": 2889712,
    "type": 1007,
    "data": {
        "joinTime": 1608390090583, // 用户加入守护团的时间，是以毫秒为单位的Unix时间
        "fansInfo": { // 用户的信息
            "userID": 1428790,
            "nickname": "柳昭郎",
            "avatar": "", // 没有用户头像
            "medal": { // 没有守护徽章信息
                "uperID": 0,
                "userID": 0,
                "clubName": "",
                "level": 0
            },
            "managerType": 0 // 没有房管类型
        },
        "uperInfo": { // 主播的信息
            "userID": 2889712,
            "nickname": "张梓义",
            "avatar": "", // 没有用户头像
            "medal": { // 没有守护徽章信息
                "uperID": 0,
                "userID": 0,
                "clubName": "",
                "level": 0
            },
            "managerType": 0 // 没有房管类型
        }
    }
}
```

#### 状态信号类型
##### 获取弹幕结束
```json
{
    "liverUID": 12345,
    "type": 2000,
    "data": {
        "liveClosed": true,
        "reason": ""
    }
}
```

`liveClosed`：直播是否结束或客户端请求[停止获取弹幕](#停止获取弹幕)

`reason`：直播正常结束或客户端请求[停止获取弹幕](#停止获取弹幕)时为空，否则为停止获取弹幕的原因

##### 直播间收到香蕉总数
```json
{
    "liverUID": 4425861,
    "type": 2001,
    "data": {
        "bananaCount": "1638"
    }
}
```

`bananaCount`：直播间收到香蕉总数

##### 在线观众和点赞数量
```json
{
    "liverUID": 4425861,
    "type": 2002,
    "data": {
        "watchingCount": "277",
        "likeCount": "2.5万",
        "likeDelta": 2
    }
}
```

`watchingCount`：直播间在线观众数量

`likeCount`：直播间点赞总数

`likeDelta`：点赞增加数量

##### 在线观众前三名
```json
{
    "liverUID": 4425861,
    "type": 2003,
    "data": [ // 最多三位观众
        {
            "userInfo": {
                "userID": 496725,
                "nickname": "病娇御姐看起来老霸道了",
                "avatar": "https://imgs.aixifan.com/style/image/201907/P044fP0S6xaP83vSsZ1RsoUmQ4Uss0Ze.jpg",
                "medal": { // 没有守护徽章
                    "uperID": 0,
                    "userID": 0,
                    "clubName": "",
                    "level": 0
                },
                "managerType": 0 // 没有房管类型
            },
            "anonymousUser": false, // 是否匿名用户
            "displaySendAmount": "486", // 赠送的全部礼物的价值，单位是AC币，注意不一定是纯数字的字符串
            "customData": "{\"userInfo\":{\"verified\":0,\"verifiedTypes\":[],\"joinUpCollege\":false},\"countInfo\":{\"fansCount\":2}}" // 用户的一些额外信息，格式为json
        }
    ]
}
```

##### 进直播间时显示的最近弹幕
```json
{
    "liverUID": 4425861,
    "type": 2004,
    "data": [ // 最多十条弹幕
        {
            "danmuInfo": {
                "sendTime": 1608456531137,
                "userInfo": {
                    "userID": 496725,
                    "nickname": "病娇御姐看起来老霸道了",
                    "avatar": "https://imgs.aixifan.com/style/image/201907/P044fP0S6xaP83vSsZ1RsoUmQ4Uss0Ze.jpg",
                    "medal": {
                        "uperID": 4425861,
                        "userID": 496725,
                        "clubName": "鸽屋咕",
                        "level": 11
                    },
                    "managerType": 0
                }
            },
            "content": "只有靠直击2，逆2了"
        }
    ]
}
```

##### 红包列表
```json
{
    "liverUID": 13945614,
    "type": 2005,
    "data": [ // 没有红包时为空
        {
            "userInfo": {
                "userID": 13240469,
                "nickname": "汤汤湯湯汤",
                "avatar": "https://imgs.aixifan.com/content/2020_09_23/1600853696165.JPG",
                "medal": { // 没有守护徽章
                    "uperID": 0,
                    "userID": 0,
                    "clubName": "",
                    "level": 0
                },
                "managerType": 0
            },
            "displayStatus": 0, // 红包状态，0是红包出现，1是可以获取红包token，2是可以抢红包
            "grabBeginTime": 1608464088394, // 开始抢红包的时间，是以毫秒为单位的Unix时间
            "getTokenLatestTime": 1608464086394, // 抢红包的用户获得token的最晚时间，是以毫秒为单位的Unix时间
            "redpackID": "c5N6p7IMyjA", // 红包ID
            "redpackBizUnit": "ztLiveAcfunRedpackGift", // 一般是"ztLiveAcfunRedpackGift"
            "redpackAmount": 99, // 红包的总价值，单位是AC币
            "settleBeginTime": 1608464148394 // 抢红包的结束时间，是以毫秒为单位的Unix时间
        }
    ]
}
```

##### Chat Call
```json
{
    "liverUID": 4425861,
    "type": 2100,
    "data": {
        "chatID": "abcd",
        "liveID": "abcde",
        "callTime": 1608464088394
    }
}
```

##### Chat Accept
```json
{
    "liverUID": 4425861,
    "type": 2101,
    "data": {
        "chatID": "abcd",
        "mediaType": 1,
        "arraySignalInfo": "abcde"
    }
}
```

##### Chat Ready
```json
{
    "liverUID": 4425861,
    "type": 2102,
    "data": {
        "chatID": "abcd",
        "guest": {
            "userID": 496725,
            "nickname": "病娇御姐看起来老霸道了",
            "avatar": "https://imgs.aixifan.com/style/image/201907/P044fP0S6xaP83vSsZ1RsoUmQ4Uss0Ze.jpg",
            "medal": {
                "uperID": 0,
                "userID": 0,
                "clubName": "",
                "level": 0
            },
            "managerType": 0
        },
        "mediaType": 1
    }
}
```

##### Chat End
```json
{
    "liverUID": 4425861,
    "type": 2103,
    "data": {
        "chatID": "abcd",
        "endType": 1
    }
}
```

#### 通知信号类型
##### 被踢出直播间
```json
{
    "liverUID": 4425861,
    "type": 3000,
    "data": {
        "kickedOutReason": "abc"
    }
}
```

`kickedOutReason`：被踢理由

##### ViolationAlert
```json
{
    "liverUID": 4425861,
    "type": 3001,
    "data": {
        "violationContent": "abc"
    }
}
```

##### 登陆帐号的房管状态
```json
{
    "liverUID": 4425861,
    "type": 3002,
    "data": {
        "managerState": 0
    }
}
```

`managerState`：0不是房管，1被添加房管，2被移除房管，3是房管

### 错误码
`10`：请求的json解析错误

`11`：请求`type`无效

`12`：请求`data`无效

`13`：处理请求时出现错误

`14`：需要登陆

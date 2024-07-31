- [流程](#流程)
  - [OBS 直播流程](#obs-直播流程)
- [WebSocket 数据格式](#websocket-数据格式)
  - [请求](#请求)
  - [响应](#响应)
    - [请求成功时的响应](#请求成功时的响应)
    - [出现错误时的响应](#出现错误时的响应)
  - [弹幕和信号](#弹幕和信号)
- [心跳包](#心跳包)
- [登陆](#登陆)
  - [请求](#请求-1)
  - [响应](#响应-1)
- [设置客户端 ID](#设置客户端-id)
  - [请求](#请求-2)
  - [响应](#响应-2)
- [请求转发数据](#请求转发数据)
  - [请求](#请求-3)
  - [响应](#响应-3)
  - [客户端接收的转发数据](#客户端接收的转发数据)
- [设置 AcFun 帐号 token](#设置-acfun-帐号-token)
  - [请求](#请求-4)
  - [响应](#响应-4)
- [扫描二维码登陆](#扫描二维码登陆)
  - [请求](#请求-5)
  - [响应](#响应-5)
    - [返回登陆二维码](#返回登陆二维码)
    - [用户扫描了登陆二维码的响应](#用户扫描了登陆二维码的响应)
    - [登陆二维码过期或者用户取消登陆的响应](#登陆二维码过期或者用户取消登陆的响应)
    - [用户成功登陆的响应](#用户成功登陆的响应)
- [命令类型](#命令类型)
  - [获取弹幕](#获取弹幕)
    - [请求](#请求-6)
    - [响应](#响应-6)
  - [停止获取弹幕](#停止获取弹幕)
    - [请求](#请求-7)
    - [响应](#响应-7)
  - [直播间观众列表](#直播间观众列表)
    - [请求](#请求-8)
    - [响应](#响应-8)
  - [礼物贡献榜](#礼物贡献榜)
    - [请求](#请求-9)
    - [响应](#响应-9)
  - [直播总结信息](#直播总结信息)
    - [请求](#请求-10)
    - [响应](#响应-10)
  - [抢红包结果](#抢红包结果)
    - [请求](#请求-11)
    - [响应](#响应-11)
  - [直播回放](#直播回放)
    - [请求](#请求-12)
    - [响应](#响应-12)
  - [全部礼物列表](#全部礼物列表)
    - [请求](#请求-13)
    - [响应](#响应-13)
  - [账户钱包](#账户钱包)
    - [请求](#请求-14)
    - [响应](#响应-14)
  - [指定用户的直播信息](#指定用户的直播信息)
    - [请求](#请求-15)
    - [响应](#响应-15)
  - [直播间列表](#直播间列表)
    - [请求](#请求-16)
    - [响应](#响应-16)
  - [上传图片](#上传图片)
    - [请求](#请求-17)
    - [响应](#响应-17)
  - [直播统计数据](#直播统计数据)
    - [请求](#请求-18)
    - [响应](#响应-18)
  - [直播预告列表](#直播预告列表)
    - [请求](#请求-19)
    - [响应](#响应-19)
  - [直播间礼物列表](#直播间礼物列表)
    - [请求](#请求-20)
    - [响应](#响应-20)
  - [指定用户的信息](#指定用户的信息)
    - [请求](#请求-21)
    - [响应](#响应-21)
  - [直播剪辑信息](#直播剪辑信息)
    - [请求](#请求-22)
    - [响应](#响应-22)
  - [登陆用户的房管列表](#登陆用户的房管列表)
    - [请求](#请求-23)
    - [响应](#响应-23)
  - [添加房管](#添加房管)
    - [请求](#请求-24)
    - [响应](#响应-24)
  - [删除房管](#删除房管)
    - [请求](#请求-25)
    - [响应](#响应-25)
  - [主播踢人记录](#主播踢人记录)
    - [请求](#请求-26)
    - [响应](#响应-26)
  - [房管踢人](#房管踢人)
    - [请求](#请求-27)
    - [响应](#响应-27)
  - [主播踢人](#主播踢人)
    - [请求](#请求-28)
    - [响应](#响应-28)
  - [登陆用户拥有的指定主播守护徽章详细信息](#登陆用户拥有的指定主播守护徽章详细信息)
    - [请求](#请求-29)
    - [响应](#响应-29)
  - [登陆用户拥有的守护徽章列表](#登陆用户拥有的守护徽章列表)
    - [请求](#请求-30)
    - [响应](#响应-30)
  - [主播守护榜](#主播守护榜)
    - [请求](#请求-31)
    - [响应](#响应-31)
  - [指定用户正在佩戴的守护徽章信息](#指定用户正在佩戴的守护徽章信息)
    - [请求](#请求-32)
    - [响应](#响应-32)
  - [佩戴守护徽章](#佩戴守护徽章)
    - [请求](#请求-33)
    - [响应](#响应-33)
  - [取消佩戴守护徽章](#取消佩戴守护徽章)
    - [请求](#请求-34)
    - [响应](#响应-34)
  - [检测开播权限](#检测开播权限)
    - [请求](#请求-35)
    - [响应](#响应-35)
  - [直播分类列表](#直播分类列表)
    - [请求](#请求-36)
    - [响应](#响应-36)
  - [推流设置](#推流设置)
    - [请求](#请求-37)
    - [响应](#响应-37)
  - [直播状态](#直播状态)
    - [请求](#请求-38)
    - [响应](#响应-38)
  - [转码信息](#转码信息)
    - [请求](#请求-39)
    - [响应](#响应-39)
  - [开始直播](#开始直播)
    - [请求](#请求-40)
    - [响应](#响应-40)
  - [停止直播](#停止直播)
    - [请求](#请求-41)
    - [响应](#响应-41)
  - [更改直播间标题和封面](#更改直播间标题和封面)
    - [请求](#请求-42)
    - [响应](#响应-42)
  - [查询是否允许观众剪辑直播录像](#查询是否允许观众剪辑直播录像)
    - [请求](#请求-43)
    - [响应](#响应-43)
  - [设置是否允许观众剪辑直播录像](#设置是否允许观众剪辑直播录像)
    - [请求](#请求-44)
    - [响应](#响应-44)
- [弹幕和信号类型](#弹幕和信号类型)
  - [弹幕类型](#弹幕类型)
    - [弹幕](#弹幕)
    - [点赞](#点赞)
    - [进入直播间](#进入直播间)
    - [关注主播](#关注主播)
    - [投蕉](#投蕉)
    - [礼物](#礼物)
    - [富文本](#富文本)
    - [加入守护团](#加入守护团)
    - [分享直播间](#分享直播间)
  - [状态信号类型](#状态信号类型)
    - [获取弹幕结束](#获取弹幕结束)
    - [获取弹幕出错](#获取弹幕出错)
    - [直播间收到香蕉总数](#直播间收到香蕉总数)
    - [在线观众和点赞数量](#在线观众和点赞数量)
    - [在线观众前三名](#在线观众前三名)
    - [进直播间时显示的最近弹幕](#进直播间时显示的最近弹幕)
    - [红包列表](#红包列表)
    - [Chat Call](#chat-call)
    - [Chat Accept](#chat-accept)
    - [Chat Ready](#chat-ready)
    - [Chat End](#chat-end)
  - [通知信号类型](#通知信号类型)
    - [被踢出直播间](#被踢出直播间)
    - [直播警告](#直播警告)
    - [登陆帐号的房管状态](#登陆帐号的房管状态)
- [错误码](#错误码)

### 流程

1. 客户端连接服务端，默认端口为 15368
2. 客户端每 5 秒发送一次[心跳包](#心跳包)给服务端，接收服务端的心跳包
3. 客户端请求[登陆](#登陆)或[扫描二维码登陆](#扫描二维码登陆)AcFun 帐号或者[设置 AcFun 帐号 token](#设置acfun帐号token)，一个连接只能同时登陆一个帐号
4. 客户端发送[命令](#命令类型)请求，接受服务端的响应

客户端请求[获取直播间弹幕](#获取弹幕)后，服务端会不断发送[弹幕和信号数据](#弹幕和信号类型)给客户端。可同时请求多个直播间的弹幕。

客户端可以请求转发数据到其他客户端，具体看[请求转发数据](#请求转发数据)。

服务端处理请求的网络超时通常是 10 秒左右。

#### OBS 直播流程

1. 客户端获取[推流设置](#推流设置)，根据返回用 OBS 设置相应的推流服务器和串流密钥并开始推流
2. 客户端每 5 秒左右获取一次[转码信息](#转码信息)（开播后可停止获取），返回不为空时请求[开始直播](#开始直播)，直播分类可以从[直播分类列表](#直播分类列表)获得，直播封面可以先[上传图片](#上传图片)到 AcFun 服务器获取图片链接
3. 直播途中可以[更改直播间标题和封面](#更改直播间标题和封面)，当要停止直播时客户端请求[停止直播](#停止直播)并停止用 OBS 推流

### WebSocket 数据格式

#### 请求

```json
{
  "type": 110,
  "requestID": "abc",
  "data": {}
}
```

`type`：请求类型

`requestID`：请求 ID，用于区分同一类型的请求，具体内容由客户端决定

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

`requestID`：请求 ID，用于区分同一类型的请求，内容和对应的请求一样

`result`：请求结果，请求成功时为 1

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

`type`：响应类型，和请求类型一样，不需要时可为 0

`requestID`：请求 ID，用于区分同一类型的请求，内容和对应的请求一样

`result`：请求结果，出现错误时（不为 1）为[错误码](#错误码)

`error`：错误信息

#### 弹幕和信号

```json
{
  "liverUID": 12345,
  "type": 1100,
  "data": {}
}
```

`liverUID`：主播 uid

`type`：弹幕和信号类型

`data`：弹幕和信号的数据

[弹幕和信号数据](#弹幕和信号类型)由服务端发送到客户端，在客户端请求[获取弹幕](#获取弹幕)后才会发送

### 心跳包

```json
{
  "type": 1
}
```

`type`：心跳包类型为 1

客户端每 5 秒发送一次心跳包给服务端，服务端收到心跳包后会发送心跳包给客户端

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

`account`：AcFun 帐号邮箱或手机号

`password`：AcFun 帐号密码

`account`或`password`为空时以匿名游客的身份登陆

#### 响应

```json
{
  "type": 2,
  "requestID": "abc",
  "result": 1,
  "data": {
    "tokenInfo": {
      "userID": 1000000083957782, // AcFun 帐号或匿名用户的 uid
      "securityKey": "1oVtckMbS958PQwD9oYA==", // 密钥
      "serviceToken": "ChRhY2Z1bi5hcGkudmlzaXRvci5zdBJwcn4Q5oc4RhQVng1kCvHAwrY4_Atih1mCLjV4Hf5O7HrdQkFHwjxQZNv0pvtU0cZhhsW1BfCPtYLvVc2DomsyqZuPkTp_AGzij4d5bnpHDlnSWAbqFmR4V09QeY_ACmrtq0VSz_eN1RV9Il7XvvUgKxoSyrmOnZEFeoExMdFq6-X8nnzoIiAMQakYBJwiJRPaToN7BdKVd33_gZ5y7Kfm1wm9PcAK7ig", // token
      "deviceID": "web_7919352416EF8C", // 设备 id
      "cookies": [] // AcFun 帐号的 cookie，匿名用户为空
    }
  }
}
```

### 设置客户端 ID

#### 请求

```json
{
  "type": 3,
  "requestID": "abc",
  "data": {
    "clientID": "defghi"
  }
}
```

`clientID`：用于识别不同的客户端，多个客户端可拥有同一个 ID，主要用于接收来自其他客户端的数据，没有设置 ID 或 ID 为空时该客户端只能接受其他客户端的广播，无法接收定向转发的数据

#### 响应

```json
{
  "type": 3,
  "requestID": "abc",
  "result": 1
}
```

### 请求转发数据

#### 请求

```json
{
  "type": 4,
  "requestID": "abc",
  "data": {
    "clientID": "defghi",
    "message": "jklmnop"
  }
}
```

`clientID`：转发的目标客户端的 ID，如果为空则向所有客户端广播

`message`：要转发的数据

#### 响应

```json
{
  "type": 4,
  "requestID": "abc",
  "result": 1
}
```

#### 客户端接收的转发数据

```json
{
  "type": 5,
  "requestID": "abc",
  "result": 1,
  "data": {
    "clientID": "defghi",
    "message": "jklmnop"
  }
}
```

`requestID`: 源请求的`requestID`

`clientID`：源客户端的 ID

`message`：转发的数据

### 设置 AcFun 帐号 token

#### 请求

```json
{
  "type": 6,
  "requestID": "abc",
  "data": {
    "userID": 1000000083957782,
    "securityKey": "1oVtckMbS958PQwD9oYA==",
    "serviceToken": "ChRhY2Z1bi5hcGkudmlzaXRvci5zdBJwcn4Q5oc4RhQVng1kCvHAwrY4_Atih1mCLjV4Hf5O7HrdQkFHwjxQZNv0pvtU0cZhhsW1BfCPtYLvVc2DomsyqZuPkTp_AGzij4d5bnpHDlnSWAbqFmR4V09QeY_ACmrtq0VSz_eN1RV9Il7XvvUgKxoSyrmOnZEFeoExMdFq6-X8nnzoIiAMQakYBJwiJRPaToN7BdKVd33_gZ5y7Kfm1wm9PcAK7ig",
    "deviceID": "web_7919352416EF8C",
    "cookies": []
  }
}
```

#### 响应

```json
{
  "type": 6,
  "requestID": "abc",
  "result": 1
}
```

### 扫描二维码登陆

二维码登陆要先发送请求，将[返回的二维码数据](#返回登陆二维码)转化为图片，用户扫描二维码时会返回[响应](#用户扫描了登陆二维码的响应)（但用户还没选择登陆）

如果登陆二维码过期或者用户取消登陆，会返回[响应](#登陆二维码过期或者用户取消登陆的响应)

用户成功登陆会返回[响应](#用户成功登陆的响应)

#### 请求

```json
{
  "type": 7,
  "requestID": "abc"
}
```

#### 响应

##### 返回登陆二维码

```json
{
  "type": 7,
  "requestID": "abc",
  "result": 1,
  "data": {
    "expireTime": 1722450851957,
    "imageData": "iVBORw0KGgoAAAANAAACmt7ZYLtr8sHsDeSB+qog4/91k18DdmrZXftOjLvJDBTvHfCt9neYMxahsjx2u7t+/7eAO7UNMuQpU1lzAYRM+vY90K1XR+UMg91u4JXE5L88AgjhgEi5d0t8wZg6IeXgs1z7FrOTA4c1mFzF7W3i6DQVi4XZa5P3rXPGPwzDerNYfZs/FsGNidleMQQgVBTMbAPcK4lg4qw1wMBo+qaZEut8BjHBi4VuR5C3av8p4yBvYwDB4ZBFa7CqsdNlQ6N2ulLICgTmtwcAldo5lDpJhUMieQhwslX5LKCTmeffMDa+Z3zCIkAHBoJTu3hIKAdEOPah2gPtcb7a9AMK3TfXPPcyPrvBIFu9LlqvxvLzte0HxETms8wdBoreNww8WfQxJwJxeF7F/gA7F4hj6xAHllGI1iRK1x3iMG6DQZkzxEGFOak0DQya5blrNBWGyVIUpoeAJtPBZVPNwgoFBhTsWC/saF/MrDh/wHLDjsBCHnVAUYOwlobevcZG9wcBOMLhgwbhgLmOQOxSTZyuhTGnDIMBQO5RdVWx2yf5OtFKkAAAAASUVORK5CYII="
  }
}
```

`expireTime`：登陆二维码失效的时间，是以毫秒为单位的 Unix time

`imageData`：Base64 编码的 PNG 二维码图片

##### 用户扫描了登陆二维码的响应

```json
{
  "type": 8,
  "requestID": "abc",
  "result": 1
}
```

##### 登陆二维码过期或者用户取消登陆的响应

```json
{
  "type": 9,
  "requestID": "abc",
  "result": 1
}
```

##### 用户成功登陆的响应

```json
{
  "type": 10,
  "requestID": "abc",
  "result": 1,
  "data": {
    "tokenInfo": {
      "userID": 103411,
      "securityKey": "+8IpTi/Xul+pkahA==",
      "serviceToken": "ChZhY2Z1bi5taWRn91bmQuYRt7ANgLQQHbmmPpeIz5yhcRoTFTELJZQeB4mnXal0X4EJE6jarSYlP1rOMPXKsHd4nQ5kg8z24ZqBgejExWj5DCJ-vDYaEtD6IgjFT54_die_3AyzbWcC8wIewQNQBxed-dFT6cEBTAB",
      "deviceID": "web_83662F48A32",
      "cookies": []
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

`liverUID`：主播的 uid

客户端可以同时请求获取多个不同主播的弹幕，重复请求同一个主播的弹幕不会重复发送弹幕

##### 响应

```json
{
  "type": 100,
  "requestID": "abc",
  "result": 1,
  "data": {
    "liverUID": 123456,
    "StreamInfo": {
      "liveID": "-ZwJdCYApS4", // 直播 ID
      "title": "不瘦十斤不改标题", // 直播间标题
      "liveStartTime": 1608433802998, // 直播开始的时间，是以毫秒为单位的 Unix time
      "panoramic": false, // 是否全景直播
      "streamList": [
        // 直播源列表
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

重复请求同一个主播的弹幕，响应的`data`字段没有`StreamInfo`字段

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
  "result": 1,
  "data": {
    "liverUID": 123456
  }
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

`liveID`：直播 ID

##### 响应

```json
{
  "type": 102,
  "requestID": "abc",
  "result": 1,
  "data": [
    {
      "userInfo": {
        "userID": 541323, // 用户 uid
        "nickname": "天然猪肉丸", // 用户昵称
        "avatar": "https://imgs.aixifan.com/content/2020_11_5/1.604508397967681E9.png", // 用户头像
        "medal": {
          // 没有守护徽章
          "uperID": 0,
          "userID": 0,
          "clubName": "",
          "level": 0
        },
        "managerType": 0 // 没有房管类型
      },
      "anonymousUser": false, // 是否匿名用户
      "displaySendAmount": "1000", // 赠送的全部礼物的价值，单位是 AC 币，注意不一定是纯数字的字符串
      "customData": "{\"userInfo\":{\"verified\":0,\"verifiedTypes\":[5],\"joinUpCollege\":true},\"countInfo\":{\"fansCount\":1398}}" // 用户的一些额外信息，格式为 json
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
        "medal": {
          // 没有守护徽章
          "uperID": 0,
          "userID": 0,
          "clubName": "",
          "level": 0
        },
        "managerType": 0 // 没有房管类型
      },
      "anonymousUser": false,
      "displaySendAmount": "1.2 万",
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
    "duration": 7565966, // 直播时长，单位为毫秒
    "likeCount": "12996", // 点赞总数
    "watchCount": "926", // 观看过直播的人数总数
    "giftCount": 10, // 直播收到的付费礼物数量，登陆主播帐号才能查询到
    "diamondCount": 50000, // 直播收到的付费礼物对应的钻石数量，100 钻石=1AC 币，登陆主播帐号才能查询到
    "bananaCount": 100 // 直播收到的香蕉数量，登陆主播帐号才能查询到
  }
}
```

#### 抢红包结果

##### 请求

```json
{
  "type": 105,
  "requestID": "abc",
  "data": {
    "liveID": "7McE2WZl9Xc",
    "redpackID": "b-D8XOlAlxI",
    "redpackBizUnit": "ztLiveAcfunRedpackGift"
  }
}
```

##### 响应

```json
{
  "type": 105,
  "requestID": "abc",
  "result": 1,
  "data": [
    {
      "userInfo": {
        "userID": 41073755,
        "nickname": "弹幕姬～",
        "avatar": "https://imgs.aixifan.com/content/2020_11_09/1604880424274.JPG",
        "medal": {
          // 没有守护徽章信息
          "uperID": 0,
          "userID": 0,
          "clubName": "",
          "level": 0
        },
        "managerType": 0 // 没有房管信息
      },
      "grabAmount": 7 // 抢到的 AC 币数量
    }
  ]
}
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
  "requestID": "abc"
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
      "giftID": 1, // 礼物 ID
      "giftName": "香蕉", // 礼物名字
      "arLiveName": "", // 不为空时礼物属于虚拟偶像区的特殊礼物
      "payWalletType": 2, // 1 为非免费礼物，2 为免费礼物
      "price": 1, // 礼物价格，非免费礼物时单位为 AC 币，免费礼物（香蕉）时为 1
      "webpPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200316101317UbXssBoH.webp", // 礼物的 webp 格式图片（动图）
      "pngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200812141711JRxMyUWH.png", // 礼物的 png 格式图片（大）
      "smallPngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200316101519KncIIcdd.png", // 礼物的 png 格式图片（小）
      "allowBatchSendSizeList": [
        // 网页或 APP 单次能够赠送的礼物数量列表，有可能为`null`
        1, 5
      ],
      "canCombo": false, // 是否能连击，一般免费礼物（香蕉）不能连击，其余能连击
      "canDraw": false, // 是否能涂鸦？
      "magicFaceID": 0,
      "vupArID": 0,
      "description": "达成蕉易，投蕉鼓励！", // 礼物的描述
      "redpackPrice": 0, // 礼物红包价格总额，单位为 AC 币
      "cornerMarkerText": ""
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
      "allowBatchSendSizeList": [1, 6, 10, 66, 233],
      "canCombo": true,
      "canDraw": true,
      "magicFaceID": 0,
      "vupArID": 0,
      "description": "哎，我跟你讲，这瓜超甜的！",
      "redpackPrice": 0,
      "cornerMarkerText": ""
    }
  ]
}
```

#### 账户钱包

##### 请求

```json
{
  "type": 108,
  "requestID": "abc"
}
```

##### 响应

```json
{
  "type": 108,
  "requestID": "abc",
  "result": 1,
  "data": {
    "acCoin": 57, // 账户 AC 币数量
    "banana": 289 // 账户香蕉数量
  }
}
```

#### 指定用户的直播信息

##### 请求

```json
{
  "type": 109,
  "requestID": "abc",
  "data": {
    "userID": 26675034
  }
}
```

该接口可能有超时等网络问题

##### 响应

```json
{
  "type": 109,
  "requestID": "abc",
  "result": 1,
  "data": {
    "profile": {
      // 用户信息
      "userID": 26675034, // 用户 uid
      "nickname": "艾栗 AIri", // 用户昵称
      "avatar": "https://tx-free-imgs.acfun.cn/content/2020_11_22/1606036415911.JPG?imageslim", // 用户头像
      "avatarFrame": "https://imgs.aixifan.com/WxlISL5vzX-6vMBv2-R3ARFn-q2iMZr-FzayAv.gif", // 用户头像挂件
      "followingCount": 109, // 用户关注数量
      "fansCount": 7090, // 用户粉丝数量
      "contributeCount": 45, // 用户投稿总数
      "signature": "盐系 Vup/红毛狐栗/语音助手/传奇丝袜朋克\n纸板屋：790088315｜微博@艾栗 AIri\nLevel up up！", // 用户签名
      "verifiedText": "AVI 联盟成员 AcFun 签约虚拟偶像 ", // 用户验证信息
      "isJoinUpCollege": true, // 用户是否加入阿普学院
      "isFollowing": true, // 登陆用户是否关注了该用户
      "isFollowed": false // 该用户是否关注了登陆帐号
    },
    "liveType": {
      // 直播类型
      "categoryID": 4, // 直播主分类 ID
      "categoryName": "虚拟偶像", // 直播主分类名字
      "subCategoryID": 403, // 直播次分类 ID
      "subCategoryName": "歌回" // 直播次分类名字
    },
    "liveID": "R3bNghsjBTI", // 直播 ID
    "streamName": "kszt_2pmWEVxgcRY", // 直播源名字
    "title": "第一届斯托瑞 DD 歌回（8 点开始）", // 直播间标题
    "liveStartTime": 1611661443616, // 直播开始的时间，是以毫秒为单位的 Unix 时间
    "portrait": false, // 是否手机直播
    "panoramic": false, // 是否全景直播
    "liveCover": "https://static.yximgs.com/bs2/ztlc/cover_R3bNghsjBTI_raw.jpg", // 直播间封面
    "onlineCount": 824, // 直播间在线人数
    "likeCount": 37669, // 直播间点赞总数
    "hasFansClub": true, // 主播是否有守护团
    "disableDanmakuShow": false, // 是否禁止显示弹幕？
    "paidShowUserBuyStatus": false // 登陆用户是否购买了付费直播？
  }
}
```

#### 直播间列表

##### 请求

```json
{
  "type": 110,
  "requestID": "abc"
}
```

##### 响应

```json
{
  "type": 110,
  "requestID": "abc",
  "result": 1,
  "data": [
    {
      "profile": {
        // 用户信息
        "userID": 378269,
        "nickname": "qyqx",
        "avatar": "https://tx-free-imgs.acfun.cn/style/image/201907/0ldW0vL9ifWM29JzsAyMlEQxdf1vRgIL.jpg?imageslim",
        "avatarFrame": "",
        "followingCount": 5,
        "fansCount": 93573,
        "contributeCount": 72,
        "signature": "不是所有人都是人",
        "verifiedText": "AcFun 游戏区官方认证 UP 主",
        "isJoinUpCollege": true,
        "isFollowing": true,
        "isFollowed": false
      },
      "liveType": {
        // 直播类型
        "categoryID": 1,
        "categoryName": "游戏直播",
        "subCategoryID": 122,
        "subCategoryName": "主机游戏"
      },
      "liveID": "_sLc9sJZIrk",
      "streamName": "kszt_wS-uzFjFxjc",
      "title": "铜之间，只玩狠牌！",
      "liveStartTime": 1611659503558,
      "portrait": false,
      "panoramic": false,
      "liveCover": "https://static.yximgs.com/bs2/ztlc/cover__sLc9sJZIrk_raw.jpg",
      "onlineCount": 1712,
      "likeCount": 46704,
      "hasFansClub": true,
      "disableDanmakuShow": false,
      "paidShowUserBuyStatus": false
    }
  ]
}
```

#### 上传图片

##### 请求

```json
{
  "type": 111,
  "requestID": "abc",
  "data": {
    "imageFile": "cdefg.jpg"
  }
}
```

`imageFile`：图片（可以是 gif）的本地路径

该接口目前不可用

##### 响应

```json
{
  "type": 111,
  "requestID": "abc",
  "result": 1,
  "data": {
    "imageURL": "https://imgs.aixifan.com/065113e-6e32-497d-ba6d-b8ca17ad77.jpg"
  }
}
```

`imageURL`：上传图片的链接

#### 直播统计数据

##### 请求

```json
{
  "type": 112,
  "requestID": "abc",
  "data": {
    "days": 20
  }
}
```

获取最近`days`日的全部直播统计数据

##### 响应

```json
{
  "type": 112,
  "requestID": "abc",
  "result": 1,
  "data": {
    "beginDate": "20210121", // 统计开始的日期
    "endDate": "20210209", // 统计结束的日期
    "overview": {
      // 全部直播的统计概况
      "duration": 17517892, // 直播时长，单位为毫秒
      "maxPopularityValue": 8,
      "watchCount": 175, // 观看过直播的人数总数
      "diamondCount": 0, // 直播收到的付费礼物对应的钻石数量，100 钻石=1AC 币
      "commentCount": 13, // 直播弹幕数量
      "bananaCount": 383 // 直播收到的香蕉数量
    },
    "liveDetail": {
      // 单场直播统计数据
      "20210128": [
        // 直播日期
        {
          "liveStartTime": 1611845023882, // 直播开始的时间，是以毫秒为单位的 Unix 时间
          "liveEndTime": 1611845099163, // 直播结束的时间，是以毫秒为单位的 Unix 时间
          "liveStat": {
            // 直播统计数据
            "duration": 75281,
            "maxPopularityValue": 5,
            "watchCount": 6,
            "diamondCount": 0,
            "commentCount": 0,
            "bananaCount": 0
          }
        }
      ]
    },
    "dailyData": [
      // 单日直播统计数据，有可能为`null`
      {
        "date": "20210128", // 直播日期
        "liveTimes": 1, // 当日直播次数
        "liveStat": {
          // 直播统计数据
          "duration": 75281,
          "maxPopularityValue": 5,
          "watchCount": 6,
          "diamondCount": 0,
          "commentCount": 0,
          "bananaCount": 0
        }
      }
    ]
  }
}
```

#### 直播预告列表

##### 请求

```json
{
  "type": 113,
  "requestID": "abc"
}
```

该接口目前不可用

##### 响应

```json
{
  "type": 113,
  "requestID": "abc",
  "result": 1,
  "data": [
    {
      "activityID": 19927, // 活动 ID
      "profile": {
        // 主播的用户信息
        "userID": 3568347,
        "nickname": "暗莉斯",
        "avatar": "https://tx-free-imgs.acfun.cn/FpUZL492VnnNC8T2XPAplXLT9eyc?imageslim",
        "avatarFrame": "https://imgs.aixifan.com/cms/2019_06_13/1560421184892.gif",
        "followingCount": 168,
        "fansCount": 10852,
        "contributeCount": 114,
        "signature": "个人势 VUP，喜欢玩游戏但很菜，其实是个染了黑毛的金渐层♪录播师傅：莉斯的小年糕 猫猫村：1072715855",
        "verifiedText": "AVI 联盟成员，AcFun 签约虚拟偶像",
        "isJoinUpCollege": true,
        "isFollowing": true,
        "isFollowed": false
      },
      "title": "暗莉斯春节新衣上线", // 预告标题
      "cover": "https://static.yximgs.com/bs2/adminBlock/treasure-1612767546220-zrEdJpue.JPG", // 预告封面
      "liveStartTime": 1613019600000, // 直播开始的时间，是以毫秒为单位的 Unix 时间
      "liveType": {
        // 直播分类
        "categoryID": 4,
        "categoryName": "虚拟偶像",
        "subCategoryID": 401,
        "subCategoryName": "聊天"
      },
      "reserve": false, // 登陆帐号是否预约了该直播
      "reserveNumber": 135 // 已预约用户的数量
    }
  ]
}
```

#### 直播间礼物列表

##### 请求

```json
{
  "type": 114,
  "requestID": "abc",
  "data": {
    "liveID": "cgbKNA8R5nY"
  }
}
```

##### 响应

```json
{
  "type": 114,
  "requestID": "abc",
  "result": 1,
  "data": [
    {
      "giftID": 16,
      "giftName": "猴岛",
      "arLiveName": "",
      "payWalletType": 1,
      "price": 2888,
      "webpPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200316111119oxpFWYGQ.webp",
      "pngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200812141000WkBybKGr.png",
      "smallPngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200316111145bJSDGWHB.png",
      "allowBatchSendSizeList": [1, 6, 10],
      "canCombo": true,
      "canDraw": false,
      "magicFaceID": 264,
      "vupArID": 0,
      "description": "我要让所有人知道，这座猴岛，被我承包了！",
      "redpackPrice": 288,
      "cornerMarkerText": ""
    },
    {
      "giftID": 36,
      "giftName": "爱你哟",
      "arLiveName": "ac102",
      "payWalletType": 1,
      "price": 5200,
      "webpPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200810163045cOrcctqJ.webp",
      "pngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200812161659mhqzRjiA.png",
      "smallPngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200810163054BlSAUWVG.png",
      "allowBatchSendSizeList": [],
      "canCombo": true,
      "canDraw": false,
      "magicFaceID": 0,
      "vupArID": 0,
      "description": "爱你哟～",
      "redpackPrice": 0,
      "cornerMarkerText": ""
    }
  ]
}
```

#### 指定用户的信息

##### 请求

```json
{
  "type": 115,
  "requestID": "abc",
  "data": {
    "userID": 26675034
  }
}
```

##### 响应

```json
{
  "type": 115,
  "requestID": "abc",
  "result": 1,
  "data": {
    "userID": 17912421,
    "nickname": "茗魂",
    "avatar": "https://tx-free-imgs.acfun.cn/content/2021_2_1/1.6121163259094958E9.png?imageslim",
    "avatarFrame": "https://imgs.aixifan.com/5qaMBO23v2-uquqYf-6rMVvq-uI7jMn-7Fr6fi.png",
    "followingCount": "253",
    "fansCount": "1.7 万",
    "contributeCount": "163",
    "signature": "身为阎王之女为了来人间考察民情，附身在了人类菲艾身上~\n录播@茗魂录播机器人\n粉丝群 481148878，欢迎来玩~",
    "verifiedText": "AVI 联盟成员，AcFun 签约虚拟偶像",
    "isJoinUpCollege": true,
    "isFollowing": true,
    "isFollowed": false,
    "liveID": "UM8fXJmceqU", // 用户正在直播时不为空
    "likeCount": 106236 // 最近一次直播的点赞总数
  }
}
```

#### 直播剪辑信息

##### 请求

```json
{
  "type": 116,
  "requestID": "abc",
  "data": {
    "liverUID": 23512715,
    "liveID": "1UiXanLh"
  }
}
```

##### 响应

```json
{
  "type": 116,
  "requestID": "abc",
  "result": 1,
  "data": {
    "status": true, // 是否允许剪辑直播录像（主播允许观众剪辑观众才能剪辑，主播直播时总是能剪辑自己的直播）
    "url": "https://onvideo.kuaishou.com/vangogh/editor/9724?source=ac", // 剪辑直播的地址，直接访问可能出现登陆问题，需要访问跳转地址
    "redirectURL": "https://onvideoapi.kuaishou.com/rest/infra/sts?authToken=ChZhY2Z1bi5taWRncm91bmQuYXBpLEmBHB6vV-m_r9Ui03LZGAC6dZGk8zfyiJ33KVWJmYKFooH-mZ1oFrzk9uE7SAO2cvem2SXh7UE-LLrl6XipyxQTyHghRaij7mnULlz5c_oJCjvVO6SW5FwfT2-MaEuvSHyafvqmlvpZIVTcO0MbnzCIgUQqizvmhPHaoXQDltf6f1xDRl8f9pu7hdgRQXEoBTAB\u0026sid=acfun.midground.api\u0026followUrl=https://onvideo.kuaishou.com/vangogh/editor/9724?source=ac" // 跳转直播剪辑的地址，访问一次后链接里的 token 就会失效
  }
}
```

#### 登陆用户的房管列表

##### 请求

```json
{
  "type": 200,
  "requestID": "abc"
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
        "nickname": "AC 娘本体",
        "avatar": "https://imgs.aixifan.com/FnlcvTfQHideC2bGgfRO2u9gfig_",
        "medal": {
          // 没有守护徽章信息
          "uperID": 0,
          "userID": 0,
          "clubName": "",
          "level": 0
        },
        "managerType": 0 // 没有房管信息
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

`managerUID`：房管的 uid

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

#### 主播踢人记录

##### 请求

```json
{
  "type": 203,
  "requestID": "abc",
  "data": {
    "liveID": "cgbKNA8R5nY"
  }
}
```

查询 liveID 指定直播的主播的踢人记录，需要[登陆](#登陆)主播的 AcFun 帐号

##### 响应

```json
{
  "type": 203,
  "requestID": "abc",
  "result": 1,
  "data": [
    {
      "userID": 45443067, // 被踢用户的 uid
      "nickname": "ACFUN-FOSS_开源⑨课", // 被踢用户的名字
      "kickTime": 1612874404648 // 用户被踢的时间，是以毫秒为单位的 Unix 时间
    }
  ]
}
```

#### 房管踢人

##### 请求

```json
{
  "type": 204,
  "requestID": "abc",
  "data": {
    "liveID": "cgbKNA8R5nY",
    "kickedUID": 12345
  }
}
```

`liveID`：主播正在直播的 liveID，需要[登陆](#登陆)帐号有对应直播间的房管权限

`kickedUID`：被踢的用户的 uid

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
    "liveID": "cgbKNA8R5nY",
    "kickedUID": 12345
  }
}
```

`liveID`：主播正在直播的 liveID，需要[登陆](#登陆)主播的 AcFun 帐号

##### 响应

```json
{
  "type": 205,
  "requestID": "abc",
  "result": 1
}
```

#### 登陆用户拥有的指定主播守护徽章详细信息

##### 请求

```json
{
  "type": 300,
  "requestID": "abc",
  "data": {
    "liverUID": 12891327
  }
}
```

##### 响应

```json
{
  "type": 300,
  "requestID": "abc",
  "result": 1,
  "data": {
    "medal": {
      "medalInfo": {
        "uperID": 12891327, // 主播 uid
        "userID": 103411, // 用户 uid
        "clubName": "白白的", // 守护徽章名字
        "level": 10 // 守护徽章等级
      },
      "uperName": "白白_Anlessya", // 主播名字
      "uperAvatar": "https://tx-free-imgs.acfun.cn/content/2020_12_27/1609078598105.JPG?imageslim", // 主播头像
      "wearMedal": false, // 用户是否佩戴主播的守护徽章
      "friendshipDegree": 18015, // 目前用户的守护徽章亲密度
      "joinClubTime": 1602752837775, // 用户加入主播守护团的时间，是以毫秒为单位的 Unix 时间
      "currentDegreeLimit": 18887, // 用户守护徽章目前等级的亲密度的上限
      "medalCount": 0 // 指定用户拥有的守护徽章数量
    },
    "medalDegree": {
      "uperID": 12891327,
      "giftDegree": 0, // 本日送直播礼物增加的亲密度
      "giftDegreeLimit": 2000, // 本日送直播礼物增加的亲密度上限
      "peachDegree": 0, // 本日投桃增加的亲密度
      "peachDegreeLimit": 520, // 本日投桃增加的亲密度上限
      "liveWatchDegree": 331, // 本日看直播时长增加的亲密度
      "liveWatchDegreeLimit": 360, // 本日看直播时长增加的亲密度上限
      "bananaDegree": 0, // 本日投蕉增加的亲密度
      "bananaDegreeLimit": 10 // 本日投蕉增加的亲密度上限
    },
    "userRank": "31" // 用户的守护徽章亲密度的排名
  }
}
```

#### 登陆用户拥有的守护徽章列表

##### 请求

```json
{
  "type": 301,
  "requestID": "abc"
}
```

`liverUID`：用于获取用户拥有的指定主播守护徽章的详细信息，可以为 0

##### 响应

```json
{
  "type": 301,
  "requestID": "abc",
  "result": 1,
  "data": [
    // 用户拥有的守护徽章列表，有可能为`null`
    {
      "medalInfo": {
        "uperID": 265135,
        "userID": 103411,
        "clubName": "处男",
        "level": 14
      },
      "uperName": "力力力力宝",
      "uperAvatar": "https://tx-free-imgs.acfun.cn/Fi1sC-eFeM6AyH7Hp7WZvUOkOULH?imageslim",
      "wearMedal": true,
      "friendshipDegree": 94486,
      "joinClubTime": 1600592466196,
      "currentDegreeLimit": 128887,
      "medalCount": 186 // 用户拥有的守护徽章数量
    }
  ]
}
```

#### 主播守护榜

##### 请求

```json
{
  "type": 302,
  "requestID": "abc",
  "data": {
    "liverUID": 26675034
  }
}
```

##### 响应

```json
{
  "type": 302,
  "requestID": "abc",
  "result": 1,
  "data": {
    "hasFansClub": true, // 主播是否有守护团
    "rankList": [
      //有可能为`null`
      {
        "profile": {
          // 用户信息
          "userID": 7755,
          "nickname": "saga-R",
          "avatar": "https://tx-free-imgs.acfun.cn/style/image/201907/w1v9fCGV7uhiNYbv5c5hAfAHehgVyHuM.jpg?imageslim",
          "avatarFrame": "http://imgs.aixifan.com/o_1cvf33pb5e1q1puncp21ku0a3446.png",
          "followingCount": 723,
          "fansCount": 23,
          "contributeCount": 20,
          "signature": "不写不写！",
          "verifiedText": "",
          "isJoinUpCollege": false,
          "isFollowing": false,
          "isFollowed": false
        },
        "friendshipDegree": 57676,
        "level": 13
      }
    ],
    "clubName": "偷芯猫",
    "medalCount": 991, // 主播守护团总人数
    "hasMedal": true, // 登陆用户是否有主播的守护徽章
    "userFriendshipDegree": 14234, // 目前登陆用户的守护徽章亲密度
    "userRank": "79" // 登陆用户的守护徽章亲密度的排名
  }
}
```

#### 指定用户正在佩戴的守护徽章信息

##### 请求

```json
{
  "type": 303,
  "requestID": "abc",
  "data": {
    "userID": 26675034
  }
}
```

##### 响应

```json
{
  "type": 303,
  "requestID": "abc",
  "result": 1,
  "data": {
    "medalInfo": {
      "uperID": 16005,
      "userID": 26675034,
      "clubName": "坏女人",
      "level": 1
    },
    "uperName": "乌拉喵",
    "uperAvatar": "https://imgs.aixifan.com/style/image/201907/4jUWfchONKiX1yLuhkGUtKfUl1YOJatr.jpg",
    "wearMedal": true,
    "friendshipDegree": 0, // 没有亲密度
    "joinClubTime": 0, // 没有用户加入守护团的时间
    "currentDegreeLimit": 0, // 没有目前等级的亲密度上限
    "medalCount": 100 // 用户拥有的守护徽章数量
  }
}
```

#### 佩戴守护徽章

##### 请求

```json
{
  "type": 304,
  "requestID": "abc",
  "data": {
    "liverUID": 26675034
  }
}
```

`liverUID`：要佩戴的守护徽章的主播 uid

##### 响应

```json
{
  "type": 304,
  "requestID": "abc",
  "result": 1
}
```

#### 取消佩戴守护徽章

##### 请求

```json
{
  "type": 305,
  "requestID": "abc"
}
```

##### 响应

```json
{
  "type": 305,
  "requestID": "abc",
  "result": 1
}
```

#### 检测开播权限

##### 请求

```json
{
  "type": 900,
  "requestID": "abc"
}
```

##### 响应

```json
{
  "type": 900,
  "requestID": "abc",
  "result": 1,
  "data": {
    "liveAuth": true
  }
}
```

`liveAuth`：能否直播

#### 直播分类列表

##### 请求

```json
{
  "type": 901,
  "requestID": "abc"
}
```

##### 响应

```json
{
  "type": 901,
  "requestID": "abc",
  "result": 1,
  "data": [
    {
      "categoryID": 1, // 直播主分类 ID
      "categoryName": "游戏直播", // 直播主分类名字
      "subCategoryID": 101, // 直播次分类 ID
      "subCategoryName": "英雄联盟" // 直播次分类名字
    }
  ]
}
```

#### 推流设置

##### 请求

```json
{
  "type": 902,
  "requestID": "abc"
}
```

##### 响应

```json
{
  "type": 902,
  "requestID": "abc",
  "result": 1,
  "data": {
    "streamName": "kszt_PYrssS_J4w", // 直播源名字
    "streamPullAddress": "https://tx-adaptive.pull.yximgs.com/livecloud/kszt_PYrssS_J4w.flv?txSecret=adfd9fcb80b9d8f6d0071ba88f33ee8b\u0026txTime=603eadd2\u0026stat=XIFGbCNUSzcScMvRvgKb%2FT%2FT2mInuvBYcy5eD%2FRbbmk%3D\u0026oidc=alihb", // 拉流地址，也就是直播源地址
    "streamPushAddress": [
      // 推流地址，目前分为阿里云和腾讯云两种，有可能为`null`
      "rtmp://aliyun-open-push.voip.yximgs.com/livecloud/kszt_PYrssS_J4w?sign=c0377c25-c6e74ddb3ea81bd98c7279d87a16ae75\u0026ks_fix_ts\u0026ks_ctx=dHRwOlBVTEw7dGZiOjE7dmVyOjYzMTtwZHk6MDt2cXQ6VU5LTk9XTjtpc1Y6ZmFsc2U7YWlkOjEwMzQxMQ%3D%3D",
      "rtmp://txyun-open-push.voip.yximgs.com/livecloud/kszt_PYrssS_J4w?sign=c0377c25-c6e74ddb3ea81bd98c7279d87a16ae75\u0026ks_fix_ts\u0026ks_ctx=dHRwOlBVTEw7dGZiOjE7dmVyOjYzMTtwZHk6MDt2cXQ6VU5LTk9XTjtpc1Y6ZmFsc2U7YWlkOjEwMzQxMQ%3D%3D"
    ],
    "panoramic": false, // 是否全景直播
    "interval": 5000, // 发送查询转码信息的时间间隔，单位为毫秒
    "rtmpServer": "rtmp://aliyun-open-push.voip.yximgs.com/livecloud", // RTMP 服务器
    "streamKey": "kszt_PYrssS_J4w?sign=c0377c25-c6e74ddb3ea81bd987279d87a16ae75\u0026ks_fix_ts\u0026ks_ctx=dHRwOlBVTEw7dGZiOjE7dmVyOjYzMTtwZHk6MDt2cXQ6VULTk9XTjtpc1Y6ZmFsc2U7YWlkOjEwMzQxMQ%3D%3D" // 直播码/串流密钥
  }
}
```

#### 直播状态

##### 请求

```json
{
  "type": 903,
  "requestID": "abc"
}
```

开播后才有返回

##### 响应

```json
{
  "type": 903,
  "requestID": "abc",
  "result": 1,
  "data": {
    "liveID": "yECC9bopbF",
    "streamName": "kszt_PYrssS_J4w",
    "title": "听歌", // 直播间标题
    "liveCover": "https://ali-live.static.yximgs.com/bs2/ztlc/cover_yECC9bopbF_compress.webp", // 直播间封面
    "liveStartTime": 1612128526972, // 直播开始的时间，是以毫秒为单位的 Unix 时间
    "panoramic": false, // 是否全景直播
    "bizUnit": "acfun", // 通常是"acfun"
    "bizCustomData": "{\"typeId\":399,\"type\":[3,399]}" // 直播分类，格式是 json
  }
}
```

#### 转码信息

##### 请求

```json
{
  "type": 904,
  "requestID": "abc",
  "data": {
    "streamName": "cdefg"
  }
}
```

`streamName`从[推流设置](#推流设置)那里获得

##### 响应

```json
{
  "type": 904,
  "requestID": "abc",
  "result": 1,
  "data": [
    {
      "streamURL": {
        "url": "https://tx-adaptive.pull.yximgs.com/livecloud/kszt_PYrssS_J4w_hd2000.flv?txSecret=75073aa3db830c6e9ab9d50c5b97640\u0026txTime=603eae0e\u0026stat=XIFGbCNUSzcScMvRvgKb%2FT%2FT2mInuvBcy5eD%2FRbbmk%3D\u0026tsc=origin\u0026oidc=alihb", // 直播源链接
        "bitrate": 2000, // 直播源码率，不一定是实际码率
        "qualityType": "HIGH", // 直播源类型
        "qualityName": "超清" // 直播源类型的中文名字
      },
      "resolution": "1280x720", // 直播视频分辨率
      "frameRate": 26, // 直播视频 FPS？
      "template": "hd2000" // 直播模板？
    }
  ]
}
```

`data`不为空说明服务器开始转码，推流成功，可以[开始直播](#开始直播)

#### 开始直播

##### 请求

```json
{
  "type": 905,
  "requestID": "abc",
  "data": {
    "title": "测试", // 直播间标题
    "coverFile": "cdefg.jpg", // 直播间封面图片（可以是 gif）的本地路径或网络链接，可以先上传图片到 AcFun 服务器获得图片链接
    "streamName": "ghijkd", // 直播源名字，从推流设置那里获得
    "portrait": false, // 是否手机直播
    "panoramic": false, // 是否全景直播
    "categoryID": 3, // 直播主分类 ID
    "subCategoryID": 399 // 直播次分类 ID
  }
}
```

##### 响应

```json
{
  "type": 905,
  "requestID": "abc",
  "result": 1,
  "data": {
    "liveID": "yECC9bopbF"
  }
}
```

#### 停止直播

##### 请求

```json
{
  "type": 906,
  "requestID": "abc",
  "data": {
    "liveID": "cdefg"
  }
}
```

`liveID`从[开始直播](#开始直播)那里获得

##### 响应

```json
{
  "type": 906,
  "requestID": "abc",
  "result": 1,
  "data": {
    "duration": 2459600, // 直播时长，单位为毫秒
    "endReason": "author stopped" // 停止直播的原因
  }
}
```

#### 更改直播间标题和封面

##### 请求

```json
{
  "type": 907,
  "requestID": "abc",
  "data": {
    "title": "测试",
    "coverFile": "cdefg.jpg",
    "liveID": "hijklmn"
  }
}
```

`title`：直播标题，为空时会导致没有标题

`coverFile`：直播间封面图片（可以是 gif）的本地路径或网络链接（可以先[上传图片](#上传图片)到 AcFun 服务器获得图片链接），为空时只改变直播间标题

`liveID`从[开始直播](#开始直播)那里获得

##### 响应

```json
{
  "type": 907,
  "requestID": "abc",
  "result": 1
}
```

#### 查询是否允许观众剪辑直播录像

##### 请求

```json
{
  "type": 908,
  "requestID": "abc"
}
```

##### 响应

```json
{
  "type": 908,
  "requestID": "abc",
  "result": 1,
  "data": {
    "canCut": true
  }
}
```

#### 设置是否允许观众剪辑直播录像

##### 请求

```json
{
  "type": 909,
  "requestID": "abc",
  "data": {
    "canCut": true
  }
}
```

登陆帐号未在直播时才能设置

##### 响应

```json
{
  "type": 909,
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
      "sendTime": 1608379795363, // 弹幕发送时间，是以毫秒为单位的 Unix 时间
      "userInfo": {
        // 发送弹幕的用户的信息
        "userID": 666609, // 用户 uid
        "nickname": "泼墨造一匹快马追回十年前姑娘", // 用户名字
        "avatar": "https://imgs.aixifan.com/content/2020_09_20/1600575703124.JPG", // 用户头像
        "medal": {
          // 用户正在佩戴的守护徽章
          "uperID": 4537972, // UP 主的 uid
          "userID": 666609, // 用户的 uid
          "clubName": "有猫饼", // 守护徽章名字
          "level": 8 // 守护徽章等级
        },
        "managerType": 0 // 用户是否房管，0 不是房管，1 是房管
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
        "medal": {
          // 没有守护徽章信息
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
    "giftDetail": {
      // 礼物详细信息
      "giftID": 12, // 礼物 ID
      "giftName": "打 Call", // 礼物名字
      "arLiveName": "", // 不为空时礼物属于虚拟偶像区的特殊礼物
      "payWalletType": 1, // 1 为非免费礼物，2 为免费礼物
      "price": 10, // 礼物价格，非免费礼物时单位为 AC 币，免费礼物（香蕉）时为 1
      "webpPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200326112443kfWQhpaG.webp", // 礼物的 webp 格式图片（动图）
      "pngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200812141131ukNHkGeU.png", // 礼物的 png 格式图片（大）
      "smallPngPic": "https://static.yximgs.com/bs2/giftCenter/giftCenter-20200316110407BdolKFLb.png", // 礼物的 png 格式图片（小）
      "allowBatchSendSizeList": [
        // 网页或 APP 单次能够赠送的礼物数量列表，有可能为`null`
        1, 6, 10, 66, 233
      ],
      "canCombo": true, // 是否能连击，一般免费礼物（香蕉）不能连击，其余能连击
      "canDraw": true, // 是否能涂鸦？
      "magicFaceID": 0,
      "vupArID": 0,
      "description": "为你转身为你爆灯，为你打 CALL 日夜不分", // 礼物的描述
      "redpackPrice": 0, // 礼物红包价格总额，单位为 AC 币
      "cornerMarkerText": ""
    },
    "count": 5, // 礼物单次赠送的数量，礼物总数是 Count * Combo
    "combo": 1, // 礼物连击数量，礼物总数是 Count * Combo
    "value": 50000, // 礼物价值，非免费礼物时单位为 AC 币*1000，免费礼物（香蕉）时单位为礼物数量
    "comboID": "FD7E07B5-DF69-4257-84BC-7FEA377E9C85", // 礼物连击 ID
    "slotDisplayDuration": 3000, // 应该是礼物动画持续的时间，单位为毫秒，送礼物后在该时间内再送一次可以实现礼物连击
    "expireDuration": 300000,
    "drawGiftInfo": {
      // 礼物涂鸦
      "screenWidth": 1440, // 手机屏幕宽度
      "screenHeight": 2560, // 手机屏幕高度
      "drawPoint": [
        // 涂鸦里各个礼物的位置，没涂鸦时为`null`
        {
          "marginLeft": 393, // 到手机屏幕左边的距离
          "marginTop": 263, // 到手机屏幕顶部的距离
          "scaleRatio": 1, // 放大倍数？
          "handup": false,
          "pointWidth": 1, // 点的宽度？
          "pointHeight": 1 // 点的高度？
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
    "sendTime": 0, // 富文本的发送时间，是以毫秒为单位的 Unix 时间，可能为 0
    "segments": [
      // 富文本各部分，类型分别是 RichTextUserInfo、RichTextPlain 或 RichTextImage
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
          "text": " 发的 5AC 币",
          "color": ""
        }
      },
      {
        "type": 1902, // RichTextImage
        "segment": {
          "pictures": [
            // 图片
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

富文本有三种：`RichTextUserInfo`、`RichTextPlain`、`RichTextImage`，对应的`type`分别为 1900、1901、1902

##### 加入守护团

```json
{
  "liverUID": 2889712,
  "type": 1007,
  "data": {
    "joinTime": 1608390090583, // 用户加入守护团的时间，是以毫秒为单位的 Unix 时间
    "fansInfo": {
      // 用户的信息
      "userID": 1428790,
      "nickname": "柳昭郎",
      "avatar": "", // 没有用户头像
      "medal": {
        // 没有守护徽章信息
        "uperID": 0,
        "userID": 0,
        "clubName": "",
        "level": 0
      },
      "managerType": 0 // 没有房管类型
    },
    "uperInfo": {
      // 主播的信息
      "userID": 2889712,
      "nickname": "张梓义",
      "avatar": "", // 没有用户头像
      "medal": {
        // 没有守护徽章信息
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

##### 分享直播间

```json
{
  "liverUID": 378269,
  "type": 1008,
  "data": {
    "danmuInfo": {
      "sendTime": 1637160244836,
      "userInfo": {
        "userID": 103411,
        "nickname": "Orzogc",
        "avatar": "https://imgs.aixifan.com/style/image/201907/vVh2gcmMBSa173bVu1bzgJ3cjQlNcPJr.jpg",
        "medal": {
          "uperID": 378269,
          "userID": 103411,
          "clubName": "QQ 星",
          "level": 14
        },
        "managerType": 0
      }
    },
    "sharePlatform": 4, // 0 为未知平台，1 为 QQ，2 为 QQ 空间，3 应该是新浪微博，4 为微信，5 为微信朋友圈，6 为 A 站动态
    "sharePlatformIcon": "https://tx-free-imgs2.acfun.cn/udata/pkg/acfun/live-share-platform/common_sharing_pop_ups_wx_3x.png" // 平台图标
  }
}
```

#### 状态信号类型

##### 获取弹幕结束

```json
{
  "liverUID": 12345,
  "type": 2000
}
```

直播结束或客户端请求[停止获取弹幕](#停止获取弹幕)时发送

##### 获取弹幕出错

```json
{
  "liverUID": 12345,
  "type": 2999,
  "data": {
    "error": "abcd"
  }
}
```

获取弹幕出现错误时发送，偶然直播结束也会出现错误

`error`：获取弹幕出错信息

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
    "likeCount": "2.5 万",
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
  "data": [
    // 最多三位观众
    {
      "userInfo": {
        "userID": 496725,
        "nickname": "病娇御姐看起来老霸道了",
        "avatar": "https://imgs.aixifan.com/style/image/201907/P044fP0S6xaP83vSsZ1RsoUmQ4Uss0Ze.jpg",
        "medal": {
          // 没有守护徽章
          "uperID": 0,
          "userID": 0,
          "clubName": "",
          "level": 0
        },
        "managerType": 0 // 没有房管类型
      },
      "anonymousUser": false, // 是否匿名用户
      "displaySendAmount": "486", // 赠送的全部礼物的价值，单位是 AC 币，注意不一定是纯数字的字符串
      "customData": "{\"userInfo\":{\"verified\":0,\"verifiedTypes\":[],\"joinUpCollege\":false},\"countInfo\":{\"fansCount\":2}}" // 用户的一些额外信息，格式为 json
    }
  ]
}
```

##### 进直播间时显示的最近弹幕

```json
{
  "liverUID": 4425861,
  "type": 2004,
  "data": [
    // 最多十条弹幕
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
      "content": "只有靠直击 2，逆 2 了"
    }
  ]
}
```

##### 红包列表

```json
{
  "liverUID": 13945614,
  "type": 2005,
  "data": [
    // 没有红包时为空
    {
      "userInfo": {
        "userID": 13240469,
        "nickname": "汤汤湯湯汤",
        "avatar": "https://imgs.aixifan.com/content/2020_09_23/1600853696165.JPG",
        "medal": {
          // 没有守护徽章
          "uperID": 0,
          "userID": 0,
          "clubName": "",
          "level": 0
        },
        "managerType": 0
      },
      "displayStatus": 0, // 红包状态，0 是红包出现，1 是可以获取红包 token，2 是可以抢红包
      "grabBeginTime": 1608464088394, // 开始抢红包的时间，是以毫秒为单位的 Unix 时间
      "getTokenLatestTime": 1608464086394, // 抢红包的用户获得 token 的最晚时间，是以毫秒为单位的 Unix 时间
      "redpackID": "c5N6p7IMyjA", // 红包 ID
      "redpackBizUnit": "ztLiveAcfunRedpackGift", // "ztLiveAcfunRedpackGift"代表的是观众，"ztLiveAcfunRedpackAuthor"代表的是主播？
      "redpackAmount": 99, // 红包的总价值，单位是 AC 币
      "settleBeginTime": 1608464148394 // 抢红包的结束时间，是以毫秒为单位的 Unix 时间
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
    "signalInfo": "abcde"
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

##### 直播警告

```json
{
  "liverUID": 4425861,
  "type": 3001,
  "data": {
    "violationContent": "abc"
  }
}
```

`violationContent`：警告内容

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

`managerState`：0 不是房管，1 被添加房管，2 被移除房管，3 是房管

### 错误码

`10`：请求的 json 解析错误

`11`：请求`type`无效

`12`：请求`data`无效

`13`：处理请求时出现错误

`14`：需要登陆

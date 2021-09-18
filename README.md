极光推送Golang客户端
===


目前支持

* Push API v3
* Schedule API v3
* Device API v3


### 安装

    go get github.com/zwczou/jpush

### 使用

1. 初始化客户端

```go
    client := jpush.New("key", "secret")
```

2. 获取推送唯一标识符 cid

```go
    cidList, err = client.PushCid(1, "push")
```

3. 推送消息

```go
    payload := &jpush.Payload{
        Platform: jpush.NewPlatform().All(),
        Audience: jpush.NewAudience().All().SetTag("abc", "ef").SetTagAnd("filmtest"),
        Notification: &jpush.Notification{
            Alert: "test",
        },
        Options: &jpush.Options{
             TimeLive:       60,
             ApnsProduction: false,
        },
    }
    msgId, err = client.Push(payload)
    // msgId, err = client.PushValidate(payload)
```


4. 创建计划任务

```
    client.ScheduleCreate
```

5. 方便扩展

如果库没有实现你想使用的方法，可以使用`client.Do`扩展

```go
// 获取任务的所有msg_id
var out struct {
  MsgIds []struct {
    MsgId string `json:"msg_id"`
  } `json:"msgids"`
}
err := client.Do(http.MethodGet, "/schedules/"+scheduleId+"/msg_ids", nil, &out)
```

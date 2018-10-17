#SC

1.初始化WebSocketClient结构体，比如：

```
h := &Handler{
}
wsc := sc.WebSocketClient{
   URL:     url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/SCIDE/SCExecutor"},
   Handler: h,
}
```

2.启动websocket客户端

```
wsc.Start()
```

3.向服务端发消息

```
wsc.Send(sc.StartContract{
   Action:"startContract",
   Contractid: "66668",
   Path: "Hello",
   Script: "contract abc{\n    export function main(arg){\n        point = JSON.parse(arg);\n        var s = 0;\n        print(point[0].score);\n        print(point.length);\n        for (var i=0;i<point.length;i++){\n            s+=point[i].score/1.0;\n        }\n        print(\"ss= \"+s);\n        return s;\n    }\n}",
   Type: "",
   Onwer: "",
})
```
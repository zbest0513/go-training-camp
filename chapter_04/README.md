# 项目概述
通知系统，采集服务器的日志，筛选出特定的异常，通过企业微信发送给指定的人(或群聊)

## 服务
系统分agent和server，2个服务

### agent
agent负责在目标机器采集日志，并按规则筛选日志，将server感兴趣的日志投递给server

### server
server负责接收来自agent发送的日志，给日志打对应的标签，发送到对应的接收人或群聊。

目录结构

```
.
└── notify
    ├── api
    │         └── server    
    │             └── service   
    ├── app
    │         ├── agent                                                     
    │         │         ├── configs
    │         │         └── internal
    │         └── server
    │             ├── cmd
    │             │         └── server
    │             ├── configs
    │             └── internal
    │                 ├── biz
    │                 ├── data                      
    │                 │         └── ent
    │                 ├── pkg
    │                 └── service
    │                     ├── handle
    │                     └── router
    ├── doc
    └── pkg
        ├── config      
        └── utils                       
35 directories
```

#目录结构

```
.
└── notify
    ├── api                                        -- 接口描述/api文档
    │         └── server    
    │             └── service   
    ├── app                                        -- 应用目录：本项目分agent和server2个应用
    │         ├── agent                            -- agent 应用根目录                         
    │         └── server                           -- server 应用根目录
    │             ├── cmd                          -- cmd 为 go build 的目录，放main.go和wire
    │             │         └── server
    │             ├── configs                      -- server 应用的配置文件目录
    │             └── internal                     -- server 应用的内部逻辑目录，对外不可见
    │                 ├── biz                      -- server biz 层 ，DO与UseCase
    │                 ├── data                     -- server data 层，PO 与持久化层
    │                 │         └── ent            -- server 持久化层应用 ent框架自动生成的代码
    │                 ├── pkg                      -- server 应用内的公共组建包
    │                 └── service                  -- server service层，我理解对应java体系的controller层
    │                     ├── handle               -- server service层的处理逻辑，负责将DTO转化成DO，校验等
    │                     └── router               -- server service的路由，绑定url与handle 
    ├── doc                                        -- 整体项目的文档，概要设计、详细设计等
    └── pkg                                        -- 整体项目的公共组件包，子项目可以本地依赖
        ├── config                                 -- 整体项目的公共配置文件
        └── utils                                  -- 整体项目的公共工具包，log/类型转换等
```
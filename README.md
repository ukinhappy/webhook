# webhook
webhook client

# 使用方法
### 1. 获取源码

```shell script
git clone git@github.com:ukinhappy/webhook.git
```
### 2.编辑

```shell script
cd webhook
go build
```

### 3. 配置config.toml

```
[http]
addr = "127.0.0.1:15890" # 服务
[[projects]]
name = "test-2" # 仓库名
branch = ["master"] # 接收指定分支的hook
event = ["push"] # 接收指定事件的hook
shellpath=""# 接收到通知后执行的shell文件


[log]
path = "" #日志路径
debug = true
maxsize = 10 #M
maxage = 7
maxbackups = 10
```

### 4. 运行
```shell script
./webhook --config=config.toml
```

### 5. 配置github webhook
```
127.0.0.1:15890/webhook_deploy
```
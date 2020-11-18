# AntFund
蚂蚁基金查询

# 初始化gomod
```
go mod init AntFund
```

# 安装支持库
```
go build ./...
```
如果不自动安装easyjson
```
go get -u github.com/mailru/easyjson/
go install  github.com/mailru/easyjson/easyjson
go build -o easyjson github.com/mailru/easyjson/easyjson
```
cd easyjson

eaeasyjson initjson_easyjson.go

跨平台编译
go tool dist list 查看支持的平台
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

conf/conf.txt是存放基金代码的配置文件，把main文件和conf文件放在一起

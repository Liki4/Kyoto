# Kyoto

随便写写
因为老是错过
写来给我推送一些京阿尼角色的生日

可以通过 TgBot 推送到 Channel

本来想所有京阿尼角色都整理成一张表
但是很多角色没有数据
于是写个用来添加生日信息的接口

哦对
现在数据库里只有我喜欢的角色
有需求就自己找资料加好了

[【盘点向】百位京阿尼动画角色の生日大全](https://www.bilibili.com/video/BV1m7411p7EG/)

以后也不一定更新
反正大概除了我也没人有这个需求了
心情好再说

## Build
```shell
go mod tidy
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build .
```

## Deploy
```shell
docker-compose up -d --build
```

## Api
`POST /birthday/add`
```
{
    "animate": "",
    "name": "",
    "date": "" // Format: %02d-%02d, like 01-01
}
```
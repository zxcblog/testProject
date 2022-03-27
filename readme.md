# 测试项目

## 检查docker-compose.yml文件错误
```shell
docker stack deploy wordpress -c=docker-compose.yml
```

## 使用docker-compose运行
```shell
cd docker-compose

docker-compose up -d
```
如果data数据被清空, 服务执行不起来时,请先将mysql服务启动, 添加数据库后在重新启动服务

## 用户登录注册
1. 用户注册后将用户名存储到 redis zset中， 用来判断用户名是否唯一
2. 将用户信息存储到redis缓存中， 通过用户名(唯一)来获取用户信息
3. 将id和用户名使用进行绑定



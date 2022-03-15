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


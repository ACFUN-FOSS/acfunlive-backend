# 测试
```
git clone https://github.com/ACFUN-FOSS/acfunlive-backend.git
cd acfunlive-backend
go build
./acfunlive-backend -debug
# 在另外一个终端
cd tests/obsClient
go build client.go
./client -account youraccount -password password -title1 标题1 -title2 标题2 -cover1 封面1.jpg -cover2 封面2.jpg
期间打印推流地址后应当及时用OBS推流
```

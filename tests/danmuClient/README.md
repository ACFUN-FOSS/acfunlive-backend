# 测试
```
git clone https://github.com/ACFUN-FOSS/acfunlive-backend.git
cd acfunlive-backend
go build
./acfunlive-backend -debug
# 在另外一个终端
cd tests/danmuClient
go build client.go
./client -account youraccount -password password -uid liveruid
```

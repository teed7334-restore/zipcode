# zipcode
台灣行政區域與郵遞區號匯入程式

## 資料夾結構
database 資料庫Driver檔

env 系統設定

hooks 呼叫第三方API

models 資料庫存取邏輯

main.go 主程式

## 程式運行原理
本程式會在你的資料庫新增一張表，叫zip_code

程式運行時，會將中華郵政提供的行政區域Open Data寫入到資料庫中

## 必須套件
本程式透過Google Protobuf 3產生所需之ResultObject，然Proto 3之後官方不支持Custom Tags，所以還需要多安裝一個寫入retags的套件

git clone https://github.com/qianlnk/protobuf.git $GOPATH/src/github.com/golang/protobuf

go install $GOPATH/src/github.com/golang/protobuf/protoc-gen-go

## 程式操作流程
1. 將./env/env.swp檔名改成env.go
2. 修改./env/env.go換成你連入的MySQL資料庫與帳號密碼
3. 到./beans底下，運行protoc --go_out=plugins=grpc+retag:. *.proto
4. go run main.go
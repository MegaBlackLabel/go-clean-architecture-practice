# go-clean-architecture-practice

## 参考にしていた記事
「速習 Go + Clean Architecture」
https://qiita.com/kurab/items/173a8ec8ea64ee09e5c3#infrastructure

## wire_gen生成
wire ./registry

## サーバ起動
go run main.go

## ユーザ登録
curl -X POST -H "Content-Type: application/json" -d '{\"Name\":\"test\"}' 127.0.0.1:8080/api/user/register
(Windowsの仕様でJSONのダブルコートを\でエスケープしている)

## ユーザ一覧
curl 127.0.0.1:8080/api/user/get

## ユーザ取得
curl 127.0.0.1:8080/api/user/get/(int)

### 環境
go version go1.19.2 linux/amd64  
### hardware
Mem 2G


### module packageダウンロード(初回のみ)
go mod init go-echo-api  
go mod tidy  

### その他
- Swagger（yaml）制御系
oapi-codegen v1.8.2  


-> openapiディレクトリ配下のtypes.gen.goとserver.gen.goをyamlから生成。
yamlを更新したら以下のShellを実行する。
oapi.sh




# 俺が考える、echoテンプレート（GOで書くAPI）

go-echo-sample

# Features

go lang を触ってみたいけどどうしたらいいか…。  
go って流行ってるけど実際にどういうふうに書くのか？  
go でAPIってどういうふうに書くのかイメージつかない…。  
docker-composeは動かしたことあるけど…。  
（docker-composeってもはや必須スキルですよね）


そんな人達のために書きました。


# Requirement

最近のdocker-compose


# Installation

dockerは公式サイトがあるので、そっちを見たほうがいいですね。

一応、これを作ったときのバージョン。
```
docker-compose version 1.29.2, build 5becea4c
```


# Usage

LinuxもしくはMacOSをベースに考えています。
```bash
git clone https://github.com/ajitama/go-echo-api.git
cd docker
docker-compose up
```


```bash
git clone https://github.com/ajitama/go-echo-api.git
docker-compose -f docker/docker-compose.yaml up
```


# Author
* ajitama

# License
ajitama  
google  


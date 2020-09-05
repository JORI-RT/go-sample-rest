# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work
#  goのrestのlibraryをインストール
RUN go get -u github.com/gorilla/mux
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD ./rest_sample /go/src/work
ENTRYPOINT ["go", "run", "main.go"] 

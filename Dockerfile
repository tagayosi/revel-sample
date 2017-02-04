FROM golang:1.7.5

RUN go get github.com/revel/revel

RUN go get github.com/revel/cmd/revel

RUN go get github.com/go-sql-driver/mysql

RUN go get github.com/jinzhu/gorm

RUN revel new revel-sample

COPY app.yaml src/revel-sample/
COPY app src/revel-sample/app/
COPY conf src/revel-sample/conf/
COPY messages src/revel-sample/messages
COPY public src/revel-sample/public

ENTRYPOINT revel run revel-sample dev 8080

EXPOSE 8080

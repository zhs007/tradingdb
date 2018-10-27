
FROM golang:1.10 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR $GOPATH/src/github.com/zhs007/tradingdb

COPY ./Gopkg.* $GOPATH/src/github.com/zhs007/tradingdb/

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure -vendor-only -v

COPY . $GOPATH/src/github.com/zhs007/tradingdb

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tradingdb . \
    && mkdir /home/tradingdb \
    && mkdir /home/tradingdb/dat \
    && mkdir /home/tradingdb/cfg \
    && cp tradingdb /home/tradingdb/ \
    && cp -r $GOPATH/src/github.com/zhs007/ankadb/www /home/tradingdb/www \
    && cp cfg/config.yaml.default /home/tradingdb/cfg/config.yaml

FROM scratch
WORKDIR /home/tradingdb
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /home/tradingdb /home/tradingdb
CMD ["./tradingdb"]

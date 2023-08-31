FROM golang:1.21

COPY . /go/src/app

WORKDIR /go/src/app/cmd/segmentation

RUN go build -o segmentation main.go

EXPOSE 8080

CMD ["./segmentation"]

FROM golang:1.19

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o awesomeProject ./main.go

CMD ["./awesomeProject"]
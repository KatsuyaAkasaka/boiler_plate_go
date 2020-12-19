FROM golang:latest
COPY ../../../ /go/src/github.com/KatsuyaAkasaka/boiler_plate_go
WORKDIR /go/src/github.com/KatsuyaAkasaka/boiler_plate_go
RUN go mod download
RUN ls
RUN go build -o main server/cmd/api/main.go
EXPOSE 3000
ENV ENV local
CMD ["./main"]

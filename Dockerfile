# golang base image
FROM golang:latest as builder

# add info 
LABEL maintainer="Jake VanCampen <jake.vancampen7@gmail.com>"

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
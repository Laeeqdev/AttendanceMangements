FROM golang:1.21.6  AS build-env
RUN apt update
RUN apt install git 
WORKDIR /go/src/github.com/Laeeqdev/AttendanceMangements/API
ADD . /go/src/github.com/Laeeqdev/AttendanceMangements/API/
RUN go mod download
CMD ["go","run","."]


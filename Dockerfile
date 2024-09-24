##
## Build
##

FROM golang:1.21.5 as builder

ARG SSH_PRIVATE_KEY
WORKDIR /app/
ADD . /app/


RUN  go mod download

ADD ./. ./
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
#RUN go build .

#RUN go build -gcflags="all=-N -l" -o access-server-debug
ENTRYPOINT CompileDaemon -build="go build -buildvcs=false ." -command="./letscollabServer"

##
## Debug
##

#FROM golang:alpine AS dev-debug

#RUN unset GOPATH

#RUN apk add build-base

#RUN go install github.com/go-delve/delve/cmd/dlv@latest

#RUN apk add libc6-compat

#WORKDIR /app

#COPY --from=builder /app/access-server-debug ./

##EXPOSE 40000 8080

#CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "--log", "/app/access-server-debug"]

##
## Deploy
##

FROM alpine as prod

WORKDIR /app

COPY --from=builder ./app/access-server ./access-server

RUN apk add libc6-compat

ENTRYPOINT [ "/app/access-server" ]

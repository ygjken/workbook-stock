FROM golang:1.16-alpine

RUN apk update
RUN apk add curl

ARG WORKDIR=/www/workbookstock
WORKDIR ${WORKDIR}
RUN mkdir -p ${WORKDIR}

RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin
RUN go get github.com/cosmtrek/air
RUN go get github.com/gin-contrib/sessions

EXPOSE 8080
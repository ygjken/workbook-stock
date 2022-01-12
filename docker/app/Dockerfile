FROM golang:1.16

RUN apt update

ARG WORKDIR=/www/workbookstock
WORKDIR ${WORKDIR}
RUN mkdir -p ${WORKDIR}

RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin
RUN go get github.com/cosmtrek/air
RUN go get github.com/gin-contrib/sessions
RUN go get golang.org/x/tools/gopls
RUN go get github.com/go-delve/delve/cmd/dlv

EXPOSE 8080
FROM golang

RUN mkdir -p /go/src//Users/wenyxu/Projects/senrui

ADD . /go/src//Users/wenyxu/Projects/senrui

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run /Users/wenyxu/Projects/senrui/todo/cmd  -watch /Users/wenyxu/Projects/senrui/todo

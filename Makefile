NAME     := greetings-server
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
server:

depinit:
	dep init

bin/$(NAME): $(SRCS)
	go build -a -tags greetings-server -installsuffix greetings-server $(LDFLAGS) -o bin/$(NAME)

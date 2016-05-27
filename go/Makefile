all: gobuild start

gobuild: goclean goinstall

gorebuild: goclean goreinstall

goclean:
	go clean -i ./...

goinstall:
	go get github.com/gin-gonic/gin
	go get gopkg.in/mgo.v2

goreinstall:
	go get -a -v github.com/gin-gonic/gin
	go get -a -v gopkg.in/mgo.v2

start:
	cd app && go run main.go

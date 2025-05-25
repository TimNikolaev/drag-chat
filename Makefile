build:
	go build -o ./.bin/chat cmd/chat/main.go

run:	build
	./.bin/chat
run:
	# gin --build . -i --bin build/go-snake
	go build . && ./go-snake

test:
	go test -v ./...
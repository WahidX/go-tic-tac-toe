run:
	# gin --build . -i --bin build/go-tic-tac-toe
	go build . && ./go-tic-tac-toe

test:
	go test -v ./...
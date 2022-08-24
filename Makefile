startdb:
	sudo systemctl start mongod

run:
	go run main.go

package:
	go mod tidy
get_dependencies:
	go get ./...

build: get_dependencies
	GOOS=darwin GOARCH=amd64 go build -o route2bimmer-mac cmd/route2bimmer.go
	GOOS=linux GOARCH=amd64 go build -o route2bimmer-linux cmd/route2bimmer.go
	GOOS=windows GOARCH=amd64 go build -o route2bimmer.exe cmd/route2bimmer.go

clean:
	rm route2bimmer-mac
	rm route2bimmer-linux
	rm route2bimmer.exe

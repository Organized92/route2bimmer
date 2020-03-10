get_dependencies:
	go get ./...

build: get_dependencies
	GOOS=darwin GOARCH=amd64 go build -o route2bimmer-mac route2bimmer/route2bimmer.go
	GOOS=linux GOARCH=amd64 go build -o route2bimmer-linux route2bimmer/route2bimmer.go
	GOOS=windows GOARCH=amd64 go build -o route2bimmer.exe route2bimmer/route2bimmer.go

clean:
	rm route2bimmer-mac
	rm route2bimmer-linux
	rm route2bimmer.exe

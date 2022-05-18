file := ./dist/kimcli
build:
	rm -rf dist
	env GOOS=windows GOARCH=amd64 go build -o $(file)-windows-amd64.exe
	env GOOS=windows GOARCH=386 go build -o $(file)-windows-i386.exe
	env GOOS=linux GOARCH=amd64 go build -o $(file)-linux-amd64
	env GOOS=linux GOARCH=386 go build -o $(file)-linux-i386
	env GOOS=darwin GOARCH=amd64 go build -o $(file)-darwin-amd64
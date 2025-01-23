all:
	go fmt main.go
	# go get -u github.com/gobuffalo/packr/v2/packr2
	GO111MODULE=on packr2
	GOOS=linux GOARCH=amd64 go build -o vim-init.linux
	# GOOS=darwin GOARCH=amd64 go build -o vim-init.darwin
	GO111MODULE=on packr2 clean

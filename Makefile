deploy:
	echo "Compiling....."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-w -s" -a -installsuffix cgo -o api.exe
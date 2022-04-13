go build -o AScanPort.exe main.go

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o AScanPort_darwin_amd64 main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o AScanPort_linux_amd64 main.go

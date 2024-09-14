# Run below if doing first time
``` bash
go mod init s3example
go get github.com/aws/aws-sdk-go-v2/aws
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3
go run main.go
```


# Run  if already set-up

```bash
go mod tidy
go run main.go
```
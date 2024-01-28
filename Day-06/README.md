# Build Commands for Various Platforms

This document provides the necessary build commands for different operating systems and architectures using Go.

```bash
# Darwin (macOS)
GOOS=darwin GOARCH=amd64 go build -o sic-darwin-amd64   # amd64
GOOS=darwin GOARCH=arm64 go build -o sic-darwin-arm64   # arm64

# Linux
GOOS=linux GOARCH=amd64 go build -o sic-linux-amd64     # amd64
GOOS=linux GOARCH=386 go build -o sic-linux-386         # 386
GOOS=linux GOARCH=arm go build -o sic-linux-arm         # arm
GOOS=linux GOARCH=arm64 go build -o sic-linux-arm64     # arm64

# Windows
GOOS=windows GOARCH=amd64 go build -o sic-windows-amd64.exe # amd64
GOOS=windows GOARCH=386 go build -o sic-windows-386.exe     # 386

# FreeBSD
GOOS=freebsd GOARCH=amd64 go build -o sic-freebsd-amd64 # amd64
GOOS=freebsd GOARCH=386 go build -o sic-freebsd-386     # 386
GOOS=freebsd GOARCH=arm go build -o sic-freebsd-arm     # arm


# Commans to archive binary files
tar -czvf myapp-linux.tar.gz myapp-linux
zip myapp-windows.zip myapp-windows.exe 

#git tag and release commands
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0
gh release create v1.0.0 ./path/to/first_asset.zip ./path/to/second_asset.zip --title "Release Title" --notes "Release notes"



manifest:
	@rsrc.exe -manifest hello.exe.manifest -o hello.exe.syso

build: manifest
	@go build -o hello.exe .
Set enviremental variable in PowerShell

$Env:GOOS = "js"; $Env:GOARCH = "wasm"

build command go build -o main.wasm

Server must recoginze wasm as a MIME Type

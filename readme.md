**Set environment variable in PowerShell :**

```

$Env:GOOS = "js"; $Env:GOARCH = "wasm"

```

**Build command :**

```

go build -o main.wasm

```

_Server must recognize WASM as a MIME Type_

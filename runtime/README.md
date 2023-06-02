# runtime [![GoDoc](https://pkg.go.dev/badge/github.com/luevano/mangal-lua-libs/runtime.svg)](https://pkg.go.dev/github.com/luevano/mangal-lua-libs/runtime)

Lua module for [runtime](https://pkg.go.dev/runtime).

## Usage

```lua
local runtime = require("runtime")
if not(runtime.goos() == "linux") then error("not linux") end
if not(runtime.goarch() == "amd64") then error("not amd64") end
```

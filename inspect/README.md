# inspect [![GoDoc](https://pkg.go.dev/badge/github.com/luevano/mangal-lua-libs/inspect.svg)](https://pkg.go.dev/badge/github.com/luevano/mangal-lua-libs/inspect)

## Usage

```lua
local inspect = require("inspect")

local table = {a={b=2}}
local result = inspect(table, {newline="", indent=""})
if not(result == "{a = {b = 2}}") then error("inspect") end
```

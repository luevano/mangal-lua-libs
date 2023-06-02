# stats [![GoDoc](https://pkg.go.dev/badge/github.com/luevano/mangal-lua-libs/stats.svg)](https://pkg.go.dev/github.com/luevano/mangal-lua-libs/stats)

Lua module for [github.com/montanaflynn/stats](https://pkg.go.dev/github.com/montanaflynn/stats).

## Usage

```lua
local stats = require("stats")

local result, _ = stats.median({0,0,10})
print(result)
-- Output: 0

local result, _ = stats.percentile({0,0,10}, 100)
print(result)
-- Output: 10
```

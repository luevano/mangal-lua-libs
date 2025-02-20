# goos [![GoDoc](https://pkg.go.dev/badge/github.com/luevano/mangal-lua-libs/goos.svg)](https://pkg.go.dev/github.com/luevano/mangal-lua-libs/goos)

Lua module for [os](https://pkg.go.dev/os).

## Usage

```lua
local goos = require("goos")

-- stat
local stat, err = goos.stat("./filename")
if err then error(err) end
print(stat.is_dir)
print(stat.size)
print(stat.mod_time)
print(stat.mode)

-- hostname
local hostname, err = goos.hostname()
if err then error(err) end
print(hostname)

-- get_pagesize
local page_size = goos.get_pagesize()
if not(page_size > 0) then error("bad pagesize") end

-- mkdir_all
goos.mkdir_all("./test/test_dir/test_dir/all")
local stat, err = goos.stat("./test/test_dir/test_dir/all")
if err then error(err) end
```

# shellescape [![GoDoc](https://pkg.go.dev/badge/github.com/luevano/mangal-lua-libs/shellescape.svg)](https://pkg.go.dev/github.com/luevano/mangal-lua-libs/shellescape)

Lua module for [github.com/alessio/shellescape](https://pkg.go.dev/github.com/alessio/shellescape).

## Usage

```lua
local shellescape = require("shellescape")

escaped = shellescape.quote("foo bar baz")
-- 'foo bar baz'

escaped_command = shellescape.quote_command({ "echo", "foo bar baz" })
-- echo 'foo bar baz'
```

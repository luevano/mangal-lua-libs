# mangal-lua-libs
[![GoDoc](https://pkg.go.dev/badge/github.com/luevano/mangal-lua-libs.svg)](https://pkg.go.dev/github.com/luevano/mangal-lua-libs)

[gopher-lua](https://github.com/yuin/gopher-lua) libs used in [mangal](https://github.com/luevano/mangal) for custom Lua scrapers.

Originally forked from [metafates/mangal-lua-libs](https://github.com/metafates/mangal-lua-libs) to fix the [headless](/headless) module and now sync'd with the original [vadv/gopher-lua-libs](https://github.com/vadv/gopher-lua-libs). Will try to keep it in sync with relevant commits.

## Index

* [base64](/base64) [encoding/base64](https://pkg.go.dev/encoding/base64) api
* [crypto](/crypto) calculate md5, sha256 hash for string
* [filepath](/filepath) path.filepath port
* [goos](/goos) os port
* [headless](/headless) headless chrome browser
* [html](/html) HTML processing
* [http](/http) http.client && http.server
* [humanize](/humanize) humanize [github.com/dustin/go-humanize](https://github.com/dustin/go-humanize) port
* [inspect](/inspect) pretty print [github.com/kikito/inspect.lua](https://github.com/kikito/inspect.lua)
* [ioutil](/ioutil) io/ioutil port
* [json](/json) json implementation
* [log](/log) log port
* [regexp](/regexp) regexp port
* [runtime](/runtime) runtime port
* [shellescape](/shellescape) shellescape <https://github.com/alessio/shellescape> port
* [stats](/stats) stats [https://github.com/montanaflynn/stats](https://github.com/montanaflynn/stats) port
* [storage](/storage) package for store persist data and share values between lua states
* [strings](/strings) strings port (utf supported)
* [template](/template) template engines
* [time](/time) time port
* [xmlpath](/xmlpath) [gopkg.in/xmlpath.v2](https://gopkg.in/xmlpath.v2) port
* [yaml](/yaml) [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) port

package plugin

import (
	"github.com/metafates/mangal-lua-libs/argparse"
	"github.com/metafates/mangal-lua-libs/aws/cloudwatch"
	"github.com/metafates/mangal-lua-libs/base64"
	"github.com/metafates/mangal-lua-libs/cert_util"
	"github.com/metafates/mangal-lua-libs/chef"
	"github.com/metafates/mangal-lua-libs/cmd"
	"github.com/metafates/mangal-lua-libs/crypto"
	"github.com/metafates/mangal-lua-libs/db"
	"github.com/metafates/mangal-lua-libs/filepath"
	"github.com/metafates/mangal-lua-libs/goos"
	"github.com/metafates/mangal-lua-libs/http"
	"github.com/metafates/mangal-lua-libs/humanize"
	"github.com/metafates/mangal-lua-libs/inspect"
	"github.com/metafates/mangal-lua-libs/ioutil"
	"github.com/metafates/mangal-lua-libs/json"
	"github.com/metafates/mangal-lua-libs/log"
	"github.com/metafates/mangal-lua-libs/pb"
	"github.com/metafates/mangal-lua-libs/pprof"
	prometheus "github.com/metafates/mangal-lua-libs/prometheus/client"
	"github.com/metafates/mangal-lua-libs/regexp"
	"github.com/metafates/mangal-lua-libs/runtime"
	"github.com/metafates/mangal-lua-libs/shellescape"
	"github.com/metafates/mangal-lua-libs/stats"
	"github.com/metafates/mangal-lua-libs/storage"
	"github.com/metafates/mangal-lua-libs/strings"
	"github.com/metafates/mangal-lua-libs/tac"
	"github.com/metafates/mangal-lua-libs/tcp"
	"github.com/metafates/mangal-lua-libs/telegram"
	"github.com/metafates/mangal-lua-libs/template"
	"github.com/metafates/mangal-lua-libs/time"
	"github.com/metafates/mangal-lua-libs/xmlpath"
	"github.com/metafates/mangal-lua-libs/yaml"
	"github.com/metafates/mangal-lua-libs/zabbix"
	lua "github.com/yuin/gopher-lua"
)

// PreloadAll preload all gopher lua packages - note it's needed here to prevent circular deps between plugin and libs
func PreloadAll(L *lua.LState) {
	Preload(L)

	argparse.Preload(L)
	base64.Preload(L)
	cert_util.Preload(L)
	chef.Preload(L)
	cloudwatch.Preload(L)
	cmd.Preload(L)
	crypto.Preload(L)
	db.Preload(L)
	filepath.Preload(L)
	goos.Preload(L)
	http.Preload(L)
	humanize.Preload(L)
	inspect.Preload(L)
	ioutil.Preload(L)
	json.Preload(L)
	log.Preload(L)
	pb.Preload(L)
	pprof.Preload(L)
	prometheus.Preload(L)
	regexp.Preload(L)
	runtime.Preload(L)
	shellescape.Preload(L)
	stats.Preload(L)
	storage.Preload(L)
	strings.Preload(L)
	tac.Preload(L)
	tcp.Preload(L)
	telegram.Preload(L)
	template.Preload(L)
	time.Preload(L)
	xmlpath.Preload(L)
	yaml.Preload(L)
	zabbix.Preload(L)
}

package plugin

import (
	"github.com/metafates/mangal-lua-libs/inspect"
	"github.com/metafates/mangal-lua-libs/tests"
	"github.com/stretchr/testify/assert"
	"testing"

	ioutil "github.com/metafates/mangal-lua-libs/ioutil"
	time "github.com/metafates/mangal-lua-libs/time"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		inspect.Preload,
		time.Preload,
		ioutil.Preload,
		Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}

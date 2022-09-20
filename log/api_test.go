package log

import (
	"github.com/metafates/mangal-lua-libs/tests"
	"github.com/stretchr/testify/assert"
	"testing"

	ioutil "github.com/metafates/mangal-lua-libs/ioutil"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		ioutil.Preload,
		Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}

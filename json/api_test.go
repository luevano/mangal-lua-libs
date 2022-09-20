package json

import (
	"github.com/metafates/mangal-lua-libs/strings"
	"github.com/metafates/mangal-lua-libs/tests"
	"github.com/stretchr/testify/assert"
	"testing"

	inspect "github.com/metafates/mangal-lua-libs/inspect"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		Preload,
		inspect.Preload,
		strings.Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}

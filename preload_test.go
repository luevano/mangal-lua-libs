package libs

import (
	"github.com/metafates/mangal-lua-libs/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreload(t *testing.T) {
	assert.NotZero(t, tests.RunLuaTestFile(t, Preload, "./preload.lua"))
}

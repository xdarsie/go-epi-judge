package smallest_subarray_covering_set_test

import (
	"os"
	"testing"

	"github.com/stefantds/go-epi-judge/config"
)

var testConfig *config.Config

func TestMain(m *testing.M) {
	var err error
	testConfig, err = config.Parse()
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}

package static

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v2/pkg/config/static"
	"gopkg.in/yaml.v2"
)

func TestConvert(t *testing.T) {
	testCases := []string{
		"./fixtures/sample01.toml",
		"./fixtures/sample02.toml",
	}

	for _, test := range testCases {
		t.Run(test, func(t *testing.T) {
			dir, err := ioutil.TempDir("", "traefik-migration-tool-static")
			require.NoError(t, err)

			t.Cleanup(func() { _ = os.RemoveAll(dir) })

			err = Convert(test, dir)
			require.NoError(t, err)

			cfgToml := static.Configuration{}
			_, err = toml.DecodeFile(filepath.Join(dir, "new-traefik.toml"), &cfgToml)
			require.NoError(t, err)

			cfgYaml := static.Configuration{}
			ymlFile, err := os.Open(filepath.Join(dir, "new-traefik.yml"))
			require.NoError(t, err)

			err = yaml.NewDecoder(ymlFile).Decode(&cfgYaml)
			require.NoError(t, err)

			assert.Equal(t, &cfgToml, &cfgYaml)
		})
	}
}

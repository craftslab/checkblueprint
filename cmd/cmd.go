package cmd

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v3"

	"github.com/craftslab/checkblueprint/config"
)

var (
	app        = kingpin.New("checkblueprint", "Check Blueprint").Version(config.Version + "-build-" + config.Build)
	configFile = app.Flag("config-file", "Config file (.yml)").Required().String()
	outputFile = app.Flag("output-file", "Output file (.json|.txt|.xlsx)").Default().String()
)

func Run() error {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	_, err := initConfig(*configFile)
	if err != nil {
		return errors.Wrap(err, "failed to init config")
	}

	return nil
}

func initConfig(name string) (*config.Config, error) {
	c := config.New()
	if c == nil {
		return &config.Config{}, errors.New("failed to new")
	}

	fi, err := os.Open(name)
	if err != nil {
		return c, errors.Wrap(err, "failed to open")
	}

	defer func() {
		_ = fi.Close()
	}()

	buf, err := io.ReadAll(fi)
	if err != nil {
		return c, errors.Wrap(err, "failed to readall")
	}

	if err := yaml.Unmarshal(buf, c); err != nil {
		return c, errors.Wrap(err, "failed to unmarshal")
	}

	return c, nil
}

package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type SystemConfig struct {
	ProjectPath string
}

func (c *SystemConfig) readConf(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return fmt.Errorf("in file %q: %v", filename, err)
	}
	return nil
}

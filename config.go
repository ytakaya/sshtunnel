package sshtunnel

import (
	"os"
	"reflect"

	"github.com/go-yaml/yaml"
)

type YAMLConfig struct {
	KeyFiles []KeyFile `yaml:"key_files"`
	Gateways []struct {
		Server       string   `yaml:"server"`
		ProxyCommand string   `yaml:"proxy_command"`
		Tunnels      []string `yaml:"tunnels"`
	} `yaml:"gateways"`
}

func (c *YAMLConfig) Equals(r *YAMLConfig) bool {
	return reflect.DeepEqual(c, r)
}

type KeyFile struct {
	Path       string
	Passphrase string
}

func (f *KeyFile) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw interface{}
	unmarshal(&raw)
	switch raw := raw.(type) {
	case string:
		*f = KeyFile{
			Path: raw,
		}
	case map[interface{}]interface{}:
		path, _ := raw["path"].(string)
		passphrase, _ := raw["passphrase"].(string)
		*f = KeyFile{
			Path:       path,
			Passphrase: passphrase,
		}
	}
	return nil
}

func LoadConfigFile(file *os.File) (*YAMLConfig, error) {
	var config YAMLConfig
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

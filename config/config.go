package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func NewConfig() *Config {
	var configFile string
	flag.StringVar(&configFile, "config", "", "")
	flag.Parse()

	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType("yaml")
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath("./conf/")
		viper.SetConfigName("config")
	}

	viper.SetDefault("ink.admin_url_path", "admin")

	conf := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}

	if conf.Ink.WorkDir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			panic(errors.Wrap(err, "init config: get current dir"))
		}
		conf.Ink.WorkDir, _ = filepath.Abs(pwd)
	} else {
		workDir, err := filepath.Abs(conf.Ink.WorkDir)
		if err != nil {
			panic(err)
		}
		conf.Ink.WorkDir = workDir
	}
	normalizeDir := func(path *string, subDir string) {
		if *path == "" {
			*path = filepath.Join(conf.Ink.WorkDir, subDir)
		} else {
			temp, err := filepath.Abs(*path)
			if err != nil {
				panic(err)
			}
			*path = temp
		}
	}
	normalizeDir(&conf.Ink.LogDir, "log")

	initDirectory(conf)
	mode = conf.Ink.Mode
	logMode = conf.Ink.LogMode
	return conf
}

func initDirectory(conf *Config) {
	mkdirFunc := func(dir string, err error) error {
		if err == nil {
			if _, err = os.Stat(dir); os.IsNotExist(err) {
				err = os.MkdirAll(dir, os.ModePerm)
			}
		}
		return err
	}
	err := mkdirFunc(conf.Ink.LogDir, nil)
	if err != nil {
		panic(fmt.Errorf("initDirectory err=%w", err))
	}
}

var (
	mode    string
	logMode LogMode
)

func IsDev() bool {
	return mode == "development"
}

func LogToConsole() bool {
	switch logMode {
	case Console:
		return true
	case File:
		return false
	default:
		return IsDev()
	}
}

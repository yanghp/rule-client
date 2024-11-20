package config

import (
	"flag"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"time"

	"github.com/DoNewsCode/core/contract"
	"github.com/knadh/koanf"
)

type ConfigReader interface {
	Cut(string) ConfigReader
	contract.ConfigAccessor
}

type Env string

func (e Env) IsDevelopment() bool {
	return e == "dev"
}

func (e Env) IsStaging() bool {
	return e == "local"
}

func (e Env) IsProduction() bool {
	return e == "prod"
}

type AppName string

func (a AppName) String() string {
	return string(a)
}

func (e Env) IsLocal() bool {
	return e == "local"
}

func (e Env) IsTesting() bool {
	return e == "testing"
}

func (e Env) IsDev() bool {
	return e == "dev"
}

func (e Env) IsProd() bool {
	return e == "prod"
}

func (e Env) String() string {
	return string(e)
}

func ProvideEnv(conf ConfigReader) Env {
	return Env(conf.String("env"))
}

func ProvideAppName(conf ConfigReader) AppName {
	return AppName(conf.String("name"))
}

var _ contract.ConfigAccessor = (*KoanfAdapter)(nil)

type KoanfAdapter struct {
	k *koanf.Koanf
}

func (k *KoanfAdapter) Cut(s string) ConfigReader {
	cut := k.k.Cut("global")
	cut.Merge(k.k.Cut(s))
	return NewKoanfAdapter(cut)
}

func NewKoanfAdapter(k *koanf.Koanf) *KoanfAdapter {
	return &KoanfAdapter{k}
}

var cgf = flag.String("config", "config.yaml", "配置文件路径")

func NewKoanfig() ConfigReader {
	flag.Parse()
	cfgFile := *cgf
	k := koanf.New(".")
	err := k.Load(file.Provider(cfgFile), yaml.Parser())
	if err != nil {
		panic(err)
	}
	return NewKoanfAdapter(k).Cut("")
}

func (k *KoanfAdapter) String(s string) string {
	return k.k.String(s)
}

func (k *KoanfAdapter) Int(s string) int {
	return k.k.Int(s)
}

func (k *KoanfAdapter) Strings(s string) []string {
	return k.k.Strings(s)
}

func (k *KoanfAdapter) Bool(s string) bool {
	return k.k.Bool(s)
}

func (k *KoanfAdapter) Get(s string) interface{} {
	return k.k.Get(s)
}

func (k *KoanfAdapter) Float64(s string) float64 {
	return k.k.Float64(s)
}

func (k *KoanfAdapter) Duration(s string) time.Duration {
	return k.k.Duration(s)
}

func (k *KoanfAdapter) Unmarshal(path string, o interface{}) error {
	return k.k.Unmarshal(path, o)
}

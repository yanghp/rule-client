package client

import (
	"testing"

	"github.com/knadh/koanf"
)

func TestMuteMap(t *testing.T) {
	k := koanf.New("")
	k.Load(Provider(map[string]interface{}{"foo": "bar"}), nil)
	if k.Get("foo") != "bar" {
		t.Error("Expected 'bar', got ", k.Get("foo"))
	}

	k = koanf.New(".")
	k.Load(Provider(map[string]interface{}{"foo": map[string]interface{}{"bar": "baz"}}), nil)
	if k.Get("foo.bar") != "baz" {
		t.Error("Expected 'baz', got ", k.Get("foo.bar"))
	}
}

package test

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/yanghp/rule-client/client"
	"github.com/yanghp/rule-client/dto"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"testing"
)

func TestRule(t *testing.T) {
	var s = []string{"etcd-1:2379", "etcd-2:2379", "etcd-3:2379"}
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: s,
	})
	if err != nil {
		t.Fatal(err)
	}

	rule, err := client.NewRuleEngine(client.WithClient(etcdClient), client.WithRulePrefix("newnovel"), client.WithLogger(log.NewJSONLogger(os.Stdout)))
	if err != nil {
		t.Fatal(err)
	}

	go rule.Watch(context.Background())

	conf, err := rule.Of("newnovel-server-dev").Payload(&dto.Payload{
		AppName: "yanghp",
	})
	if err != nil {
		t.Fatal(err)
	}
	var rest map[string]interface{}
	err = conf.Unmarshal("", &rest)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rest)

}

package repository

import (
	"fmt"
	"testing"
)

func TestRepositoryConfig(t *testing.T) {
	var a uint8
	var b int = 257
	a = uint8(b)

	fmt.Println(a)

	//k := koanf.New(".")
	//f := file.Provider("../../config/kitty.yaml")
	//err := k.Load(f, yaml.Parser())
	//if err != nil {
	//	log.Fatalf("error loading config: %v", err)
	//}
	//c := config.NewKoanfAdapter(k)
	//
	////初始化
	//env := c.String("global.env")
	//logger := kitty_log.NewLogger(config.Env(env))
	//etcdAddrs := c.String("app.etcd.addrs")
	//_ = logger.Log("etcd.addrs", etcdAddrs)
	//
	////etcd 实例
	//ctx, cancel := context.WithCancel(context.Background())
	//client, err := clientv3.New(clientv3.Config{
	//	Endpoints: c.Strings("app.etcd.addrs"),
	//	Context:   ctx,
	//})
	//logger.Log("etcd operate ", "start")
	//if err != nil {
	//	logger.Log("intt etcd error", err.Error())
	//}
	//defer func() {
	//	err := client.Close()
	//	cancel()
	//	logger.Log("etcd err", err)
	//}()
	//cc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//value, err := client.Get(cc, "test_key1")
	//cancel()
	//if err != nil {
	//	logger.Log("get key error , error", err.Error())
	//	return
	//}
	//logger.Log("before , test_key1", value)
	//cc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//cancel()
	//client.Put(cc, "test_key1", "abcdefgh")
	//cc2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//resp, _ := client.Get(cc2, "test_key1")
	//for _, ev := range resp.Kvs {
	//	logger.Log(ev.Key, ev.Value)
	//}
}

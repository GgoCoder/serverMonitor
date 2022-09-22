package serviceregister

import (
	"context"
	"encoding/json"
	"fmt"
	"serverMonitor/internal/serviceRegister/pkg/config"
	"serverMonitor/pkg/etcd"
	"serverMonitor/pkg/typed"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type HandleMicroService interface{}

func Register(serviceName, addr string, port uint16) {
	newService := &typed.MicroService{
		ServiceName: serviceName,
		Addr:        addr,
		Port:        port,
	}
	etcdClient := etcd.GetClient(config.ServiceRegisterConfig.Etcd.EndPoints, 5*time.Second)
	defer etcdClient.Close()
	data, err := json.Marshal(newService)
	if err != nil {
		fmt.Printf("failed marshal service, err: %+v\n", err)
	}
	//TODO etcd 连接不上卡住不报错 不超时 status
	_, err = etcdClient.Put(context.Background(), "/"+serviceName, string(data), clientv3.WithPrevKV())
	if err != nil {
		fmt.Printf("failed to register micro service {%s}, err: %+v\n", serviceName, err)
	}
	fmt.Println("register successfully!")
}

func DeRegister(serviceName string) {
	etcdClient := etcd.GetClient(config.ServiceRegisterConfig.Etcd.EndPoints, 5*time.Second)
	defer etcdClient.Close()
	_, err := etcdClient.Delete(context.Background(), serviceName)
	if err != nil {
		fmt.Printf("failed to delete register service {%s}, err: %+v\n", serviceName, err)
	}

}

//考虑做成接口，每个服务都实现这个接口
func DiscoverServices(serviceName string, service chan *typed.MicroService) {

	etcdClient := etcd.GetClient(config.ServiceRegisterConfig.Etcd.EndPoints, 5*time.Second)
	serviceChan := etcdClient.Watch(context.Background(), serviceName)
	for {
		select {
		case services := <-serviceChan:
			for _, s := range services.Events {
				switch s.Type {
				case mvccpb.PUT:
					fmt.Printf("service add/update, key: %s, service:%+v\n", string(s.Kv.Key), string(s.Kv.Value))
				case mvccpb.DELETE:
					fmt.Printf("service delete, key: %s, service:%+v\n", string(s.Kv.Key), string(s.Kv.Value))
				}
			}
		}
	}
}

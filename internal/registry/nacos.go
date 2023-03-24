package registry

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"helloworld/internal/conf"
)


func NewNacosRegistry(c *conf.Server) vo.NacosClientParam {
	serverConfigs := []constant.ServerConfig{
		constant.ServerConfig{
			IpAddr:      "172.16.3.85",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:     "http",
		},
	}

	clientConfig := &constant.ClientConfig{
		NamespaceId:         "dev",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		//LogDir:              "/tmp/nacos/log",
		//CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "error",
	}

	return vo.NacosClientParam{
		ClientConfig:  clientConfig,
		ServerConfigs: serverConfigs,
	}
}


// NewDiscovery nacos服务发现注入
func NewDiscovery(param vo.NacosClientParam) registry.Discovery {
	client, _ := clients.NewNamingClient(param)


	return nacos.New(client)
}

// NewRegistrar 服务注册业务注入
func NewRegistrar(param vo.NacosClientParam) registry.Registrar {
	fmt.Printf("param = %+v", param)
	iClient, err := clients.NewNamingClient(param)
	if err != nil {
		fmt.Println("err====", err)
	}

	return nacos.New(iClient)
}
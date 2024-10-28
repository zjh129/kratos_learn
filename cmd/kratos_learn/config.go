package main

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/config/apollo/v2"
	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	knacos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	consul_api "github.com/hashicorp/consul/api"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"kratos_learn/internal/conf"
	"strings"
)

// initConfig init the config
func initConfig() config.Config {
	dirverC := config.New(
		config.WithSource(
			file.NewSource(fmt.Sprintf("%s/config_driver.yaml", flagconf)),
		),
	)
	defer dirverC.Close()

	if err := dirverC.Load(); err != nil {
		panic(err)
	}
	driverConfig := conf.ConfDriver{}
	if err := dirverC.Scan(&driverConfig); err != nil {
		panic(err)
	}
	switch driverConfig.Type {
	case "file":
		// 读取本地文件
		return config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
	case "apollo":
		if driverConfig.Apollo == nil {
			panic("apollo config is nil")
		}
		return config.New(
			config.WithSource(
				apollo.NewSource(
					apollo.WithAppID(driverConfig.Apollo.AppId),
					apollo.WithCluster(driverConfig.Apollo.Cluster),
					apollo.WithEndpoint(driverConfig.Apollo.Endpoint),
					apollo.WithNamespace(driverConfig.Apollo.NamespaceName),
					apollo.WithEnableBackup(),
					apollo.WithSecret(driverConfig.Apollo.Secret),
				),
			),
		)
	case "consul":
		if driverConfig.Consul == nil {
			panic("consul config is nil")
		}
		// 读取consul配置
		consulClient, err := consul_api.NewClient(&consul_api.Config{
			Address:    driverConfig.Consul.Address,
			Scheme:     driverConfig.Consul.Scheme,
			PathPrefix: driverConfig.Consul.PathPrefix,
			Token:      driverConfig.Consul.Token,
		})
		if err != nil {
			panic(err)
		}
		cs, err := consul.New(consulClient, consul.WithPath(driverConfig.Consul.Path))
		if err != nil {
			panic(err)
		}
		return config.New(config.WithSource(cs))
	case "nacos":
		if driverConfig.Nacos == nil {
			panic("nacos config is nil")
		}
		sc := []constant.ServerConfig{}
		for _, addr := range driverConfig.Nacos.Addrs {
			addrArr := strings.Split(addr, ":")
			if len(addrArr) == 1 {
				sc = append(sc, *constant.NewServerConfig(addrArr[0], 8848))
			} else {
				var port uint64
				fmt.Sscanf(addrArr[1], "%d", &port)
				sc = append(sc, *constant.NewServerConfig(addrArr[0], port))
			}
		}

		cc := &constant.ClientConfig{
			NamespaceId:         driverConfig.Nacos.Namespace,
			Username:            driverConfig.Nacos.Username,
			Password:            driverConfig.Nacos.Password,
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogDir:              "/tmp/nacos/log",
			CacheDir:            "/tmp/nacos/cache",
			LogLevel:            "debug",
		}

		// a more graceful way to create naming client
		client, err := clients.NewConfigClient(
			vo.NacosClientParam{
				ClientConfig:  cc,
				ServerConfigs: sc,
			},
		)
		if err != nil {
			panic(err)
		}
		return config.New(
			config.WithSource(
				knacos.NewConfigSource(
					client,
					knacos.WithGroup(driverConfig.Nacos.Group),
					knacos.WithDataID(driverConfig.Nacos.DataId),
				),
			),
		)
	default:
		panic("unknown config driver")
	}
}

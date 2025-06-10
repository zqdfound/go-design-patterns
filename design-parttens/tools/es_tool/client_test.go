package estool

import (
	"fmt"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/spf13/viper"
)

func TestLink(t *testing.T) {

	// 初始化 Viper
	viper.SetConfigName("es_config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			viper.GetString("address"),
		},
	}

	// 创建客户端连接
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("elasticsearch.NewTypedClient failed, err:%v\n", err)
		return
	}
	fmt.Print("======链接信息======")
	fmt.Print(client)
}

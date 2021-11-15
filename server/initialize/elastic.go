package initialize

import (
	"github.com/olivere/elastic/v7"
	"mall.com/global"
)

func Elastic() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200/"))
	if err != nil {
		panic(err)
	}
	global.Es = client
}

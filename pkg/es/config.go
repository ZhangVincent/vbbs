package es

import (
	"errors"
	"sync"
	"vbbs/pkg/config"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"vbbs/pkg/simple/common/strs"
)

var (
	client      *elastic.Client
	index       string
	once        sync.Once
	errNoConfig = errors.New("es config not found. ")
)

func initClient() *elastic.Client {
	once.Do(func() {
		var err error
		if !strs.IsAnyBlank(config.Instance.Es.Url, config.Instance.Es.Index) {
			index = config.Instance.Es.Index
			client, err = elastic.NewClient(
				elastic.SetURL(config.Instance.Es.Url),
				elastic.SetHealthcheck(false),
				elastic.SetSniff(false),
			)
		} else {
			err = errNoConfig
		}
		if err != nil {
			logrus.Error(err)
		}
	})
	return client
}

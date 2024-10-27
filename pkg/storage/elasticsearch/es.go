package es

import (
	"crypto/tls"
	logger "ddCode-server/pkg/zlogs"
	"net/http"
	"time"

	"github.com/olivere/elastic/v7"
)

type Config struct {
	User  string   // 用户名
	Pwd   string   // 密码
	Addrs []string // 地址串
}

type OptionFunc func(*Elastic)

func WithLogger(logger logger.Logger) func(*Elastic) {
	return func(e *Elastic) {
		e.loggeer = logger
	}
}

// Elastic 封装的ES操作器.
type Elastic struct {
	es      *elastic.Client
	loggeer logger.Logger
}

type Filters struct {
	T        interface{}
	Qu       string
	Wildcard string
}

func NewElastic(conf *Config, ops ...OptionFunc) (*Elastic, error) {
	httpCli := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // nolint
			MaxIdleConnsPerHost: 4096,
		},
		Timeout: 60 * time.Second,
	}
	esOptions := []elastic.ClientOptionFunc{
		elastic.SetURL(conf.Addrs...),
		elastic.SetSniff(false),
		elastic.SetHttpClient(httpCli),
		elastic.SetHealthcheck(false),
	}
	if conf.Pwd != "" && conf.User != "" {
		esOptions = append(esOptions, elastic.SetBasicAuth(conf.User, conf.Pwd))
	}
	es, err := elastic.NewClient(esOptions...)
	if err != nil {
		return nil, err
	}

	ecli := &Elastic{
		es: es,
	}

	for _, op := range ops {
		op(ecli)
	}
	return ecli, nil
}

func (e *Elastic) GetClient() *elastic.Client {
	return e.es
}

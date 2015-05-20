package app

import (
	"net/http"

	"github.com/tukdesk/tuksearch/app/api"

	_ "github.com/tukdesk/tuksearch/context"

	"github.com/hprose/hprose-go/hprose"
	"github.com/tukdesk/ledisdbcli"
	"github.com/tukdesk/tuksearch/bleve"

	_ "github.com/tukdesk/bleve-ledisdb-storage"
	_ "github.com/wangbin/jiebago/tokenizers"
)

type App struct {
	cfg     Config
	service *hprose.HttpService
}

func New(cfg Config) (*App, error) {
	// index mapping
	defaultIndexMapping := bleve.NewIndexMapping()

	err := defaultIndexMapping.AddCustomTokenizer("jieba",
		map[string]interface{}{
			"file": cfg.IndexJieba.DictPath,
			"type": "jieba",
		})

	if err != nil {
		return nil, err
	}

	err = defaultIndexMapping.AddCustomAnalyzer("jieba",
		map[string]interface{}{
			"type":      "custom",
			"tokenizer": "jieba",
			"token_filters": []string{
				"possessive_en",
				"to_lower",
				"stop_en",
			},
		})

	if err != nil {
		return nil, err
	}

	defaultIndexMapping.DefaultAnalyzer = "jieba"

	// index store
	clientConfig := ledisdbcli.Config{
		Addr:    cfg.IndexStore.Addr,
		DBIndex: cfg.IndexStore.DBIndex,
	}
	client, err := ledisdbcli.New(clientConfig)
	if err != nil {
		return nil, err
	}

	// init index

	// apis
	global := api.Global{}
	global.DefaultIndexMapping = defaultIndexMapping
	global.IndexStoreClient = client

	apis := api.New(global)
	service := hprose.NewHttpService()
	service.DebugEnabled = true
	service.AddMethods(apis)
	service.ServiceEvent = myServiceEvent{}

	return &App{
		cfg:     cfg,
		service: service,
	}, nil
}

func (this *App) Run() error {
	return http.ListenAndServe(this.cfg.Addr, this.service)
}

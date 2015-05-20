package app

type Config struct {
	Addr       string           `json:"addr"`
	DataDir    string           `json:"data_dir"`
	IndexStore IndexStoreConfig `json:"index_store"`
	IndexJieba IndexJiebaConfig `json:"index_jieba"`
}

type IndexStoreConfig struct {
	Addr    string `json:"addr"`
	DBIndex int    `json:"db_index"`
}

type IndexJiebaConfig struct {
	DictPath string `json:"dict_path"`
}

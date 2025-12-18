package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}
type KafkaConf struct {
	Address string `ini:"address"`
	MaxSize int `ini:"chan_max_size"`
}
type EtcdConf struct {
	Address string `ini:"address"`
	TimeOut int    `ini:"timeout"`
	Key     string `ini:"collect_log_key"`
}
type TaillogConf struct {
	FileName string `ini:"filename"`
}

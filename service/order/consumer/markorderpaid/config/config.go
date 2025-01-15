package config

type Config struct {
	Mysql struct {
		DataSource string
	}
	Mq struct {
		NameServers []string
	}
}

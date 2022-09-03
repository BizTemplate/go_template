package core

type BaseConfig struct {
	dbConfig    DbConfig
	redisConfig RedisConfig
	adminConfig AdminConfig
}

type DbConfig struct {
	host       string `json:"string" :"host"`
	username   string `json:"string" :"username"`
	password   string `json:"string" :"password"`
	poolSize   int    `json:"string":"poolSize"`
	mainSchema string `json:"string":"mainSchema"`
}

type RedisConfig struct {
	host string
}

type AdminConfig struct {
	adminPort int `json:"int":'adminPort'`
}

type JobConfig struct {
}

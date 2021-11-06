package config

import "orderserve/dal"

func InitRrs() {
	mysqlConf := Conf.Mysql
	redisConf := Conf.Redis
	dal.SqlEngine(mysqlConf.MysqlAddr, mysqlConf.MysqlDriver, mysqlConf.MysqlShow)
	dal.InitRedisStore(redisConf.RedisAddr, redisConf.RedisPwd)
}

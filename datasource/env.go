package datasource

import "os"

type EnvValues struct {
	RedisAddr string
	RedisPass string
	DbAddr    string
}

var localEnvCache *EnvValues

func GetEnvs() EnvValues {
	if localEnvCache != nil {
		return *localEnvCache
	}
	target := EnvValues{
		RedisAddr: getEnv("CHAT_REDIS_ADDR"),
		RedisPass: getEnv("CHAT_REDIS_PASS"),
		DbAddr:    getEnv("CHAT_RDS_ADDR"),
	}
	localEnvCache = &target
	return target
}

func getEnv(target string) string {
	result, hasSet := os.LookupEnv(target)
	if !hasSet {
		panic(target + " has not set")
	}
	return result
}

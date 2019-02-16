package datasource

import "os"

type EnvValues struct {
	redisAddr string
	redisPass string
	dbAddr    string
}

var localEnvCache *EnvValues

func GetEnvs() EnvValues {
	if localEnvCache != nil {
		return *localEnvCache
	}
	target := EnvValues{
		redisAddr: getEnv("CHAT_REDIS_ADDR"),
		redisPass: getSecretEnv("CHAT_REDIS_PASS"),
		dbAddr:    getSecretEnv("CHAT_RDS_ADDR"),
	}
	localEnvCache = &target
	return target
}

func getSecretEnv(target string) string {
	result, hasSet := os.LookupEnv(target)
	if !hasSet {
		panic(target + " has not set")
	}

	println(target + " => xxxxxxxx")

	return result
}

func getEnv(target string) string {
	result, hasSet := os.LookupEnv(target)
	if !hasSet {
		panic(target + " has not set")
	}

	println(target + " => " + result)

	return result
}

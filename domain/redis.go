package domain

type ClientRedis interface {
	GetRedisValue(string) ([]byte, error)
	SetRedisValue(string, interface{}, int)
	DelRedisValue(string) error
}

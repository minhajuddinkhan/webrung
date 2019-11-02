package auth

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung"
)

//Redis Redis
type Redis interface {
	Ping() error
}

type authRedis struct {
	url string
}

//NewAuthRedis returns connection of the authentication redis client
func NewAuthRedis(conf *webrung.Conf) (Redis, error) {
	if conf == nil {
		return nil, fmt.Errorf("redis err: nil configuration provided")
	}
	return &authRedis{url: conf.AuthRedis.RedisURL}, nil
}

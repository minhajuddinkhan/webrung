package socket

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung"
)

//Redis Redis
type Redis interface {
	Ping() error
}

type socketRedis struct {
	url string
}

//NewSocketRedis returns connection of the socket redis client
func NewSocketRedis(conf *webrung.Conf) (Redis, error) {
	if conf == nil {
		return nil, fmt.Errorf("redis err: nil configuration provided")
	}
	return &socketRedis{url: conf.SocketRedis.RedisURL}, nil
}

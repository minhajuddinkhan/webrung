package webrung_test

import (
	"os"
	"testing"

	"github.com/minhajuddinkhan/webrung/cache/socket"

	"github.com/minhajuddinkhan/webrung"
	"github.com/stretchr/testify/assert"
)

func TestSocketRedis_PingShouldNotError(t *testing.T) {

	url := os.Getenv("SOCKET_REDIS_URL")
	r, err := socket.NewSocketRedis(&webrung.Conf{
		SocketRedis: webrung.Redis{
			RedisURL: url,
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, r)
	err = r.Ping()
	assert.Nil(t, err)
}

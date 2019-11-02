package webrung_test

import (
	"os"
	"testing"

	"github.com/minhajuddinkhan/webrung"
	"github.com/minhajuddinkhan/webrung/cache/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthRedis_PingShouldNotError(t *testing.T) {

	url := os.Getenv("AUTH_REDIS_URL")
	r, err := auth.NewAuthRedis(&webrung.Conf{
		AuthRedis: webrung.AuthRedis{
			RedisURL: url,
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, r)
	err = r.Ping()
	assert.NotNil(t, err)
}

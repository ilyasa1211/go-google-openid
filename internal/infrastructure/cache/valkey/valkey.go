package valkey

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	config "github.com/ilyasa1211/go-google-openid/internal/config/cache"
	"github.com/valkey-io/valkey-go"
)

var conf *config.ValkeyConf = config.NewValkeyConf()

func NewValkeyClient() *valkey.Client {
	address := strings.Join([]string{conf.Host, conf.Port}, ":")
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{address},
	})

	if err != nil {
		log.Fatalln("Failed to connect valkey: ", err)
	}

	return &client
}

type CacheImpl struct {
	Client *valkey.Client
}

func NewCacheImpl(c *valkey.Client) *CacheImpl {
	return &CacheImpl{c}
}

func (a *CacheImpl) Set(k string, v string) {
	cmd := (*a.Client).B().Set().Key(k).Value(v).Ex(time.Minute * 5).Build()

	ctx := context.Background()

	err := (*a.Client).Do(ctx, cmd).Error()

	if err != nil {
		log.Fatalln("Failed to set key")
	}
}

func (a *CacheImpl) Get(k string) string {
	cmd := (*a.Client).B().Get().Key(k).Build()

	ctx := context.Background()

	v, err := (*a.Client).Do(ctx, cmd).ToString()

	if err != nil {
		fmt.Println("error get cache: ", err)
		return ""
	}

	return v
}

func (a *CacheImpl) Del(k string) {
	cmd := (*a.Client).B().Del().Key(k).Build()

	ctx := context.Background()

	if err := (*a.Client).Do(ctx, cmd).Error(); err != nil {
		log.Fatalln("error get cache: ", err)
	}
}

package instance

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/zakirkun/ehe/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAppContext struct {
	Cfg config.Config
	Ctx context.Context
}
type IAppLayer interface {
	MongoSetup() *mongo.Client
	RedisSetup() *redis.Client
	WebServerSetup() http.Server
}

func NewInstance(cfg config.Config, ctx context.Context) IAppContext {
	return IAppContext{Cfg: cfg, Ctx: ctx}
}

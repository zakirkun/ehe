package services

import (
	"context"

	"github.com/zakirkun/ehe/app/domain/contract"
	"go.mongodb.org/mongo-driver/mongo"
)

type Opts struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthServices(opt Opts) contract.IAuthService {
	return &authServicesContext{collection: opt.collection, ctx: opt.ctx}
}

func NewUsersServices(opt Opts) contract.IUserService {
	return &usersServiceContext{collection: opt.collection, ctx: opt.ctx}
}

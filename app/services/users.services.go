package services

import (
	"context"
	"strings"

	"github.com/zakirkun/ehe/app/domain/contract"
	"github.com/zakirkun/ehe/app/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type usersServiceContext struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUsersServices(collection *mongo.Collection, ctx context.Context) contract.IUserService {
	return &usersServiceContext{collection: collection, ctx: ctx}
}

func (r *usersServiceContext) FindUserById(id string) (*models.DBResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	var user *models.DBResponse

	query := bson.M{"_id": oid}
	err := r.collection.FindOne(r.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

func (r *usersServiceContext) FindUserByEmail(email string) (*models.DBResponse, error) {
	var user *models.DBResponse

	query := bson.M{"email": strings.ToLower(email)}
	err := r.collection.FindOne(r.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

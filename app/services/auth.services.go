package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/zakirkun/ehe/app/domain/models"
	"github.com/zakirkun/ehe/app/domain/types"
	"github.com/zakirkun/ehe/app/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type authServicesContext struct {
	collection *mongo.Collection
	ctx        context.Context
}

func (s *authServicesContext) SignUpUser(user *types.SignUpInput) (*models.DBResponse, error) {

	// input users
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = true
	user.Role = "user"

	hashedPassword, _ := helper.HashPassword(user.Password)
	user.Password = hashedPassword
	res, err := s.collection.InsertOne(s.ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}

	// Create a unique index for the email field
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := s.collection.Indexes().CreateOne(s.ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = s.collection.FindOne(s.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil

}

func (s *authServicesContext) SignInUser(*types.SignInInput) (*models.DBResponse, error) {
	return nil, nil
}

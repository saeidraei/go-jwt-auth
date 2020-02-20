package userRW

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"github.com/saeidraei/go-jwt-auth/domain"
	"github.com/saeidraei/go-jwt-auth/uc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type rw struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func New() uc.UserRW {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s",
		viper.GetString("mongo.user"),
		viper.GetString("mongo.password"),
		viper.GetString("mongo.host"),
		viper.GetString("mongo.port"))))
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	collection := client.Database("clean").Collection("users")
	return rw{
		client:     client,
		collection: collection,
	}
}

func (rw rw) Create(username, email, password string) (*domain.User, error) {
	if _, err := rw.GetByEmail(email); err == nil {
		return nil, uc.ErrAlreadyInUse
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := rw.collection.InsertOne(ctx, bson.M{"Name": username, "Email": email, "Password": password})
	if err != nil {
		panic(err)
	}
	return rw.GetByEmail(email)
}

func (rw rw) GetByName(userName string) (*domain.User, error) {
	var user domain.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res := rw.collection.FindOne(ctx, bson.M{"Name": userName})
	err := res.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (rw rw) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res := rw.collection.FindOne(ctx, bson.M{"Email": email})
	err := res.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (rw rw) GetByEmailAndPassword(email, password string) (*domain.User, error) {
	var user domain.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res := rw.collection.FindOne(ctx, bson.M{"Email": email, "Password": password})
	err := res.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (rw rw) Save(user domain.User) error {
	//if user, _ := rw.GetByName(user.Name); user == nil {
	//	return uc.ErrNotFound
	//}
	//
	//user.UpdatedAt = time.Now()
	//rw.store.Store(user.Name, user)

	return nil
}

package userRW

import (
	"context"
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
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:testpassword@mongo:27017"))
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
	if _, err := rw.GetByName(username); err == nil {
		return nil, uc.ErrAlreadyInUse
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := rw.collection.InsertOne(ctx, bson.M{"Name": username, "Email": email, "password": password})
	if err != nil {
		panic(err)
	}
	return rw.GetByName(username)
}

func (rw rw) GetByName(userName string) (*domain.User, error) {
	//value, ok := rw.store.Load(userName)
	//if !ok {
	//	return nil, uc.ErrNotFound
	//}
	//
	//user, ok := value.(domain.User)
	//if !ok {
	//	return nil, errors.New("not a user stored at key")
	//}
	var user domain.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res := rw.collection.FindOne(ctx, bson.M{"Name": userName})
	err := res.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (rw rw) GetByEmailAndPassword(email, password string) (*domain.User, error) {
	var user domain.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res := rw.collection.FindOne(ctx, bson.M{"Email": email, "password": password})
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

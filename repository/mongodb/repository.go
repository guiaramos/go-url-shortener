package mongodb

import (
	"context"
	"time"

	"github.com/guiaramos/go-url-shortener/shortener"
	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

// NewMongoRepository create a new MongoDB redirect repository
func NewMongoRepository(url, db string, timeout int) (shortener.RedirectRepository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(timeout) * time.Second,
		database: db,
	}

	client, err := newMongoClient(url, timeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepository")
	}

	repo.client = client
	return repo, nil
}

func newMongoClient(url string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, err
}

func (r *mongoRepository) Find(code string) (*shortener.Redirect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	coll := r.client.Database(r.database).Collection("redirects")

	redirect := &shortener.Redirect{}
	filter := bson.M{"code": code}

	err := coll.FindOne(ctx, filter).Decode(&redirect)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(shortener.ErrRedirectNotFound, "repository.Redirect.Find")
		}
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}

	return redirect, nil
}

func (r *mongoRepository) Store(redirect *shortener.Redirect) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	coll := r.client.Database(r.database).Collection("redirects")

	_, err := coll.InsertOne(
		ctx,
		bson.M{
			"code":       redirect.Code,
			"url":        redirect.URL,
			"created_at": redirect.CreatedAt,
		},
	)

	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}

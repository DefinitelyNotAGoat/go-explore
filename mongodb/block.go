package mongodb

import (
	"context"
	"fmt"
	"time"

	goTezos "github.com/DefinitelyNotAGoat/go-tezos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo structure to handle db related stuff
type Mongo struct {
	client      *mongo.Client
	db          *mongo.Database
	collections map[string]*mongo.Collection
}

// NewMongoService dials and returns a new mongo object with a client
func NewMongoService(dial string, db string) (*Mongo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dial))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	database := client.Database(db)

	mg := Mongo{client: client, db: database}

	collections := make(map[string]*mongo.Collection)
	collections["blocks"] = mg.db.Collection("blocks")

	mg.collections = collections

	return &mg, nil
}

func (m *Mongo) StateCheck(errch chan error) {
	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{})
	go func() {
		select {
		case <-ticker.C:
			ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
			err := m.client.Ping(ctx, readpref.Primary())
			if err != nil {
				errch <- fmt.Errorf("could not ping mongo: %v", err)
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}()
}

// InsertBlock inserts a block into a mongo collection
func (m *Mongo) InsertBlock(block goTezos.Block) error {

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	_, err := m.collections["blocks"].InsertOne(ctx, block)
	if err != nil {
		return err
	}
	return nil
}

package mg

import (
	"FractalCompression/internal/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type MG struct {
	client   *mongo.Client
	database string
}

func NewMG(cnf config.CompressionConfig) (*MG, error) {
	uri := fmt.Sprintf("mongodb://%s:%d", cnf.DC.Host, cnf.DC.Port)
	log.Print(uri)
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	return &MG{client: client, database: cnf.DC.Database}, nil
}

func (mg *MG) GetNames(tableName string) ([]string, []string, error) {
	//get names of field
	command := bson.D{}
	mg.client.Database(tableName).RunCommand(context.Background(), command, nil)
	//get types of field
	return nil, nil, nil
}

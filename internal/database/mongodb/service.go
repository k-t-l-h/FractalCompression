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
	command := "\n  \"mapreduce\" : \"" + tableName + "\"," +
		"\n  \"map\" : function() {\n    for (var key in this) " +
		"{ emit(key, null); }\n  },\n  \"reduce\" : function(key, stuff) " +
		"{ return null; }, \n  \"out\": \" "+ tableName +"\" + \"_keys\"\n"

	res := mg.client.Database(tableName).RunCommand(context.Background(), command, nil)
	log.Print(res)
	res2, err := mg.client.Database(mg.database).
		Collection(tableName+"_keys").
		Distinct(context.Background(), "_id", bson.D{})
	//get types of field
	log.Print(res2, err)
	var names []string

	for _, x := range res2 {
		names = append(names,  fmt.Sprintf("%v", x))
	}
	log.Print(names)
	return nil, nil, nil
}


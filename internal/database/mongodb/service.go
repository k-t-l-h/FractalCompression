package mg

import (
	"FractalCompression/internal/config"
	"context"
	"fmt"
	"github.com/pkg/errors"
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

	comm := bson.D{bson.E{"mapreduce", tableName},
		bson.E{"map", "function(){for (var key in this) { emit(key, null); }"},
		bson.E{"reduce", "function(key, stuff) { return null; }"},
		bson.E{"out", tableName + "_keys"}}

	tag := mg.client.Database(tableName).RunCommand(context.Background(), comm, nil)

	if err := tag.Err(); err != nil {
		return nil, nil, errors.Wrap(err, "error while getting columns names 1")
	}

	rows, err := mg.client.Database(mg.database).
		Collection(tableName+"_keys").
		Distinct(context.Background(), "_id", bson.D{})
	if err != nil {
		return nil, nil, errors.Wrap(err, "error while getting columns names 2")
	}

	var names []string

	for _, x := range rows {
		names = append(names, fmt.Sprintf("%v", x))
	}
	//TODO: обеспечить проверку типа
	for _, row := range names {
		if row != "type" {
			stage := bson.D{
				{Key: "$project", Value: bson.D{
					bson.E{Key: row,
						Value: bson.E{Key: "$type",
							Value: fmt.Sprintf("$%s", row)},
					},
				},
				},
			}
			it, _ := mg.client.Database(mg.database).Collection(tableName).
				Aggregate(context.TODO(), mongo.Pipeline{stage})

			for it.Next(context.Background()) {
				//log.Print(it.Current.String())
			}
		}
	}

	return names, nil, nil
}

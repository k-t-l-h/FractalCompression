package mg

import "go.mongodb.org/mongo-driver/bson/primitive"

type Names struct {
	ID    primitive.ObjectID `bson:"_id"`
	Names []string           `bson:"names"`
}

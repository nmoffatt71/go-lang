package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Medication struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	PatientID  primitive.ObjectID `bson:"patient_id"`
	ProviderID primitive.ObjectID `bson:"provider_id"`
}

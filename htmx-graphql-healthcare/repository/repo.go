package repository

import (
	"context"
	"htmx-graphql-healthcare/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRepo struct {
	DB *mongo.Database
}

func NewMongoRepo(db *mongo.Database) *MongoRepo {
	return &MongoRepo{DB: db}
}

func (r *MongoRepo) GetMedications(ctx context.Context) ([]*model.Medication, error) {
	cur, err := r.DB.Collection("medications").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var meds []*model.Medication
	if err := cur.All(ctx, &meds); err != nil {
		return nil, err
	}
	return meds, nil
}

func (r *MongoRepo) AddMedication(ctx context.Context, med *model.Medication) error {
	_, err := r.DB.Collection("medications").InsertOne(ctx, med)
	return err
}

func (r *MongoRepo) GetPatientByID(ctx context.Context, id primitive.ObjectID) (*model.Patient, error) {
	var p model.Patient
	err := r.DB.Collection("patients").FindOne(ctx, bson.M{"_id": id}).Decode(&p)
	return &p, err
}

func (r *MongoRepo) GetProviderByID(ctx context.Context, id primitive.ObjectID) (*model.Provider, error) {
	var p model.Provider
	err := r.DB.Collection("providers").FindOne(ctx, bson.M{"_id": id}).Decode(&p)
	return &p, err
}

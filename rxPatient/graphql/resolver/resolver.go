package resolver

import (
	"context"
	"rxPatient/graphql/generated"
	"rxPatient/graphql/model"
	"rxPatient/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Resolver struct{}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver Resolver

func (r *queryResolver) Patients(ctx context.Context) ([]*model.Patient, error) {
	cursor, err := db.PatientCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var patients []*model.Patient
	if err = cursor.All(ctx, &patients); err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *queryResolver) SearchPatientsByLastname(ctx context.Context, lastname string) ([]*model.Patient, error) {
	filter := bson.M{"lastname": bson.M{"$regex": lastname, "$options": "i"}}
	cursor, err := db.PatientCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var patients []*model.Patient
	if err = cursor.All(ctx, &patients); err != nil {
		return nil, err
	}
	return patients, nil
}

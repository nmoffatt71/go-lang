package resolver

import (
	"context"
	"patient-manager/graphql/generated"
	"patient-manager/graphql/model"
	"patient-manager/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

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

type mutationResolver Resolver

func (r *mutationResolver) CreatePatient(ctx context.Context, firstname string, lastname string, middlename *string, dateofbirth string, gender string) (*model.Patient, error) {
	patient := model.Patient{
		Firstname:  firstname,
		Lastname:   lastname,
		Middlename: middlename,
		Dateofbirth: dateofbirth,
		Gender:     gender,
	}
	_, err := db.PatientCollection.InsertOne(ctx, patient)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *Resolver) SearchPatientsResolver(ctx context.Context, search string) ([]*model.Patient, error) {
	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"firstname": bson.M{"$regex": search, "$options": "i"}},
				{"lastname": bson.M{"$regex": search, "$options": "i"}},
				{"middlename": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}
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
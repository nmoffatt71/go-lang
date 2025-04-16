package service

import (
	"context"
	"htmx-graphql-healthcare/model"
	"htmx-graphql-healthcare/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	Repo *repository.MongoRepo
}

func NewService(repo *repository.MongoRepo) *Service {
	return &Service{Repo: repo}
}

type MedicationWithDetails struct {
	Name     string
	Patient  *model.Patient
	Provider *model.Provider
}

func (s *Service) GetMedications(ctx context.Context) ([]*MedicationWithDetails, error) {
	meds, err := s.Repo.GetMedications(ctx)
	if err != nil {
		return nil, err
	}
	var result []*MedicationWithDetails
	for _, m := range meds {
		patient, _ := s.Repo.GetPatientByID(ctx, m.PatientID)
		provider, _ := s.Repo.GetProviderByID(ctx, m.ProviderID)
		result = append(result, &MedicationWithDetails{
			Name:     m.Name,
			Patient:  patient,
			Provider: provider,
		})
	}
	return result, nil
}

func (s *Service) AddMedication(ctx context.Context, name, patientID, providerID string) (*MedicationWithDetails, error) {
	pid, _ := primitive.ObjectIDFromHex(patientID)
	prid, _ := primitive.ObjectIDFromHex(providerID)
	med := &model.Medication{
		Name:       name,
		PatientID:  pid,
		ProviderID: prid,
	}
	err := s.Repo.AddMedication(ctx, med)
	if err != nil {
		return nil, err
	}
	patient, _ := s.Repo.GetPatientByID(ctx, pid)
	provider, _ := s.Repo.GetProviderByID(ctx, prid)
	return &MedicationWithDetails{
		Name:     name,
		Patient:  patient,
		Provider: provider,
	}, nil
}

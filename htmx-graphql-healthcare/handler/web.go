package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"htmx-graphql-healthcare/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
}

func GetMedicationsHandler(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		meds, _ := svc.GetMedications(r.Context())
		for _, med := range meds {
			fmt.Fprintf(w, "<div>%s prescribed by %s to %s</div>", med.Name, med.Provider.Name, med.Patient.Name)
		}
	}
}

func AddMedicationHandler(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.FormValue("name")
		patientId := r.FormValue("patientId")
		providerId := r.FormValue("providerId")
		med, _ := svc.AddMedication(r.Context(), name, patientId, providerId)
		fmt.Fprintf(w, "<div>%s prescribed by %s to %s</div>", med.Name, med.Provider.Name, med.Patient.Name)
	}
}

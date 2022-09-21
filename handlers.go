package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) allReportsHandler(w http.ResponseWriter, r *http.Request) {
	reports, err := s.DB.GetAllReports()
	if err != nil {
		respondWithError(w, err)
		return
	}
	respondWithJSON(w, http.StatusOK, reports)
}

func (s *Server) reportHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	r_ID, err := strconv.Atoi(vars["r_ID"])
	if err != nil {
		respondWithError(w, err)
		return
	}
	report, err := s.DB.GetReport(r_ID)
	if err != nil {
		respondWithError(w, err)
		return
	}
	respondWithJSON(w, http.StatusOK, report)
}

func (s *Server) addReportHandler(w http.ResponseWriter, r *http.Request) {
	certname := r.FormValue("certname")
	environment := r.FormValue("environment")
	status := r.FormValue("status")
	time := r.FormValue("time")
	transaction_uuid := r.FormValue("transaction_uuid")

	err := s.DB.AddReport(certname, environment, status, time, transaction_uuid)
	if err != nil {
		respondWithError(w, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}

func (s *Server) removeReportHandler(w http.ResponseWriter, r *http.Request) {
	transaction_uuid := r.FormValue("transaction_uuid")

	err := s.DB.RemoveReport(transaction_uuid)
	if err != nil {
		respondWithError(w, err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}

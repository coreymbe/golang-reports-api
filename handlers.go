package main

import (
	"encoding/json"
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
	var req Report
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, err)
		return
	}

	req_err := s.DB.AddReport(req.Certname, req.Environment, req.Status, req.Time, req.TransactionUUID)
	if req_err != nil {
		respondWithError(w, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}

func (s *Server) removeReportHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	r_ID, err := strconv.Atoi(vars["r_ID"])
	if err != nil {
		respondWithError(w, err)
		return
	}
	del_err := s.DB.RemoveReport(r_ID)
	if del_err != nil {
		respondWithError(w, err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}

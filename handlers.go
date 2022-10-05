package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func (s *Server) authHandler(w http.ResponseWriter, r *http.Request) {
	var req User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		msg := &AuthError{
			Error: "Must provide username and password to generate auth token!",
		}
		respondWithJSON(w, http.StatusUnauthorized, msg)
		return
	}

	user, err := s.DB.UserExists(req.Name, req.Password)
	if err != nil {
		respondWithError(w, err)
		return
	}

	if user.Name == req.Name && user.Password == req.Password {
		expiresAt := (time.Now().Add(time.Hour * 12)).Unix()

		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = &AuthTokenClaim{
			&jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Unix(expiresAt, 0)),
			},
		}

		tokenString, err := token.SignedString([]byte(os.Getenv("JWTSecretKey")))
		if err != nil {
			respondWithError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		auth_token := &AuthToken{
			Token:     tokenString,
			TokenType: "Bearer",
			ExpiresAt: time.Unix(expiresAt, 0),
		}

		respondWithJSON(w, http.StatusOK, auth_token)
	} else {
		msg := &AuthError{
			Error: "Username or Password is incorrect!",
		}
		respondWithJSON(w, http.StatusUnauthorized, msg)
	}
}

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
	authHeader := r.Header.Get("authorization")
	if authHeader != "" {
		authToken := strings.Split(authHeader, " ")

		valid := authCheck(authToken[1])
		if valid {
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
		} else {
			msg := &AuthError{
				Error: "Token is not valid!",
			}
			respondWithJSON(w, http.StatusUnauthorized, msg)
		}
	} else {
		msg := &AuthError{
			Error: "This API requires a token for authentication.",
		}
		respondWithJSON(w, http.StatusUnauthorized, msg)
	}
}

func (s *Server) removeReportHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("authorization")
	if authHeader != "" {
		authToken := strings.Split(authHeader, " ")

		valid := authCheck(authToken[1])
		if valid {
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
		} else {
			msg := &AuthError{
				Error: "Token is not valid!",
			}
			respondWithJSON(w, http.StatusUnauthorized, msg)
			return
		}
	} else {
		msg := &AuthError{
			Error: "This API requires a token for authentication.",
		}
		respondWithJSON(w, http.StatusUnauthorized, msg)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}

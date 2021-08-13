package handlers

import (
	"log"
	"net/http"
	"fmt"
)

type Version struct{
	l *log.Logger
}

func NewVersion(l *log.Logger) *Version {
	return &Version{l}
}

func (v Version) GetVersion(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Zuri Chat API - Version 0.0001\n")
}


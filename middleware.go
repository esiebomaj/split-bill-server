package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GroupHandler func(w http.ResponseWriter, r *http.Request, group *Group)

func (b *AppBackend) AuthMiddleWare(handler GroupHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey := r.Header.Get("Secret-Code")
		if apiKey == "" {
			HandleErrorJson(w, 401, "Secret-Code header required")
			return
		}

		groupIDStr := chi.URLParam(r, "groupID")
		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			HandleErrorJson(w, 400, fmt.Sprintf("Could not retrieve group %v", err))
			return
		}

		group := Group{}
		err = b.GetAllGroupById(uint(groupID), &group)
		if err != nil {
			HandleErrorJson(w, 400, fmt.Sprintf("Could not retrieve group %v", err))
			return
		}

		if group.SecretCode != apiKey {
			HandleErrorJson(w, 401, "Incorrect secret code")
			return
		}

		handler(w, r, &group)
	}
}

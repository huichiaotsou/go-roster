package handler

import "net/http"

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request)    {}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {}

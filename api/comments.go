package api

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/holacloud/store"
)

type Comment struct {
	store.Id
	PageID    string    `json:"page_id"`
	ParentID  string    `json:"parent_id,omitempty"`
	Author    string    `json:"author"`
	Email     string    `json:"email,omitempty"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentsAPI struct {
	Store store.Storer[*Comment]
}

func NewCommentsAPI(dataDir string) (*CommentsAPI, error) {
	s, err := store.NewStoreDiskCached[*Comment](dataDir)
	if err != nil {
		return nil, fmt.Errorf("create comments store: %w", err)
	}
	return &CommentsAPI{Store: s}, nil
}

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (a *CommentsAPI) List(w http.ResponseWriter, r *http.Request) {
	pageID := r.URL.Query().Get("page_id")
	if pageID == "" {
		http.Error(w, `{"error":"page_id is required"}`, http.StatusBadRequest)
		return
	}

	all, err := a.Store.List(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	comments := make([]*Comment, 0)
	for _, c := range all {
		if (*c).PageID == pageID {
			comments = append(comments, *c)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"comments": comments,
	})
}

type CreateCommentRequest struct {
	PageID   string `json:"page_id"`
	ParentID string `json:"parent_id,omitempty"`
	Author   string `json:"author"`
	Email    string `json:"email,omitempty"`
	Content  string `json:"content"`
}

func (a *CommentsAPI) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"invalid JSON: %s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	if req.PageID == "" {
		http.Error(w, `{"error":"page_id is required"}`, http.StatusBadRequest)
		return
	}
	if req.Author == "" {
		http.Error(w, `{"error":"author is required"}`, http.StatusBadRequest)
		return
	}
	if req.Content == "" {
		http.Error(w, `{"error":"content is required"}`, http.StatusBadRequest)
		return
	}

	comment := &Comment{
		Id:        *store.NewId(generateID()),
		PageID:    req.PageID,
		ParentID:  req.ParentID,
		Author:    req.Author,
		Email:     req.Email,
		Content:   req.Content,
		CreatedAt: time.Now().UTC(),
	}

	if err := a.Store.Put(r.Context(), &comment); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"comment": comment,
	})
}

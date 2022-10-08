package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/saitooooooo/go-api-intermediate/apperrors"
	"github.com/saitooooooo/go-api-intermediate/controllers/services"
	"github.com/saitooooooo/go-api-intermediate/models"
)

// Comment用のコントローラ構造体
type CommentController struct {
	service services.CommentServicer
}

// コンストラクタ関数
func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}

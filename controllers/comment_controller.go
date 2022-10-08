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
		apperrors.ErrorHandler(w, req, err)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(comment)
}

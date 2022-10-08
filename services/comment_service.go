package services

import (
	"github.com/saitooooooo/go-api-intermediate/apperrors"
	"github.com/saitooooooo/go-api-intermediate/models"
	"github.com/saitooooooo/go-api-intermediate/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}

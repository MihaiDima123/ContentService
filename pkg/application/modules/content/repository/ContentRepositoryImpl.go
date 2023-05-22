package repository

import (
	"contentservice/pkg/application/customErrors"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/application/modules/content/interfaces"
	"contentservice/pkg/interfaces/restful"
	"errors"
	"gorm.io/gorm"
)

type ContentRepositoryImpl struct {
	dbConn *gorm.DB
}

func NewContentRepository() interfaces.ContentRepository {
	return new(ContentRepositoryImpl)
}

func (cr *ContentRepositoryImpl) Configure(configuration restful.RepositoryConfiguration) {
	cr.dbConn = configuration.Connection
}

func (cr *ContentRepositoryImpl) Create(post post_entities.Post) (int64, customErrors.DbError) {
	result := cr.dbConn.Create(&post)

	if result.Error != nil {
		return post.ID, customErrors.ResourceNotCreatedError
	}

	return post.ID, nil
}

func (cr *ContentRepositoryImpl) GetById(id int64) (*post_entities.Post, customErrors.DbError) {
	post := new(post_entities.Post)

	err := cr.dbConn.Where("id = ?", id).First(post).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, customErrors.ResourceNotFoundError
	}

	return post, nil
}

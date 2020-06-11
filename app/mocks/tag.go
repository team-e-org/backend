package mocks

import (
	"app/models"
	"app/repository"
)

type TagMock struct {
	ExpectedTag *models.Tag
}

func NewTagRepository() repository.TagRepository {
	return &TagMock{}
}

func (m *TagMock) AddTag(tag *db.Tag) error {
	m.ExpectedTag = tag
	return nil
}

func (m *TagMock) GetTag(tagID int) (*models.Tag, error) {
	return m.ExpectedTag, nil
}

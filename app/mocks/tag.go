package mocks

import (
	"app/models/db"
	"app/repository"
)

type TagMock struct {
	ExpectedTag *db.Tag
}

func NewTagRepository() repository.TagRepository {
	return &TagMock{}
}

func (m *TagMock) CreateTag(tag *db.Tag) error {
	m.ExpectedTag = tag
	return nil
}

func (m *TagMock) GetTag(tagID int) (*db.Tag, error) {
	return m.ExpectedTag, nil
}

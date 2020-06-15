package mocks

import (
	"app/models"
	"app/repository"
	"errors"
)

type TagMock struct {
	ExpectedTags []*models.Tag
	PinTagMapper map[int][]int // map[pinID][]tagID
}

func NewTagRepository() repository.TagRepository {
	return &TagMock{
		ExpectedTags: make([]*models.Tag, 0),
		PinTagMapper: make(map[int][]int),
	}
}

func (m *TagMock) CreateTag(tag *models.Tag) error {
	m.ExpectedTags = append(m.ExpectedTags, tag)
	return nil
}

func (m *TagMock) GetTag(tagID int) (*models.Tag, error) {
	for _, t := range m.ExpectedTags {
		if t.ID == tagID {
			return t, nil
		}
	}
	return nil, noTagError()
}

func (m *TagMock) AttachTagToPin(tagID int, pinID int) error {
	m.PinTagMapper[pinID] = append(m.PinTagMapper[pinID], tagID)
	return nil
}

func (m *TagMock) GetTagsByPinID(pinID int) ([]*models.Tag, error) {
	tags := make([]*models.Tag, 0, len(m.PinTagMapper))
	for _, id := range m.PinTagMapper[pinID] {
		for _, t := range m.ExpectedTags {
			if t.ID == id {
				tags = append(tags, t)
			}
		}
	}
	return tags, nil
}

func noTagError() error {
	return errors.New("An error occurred, the tag does not exist")
}

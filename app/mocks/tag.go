package mocks

import (
	"app/models"
	"app/repository"
	"errors"
)

type PinID int
type TagID int

type TagMock struct {
	ExpectedTags []*models.Tag
	PinTagMapper map[PinID][]TagID // map[pinID][]tagID
}

func NewTagRepository() repository.TagRepository {
	return &TagMock{}
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
	m.PinTagMapper[PinID(pinID)] = append(m.PinTagMapper[PinID(pinID)], TagID(tagID))
	return nil
}

func (m *TagMock) GetTagsByPinID(pinID int) ([]*models.Tag, error) {
	tags := make([]*models.Tag, 0, len(m.PinTagMapper))
	for _, id := range m.PinTagMapper[PinID(pinID)] {
		for _, t := range m.ExpectedTags {
			if TagID(t.ID) == id {
				tags = append(tags, t)
			}
		}
	}
	return tags, nil
}

func noTagError() error {
	return errors.New("An error occurred, the tag does not exist")
}

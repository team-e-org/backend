package usecase

import (
	"app/db"
	"app/helpers"
	"app/logs"
	"app/models"
	"database/sql"
)

func GetTagsByPinID(data db.DataStorageInterface, pinID int) ([]*models.Tag, helpers.AppError) {
	tags, err := data.Tags().GetTagsByPinID(pinID)
	if err == sql.ErrNoRows {
		logs.Info("No tags found with pinID: %d", pinID)
		return []*models.Tag{}, nil
	}
	if err != nil {
		logs.Error("An error occurred while getting tags of pin %d: %v", pinID, err)
		err := helpers.NewInternalServerError(err)
		return nil, err
	}

	return tags, nil
}

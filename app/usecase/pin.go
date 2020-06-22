package usecase

import (
	"app/db"
	"app/helpers"
	"app/logs"
	"app/models"
	"database/sql"
)

func GetPinsByBoardID(data *db.DataStorage, userID int, boardID int, page int) ([]*models.Pin, helpers.AppError) {
	pins, err := data.Pins.GetPinsByBoardID(boardID, page)
	if err != nil {
		logs.Error("An error occurred while getting pins in board: %v", err)
		err := helpers.NewInternalServerError(err)
		return nil, err
	}

	pins = removePrivatePin(pins, userID)
	return pins, nil
}

func ServePin(data *db.DataStorage, pinID int, userID int) (*models.Pin, helpers.AppError) {
	pin, err := data.Pins.GetPin(pinID)
	if err == sql.ErrNoRows {
		logs.Error("Pin not found in database: %v", pinID)
		err := helpers.NewNotFound(err)
		return nil, err
	}
	if err != nil {
		logs.Error("An error occurred while getting pin from database: %v", err)
		err := helpers.NewInternalServerError(err)
		return nil, err
	}

	if pin.IsPrivate && *pin.UserID != userID {
		logs.Error("Pin not found in database: %v", pinID)
		err := helpers.NewNotFound(err)
		return nil, err
	}

	return pin, nil
}

func CreatePin(data *db.DataStorage, pin *models.Pin, userID int, boardID int) (*models.Pin, helpers.AppError) {
	pin, err := data.Pins.CreatePin(pin, boardID)
	if err != nil {
		logs.Error("Creating pin: %v", err)
		err := helpers.NewInternalServerError(err)
		return nil, err
	}

	err = data.BoardsPins.CreateBoardPin(boardID, pin.ID)
	if err != nil {
		logs.Error("Creating board_pin: %v", err)
		err := helpers.NewInternalServerError(err)
		return nil, err
	}

	return pin, nil
}

func removePrivatePin(pins []*models.Pin, userID int) []*models.Pin {
	for i, pin := range pins {
		if pin.IsPrivate && *pin.UserID != userID {
			pins = append(pins[:i], pins[i+1:]...)
		}
	}

	return pins
}

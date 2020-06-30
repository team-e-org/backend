package usecase

import (
	"app/db"
	"app/helpers"
	"app/logs"
	"app/models"
	"database/sql"
	"fmt"
)

func GetPinsByBoardID(data db.DataStorageInterface, userID int, boardID int, page int) ([]*models.Pin, helpers.AppError) {
	pins, err := data.Pins().GetPinsByBoardID(boardID, page)
	if err != nil {
		logs.Error("An error occurred while getting pins in board: %v", err)
		err := helpers.NewInternalServerError(err)
		return nil, err
	}

	pins = removePrivatePin(pins, userID)
	return pins, nil
}

func ServePin(data db.DataStorageInterface, pinID int, userID int) (*models.Pin, helpers.AppError) {
	data.AWSS3()
	pin, err := data.Pins().GetPin(pinID)
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

func GetPins(data db.DataStorageInterface, page int) ([]*models.Pin, helpers.AppError) {
	pins, err := data.Pins().GetPins(page)
	if err != nil {
		logs.Error("An error occurred while getting pins in board: %v", err)
		err := helpers.NewInternalServerError(err)
		return nil, err
	}

	return pins, nil
}

func CreatePin(data db.DataStorageInterface, pin *models.Pin, boardID int) (*models.Pin, helpers.AppError) {
	pin, err := data.Pins().CreatePin(pin, boardID)
	if err != nil {
		logs.Error("Creating pin: %v", err)
		err := helpers.NewInternalServerError(err)
		return nil, err
	}

	return pin, nil
}

func UpdatePin(data db.DataStorageInterface, newPin *models.Pin, userID int) (*models.Pin, helpers.AppError) {

	pin, err := data.Pins().GetPin(newPin.ID)
	if err == sql.ErrNoRows {
		logs.Error("Pin not found in database: %v", newPin.ID)
		err := helpers.NewNotFound(err)
		return nil, err
	}
	if err != nil {
		logs.Error("An error occurred while getting pin from database: %v", err)
		err := helpers.NewInternalServerError(err)
		return nil, err
	}

	if *pin.UserID != userID {
		logs.Error("Not user's pin error")
		err := helpers.NewUnauthorized(fmt.Errorf("Not user's pin error"))
		return nil, err
	}

	pin.Title = newPin.Title
	pin.Description = newPin.Description
	pin.URL = newPin.URL

	err = data.Pins().UpdatePin(pin)

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

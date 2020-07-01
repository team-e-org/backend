package usecase

import (
	"app/db"
	"app/helpers"
	"app/logs"
	"app/models"
	"context"
	"database/sql"
	"fmt"
	"mime/multipart"
	"time"

	"golang.org/x/sync/errgroup"
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

func uploadImageToS3(ctx context.Context, data db.DataStorageInterface, file multipart.File, fileName string, contentType string, userID int) error {
	i := 0
	err := fmt.Errorf("Uploading file to S3 failed")
	fmt.Printf("ctx: %v\n", ctx)
	for {
		select {
		case <-ctx.Done():
			return err
		default:
			if err := data.AWSS3().UploadImage(file, fileName, contentType, userID); err != nil {
				i++
				logs.Error("%v, %d", err, i)
				continue
			}
			logs.Info("Uploading file to S3 succeeded")
			return nil
		}
	}
}

func createPin(ctx context.Context, data db.DataStorageInterface, pin *models.Pin, boardID int, pinIDCh chan int) error {
	i := 0
	err := fmt.Errorf("Inserting pin column into DB failed")
	for {
		select {
		case <-ctx.Done():
			return err
		default:
			pin, err := data.Pins().CreatePin(pin, boardID)
			if err != nil {
				i++
				logs.Error("%v, %d", err, i)
				continue
			}
			pinIDCh <- pin.ID
			logs.Info("Inserting pin column into DB succeeded")
			return nil
		}
	}
}

func CreatePin(data db.DataStorageInterface, pin *models.Pin, file multipart.File, fileName string, contentType string, userID int, boardID int) (*models.Pin, helpers.AppError) {

	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pinIDCh := make(chan int, 1)

	eg.Go(func() error {
		return uploadImageToS3(ctx, data, file, fileName, contentType, userID)
	})

	eg.Go(func() error {
		return createPin(ctx, data, pin, boardID, pinIDCh)
	})

	if err := eg.Wait(); err != nil {
		logs.Error("CreatePin failed: %v", err)
		err := helpers.NewInternalServerError(err)
		return nil, err
	}

	pinID := <-pinIDCh
	pin.ID = pinID

	logs.Info("New pin inserted: %v", pin)

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

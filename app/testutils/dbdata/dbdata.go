package dbdata

import (
	"app/models"
)

var BaseUser = &models.User{
	ID:             0,
	Name:           "test user",
	Email:          "basic@user.com",
	HashedPassword: "$2a$10$Ga.kdsed5IQEng6SRNlq3.u14XlcF61MfmPyJCHPckvFxYk0HGbsu", //asdf1234
}

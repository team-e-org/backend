package dbdata

import "app/models"

var BaseUser = &models.User{
	ID:             1,
	Name:           "current user",
	Email:          "current_user@email.com",
	Icon:           "test icon",
	HashedPassword: "$2a$10$SOWUFP.hkVI0CrCJyfh5vuf/Gu.SDpv6Y2DYZ/Dbwyr.AKtlAldFe", // before hashing: password
}

var BaseUserPassword = "password"

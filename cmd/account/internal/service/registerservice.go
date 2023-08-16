package service

import "errors"

func Register(accountName, password string) error {
	if accountName == "leebai" && password == "123456" {
		return nil
	}

	return errors.New("invalid account or password")
}

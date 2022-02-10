package database

import "fileserver/model"

func StoreNewUser(userinfo *model.Userinfo, password *model.UserPassword) error {

	return nil
}

func CheckAvailableAccount(account string) bool {
	return true
}

func CheckUserPassword(account, password string) bool {
	return true
}

package handler

import (
	"fileserver/cache"
	"fileserver/database"
	"fileserver/model"
	"fileserver/utils"
	"net/http"
)

type LanderRequest struct {
	Account   string
	Username  string
	Password  string
	InviterID int64
}

func GetLanderRequest(r *http.Request) (*LanderRequest, *model.ResponseCode) {
	var req = new(LanderRequest)
	var ok = true
	req.Account, ok = utils.GetFormString(r, "account")
	if !ok {
		return nil, model.ResponseCodeMissingParam
	}
	req.Username, ok = utils.GetFormString(r, "username")
	if !ok {
		req.Username = req.Account
	}
	req.Password, ok = utils.GetFormString(r, "password")
	if !ok {
		return nil, model.ResponseCodeMissingParam
	}
	req.InviterID, _ = utils.GetFormInt64(r, "inviter_id")

	return req, nil
}

func Regestry(r *http.Request) *model.ResponseCode {
	r.ParseForm()
	req, resCode := GetLanderRequest(r)
	if resCode != nil {
		return resCode
	}

	if !database.CheckAvailableAccount(req.Account) {
		return model.ResponseCodeDuplicateAccount
	}

	userinfo, userPassword := InitUser()

	if err := database.StoreNewUser(userinfo, userPassword); err != nil {
		return model.ResponseCodeCreateAccountFailed
	}

	return model.ResponseCodeOK
}

func Login(r *http.Request) *model.ResponseCode {
	r.ParseForm()
	req, resCode := GetLanderRequest(r)
	if resCode != nil {
		return resCode
	}

	if !database.CheckUserPassword(req.Account, req.Password) {
		return model.ResponseCodeWrongAccountOrPassword
	}

	// TODO token

	return model.ResponseCodeOK.SetSuccessResponse("OK") // token struct json string
}

// CheckAlive is default request, don't need to catch this action
func CheckAlive(r *http.Request) *model.ResponseCode {
	r.ParseForm()
	token, ok := utils.GetFormString(r, "token")
	if !ok {
		return model.ResponseCodeMissingParam
	}

	if !cache.CheckAlive(token) {
		return model.ResponseCodeLoginExpire
	}

	return model.ResponseCodeOK
}

func InitUser() (*model.Userinfo, *model.UserPassword) {
	return nil, nil
}

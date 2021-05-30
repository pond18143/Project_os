package iojwt

import (
	mssql "project_os/database/mssql"
)

func getLogin(request credentials) (detail roleLogin, err error) {

	if err = mssql.DB.
		Table("shopping.dbo.login").
		Where("user_name = ?", request.UserName).
		Find(&detail).Error; err != nil {
		return
	}

	return
}

func createUser(request inputRegister) (result messageResponse, err error) {
	if err = mssql.DB.Table("shopping.dbo.login").
		Create(&request).Error; err != nil {
		return
	}

	return
}
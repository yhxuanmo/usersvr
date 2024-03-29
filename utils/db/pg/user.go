package pg

import "usersvr/utils/db/model"

func Register(Email, Password string)(*model.User, error){
	user := &model.User{
		Email: Email,
		Password: Password,
	}
	res := PgClient.Create(&user)
	if err := res.Error; err != nil{
		return user, err
	}else {
		return user, nil
	}
}

func FindOneUser(Email string)(*model.User, error){
	var user model.User
	err := PgClient.Where("email = ?", Email).Find(&user).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}


func UpdateUser(User *model.User, UpdataInfo map[string]interface{}) (*model.User, error) {
	err := PgClient.Model(&User).Updates(UpdataInfo).Error
	if err != nil {
		return User, err
	}
	return User, nil
}

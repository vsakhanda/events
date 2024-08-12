package models

type Auth struct {
	ID       int
	Username string
	Password string
}

//
//func CheckAuth(username, password string) (bool, error) {
//	var auth Auth
//	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return false, err
//	}
//
//	if auth.ID > 0 {
//		return true, nil
//	}
//
//	return false, nil
//}

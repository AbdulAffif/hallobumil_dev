package repository

import (
	"github.com/google/uuid"
	"log"
	"time"

	"github.com/AbdulAffif/hallobumil_dev/api/entity"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	InsertUserData(user entity.User_data) (entity.User_data,error)
	InsertPregnancy(param entity.Ud_pregnancy)(entity.Ud_pregnancy,error)
	UpdateUser(user entity.User) entity.Result
	VerifyCredEmail(email string, password string) (entity.User,error)
	VerifyCredPhone(phone string, password string) (entity.User,error)
	IsDuplicateEmail(email string) (tx *gorm.DB)
	GetUser(userID uint64) (entity.User,error)
	GetUserData(userID uint64) (entity.User_data,error)
	GetPregnancy(userID uint64)(entity.Ud_pregnancy,error)
	InsertAccess(acc entity.Ud_access) entity.Ud_access
	UpdateAccess(acc entity.Ud_access) entity.Ud_access
	FindToken(userID uint64, identifier string) (entity.Ud_access)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}
func (db *userConnection) InsertUser(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	user.Created_date = time.Now()
	user.Last_update = time.Now()
	user.Otp_verifivation_code_created_date = time.Now()
	user.Phone_otp_code_created_at = time.Now()
	user.Verification_code = uuid.New().String()
	user.Is_verified = false
	user.Is_new = true
	user.Is_deleted = false
	var rst = db.connection.Create(&user)
	return user, rst.Error
}
func (db *userConnection) InsertUserData(userData entity.User_data) (entity.User_data,error) {
	userData.Is_deleted = false
	userData.Is_migrate = true
	userData.Iteration = 1
	userData.Last_update = time.Now()
	userData.Iteration_pre = 0
	userData.Current_state = 1
	err := db.connection.Create(&userData)
	return userData, err.Error
}
func (db *userConnection)InsertPregnancy(param entity.Ud_pregnancy)(entity.Ud_pregnancy,error)  {
	param.Iteration = 1
	err := db.connection.Create(&param)
	return param, err.Error
}
func (db *userConnection) UpdateUser(user entity.User) entity.Result {
	var result entity.Result

	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	user.Created_date = time.Now()
	user.Last_update = time.Now()
	db.connection.Save(&user)
	return result
}
func (db *userConnection) VerifyCredEmail(email string, password string) (entity.User,error) {
	var result entity.User
	res := db.connection.Where("email = ? and is_deleted = ?", email, false).Order("id").Take(&result)
	//res := db.connection.Table("users").Select("users.id as id, users.email as email, users.is_verified as is_verified, users.status as status,user_data.name as name, user_data.profile_picture as profile_picture, user_data.birthdate as birthdate, user_data.height as height, user_data.iteration as iteration").Joins("left join user_data on users.id = user_data.id_user").Where("users.email = ? ", email).Scan(&result)
	return result, res.Error
}
func (db *userConnection) VerifyCredPhone(phone string, password string) (entity.User,error) {
	var result entity.User
	res := db.connection.Where("phone = ? and is_deleted = ?", phone, false).Order("id").Take(&result)
	//res := db.connection.Model(entity.User{}).Joins("left join user_data on [users].id = user_data.id_user").Where("user_data.phone = ? AND user_data.is_deleted = ?", phone, false).Scan(&result)
	return result, res.Error
}
func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ? and is_deleted = ?", email, false).Order("id").Take(&user)
}

func (db *userConnection) GetUser(userID uint64) (entity.User,error) {
	var user entity.User
	res := db.connection.Find(&user, userID)
	return user,res.Error
}
func (db *userConnection) GetUserData(userID uint64) (entity.User_data,error) {
	var userData entity.User_data
	res:= db.connection.Where("id_user = ?",userID).Order("id").Take(&userData)
	return userData,res.Error
}
func (db *userConnection) GetPregnancy(userID uint64) (entity.Ud_pregnancy,error) {
	var userPreg entity.Ud_pregnancy
	res:= db.connection.Where("id_user = ?",userID).Order("id").Take(&userPreg)
	return userPreg,res.Error
}
func (db *userConnection) InsertAccess(acc entity.Ud_access) entity.Ud_access {
	acc.Token = uuid.New().String()
	acc.Last_update= time.Now()
	db.connection.Create(&acc)
	return acc
}
func (db *userConnection) UpdateAccess(acc entity.Ud_access) entity.Ud_access {
	acc.Token = uuid.New().String()
	acc.Last_update = time.Now()
	db.connection.Save(&acc)
	return acc
}
func (db *userConnection) FindToken(userID uint64, identifier string) (entity.Ud_access) {
	var acc entity.Ud_access
	db.connection.Where("id_user = ? and identifier = ? and is_deleted = ?", userID, identifier, false).Order("id").Take(&acc)
	return acc
}
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

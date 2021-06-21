package service

import (
	"github.com/AbdulAffif/hallobumil_dev/api/dto"
	"github.com/AbdulAffif/hallobumil_dev/api/entity"
	"github.com/AbdulAffif/hallobumil_dev/api/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type AuthService interface {
	VerifyEmail(email string, password string, identifier string) interface{}
	VerifyPhone(phone string, password string,identifier string) interface{}
	CreateUser(user dto.RegisterDTO) entity.JsonRegister
	IsDuplicateEmail(email string) bool
	Ping()bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyEmail(email string, password string, identifier string) interface{} {
	var res entity.Result
	user, errUser := service.userRepository.VerifyCredEmail(email, password)
	if errUser != nil{
		log.Println("login user error",email)
	}else{
		ud, errUD := service.userRepository.GetUserData(user.ID)
		if errUD != nil {
			log.Println("login user data error")
		}
		preg, errPrag := service.userRepository.GetPregnancy(user.ID)
		if errPrag != nil {
			log.Println("login pregnancy error")
		}
		dataAcc := entity.Ud_access{}
		dataAcc.Id_user = user.ID
		dataAcc.Identifier = identifier
		token := service.getToken(dataAcc)
		res.ID = user.ID
		res.Name = ud.Name
		res.Email = user.Email
		res.Birthdate = ud.Birthdate
		res.Height = ud.Height
		res.Phone = ud.Phone
		res.Iteration = ud.Iteration
		res.Is_verified = user.Is_verified
		res.ProfilePicture = ud.ProfilePicture
		res.Status = user.Status
		respone :=entity.JsonRegister{
			res,
			preg,
			token,
			"",
			time.Now(),
		}
		comparePassword := comparePassword(user.Password, []byte(password))
		if user.Email == email && comparePassword {
			return respone
		}
	}

	return false

}
func (service *authService) VerifyPhone(phone string, password string,identifier string) interface{} {
	var res entity.Result
	user, errUser := service.userRepository.VerifyCredPhone(phone, password)
	if errUser != nil{
		log.Println("login user error ",phone)
	}else {
		ud, errUD := service.userRepository.GetUserData(user.ID)
		if errUD != nil {
			log.Println("login user data error")
		}
		preg, errPrag := service.userRepository.GetPregnancy(user.ID)
		if errPrag != nil {
			log.Println("login pregnancy error")
		}
		dataAcc := entity.Ud_access{}
		dataAcc.Id_user = user.ID
		dataAcc.Identifier = identifier
		token := service.getToken(dataAcc)
		res.ID = user.ID
		res.Name = ud.Name
		res.Email = user.Email
		res.Birthdate = ud.Birthdate
		res.Height = ud.Height
		res.Phone = ud.Phone
		res.Iteration = ud.Iteration
		res.Is_verified = user.Is_verified
		res.ProfilePicture = ud.ProfilePicture
		res.Status = user.Status
		respone :=entity.JsonRegister{
			res,
			preg,
			token,
			"",
			time.Now(),
		}
		comparePassword := comparePassword(user.Password, []byte(password))
		if user.Phone == phone && comparePassword {
			return respone
		}
	}

	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.JsonRegister {
	var res entity.Result
	userToCreate := entity.User{}
	userToCreate.Email = user.Email
	userToCreate.Password = user.Password
	userToCreate.Phone = user.Phone
	userToCreate.RegisterOsType = user.RegisterOsType
	//err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	//if err != nil {
	//	log.Fatalf("failed map %v", err)
	//}
	userToCreate.Phone = fixPhone(user.Phone)
	ius , errius := service.userRepository.InsertUser(userToCreate)
	if errius != nil {
		log.Println("insert data user error")
	}
	UDToCreate := entity.User_data{}
	UDToCreate.Id_user = ius.ID
	UDToCreate.Name = user.Name
	UDToCreate.Phone = fixPhone(user.Phone)
	UDToCreate.Height = user.Height
	UDToCreate.Birthdate, _ = time.Parse("2006-01-02",user.Birthdate)
	//errUD := smapping.FillStruct(&UDToCreate, smapping.MapFields(&user))
	//if errUD != nil {
	//	log.Fatalf("failed map %v", errUD)
	//}
	iud , erriud := service.userRepository.InsertUserData(UDToCreate)
	if erriud != nil {
		log.Println("insert data user data error")
	}

	dataPregnancy := entity.Ud_pregnancy{}
	dataPregnancy.IdUser = ius.ID
	dataPregnancy.Expected_date , _ = time.Parse("2006-04-02",user.ExpectedDate)
	ipreg , errPreg := service.userRepository.InsertPregnancy(dataPregnancy)
	if errPreg != nil {
		log.Println("insert pragnancy error")
	}
	dataAcc := entity.Ud_access{}
	dataAcc.Id_user = ius.ID
	dataAcc.Identifier = user.Identifier
	token := service.getToken(dataAcc)
	res.ID = ius.ID
	res.Name = iud.Name
	res.Email = ius.Email
	res.Birthdate = iud.Birthdate
	res.Height = iud.Height
	res.Phone = iud.Phone
	res.Iteration = iud.Iteration
	res.Is_verified = ius.Is_verified
	res.ProfilePicture = iud.ProfilePicture
	res.Status = ius.Status
	respone :=entity.JsonRegister{
		res,
		ipreg,
		token,
		"",
		time.Now(),
	}
	return respone
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
func (service *authService) Ping()bool  {
	res := service.userRepository.IsDuplicateEmail("tes")
	return !(res.Error == nil)
}
func fixPhone(phone string) string  {
	p := []rune(phone)
	twoNum := string(p[0:2])

	if twoNum == "62" {
		phone = "0" + string(p[2:len(p)-1])
	} else if twoNum == "+6" {
		phone = "0" + string(p[3:len(p)-2])
	}
	return phone
}
func (service *authService) getToken(dataAcc entity.Ud_access) string  {
	var token string
	cekToken := service.userRepository.FindToken(dataAcc.Id_user,dataAcc.Identifier)
	if cekToken.Token == ""  {
		errToken := service.userRepository.InsertAccess(dataAcc)
		token = errToken.Token
	}else {
		dataAcc.ID = cekToken.ID
		errToken := service.userRepository.UpdateAccess(dataAcc)
		token = errToken.Token
	}

	return token
}
func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

package service

import (
	structs "Golangcrud/Structs"
	"Golangcrud/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt     = "24729384865hehtwe"
	siginKey = "ioywf8iy283rkjfk"
	tokenTTL = 12 * time.Hour
)

type Authservice struct {
	repo repository.Authorization
}

type StructClaims struct {
	jwt.StandardClaims
	UserId int `json:"userid"`
}

func newAuthService(repo repository.Authorization) *Authservice {
	return &Authservice{repo: repo}
}

func (s *Authservice) GenerateToken(username, password string) (string, error) {

	user, err := s.repo.Getuser(username, generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &StructClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, user.Id,
	})

	return token.SignedString([]byte(siginKey))
}


func (s *Authservice)Checkdatabaseusertoken(token string)(structs.User,error){
	return s.repo.Checkdatabaseusertoken(token)
}


func (s *Authservice) CreateUser(user structs.User) (string, error) {
	user.Id = rand.Intn(10000)
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *Authservice) ForgotMypassword(myuser string,username structs.Fpasswordstruct) (string, error) {
	if myuser != username.Username{ 
		return "",errors.New("this is not your user")
	}
	token := geneRateForgotpasswordHash();
	return s.repo.ForgotMypassword(token,myuser,username);
}

func(s *Authservice)ChangePassword(user structs.User,password structs.Newpassword)(string,error){ 
	username := user.Username;
	if password.Password != password.Repeatepassword{
		return "",errors.New("Passwords need to be same")
	} 
	newpassword := generatePasswordHash(password.Password);
	return s.repo.ChangePassword(username,newpassword);
}


func (s *Authservice) Parsetoken(token string) (int, error) {
	jwttoken, err := jwt.ParseWithClaims(token, &StructClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(siginKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := jwttoken.Claims.(*StructClaims)

	if !ok {
		return 0, errors.New("token claims are not type structs claims")
	}
	return claims.UserId, nil
}
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}



func geneRateForgotpasswordHash()string{
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
 
    s := make([]rune, 10)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}

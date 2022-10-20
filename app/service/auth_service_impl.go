package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kurir-go/app/dto"
	"github.com/kurir-go/app/model"
	"github.com/kurir-go/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepository repository.UserRepository
	secretKey      string
	issuer         string
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepository: userRepo,
		issuer:         "kurirgo",
		secretKey:      getSecretKey(),
	}
}

type jwtCustomClaim struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "default_key"
	}
	return secretKey
}

func comparePassword(hashedPass string, plainPassword []byte) bool {
	byteHash := []byte(hashedPass)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func hashAndSalt(password []byte) (string, error) {
	cost := 6
	hash, err := bcrypt.GenerateFromPassword(password, cost)

	return string(hash), err
}

func (s *AuthServiceImpl) Login(username string, password string) (dto.LoginResponse, error) {
	log.Printf("Start login service")
	user := s.userRepository.FindByUsername(username)

	if (user == dto.User{}) {
		log.Printf("Error] Login service: user not found")
		return dto.LoginResponse{}, errors.New("user not found")
	}

	log.Printf("Login service: compare password")
	comparedPassword := comparePassword(user.Password, []byte(password))

	if !comparedPassword {
		log.Printf("Error] Login service: username/password mismatch")
		return dto.LoginResponse{}, errors.New("username/password mismatch")
	}

	tokenExpire := time.Now().Add(time.Hour * 5)

	claims := &jwtCustomClaim{
		UserID:   user.ID,
		Username: user.Username,
		Name:     user.Name,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpire),
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	log.Printf("Login service: signing token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		log.Printf("Error] Login service: error signing token")
		return dto.LoginResponse{}, err
	}

	res := dto.LoginResponse{
		Token:       t,
		Username:    user.Username,
		Name:        user.Name,
		Role:        user.Role,
		TokenExpire: tokenExpire,
	}

	log.Printf("Login service: Login service finish")
	return res, nil
}

func (s *AuthServiceImpl) Register(username string, name string, password string, role string) error {
	log.Printf("Start register service ...")
	isUsernameExist := s.userRepository.CheckUsernameExist(username)

	if isUsernameExist {
		log.Printf("Register service: username already exists")
		return errors.New("username already exists")
	}

	hashedPassword, errHash := hashAndSalt([]byte(password))

	if errHash != nil {
		log.Printf("Register service: fail to hash password: %s", errHash.Error())
		return errHash
	}

	var newUser model.User

	newUser.Username = username
	newUser.Name = name
	newUser.Password = hashedPassword
	newUser.Role = role
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	errSave := s.userRepository.SaveUser(&newUser)

	if errSave != nil {
		log.Printf("Register service: fail to save user: %s", errSave.Error())
		return errSave
	}

	return nil
}

func (s *AuthServiceImpl) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}

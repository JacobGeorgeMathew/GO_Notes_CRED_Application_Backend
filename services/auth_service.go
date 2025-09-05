// package services

// import (
// 	"errors"
// 	"time"

// 	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/models"
// 	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/repository"
// 	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/utils"

// 	"golang.org/x/crypto/bcrypt"
// )

// type AuthService interface {
// 	Register(req *models.RegisterRequest) (*models.AuthResponse, error)
// 	Login(req *models.LoginRequest) (*models.AuthResponse, error)
// }

// type authService struct {
// 	userRepo repository.UserRepository
// }

// func NewAuthService(userRepo repository.UserRepository) AuthService {
// 	return &authService{userRepo: userRepo}
// }

// func (s *authService) Register(req *models.RegisterRequest) (*models.AuthResponse, error) {
// 	// Check if user already exists
// 	existingUser, _ := s.userRepo.GetUserByEmail(req.Email)
// 	if existingUser != nil {
// 		return nil, errors.New("user already exists")
// 	}

// 	// Hash password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create user
// 	user := &models.User{
// 		Username:  req.Username,
// 		Email:     req.Email,
// 		Password:  string(hashedPassword),
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	err = s.userRepo.CreateUser(user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Generate JWT token
// 	token, err := utils.GenerateJWT(user.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Remove password from response
// 	user.Password = ""

// 	return &models.AuthResponse{
// 		Token: token,
// 		User:  *user,
// 	}, nil
// }

// func (s *authService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
// 	// Get user by email
// 	user, err := s.userRepo.GetUserByEmail(req.Email)
// 	if err != nil {
// 		return nil, errors.New("invalid credentials")
// 	}

// 	// Check password
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
// 	if err != nil {
// 		return nil, errors.New("invalid credentials")
// 	}

// 	// Generate JWT token
// 	token, err := utils.GenerateJWT(user.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Remove password from response
// 	user.Password = ""

// 	return &models.AuthResponse{
// 		Token: token,
// 		User:  *user,
// 	}, nil
// }

// 3. UPDATE services/auth_service.go - Add GetUserByID method
package services

import (
	"errors"
	"time"

	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/models"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/repository"
	"github.com/JacobGeorgeMathew/GO_Notes_CRED_Application_Backend/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req *models.RegisterRequest) (*models.AuthResponse, error)
	Login(req *models.LoginRequest) (*models.AuthResponse, error)
	GetUserByID(userID int) (*models.User, error) // 🔥 NEW METHOD
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(req *models.RegisterRequest) (*models.AuthResponse, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.GetUserByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	// Remove password from response
	user.Password = ""

	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *authService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	// Remove password from response
	user.Password = ""

	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

// 🔥 NEW: Get user by ID for /me endpoint
func (s *authService) GetUserByID(userID int) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	
	// Remove password from response
	user.Password = ""
	return user, nil
}

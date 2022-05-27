package api

import (
	"context"
	"crypto/rand"
	pbPost "dislinkt/common/proto/post_service"
	pb "dislinkt/common/proto/user_service"
	pbUser "dislinkt/common/proto/user_service"
	"dislinkt/user_service/application"
	"dislinkt/user_service/auth"
	"dislinkt/user_service/domain"
	"encoding/hex"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"unicode"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service     *application.UserService
	mailService *application.MailService
	jwtManager  *auth.JWTManager
	userClient  pbUser.UserServiceClient
	postClient  pbPost.PostServiceClient
	validate    *validator.Validate
}

func NewUserHandler(service *application.UserService, mailService *application.MailService, jwtManager *auth.JWTManager, userClient pbUser.UserServiceClient,
	postClient pbPost.PostServiceClient) *UserHandler {
	return &UserHandler{
		service:     service,
		mailService: mailService,
		jwtManager:  jwtManager,
		userClient:  userClient,
		postClient:  postClient,
		validate:    domain.NewUserValidator(),
	}
}

func (handler *UserHandler) FindByUsername(ctx context.Context, request *pb.FindByUsernameRequest) (*pb.FindByUsernameResponse, error) {

	User, err := handler.service.Find(request.Username)
	if err != nil {
		return nil, err
	}
	UserPb := mapUserToPb(User)
	response := &pb.FindByUsernameResponse{
		User: UserPb,
	}
	return response, err
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	userId := request.Id
	User, err := handler.service.Get(userId)
	if err != nil {
		return nil, err
	}
	UserPb := mapUserToPb(User)
	response := &pb.GetResponse{
		User: UserPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, User := range Users {
		current := mapUserToPb(User)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) GetAllPublic(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, User := range Users {
		if !User.Private {
			current := mapUserToPb(User)
			response.Users = append(response.Users, current)
		}
	}
	return response, nil
}

func (handler *UserHandler) GetAllUsernames(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllUsernamesResponse, error) {
	Users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllUsernamesResponse{
		Usernames: []string{},
	}
	for _, User := range Users {
		response.Usernames = append(response.Usernames, User.Username)
	}
	return response, nil
}

func (handler UserHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.User.Password), bcrypt.DefaultCost)
	user := mapPbToUser(request.User)

	if !isValid(request.User.Password) {
		return nil, status.Errorf(codes.InvalidArgument, "password failed: %v", err)
	}

	if !user.DateOfBirth.Before(time.Now()) {
		return nil, status.Errorf(codes.InvalidArgument, "date of birth failed: %v", err)
	}

	user.HashedPassword = string(hashedPassword)
	if err := handler.validate.Struct(user); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	err = handler.service.Create(user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		User: mapUserToPb(user),
	}, nil
}

func (handler UserHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	user := mapPbToUser(request.User)
	userId := request.Id
	oldUser, _ := handler.service.Get(userId)
	user.HashedPassword = oldUser.HashedPassword
	user.Role = oldUser.Role
	user.Activated = oldUser.Activated
	user.Private = oldUser.Private
	if err := handler.validate.Struct(user); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	err := handler.service.Update(userId, user)
	if err != nil {
		return nil, err
	}

	if oldUser.FirstName != user.FirstName || oldUser.LastName != user.LastName {
		_, err = handler.postClient.UpdateUser(context.Background(), &pbPost.UpdateUserRequest{
			User: &pbPost.User{
				Username:  user.Username,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		})
		if err != nil {
			handler.service.Update(userId, oldUser)
			return nil, err
		}
	}

	return &pb.UpdateResponse{
		User: mapUserToPb(user),
	}, nil
}

func (handler *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := handler.service.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	if user.Activated == false {
		return nil, status.Errorf(codes.Internal, "You need to activate account first!")
	}
	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := handler.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler UserHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	checkUser, err := handler.service.Find(request.User.Username)
	if checkUser != nil {
		return nil, status.Errorf(codes.Internal, "username already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.User.Password), bcrypt.DefaultCost)
	user := mapPbToUser(request.User)

	if !isValid(request.User.Password) {
		return nil, status.Errorf(codes.InvalidArgument, "password failed: %v", err)
	}

	if !user.DateOfBirth.Before(time.Now()) {
		return nil, status.Errorf(codes.InvalidArgument, "date of birth failed: %v", err)
	}

	user.HashedPassword = string(hashedPassword)
	user.Token = GenerateSecureToken(32)
	handler.mailService.SendActivationEmail(user.Token, "http://localhost:8000/activate/")
	if err := handler.validate.Struct(user); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	err = handler.service.Create(user)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		User: mapUserToPb(user),
	}, nil
}

func (handler UserHandler) ActivateAccount(ctx context.Context, request *pb.ActivateRequest) (*pb.ActivateResponse, error) {
	return &pb.ActivateResponse{
		User: mapUserToPb(handler.service.ActivateAccount(request.Token)),
	}, nil
}

func (handler UserHandler) PasswordlessLoginDemand(ctx context.Context, request *pb.PasswordlessLoginDemandRequest) (*pb.PasswordlessLoginDemandResponse, error) {
	user, _ := handler.service.Find(request.Username)
	user.PasswordToken = GenerateSecureToken(32)
	handler.service.Update(user.Id.Hex(), user)
	handler.mailService.SendActivationEmail(user.PasswordToken, "http://localhost:4200/redirect/")
	return &pb.PasswordlessLoginDemandResponse{
		Email: user.Email,
	}, nil
}

func (handler UserHandler) PasswordlessLogin(ctx context.Context, request *pb.PasswordlesLoginRequest) (*pb.LoginResponse, error) {
	user, _ := handler.service.PasswordlessLogin(request.Token)
	if user.PasswordToken == "" {
		return nil, status.Errorf(codes.Internal, "Passwordless login request doesn't exist!")
	}
	token, err := handler.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	user.PasswordToken = ""
	handler.service.Update(user.Id.Hex(), user)

	return &pb.LoginResponse{AccessToken: token}, nil
}

func (handler UserHandler) RecoverAccountDemand(ctx context.Context, request *pb.RecoverAccountDemandRequest) (*pb.RecoverAccountDemandResponse, error) {
	user, err := handler.service.FindByEmail(request.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "no account with said email")
	}
	user.RecoveryToken = GenerateSecureToken(32)
	handler.service.Update(user.Id.Hex(), user)
	//TODO
	//front da se pogodi
	handler.mailService.SendActivationEmail(user.RecoveryToken, "http://localhost:4200/recover/")
	return &pb.RecoverAccountDemandResponse{}, nil
}

func (handler UserHandler) RecoverAccount(ctx context.Context, request *pb.RecoverAccountRequest) (*pb.RecoverAccountResponse, error) {
	_, err := handler.service.RecoverAccount(request.Token, request.NewPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "recovery account error")
	}
	return &pb.RecoverAccountResponse{}, nil
}

func (handler UserHandler) ChangePassword(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	if !isValid(request.NewPassword) {
		return nil, status.Errorf(codes.InvalidArgument, "new password inadequate")
	}
	err := handler.service.ChangePassword(request.Username, request.NewPassword, request.OldPassword)
	return &pb.ChangePasswordResponse{}, err
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func isValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 4 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

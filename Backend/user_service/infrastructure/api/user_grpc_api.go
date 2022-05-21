package api

import (
	"context"
	"crypto/rand"
	pb "dislinkt/common/proto/user_service"
	pbUser "dislinkt/common/proto/user_service"
	"dislinkt/user_service/application"
	"dislinkt/user_service/auth"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service     *application.UserService
	mailService *application.MailService
	jwtManager  *auth.JWTManager
	userClient  pbUser.UserServiceClient
}

func NewUserHandler(service *application.UserService, mailService *application.MailService, jwtManager *auth.JWTManager, userClient pbUser.UserServiceClient) *UserHandler {
	return &UserHandler{
		service:     service,
		mailService: mailService,
		jwtManager:  jwtManager,
		userClient:  userClient,
	}
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

func (handler UserHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.User.Password), bcrypt.DefaultCost)
	user := mapPbToUser(request.User)
	user.HashedPassword = string(hashedPassword)
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
	err := handler.service.Update(userId, user)
	if err != nil {
		return nil, err
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
	user.HashedPassword = string(hashedPassword)
	user.Token = GenerateSecureToken(32)
	handler.mailService.SendActivationEmail(user.Token, "http://localhost:8000/activate/")
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
	handler.mailService.SendActivationEmail(user.PasswordToken, "http://localhost:8000/login/")
	return &pb.PasswordlessLoginDemandResponse{
		Email: user.Email,
	}, nil
}

func (handler UserHandler) PasswordlessLogin(ctx context.Context, request *pb.PasswordlesLoginRequest) (*pb.LoginResponse, error) {
	user, _ := handler.service.PasswordlessLogin(request.Token)
	token, err := handler.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	user.PasswordToken = ""
	handler.service.Update(user.Id.Hex(), user)

	return &pb.LoginResponse{AccessToken: token}, nil
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

package controller

import (
	"Identity/internal/model"
	userapiv1 "Identity/internal/proto/identity/userapi/v1"
	"Identity/internal/service"
	"Identity/pkg/jwt"
	"Identity/pkg/limiter"
	"context"
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/casbin/casbin"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	iplimiter      *limiter.IPLimiter
	enforcer       *casbin.Enforcer
	jwt            *jwt.JwtToken
	userService    service.UserService
	sessionService service.SessionService
	userapiv1.UnimplementedUserServiceServer
}

var _ userapiv1.UserServiceServer = (*UserServiceServer)(nil)

func (s *UserServiceServer) CreateUser(ctx context.Context, req *userapiv1.CreateUserRequest) (*userapiv1.CreateUserResponse, error) {
	peer, ok := peer.FromContext(ctx)
	if ok {
		clientIP := peer.Addr.String()
		lim := s.iplimiter.Limiter(clientIP)
		if !lim.Allow() {
			return nil, status.Error(codes.Code(code.Code_FAILED_PRECONDITION), "too many requests")
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	token, found := md["authorization"]
	if !found || len(token) != 1 {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	tc, err := s.jwt.ExtractClaims(model.JwtToken(token[0]))
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	err = s.sessionService.CheckSession(ctx, tc.ID)
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "log in first")
	}

	s.enforcer.LoadPolicy()
	if !s.enforcer.Enforce(tc.Username, "users", "create") {
		return nil, status.Error(codes.Code(code.Code_PERMISSION_DENIED), "not allowed")
	}

	ru := model.RawUser{
		UUN:      req.Uun,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     req.Role,
	}

	err = s.userService.Create(ctx, ru)
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_INTERNAL), "something went wrong")
	}

	return nil, status.Error(codes.Code(code.Code_OK), "created")
}

func (s *UserServiceServer) UserByID(ctx context.Context, req *userapiv1.UserByIDRequest) (*userapiv1.UserByIDResponse, error) {
	peer, ok := peer.FromContext(ctx)
	if ok {
		clientIP := peer.Addr.String()
		lim := s.iplimiter.Limiter(clientIP)
		if !lim.Allow() {
			return nil, status.Error(codes.Code(code.Code_FAILED_PRECONDITION), "too many requests")
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	token, found := md["authorization"]
	if !found || len(token) != 1 {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	tc, err := s.jwt.ExtractClaims(model.JwtToken(token[0]))
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	err = s.sessionService.CheckSession(ctx, tc.ID)
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "log in first")
	}

	s.enforcer.LoadPolicy()
	if !s.enforcer.Enforce(tc.Username, "users", "read") {
		return nil, status.Error(codes.Code(code.Code_PERMISSION_DENIED), "not allowed")
	}

	sId, _ := snowflake.ParseString(req.Id)
	userInfo, err := s.userService.ReadByID(ctx, model.ID(sId))
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_INTERNAL), "something went wrong")
	}

	return &userapiv1.UserByIDResponse{
		Id:           fmt.Sprint(userInfo.ID),
		Uun:          userInfo.UUN,
		Username:     userInfo.Username,
		Email:        userInfo.Email,
		CreationTime: userInfo.Created_at.String(),
		Role:         userInfo.Role,
		Status:       userInfo.Status,
	}, status.Error(codes.Code(code.Code_OK), "read")
}

func (s *UserServiceServer) ChangePermission(ctx context.Context, req *userapiv1.ChangePermissionRequest) (*userapiv1.ChangePermissionResponse, error) {
	peer, ok := peer.FromContext(ctx)
	if ok {
		clientIP := peer.Addr.String()
		lim := s.iplimiter.Limiter(clientIP)
		if !lim.Allow() {
			return nil, status.Error(codes.Code(code.Code_FAILED_PRECONDITION), "too many requests")
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	token, found := md["authorization"]
	if !found || len(token) != 1 {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	tc, err := s.jwt.ExtractClaims(model.JwtToken(token[0]))
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	err = s.sessionService.CheckSession(ctx, tc.ID)
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "log in first")
	}

	s.enforcer.LoadPolicy()
	if !s.enforcer.Enforce(tc.Username, "permissions", "update") {
		return nil, status.Error(codes.Code(code.Code_PERMISSION_DENIED), "not allowed")
	}

	s.userService.ChangePermission(ctx, req.Username, req.NewRole)

	return nil, status.Error(codes.Code(code.Code_OK), "role changed, permissions affected")
}

func (s *UserServiceServer) ChangeStatus(ctx context.Context, req *userapiv1.ChangeStatusRequest) (*userapiv1.ChangeStatusResponse, error) {
	peer, ok := peer.FromContext(ctx)
	if ok {
		clientIP := peer.Addr.String()
		lim := s.iplimiter.Limiter(clientIP)
		if !lim.Allow() {
			return nil, status.Error(codes.Code(code.Code_FAILED_PRECONDITION), "too many requests")
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	token, found := md["authorization"]
	if !found || len(token) != 1 {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	tc, err := s.jwt.ExtractClaims(model.JwtToken(token[0]))
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	err = s.sessionService.CheckSession(ctx, tc.ID)
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "log in first")
	}

	s.enforcer.LoadPolicy()
	if !s.enforcer.Enforce(tc.Username, "status", "update") {
		return nil, status.Error(codes.Code(code.Code_PERMISSION_DENIED), "not allowed")
	}

	if req.Status != "suspend" && req.Status != "active" {
		return nil, status.Error(codes.Code(code.Code_INVALID_ARGUMENT), "bad status")
	}

	sId, _ := snowflake.ParseString(req.Id)
	err = s.userService.UpdateByID(ctx, model.UserInfo{
		ID:     model.ID(sId),
		Status: req.Status,
	})
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_INTERNAL), "something went wrong")
	}

	return nil, status.Error(codes.Code(code.Code_OK), "status updated successfully")
}

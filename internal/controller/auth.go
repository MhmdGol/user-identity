package controller

import (
	"Identity/internal/model"
	authapiv1 "Identity/internal/proto/identity/authapi/v1"
	"Identity/internal/service"
	"Identity/pkg/limiter"
	"context"

	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	authService service.AuthService
	iplimiter   *limiter.IPLimiter
	authapiv1.UnimplementedAuthServiceServer
}

// var _ authapiv1.AuthServiceServer = (*AuthServiceServer)(nil)

func (s *AuthServiceServer) Login(ctx context.Context, req *authapiv1.LoginRequest) (*authapiv1.LoginResponse, error) {
	peer, ok := peer.FromContext(ctx)
	if ok {
		clientIP := peer.Addr.String()
		lim := s.iplimiter.Limiter(clientIP)
		if !lim.Allow() {
			return nil, status.Error(codes.Code(code.Code_FAILED_PRECONDITION), "too many requests")
		}
	}

	l := model.LoginInfo{
		Username: req.Username,
		Password: req.Password,
	}

	t, _ := s.authService.Login(ctx, l)

	return &authapiv1.LoginResponse{
		Token: string(t),
	}, status.Error(codes.Code(code.Code_OK), "logged in")
}

func (s *AuthServiceServer) Logout(ctx context.Context, req *authapiv1.LogoutRequest) (*authapiv1.LogoutResponse, error) {
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

	err := s.authService.Logout(ctx, model.JwtToken(token[0]))
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_INTERNAL), "something went wrong")
	}

	return nil, status.Error(codes.Code(code.Code_OK), "logged out")
}

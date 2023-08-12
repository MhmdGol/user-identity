package controller

import (
	"Identity/internal/model"
	authapiv1 "Identity/internal/proto/identity/authapi/v1"
	"Identity/internal/service"
	"Identity/pkg/limiter"
	"context"

	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
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
	l := model.LoginInfo{
		Username: req.Username,
		Password: req.Password,
	}

	peer, ok := peer.FromContext(ctx)
	if ok {
		clientIP := peer.Addr.String()
		lim := s.iplimiter.Limiter(clientIP)
		if !lim.Allow() {
			return nil, status.Error(codes.Code(code.Code_FAILED_PRECONDITION), "too many requests")
		}
	}

	t, _ := s.authService.Login(ctx, l)

	return &authapiv1.LoginResponse{
		Token: string(t),
	}, nil

}

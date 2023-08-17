package controller

import (
	"Identity/internal/model"
	authapiv1 "Identity/internal/proto/identity/authapi/v1"
	"Identity/internal/service"
	"Identity/pkg/jwt"
	"Identity/pkg/limiter"
	"context"

	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	authService    service.AuthService
	sessionService service.SessionService
	iplimiter      *limiter.IPLimiter
	jwt            *jwt.JwtToken
	authapiv1.UnimplementedAuthServiceServer
}

var _ authapiv1.AuthServiceServer = (*AuthServiceServer)(nil)

func NewAuthServiceServer(
	as service.AuthService,
	ss service.SessionService,
	il *limiter.IPLimiter,
	jt *jwt.JwtToken,
) *AuthServiceServer {
	return &AuthServiceServer{
		authService:    as,
		sessionService: ss,
		iplimiter:      il,
		jwt:            jt,
	}
}

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

	tc, err := s.jwt.ExtractClaims(model.JwtToken(token[0]))
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_UNAUTHENTICATED), "missing token")
	}

	err = s.authService.Logout(ctx, tc.ID)
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_INTERNAL), "something went wrong")
	}

	return nil, status.Error(codes.Code(code.Code_OK), "logged out")
}

func (s *AuthServiceServer) ChangePassword(ctx context.Context, req *authapiv1.ChangePasswordRequest) (*authapiv1.ChangePasswordResponse, error) {
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

	err = s.authService.UpdatePassword(ctx, model.UpdatePassword{
		Username: req.Username,
		OldPass:  req.OldPassword,
		NewPass:  req.NewPassword,
	})
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_INTERNAL), "something went wrong")
	}

	return nil, status.Error(codes.Code(code.Code_OK), "password changed successfully")
}

func (s *AuthServiceServer) PasswordRecovery(ctx context.Context, req *authapiv1.PasswordRecoveryRequest) (*authapiv1.PasswordRecoveryResponse, error) {
	peer, ok := peer.FromContext(ctx)
	if ok {
		clientIP := peer.Addr.String()
		lim := s.iplimiter.Limiter(clientIP)
		if !lim.Allow() {
			return nil, status.Error(codes.Code(code.Code_FAILED_PRECONDITION), "too many requests")
		}
	}

	err := s.authService.PasswordRecovery(ctx, req.Username)
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_FAILED_PRECONDITION), "something went wrong")
	}

	return nil, status.Error(codes.Code(code.Code_OK), "password recovery email sent")

}

func (s *AuthServiceServer) ResetPassword(ctx context.Context, req *authapiv1.ResetPasswordRequest) (*authapiv1.ResetPasswordResponse, error) {
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

	err = s.authService.ResetPassword(ctx, tc.Username, req.NewPassword)
	if err != nil {
		return nil, status.Error(codes.Code(code.Code_FAILED_PRECONDITION), "something went wrong")
	}

	return nil, status.Error(codes.Code(code.Code_OK), "password been reset")
}

func (s *AuthServiceServer) TwoFactorAuthentication(ctx context.Context, req *authapiv1.TwoFactorAuthenticationRequest) (*authapiv1.TwoFactorAuthenticationResponse, error) {
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

	t, _ := s.authService.TwoFactorAuth(ctx, l, req.Code)

	return &authapiv1.TwoFactorAuthenticationResponse{
		Token: string(t),
	}, status.Error(codes.Code(code.Code_OK), "logged in")
}

package handler

import (
	"context"

	"txn-engine-phase-2/admin-service/internal/service"
	adminpb "txn-engine-phase-2/proto/gen/go/admin"
)

type GRPCHandler struct {
	adminpb.UnimplementedAuthServiceServer
	svc *service.AuthService
}

func NewGRPCHandler(s *service.AuthService) *GRPCHandler {
	return &GRPCHandler{svc: s}
}

func (h *GRPCHandler) RegisterAdmin(ctx context.Context, req *adminpb.RegisterRequest) (*adminpb.RegisterResponse, error) {
	id, err := h.svc.Register(ctx, req.Name, req.Email, req.Mobile, req.Password, req.Role)
	if err != nil {
		return &adminpb.RegisterResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &adminpb.RegisterResponse{
		Success: true,
		Message: "admin created",
		Id:      id,
	}, nil
}

func (h *GRPCHandler) LoginAdmin(ctx context.Context, req *adminpb.LoginRequest) (*adminpb.LoginResponse, error) {
	otp, err := h.svc.Login(ctx, req.Mobile)
	if err != nil {
		return &adminpb.LoginResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &adminpb.LoginResponse{
		Success: true,
		Message: "OTP sent",
		Otp:     otp,
	}, nil
}

func (h *GRPCHandler) VerifyOtp(ctx context.Context, req *adminpb.VerifyOtpRequest) (*adminpb.VerifyOtpResponse, error) {
	at, rt, id, err := h.svc.VerifyOTP(ctx, req.Mobile, req.Otp)
	if err != nil {
		return &adminpb.VerifyOtpResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &adminpb.VerifyOtpResponse{
		Success:      true,
		Message:      "Login successful",
		Token:        at,
		RefreshToken: rt,
		Id:           id,
	}, nil
}

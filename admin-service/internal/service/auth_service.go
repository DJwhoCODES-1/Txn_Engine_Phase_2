package service

import (
	"context"
	"errors"
	"time"

	"txn-engine-phase-2/admin-service/internal/model"
	"txn-engine-phase-2/admin-service/internal/repository"
	"txn-engine-phase-2/admin-service/internal/utils"
)

type AuthService struct {
	repo      *repository.AdminRepository
	jwtSecret string
}

func NewAuthService(r *repository.AdminRepository, secret string) *AuthService {
	return &AuthService{repo: r, jwtSecret: secret}
}

func (s *AuthService) Register(ctx context.Context, name, email, mobile, password, role string) (string, error) {
	// check if already exists
	_, err := s.repo.FindByMobile(ctx, mobile)
	if err == nil {
		return "", errors.New("admin already exists")
	}

	admin := &model.Admin{
		Name:      name,
		Email:     email,
		Mobile:    mobile,
		Password:  password, // ⚠️ hash later
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := s.repo.Create(ctx, admin)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *AuthService) Login(ctx context.Context, mobile string) (string, error) {
	_, err := s.repo.FindByMobile(ctx, mobile)
	if err != nil {
		return "", errors.New("admin not found")
	}

	otp := utils.GenerateOTP()

	if err := s.repo.UpdateOTP(ctx, mobile, otp); err != nil {
		return "", err
	}

	return otp, nil
}

func (s *AuthService) VerifyOTP(ctx context.Context, mobile, otp string) (string, string, string, error) {
	admin, err := s.repo.FindByMobile(ctx, mobile)
	if err != nil {
		return "", "", "", errors.New("admin not found")
	}

	if admin.OTP != otp {
		return "", "", "", errors.New("invalid otp")
	}

	if time.Since(admin.OTPTimestamp) > 5*time.Minute {
		return "", "", "", errors.New("otp expired")
	}

	_ = s.repo.MarkOTPVerified(ctx, mobile)

	at, rt, err := utils.GenerateTokens(s.jwtSecret, admin.ID.Hex())
	if err != nil {
		return "", "", "", err
	}

	return at, rt, admin.ID.Hex(), nil
}

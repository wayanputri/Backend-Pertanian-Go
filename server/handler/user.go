package handler

import (
	"backend/app/jwt_middle"
	"backend/pb"
	"backend/pkg/models"
	"context"
	"fmt"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	email, password, id, role, err := s.GetProvider().GetUserByEmail(req.GetEmail())
	if err != nil {
		return &pb.LoginUserResponse{
			Token:   "",
			Status:  500,
			Fail:    true,
			Message: err.Error(),
		}, err
	}
	fmt.Println("email", req.GetEmail(), ",", email)
	fmt.Println("password", req.GetPassword(), ",", password)

	// isAdmin := false
	// if role == "admin" {
	// 	isAdmin = true
	// }

	// Verifikasi kredensial user (dalam contoh ini hanya pengecekan sederhana)
	if req.GetEmail() == email && req.GetPassword() == password {
		// Jika username dan password valid, buatkan JWT
		token, err := jwt_middle.CreateToken(id, role) // ID user sebagai contoh
		if err != nil {
			return &pb.LoginUserResponse{
				Token:   "",
				Status:  500,
				Fail:    true,
				Message: err.Error(),
			}, fmt.Errorf("failed to generate token: %v", err)
		}
		return &pb.LoginUserResponse{
			Token:   token,
			Status:  200,
			Fail:    false,
			Message: "login successfully",
		}, nil
	}

	return &pb.LoginUserResponse{
		Token:   "",
		Status:  400,
		Fail:    true,
		Message: "invalid credentials",
	}, fmt.Errorf("invalid credentials")
}

func (s *Server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {

	// Simulasikan penyimpanan user ke database
	// (In real scenario, you would store the user in a database)

	if req.Email == "" {
		return &pb.RegisterUserResponse{
			Id:      0,
			Status:  400,
			Fail:    true,
			Message: "failed register user, email wajib diisi",
			Token:   "",
		}, fmt.Errorf("failed register user, email wajib diisi")
	}

	if req.Password == "" {
		return &pb.RegisterUserResponse{
			Id:      0,
			Status:  400,
			Fail:    true,
			Message: "failed register user, password wajib diisi",
			Token:   "",
		}, fmt.Errorf("failed register user, Password wajib diisi")
	}
	if req.Name == "" {
		return &pb.RegisterUserResponse{
			Id:      0,
			Status:  400,
			Fail:    true,
			Message: "failed register user, nama wajib diisi",
			Token:   "",
		}, fmt.Errorf("failed register user, nama wajib diisi")
	}

	data := &models.User{
		Name:            req.Name,
		Email:           req.Email,
		Password:        req.Password,
		Role:            "user",
		LandArea:        req.LandArea,
		TypeOfLivestock: req.TypeOfLivestock,
		Address:         req.Address,
	}
	id, err := s.GetProvider().CreateUser(data)
	if err != nil {
		return nil, err
	}

	// Buat JWT token
	token, err := jwt_middle.CreateToken(id, "user")
	if err != nil {
		return &pb.RegisterUserResponse{
			Id:      0,
			Status:  500,
			Fail:    true,
			Message: err.Error(),
			Token:   "",
		}, fmt.Errorf("failed to generate token: %v", err)
	}

	return &pb.RegisterUserResponse{
		Id:      int64(id),
		Status:  200,
		Fail:    false,
		Message: "Register Successfully",
		Token:   token,
	}, nil
}

func (s *Server) GetUserAll(ctx context.Context, _ *pb.Empty) (*pb.GetUserResponse, error) {
	fmt.Println("GetUser Start")
	fmt.Println("ctx", ctx)
	fmt.Println("ctx println")

	userID, role, err := jwt_middle.ExtractTokenUserId(ctx)
	if err != nil {
		return &pb.GetUserResponse{
			Status:  400,
			Fail:    true,
			Message: err.Error(),
		}, err
	}
	if role != "admin" {
		return &pb.GetUserResponse{
			Status:  400,
			Fail:    true,
			Message: "hanya admin yg bisa melihat seluruh user",
		}, err

	}

	fmt.Printf("User ID: %d Role:%s\n", userID, role)
	data, err := s.GetProvider().GetUserAll()
	if err != nil {
		return &pb.GetUserResponse{
			Status:  500,
			Fail:    true,
			Message: "failed get User all",
		}, err
	}

	var userData []*pb.User
	for _, v := range data {
		var user pb.User
		user.Id = int64(v.ID)
		user.Email = v.Email
		user.Name = v.Name
		user.Role = v.Role
		user.Address = v.Address
		user.TypeOfLivestock = v.TypeOfLivestock
		user.LandArea = v.LandArea

		userData = append(userData, &user)
	}
	return &pb.GetUserResponse{
		Status:  200,
		Fail:    false,
		Message: "success get user all",
		Data:    userData,
	}, nil
}

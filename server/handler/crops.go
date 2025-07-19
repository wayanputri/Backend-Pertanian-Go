package handler

import (
	"backend/app/jwt_middle"
	"backend/pb"
	"backend/pkg/models"
	"context"
	"errors"
	"fmt"
	"time"
)

func (s *Server) CreateCrops(ctx context.Context, req *pb.Crops) (*pb.GlobalResponse, error) {

	if req.LuasLahan == "" {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "luas lahan tidak boleh kosong",
		}, errors.New("luas lahan tidak boleh kosong")
	}
	if req.Jenis == "" {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "spesies tidak boleh kosong",
		}, errors.New("spesies tidak boleh kosong")
	}

	userID, _, err := jwt_middle.ExtractTokenUserId(ctx)
	if err != nil {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: err.Error(),
		}, err
	}

	idFarmArea, err := s.provider.InsertArea(&models.FarmArea{
		Nama: "Crops Area",
		Type: "Crops",
		Wide: req.GetLuasLahan(),
	})
	if err != nil {
		return &pb.GlobalResponse{
			Status:  500,
			Failed:  true,
			Message: "failed insert data farm area",
		}, err
	}
	now := time.Now()
	err = s.provider.InsertCrops(&models.Crops{
		Spesies:            req.GetJenis(),
		Age:                req.GetUmur(),
		WaterNeeds:         req.GetKebutuhanAir(),
		FarmAreaID:         idFarmArea,
		PestControl:        req.GetPengendalianHama(),
		OrganicFertilizer:  req.GetPupukOrganik(),
		ChemicalFertilizer: req.GetPupukKimia(),
		UserID:             userID,
		Reference:          fmt.Sprintf("TIOA%d%03d", now.Unix(), now.Nanosecond()/1e6%1000),
	})
	if err != nil {
		return &pb.GlobalResponse{
			Status:  500,
			Failed:  true,
			Message: "failed insert data crops",
		}, err
	}

	return &pb.GlobalResponse{
		Status:  200,
		Failed:  false,
		Message: "success create data crops",
	}, nil
}

func (s *Server) GetCropsAll(ctx context.Context, _ *pb.Empty) (*pb.CropsResponse, error) {

	userID, role, err := jwt_middle.ExtractTokenUserId(ctx)
	if err != nil {
		return &pb.CropsResponse{
			Status:  400,
			Failed:  true,
			Message: err.Error(),
		}, err
	}

	data, err := s.provider.GetCropsAll(uint64(userID), role)
	if err != nil {
		return &pb.CropsResponse{
			Status:  500,
			Failed:  true,
			Message: err.Error(),
		}, err
	}
	return &pb.CropsResponse{
		Status:  200,
		Failed:  false,
		Message: "get crops successfully",
		Data:    data,
	}, nil
}

func (s *Server) GetCropsDetail(ctx context.Context, req *pb.CropsDetailRequest) (*pb.CropsDetailResponse, error) {
	if req.Reference == "" {
		return &pb.CropsDetailResponse{
			Status:  400,
			Failed:  true,
			Message: "reference harus diisi",
		}, errors.New("reference harus diisi")
	}

	userID, role, err := jwt_middle.ExtractTokenUserId(ctx)
	if err != nil {
		return &pb.CropsDetailResponse{
			Status:  400,
			Failed:  true,
			Message: err.Error(),
		}, err
	}

	data, err := s.provider.GetCropDetail(userID, role, req.Reference)
	if err != nil {
		return &pb.CropsDetailResponse{
			Status:  500,
			Failed:  true,
			Message: err.Error(),
		}, err
	}

	return &pb.CropsDetailResponse{
		Status:  200,
		Failed:  false,
		Message: "get crop detail successfully",
		Data:    data,
	}, nil

}

func (s *Server) UpdateCropsDetail(ctx context.Context, req *pb.Crops) (*pb.GlobalResponse, error) {
	if req.Id == 0 {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "id tidak boleh kosong",
		}, errors.New("id tidak boleh kosong")
	}

	userID, role, err := jwt_middle.ExtractTokenUserId(ctx)
	if err != nil {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: err.Error(),
		}, err
	}
	fmt.Println("request", req)
	err = s.provider.UpdateCrops(models.Crops{
		ID:                 int(req.GetId()),
		Spesies:            req.GetJenis(),
		Age:                req.GetUmur(),
		WaterNeeds:         req.GetKebutuhanAir(),
		PestControl:        req.GetPengendalianHama(),
		OrganicFertilizer:  req.GetPupukOrganik(),
		ChemicalFertilizer: req.GetPupukKimia(),
		UserID:             userID,
	}, userID, role, req.GetLuasLahan())
	if err != nil {
		return &pb.GlobalResponse{
			Status:  500,
			Failed:  true,
			Message: err.Error(),
		}, err
	}

	return &pb.GlobalResponse{
		Status:  200,
		Failed:  false,
		Message: "update crops successfully",
	}, err

}

func (s *Server) DeleteCropsDetail(ctx context.Context, req *pb.DeleteCropsRequest) (*pb.GlobalResponse, error) {
	if req.Id == 0 {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "id tidak boleh kosong",
		}, errors.New("id tidak boleh kosong")
	}
	userID, role, err := jwt_middle.ExtractTokenUserId(ctx)
	if err != nil {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: err.Error(),
		}, err
	}
	err = s.provider.DeleteCrops(int(req.GetId()), userID, role)
	if err != nil {
		return &pb.GlobalResponse{
			Status:  500,
			Failed:  true,
			Message: err.Error(),
		}, err
	}
	return &pb.GlobalResponse{
		Status:  200,
		Failed:  false,
		Message: "delete crops successfully",
	}, err
}

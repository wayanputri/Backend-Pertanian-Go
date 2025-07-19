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

func (s *Server) CreateLifeStock(ctx context.Context, req *pb.CreateLifeStocks) (*pb.GlobalResponse, error) {
	if req.JenisHewan == "" {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "jenis hewan tidak boleh kosong",
		}, errors.New("jenis hewan tidak boleh kosong")
	}
	if req.UmurHewan == "" {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "umur hewan tidak boleh kosong",
		}, errors.New("umur hewan tidak boleh kosong")
	}

	if req.Quantity == 0 {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "quantity tidak boleh kosong",
		}, errors.New("quantity tidak boleh kosong")
	}
	if req.KebutuhanPakan == "" {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "kebutuhan pakan tidak boleh kosong",
		}, errors.New("kebutuhan pakan tidak boleh kosong")
	}
	if req.PersentaseKesehatan == "" {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "persentase kesehatan tidak boleh kosong",
		}, errors.New("persentase kesehatan tidak boleh kosong")
	}
	if req.LuasKandang == "" {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "luas kandang tidak boleh kosong",
		}, errors.New("luas kandang tidak boleh kosong")
	}

	if req.JenisMakanan == "" {
		return &pb.GlobalResponse{
			Status:  400,
			Failed:  true,
			Message: "jenis makanan tidak boleh kosong",
		}, errors.New("jenis makanan tidak boleh kosong")
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
		Nama: "Life Stock Area",
		Type: "Life Stock",
		Wide: req.LuasKandang,
	})
	if err != nil {
		return &pb.GlobalResponse{
			Status:  500,
			Failed:  true,
			Message: "failed insert data farm area",
		}, err
	}
	now := time.Now()
	err = s.provider.InsertLifeStock(&models.Livestocks{
		Spesies:        req.GetJenisHewan(),
		Age:            req.GetUmurHewan(),
		HealthStatus:   req.GetPersentaseKesehatan(),
		KebutuhanMakan: req.GetKebutuhanPakan(),
		TypeOfFood:     req.GetJenisMakanan(),
		Quantity:       int(req.GetQuantity()),
		FarmAreaID:     idFarmArea,
		UserID:         userID,
		Reference:      fmt.Sprintf("JIOA%d%03d", now.Unix(), now.Nanosecond()/1e6%1000),
	})
	if err != nil {
		return &pb.GlobalResponse{
			Status:  500,
			Failed:  true,
			Message: "failed insert data life stock",
		}, err
	}

	return &pb.GlobalResponse{
		Status:  200,
		Failed:  false,
		Message: "success create data life stock",
	}, nil
}

func (s *Server) GetLifeStockAll(ctx context.Context, _ *pb.Empty) (*pb.LifeStockAllResponse, error) {

	userID, role, err := jwt_middle.ExtractTokenUserId(ctx)
	if err != nil {
		return &pb.LifeStockAllResponse{
			Status:  400,
			Failed:  true,
			Message: err.Error(),
		}, err
	}

	data, err := s.provider.GetLifeStockAll(uint64(userID), role)
	if err != nil {
		return &pb.LifeStockAllResponse{
			Status:  500,
			Failed:  true,
			Message: err.Error(),
		}, err
	}
	return &pb.LifeStockAllResponse{
		Status:  200,
		Failed:  false,
		Message: "get life stock all successfully",
		Data:    data,
	}, nil
}

func (s *Server) GetLifeStockDetail(ctx context.Context, req *pb.LifeStockDetailRequest) (*pb.LifeStockDetailResponse, error) {

	userID, role, err := jwt_middle.ExtractTokenUserId(ctx)
	if err != nil {
		return &pb.LifeStockDetailResponse{
			Status:  400,
			Failed:  true,
			Message: err.Error(),
		}, err
	}
	data, err := s.provider.GetLifeStockDetail(userID, role, req.Reference)
	if err != nil {
		return &pb.LifeStockDetailResponse{
			Status:  500,
			Failed:  true,
			Message: err.Error(),
		}, err
	}

	return &pb.LifeStockDetailResponse{
		Status:  200,
		Failed:  false,
		Message: "get life stock detail successfully",
		Data:    data,
	}, nil
}

func (s *Server) DeleteLifeStockDetail(ctx context.Context, req *pb.DeleteLifeStockRequest) (*pb.GlobalResponse, error) {
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
	err = s.provider.DeleteLifeStock(int(req.GetId()), userID, role)
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
		Message: "delete life stock successfully",
	}, err
}

func (s *Server) UpdateLifeStockDetail(ctx context.Context, req *pb.CreateLifeStocks) (*pb.GlobalResponse, error) {
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
	err = s.provider.UpdateLifeStock(models.Livestocks{
		ID:             int(req.Id),
		Spesies:        req.GetJenisHewan(),
		Age:            req.GetUmurHewan(),
		HealthStatus:   req.GetPersentaseKesehatan(),
		KebutuhanMakan: req.GetKebutuhanPakan(),
		TypeOfFood:     req.GetJenisMakanan(),
		Quantity:       int(req.GetQuantity()),
	}, userID, role, req.LuasKandang)
	if err != nil {
		return &pb.GlobalResponse{
			Status:  500,
			Failed:  true,
			Message: err.Error(),
		}, err
	}
	return &pb.GlobalResponse{
		Status:  200,
		Failed:  true,
		Message: "update life stock successfully",
	}, nil

}

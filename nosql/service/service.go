package service

import (
	"context"
	"fmt"
	"nosql/models"
	"nosql/pkg/helper"
	"nosql/pkg/logger"
	"nosql/storage"

	db "go.mongodb.org/mongo-driver/mongo"
)

type productService struct {
	logger  logger.Logger
	storage storage.StorageI
}

func NewProductService(db *db.Database, log logger.Logger) *productService {
	return &productService{
		storage: storage.NewProductStorage(db),
		logger:  log,
	}
}

func (ps *productService) Create(ctx context.Context, req *models.Product) (*models.CreateResponse, error) {
	fmt.Println("REQ", req)
	ID, err := ps.storage.Product().Create(req)
	fmt.Println("ID", ID)

	if err != nil {
		ps.logger.Error("error while creating product", logger.Error(err))
		return nil, helper.HandleError(ps.logger, err, "error while creating product", req)
	}
	return &models.CreateResponse{
		ID: ID,
	}, nil
}

func (ps *productService) Get(ctx context.Context, req *models.GetRequest) (*models.GetProductResponse, error) {

	response, err := ps.storage.Product().Get(req.ID)

	if err != nil {
		return nil, helper.HandleError(ps.logger, err, "error while getting product", req)
	}
	return &models.GetProductResponse{
		Product: response,
	}, nil
}

func (ps *productService) GetAll(ctx context.Context, req *models.GetAllProductsRequest) (*models.GetAllProductsResponse, error) {

	response, count, err := ps.storage.Product().GetAll(req.Page, req.Limit, req.Name)

	if err != nil {
		return nil, helper.HandleError(ps.logger, err, "error while getting product", req)
	}
	return &models.GetAllProductsResponse{
		Count:    count,
		Products: response,
	}, nil
}

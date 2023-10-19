package grpc

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/matheuslssilva/fullcycle-grpc-products-challenge/application/grpc/pb"
	"github.com/matheuslssilva/fullcycle-grpc-products-challenge/application/usecase"
	"github.com/matheuslssilva/fullcycle-grpc-products-challenge/domain/model"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.ProductCreation) (*pb.ProductCreatedResult, error) {
	value := stringToFloat64(in.Value)
	product, err := p.ProductUseCase.Create(in.Name, value)
	if err != nil {
		return &pb.ProductCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.ProductCreatedResult{
		Id:     product.ID,
		Status: "created",
	}, nil
}

func (p *ProductGrpcService) FindAll(ctx context.Context, in *pb.ProductFindAllRequest) (*pb.ProductFindAllResult, error) {
	products, err := p.ProductUseCase.FindAll()
	if err != nil {
		return &pb.ProductFindAllResult{
			Status: "cannot get data",
			Error:  err.Error(),
		}, err
	}

	finalResult := toFinalResult(products)

	return finalResult, nil
}

func stringToFloat64(v string) float64 {
	result, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func toFinalResult(data []*model.Product) *pb.ProductFindAllResult {
	result := pb.ProductFindAllResult{}

	for _, product := range data {
		result.Products = append(result.Products, &pb.ProductInfo{
			Id:    product.ID,
			Name:  product.Name,
			Value: fmt.Sprintf("%.2f", product.Value),
		})
	}

	return &result
}

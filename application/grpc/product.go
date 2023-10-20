package grpc

import (
	"context"
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

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	product, err := p.ProductUseCase.Create(in.Name, in.Description, in.Price)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products, err := p.ProductUseCase.FindAll()
	if err != nil {
		return nil, err
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

func toFinalResult(data []*model.Product) *pb.FindProductsResponse {
	result := pb.FindProductsResponse{}

	for _, product := range data {
		result.Products = append(result.Products, &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}

	return &result
}

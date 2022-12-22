package service

import (
	"context"
	"log"
	"strconv"
	"unary_grpc/helpers"
	productPb "unary_grpc/pb/product"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
	Client *firestore.Client
}

func (p *ProductService) GetProducts(ctx context.Context, empty *productPb.Empty) (*productPb.Products, error) {
	colName := helpers.GetEnv()

	var products []*productPb.Product

	iter := p.Client.Collection(colName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var product *productPb.Product
		err = doc.DataTo(&product)
		if err != nil {
			return nil, err
		}
		// log.Println(category["Id"].(int64))
		// log.Println(product)
		products = append(products, product)
	}
	// category := doc.Data()["ProductCategory"].(map[string]interface{})

	// var product *productPb.Product
	// result, err := p.Client.Collection(colName).Doc("2").Get(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// log.Println(product.Data())
	// log.Println(category["Id"].(int64))
	// result.DataTo(&product)
	// log.Println(product)

	log.Println(products)

	return nil, nil
}

func (p *ProductService) GetProduct(ctx context.Context, id *productPb.Id) (*productPb.Product, error) {
	colName := helpers.GetEnv()

	var product *productPb.Product
	idString := strconv.FormatInt(int64(id.GetId()), 10)

	result, err := p.Client.Collection(colName).Doc(idString).Get(ctx)
	if err != nil {
		return nil, err
	}

	result.DataTo(&product)

	return product, nil
}

func (p *ProductService) CreateProduct(ctx context.Context, product *productPb.Product) (*productPb.Id, error) {
	colName := helpers.GetEnv()

	var Response productPb.Id

	product = &productPb.Product{
		Id:    product.GetId(),
		Name:  product.GetName(),
		Price: product.GetPrice(),
		Stock: product.GetStock(),
		ProductCategory: &productPb.ProductCategory{
			Id:   product.ProductCategory.GetId(),
			Name: product.ProductCategory.GetName(),
		},
	}

	idString := strconv.FormatInt(int64(product.GetId()), 10)

	_, err := p.Client.Collection(colName).Doc(idString).Set(
		ctx,
		&product,
	)
	if err != nil {
		log.Printf("Failed adding a new product: %v", err)
		return nil, err
	}

	Response.Id = product.Id
	return &Response, nil
}

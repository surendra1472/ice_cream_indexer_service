package es

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"ic-indexer-service/app/model/request"
	"log"
	"strconv"
)

type IcecreamEsHelper interface {
	GenerateSearchQuery(context.Context, request.IcecreamFilter) (*elastic.BoolQuery, error)
}

func NewIcecreamEsHelper() IcecreamEsHelper {
	return &icecreamEsHelper{}
}

type icecreamEsHelper struct {
}

func (ie icecreamEsHelper) GenerateSearchQuery(ctx context.Context, params request.IcecreamFilter) (*elastic.BoolQuery, error) {

	query := elastic.NewBoolQuery()

	ie.queryWithId(ctx, query, params)
	ie.queryWithProductId(ctx, query, params)
	ie.queryWithName(ctx, query, params)

	printQuery(ctx, query)
	return query, nil
}

func printQuery(ctx context.Context, query *elastic.BoolQuery) {
	src, err := query.Source()
	if err != nil {
		log.Print(ctx, err.Error())
		return
	}
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		log.Fatal(ctx, err.Error())
		return
	}
	log.Print(ctx, "Search Query : ", string(data))
}

func (ie icecreamEsHelper) queryWithId(ctx context.Context, query *elastic.BoolQuery, params request.IcecreamFilter) {
	if params.Id != 0 {
		query.Must(elastic.NewTermQuery("id", strconv.FormatInt(params.Id, 10)))
	}
}

func (ie icecreamEsHelper) queryWithProductId(ctx context.Context, query *elastic.BoolQuery, params request.IcecreamFilter) {
	if params.ProductId != "" {
		query.Must(elastic.NewTermQuery("product_id", params.ProductId))
	}
}

func (ie icecreamEsHelper) queryWithName(ctx context.Context, query *elastic.BoolQuery, params request.IcecreamFilter) {
	if params.Name != "" {
		query.Must(elastic.NewTermQuery("name", params.Name))
	}
}

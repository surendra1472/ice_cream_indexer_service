package dataaccessor

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"ic-indexer-service/app/builder"
	"ic-indexer-service/app/config"
	"ic-indexer-service/app/helpers/es"
	"ic-indexer-service/app/model/bo"
	"ic-indexer-service/app/model/request"
	"ic-indexer-service/app/model/response"
	"log"
	"net/http"
	"strconv"
)

//go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=IceCreamDataAccessor"
type IceCreamDataAccessor interface {
	CreateOrReplaceIcecream(context.Context, bo.ESIcecream) error
	DeleteIcecream(context.Context, request.IcecreamDelete) error
	GetIcecreams(context.Context, request.IcecreamFilter) (response.BulkIcecreamIndexResponse, error)
	GetIcecreamByKey(context.Context, request.IcecreamIndexRequest) (bo.ESIcecream, error)
}

func NewIceCreamDataAccessor() IceCreamDataAccessor {
	return &icecreamDataAccessorService{
		index: config.GetConfig().ElasticSearch.Indices.IcecreamIndex,
	}
}

type icecreamDataAccessorService struct {
	index config.ESIndex
}

func (icdas icecreamDataAccessorService) GetIcecreamByKey(ctx context.Context, icecreamRequest request.IcecreamIndexRequest) (bo.ESIcecream, error) {

	var searchQuery *elastic.BoolQuery
	var err error
	var icecream bo.ESIcecream

	if icecreamRequest.Id.Set {
		searchQuery, err = es.NewIcecreamEsHelper().GenerateSearchQuery(ctx, request.IcecreamFilter{Id: *icecreamRequest.Id.Value})

		if err != nil {
			log.Fatal(ctx, "failed to generate search query, error", err.Error())
			return bo.ESIcecream{}, err
		}
		searchService := config.GetESConnection().Search().
			Index(icdas.index.IndexName).
			Query(searchQuery).
			Pretty(true)

		searchResult, err := searchService.Do(ctx)
		if err != nil {
			log.Print(ctx, err.Error())
			return bo.ESIcecream{}, err
		}
		for _, hit := range searchResult.Hits.Hits {
			bytes := []byte(*hit.Source)
			err = json.Unmarshal(bytes, &icecream)
			if err != nil {
				log.Print(ctx, err.Error())
				return bo.ESIcecream{}, err
			}
		}
	}
	return icecream, nil
}

func (icdas icecreamDataAccessorService) GetIcecreams(ctx context.Context, icecreamRequest request.IcecreamFilter) (response.BulkIcecreamIndexResponse, error) {
	var totalResults int64

	err, searchResult := icdas.getIcecreamSearchResults(ctx, icecreamRequest)
	if err != nil {
		log.Fatal(ctx, err.Error())
		return response.BulkIcecreamIndexResponse{Icecreams: nil, TotalResults: 0}, err
	}

	totalResults = searchResult.TotalHits()
	icecreams, err := icdas.extractIcecreamFromSearchResult(searchResult, err)
	if err != nil {
		log.Fatal(ctx, err.Error())
		return response.BulkIcecreamIndexResponse{Icecreams: nil, TotalResults: totalResults}, err
	}

	return response.BulkIcecreamIndexResponse{Icecreams: icecreams, TotalResults: totalResults}, nil
}

func (icdas icecreamDataAccessorService) getIcecream(hit *elastic.SearchHit, err error) (bo.ESIcecream, error) {
	bytes := []byte(*hit.Source)
	var icecream bo.ESIcecream
	err = json.Unmarshal(bytes, &icecream)
	return icecream, err
}

func (icdas icecreamDataAccessorService) extractIcecreamFromSearchResult(searchResult *elastic.SearchResult, err error) ([]response.IcecreamIndexResponse, error) {
	icecreams := []response.IcecreamIndexResponse{}
	for _, hit := range searchResult.Hits.Hits {
		icecream, err := icdas.getIcecream(hit, err)
		if err != nil {
			return nil, err
		}
		icecreams = append(icecreams, builder.IcecreamIndexResponseBuilder(icecream))
	}
	return icecreams, nil
}

func (icdas icecreamDataAccessorService) getIcecreamSearchResults(ctx context.Context, IcecreamRequest request.IcecreamFilter) (error, *elastic.SearchResult) {
	searchQuery, err := es.NewIcecreamEsHelper().GenerateSearchQuery(ctx, IcecreamRequest)
	if err != nil {
		log.Print(ctx, "failed to generate search query, error", err.Error())
		return err, nil
	}
	searchService := config.GetESConnection().Search().
		Index(icdas.index.IndexName).
		Query(searchQuery).
		Pretty(true)

	searchResult, err := searchService.Do(ctx)
	return err, searchResult
}

func (icdas icecreamDataAccessorService) CreateOrReplaceIcecream(ctx context.Context, icecream bo.ESIcecream) error {
	log.Print(ctx, "Updating IcecreamIndex")
	updatedDocJSONByte, err := json.Marshal(icecream)
	if err != nil {
		return err
	}
	updatedDocJSONString := string(updatedDocJSONByte)
	log.Print(ctx, "Updating index with : ", updatedDocJSONString)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(updatedDocJSONString), &data)
	if err != nil {
		return err
	}

	_, err = config.GetESConnection().Index().
		Index(icdas.index.IndexName).
		Type(icdas.index.Type).
		Id(strconv.Itoa(int(*icecream.Id))).
		BodyJson(icecream).
		Do(ctx)

	log.Print(ctx, "Updated IcecreamIndex. err :", err)
	return err
}

func (icdas icecreamDataAccessorService) DeleteIcecream(ctx context.Context, icecream request.IcecreamDelete) error {
	log.Print(ctx, "Deleting IcecreamIndex")

	res, err := config.GetESConnection().Delete().Index(config.GetConfig().ElasticSearch.Indices.IcecreamIndex.IndexName).
		Type("icecream").Id(icecream.ProductId).Do(ctx)
	if res != nil && err != nil {
		if esErr, ok := err.(*elastic.Error); ok && esErr.Status == http.StatusNotFound {
			log.Print(ctx, "Data was already deleted")
			return nil
		}
	}
	return err
}

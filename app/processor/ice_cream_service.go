package processor

import (
	"context"
	"fmt"
	"ic-indexer-service/app/api/dataaccessor"
	"ic-indexer-service/app/model/bo"
	"ic-indexer-service/app/model/request"
	"ic-indexer-service/app/model/response"
	"log"
)

type IcecreamIndexer interface {
	PartialUpdate(context.Context, request.IcecreamIndexRequest) error
	DeleteIcecream(context.Context, request.IcecreamDelete) error
	Search(context.Context, request.IcecreamFilter) (response.BulkIcecreamIndexResponse, error)
}

type icecreamIndexerService struct {
	das           dataaccessor.IceCreamDataAccessor
	updateHandler UpdateHandler
}

func GetNewIcecreamIndexerService(das dataaccessor.IceCreamDataAccessor, updateHandler UpdateHandler) IcecreamIndexer {
	return &icecreamIndexerService{das: das, updateHandler: updateHandler}
}

func (iis icecreamIndexerService) Search(ctx context.Context, params request.IcecreamFilter) (response.BulkIcecreamIndexResponse, error) {
	return iis.das.GetIcecreams(ctx, params)
}

func (iis icecreamIndexerService) DeleteIcecream(ctx context.Context, params request.IcecreamDelete) error {
	return iis.das.DeleteIcecream(ctx, params)
}

func (iis icecreamIndexerService) PartialUpdate(ctx context.Context, icecreamRequest request.IcecreamIndexRequest) error {
	log.Print(ctx, "The update params  is :", fmt.Sprintf("%+v", icecreamRequest))

	oldIcecream, err := iis.das.GetIcecreamByKey(ctx, icecreamRequest)
	var newEsIcecream bo.ESIcecream
	if err != nil {
		return err
	}
	if oldIcecream.Id != nil {
		log.Print(ctx, "The existing icecream is :", fmt.Sprintf("%+v", oldIcecream))
	}

	err = iis.updateHandler.UpdateIcecreamDetails(ctx, &newEsIcecream, icecreamRequest)
	if err != nil {
		log.Print(ctx, "Failed to update the icecream. Error is: ", err)
		return err
	}
	log.Print(ctx, "The new updated icecream is :", fmt.Sprintf("%+v", newEsIcecream))
	err = iis.updateSources(ctx, newEsIcecream)

	if err != nil {
		return err
	}
	return err
}

func (iis icecreamIndexerService) updateSources(ctx context.Context, icecream bo.ESIcecream) error {
	err := iis.das.CreateOrReplaceIcecream(ctx, icecream)
	if err != nil {
		log.Print(ctx, "Failed to update elastic search. Err is :", err)
		return err
	}
	return nil
}

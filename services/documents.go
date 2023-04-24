package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type DocumentService struct {
	Options []option.RequestOption
}

func NewDocumentService(opts ...option.RequestOption) (r *DocumentService) {
	r = &DocumentService{}
	r.Options = opts
	return
}

// Create a document.
func (r *DocumentService) New(ctx context.Context, documentable_type requests.DocumentNewParamsDocumentableType, documentable_id string, body *requests.DocumentNewParams, opts ...option.RequestOption) (res *responses.Document, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/documents", documentable_type, documentable_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get an existing document.
func (r *DocumentService) Get(ctx context.Context, documentable_type requests.DocumentGetParamsDocumentableType, documentable_id string, id string, opts ...option.RequestOption) (res *responses.Document, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/documents/%s", documentable_type, documentable_id, id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Get a list of documents.
func (r *DocumentService) List(ctx context.Context, documentable_type requests.DocumentListParamsDocumentableType, documentable_id string, query *requests.DocumentListParams, opts ...option.RequestOption) (res *responses.Page[responses.Document], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/%s/%s/documents", documentable_type, documentable_id)
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of documents.
func (r *DocumentService) ListAutoPager(ctx context.Context, documentable_type requests.DocumentListParamsDocumentableType, documentable_id string, query *requests.DocumentListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Document] {
	return responses.NewPageAutoPager(r.List(ctx, documentable_type, documentable_id, query, opts...))
}

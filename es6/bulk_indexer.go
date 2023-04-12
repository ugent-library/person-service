package es6

import (
	"bytes"
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esutil"
	"github.com/ugent-library/people/models"
)

type bulkIndexer[T any] struct {
	bi         esutil.BulkIndexer
	docFn      func(T) (string, []byte, error)
	indexErrFn func(string, error)
}

func newBulkIndexer[T any](client *elasticsearch.Client, index string, docFn func(T) (string, []byte, error), config models.BulkIndexerConfig) (*bulkIndexer[T], error) {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: client,
		Index:  index,
		OnError: func(ctx context.Context, err error) {
			// TODO wrap error
			config.OnError(err)
		},
		// TODO appropriate place for this? without this a controller may search
		// too soon, and see no results
		Refresh: "wait_for",
	})

	if err != nil {
		return nil, err
	}

	return &bulkIndexer[T]{
		bi:         bi,
		docFn:      docFn,
		indexErrFn: config.OnIndexError,
	}, nil
}

func (b *bulkIndexer[T]) Index(ctx context.Context, t T) error {
	id, doc, err := b.docFn(t)
	if err != nil {
		return err
	}

	err = b.bi.Add(
		ctx,
		esutil.BulkIndexerItem{
			Action:       "index",
			DocumentID:   id,
			DocumentType: "_doc",
			Body:         bytes.NewReader(doc),
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
				// TODO wrap error
				if err == nil {
					err = fmt.Errorf("%+v", res.Error)
				}
				b.indexErrFn(item.DocumentID, err)
			},
		},
	)

	return err
}

func (b *bulkIndexer[T]) Close(ctx context.Context) error {
	// TODO wrap error
	return b.bi.Close(ctx)
}
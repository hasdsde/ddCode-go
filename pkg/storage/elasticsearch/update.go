package es

import "context"

// UpdateCoverDocByID 覆盖更新doc
func (e *Elastic) UpdateCoverDocByID(ctx context.Context, index, id string, doc interface{}) error {
	_, err := e.es.Update().Index(index).Id(id).Doc(doc).DocAsUpsert(true).Do(ctx)
	return err
}

// UpdateDocByID 更新doc，不新增
func (e *Elastic) UpdateDocByID(ctx context.Context, index, id string, docAsUpsert bool, doc interface{}) error {
	_, err := e.es.Update().Index(index).Id(id).Doc(doc).DocAsUpsert(docAsUpsert).Do(ctx)
	return err
}

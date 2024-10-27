package es

import (
	"context"
	logger "ddCode-server/pkg/zlogs"
	"encoding/json"

	"github.com/olivere/elastic/v7"
)

// GetDocByID 通过ID获取doc
func (e *Elastic) GetDocByID(ctx context.Context, index, id string, dist interface{}) error {
	doc, err := e.es.Get().Index(index).Id(id).Do(ctx) // nolint
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			e.loggeer.Infof("Document not found", err, logger.MakeField("ID", id))
		case elastic.IsTimeout(err):
			e.loggeer.Errorf("Timeout retrieving document", err, logger.MakeField("ID", id))
		case elastic.IsConnErr(err):
			e.loggeer.Errorf("Connection problem", err, logger.MakeField("ID", id))
		}
		return err
	}
	return json.Unmarshal(doc.Source, dist)
}

type SearchOption func(*elastic.SearchService)

func SearchWithSort(field string, ascending bool) SearchOption {
	return func(s *elastic.SearchService) {
		s.Sort(field, ascending)
	}
}

// SearchWithPage
// page: 从第0页开始
func SearchWithPage(page, size int) SearchOption {
	if page < 0 {
		page = 0
	}
	from := page * size
	return func(s *elastic.SearchService) {
		s.From(from).Size(size)
	}
}

// SearchDatasByQuery
func (e *Elastic) SearchDocsByTermQuery(ctx context.Context, index string,
	tqs map[string]interface{}, ops ...SearchOption) ([][]byte, error) {
	searchService := e.es.Search().Index(index)
	for k, v := range tqs {
		searchService.Query(elastic.NewTermQuery(k, v))
	}
	for _, op := range ops {
		op(searchService)
	}
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			e.loggeer.Infof("Document not found: %v", err)
			return nil, nil
		case elastic.IsTimeout(err):
			e.loggeer.Errorf("Timeout retrieving document: %v", err)
		case elastic.IsConnErr(err):
			e.loggeer.Errorf("Connection problem: %v", err)
		}
		return nil, err
	}

	if searchResult.Hits.TotalHits.Value < 1 {
		e.loggeer.Info("not found recode", logger.MakeField("index", index))
		return [][]byte{}, nil
	}
	dst := make([][]byte, 0, searchResult.Hits.TotalHits.Value)
	for _, hit := range searchResult.Hits.Hits {
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			e.loggeer.Errorf("data marshal", err, logger.MakeField("index", hit.Index), logger.MakeField("id", hit.Id))
			continue
		}
		dst = append(dst, data)
	}
	return dst, nil
}

// SearchDocByTermQuery 匹配查询一个doc
// tar: 结果对象指针
// return: docID, err
func (e *Elastic) SearchDocByTermQuery(ctx context.Context, index string,
	tqs map[string]interface{}, tar interface{}, ops ...SearchOption) (string, error) {
	searchService := e.es.Search().Index(index)
	querys := []elastic.Query{}
	for k, v := range tqs {
		querys = append(querys, elastic.NewTermQuery(k, v))
	}
	boolQuery := elastic.NewBoolQuery().Must().Must(querys...)
	searchResult, err := searchService.Query(boolQuery).Do(ctx)
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			e.loggeer.Infof("Document not found: %v", err)
			return "", nil
		case elastic.IsTimeout(err):
			e.loggeer.Errorf("Timeout retrieving document: %v", err)
		case elastic.IsConnErr(err):
			e.loggeer.Errorf("Connection problem: %v", err)
		}
		return "", err
	}
	if searchResult.Hits.TotalHits.Value < 1 {
		e.loggeer.Info("not found recode", logger.MakeField("index", index))
		return "", nil
	}
	hit := searchResult.Hits.Hits[0]
	if err := json.Unmarshal(hit.Source, tar); err != nil {
		e.loggeer.Errorf("docs Unmarshal obj", err, logger.MakeField("index", hit.Index), logger.MakeField("id", hit.Id))
		return hit.Id, err
	}
	return hit.Id, nil
}

// SearchDocByTermMatch 匹配查询一个doc
// tar: 结果对象指针
// return: docID, err
func (e *Elastic) SearchDocByTermMatch(ctx context.Context, index string,
	tqs map[string]interface{}, tar interface{}, ops ...SearchOption) (string, error) {
	searchService := e.es.Search().Index(index)
	querys := []elastic.Query{}
	for k, v := range tqs {
		querys = append(querys, elastic.NewMatchQuery(k, v))
	}
	boolQuery := elastic.NewBoolQuery().Must().Must(querys...)
	searchResult, err := searchService.Query(boolQuery).Do(ctx)
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			e.loggeer.Infof("Document not found: %v", err)
			return "", nil
		case elastic.IsTimeout(err):
			e.loggeer.Errorf("Timeout retrieving document: %v", err)
		case elastic.IsConnErr(err):
			e.loggeer.Errorf("Connection problem: %v", err)
		}
		return "", err
	}
	if searchResult.Hits.TotalHits.Value < 1 {
		e.loggeer.Info("not found recode", logger.MakeField("index", index))
		return "", nil
	}
	hit := searchResult.Hits.Hits[0]
	if err := json.Unmarshal(hit.Source, tar); err != nil {
		e.loggeer.Errorf("docs Unmarshal obj", err, logger.MakeField("index", hit.Index), logger.MakeField("id", hit.Id))
		return hit.Id, err
	}
	return hit.Id, nil
}

// SearchDocsByTermScrollQuery 通过 scrollId 滚动查询数据

func (e *Elastic) SearchDocsByTermScrollQuery(ctx context.Context, index, scrollID string, pageSize int, filters []*Filters) (*elastic.SearchResult, error) {
	var boolQuery *elastic.BoolQuery
	if len(filters) > 0 {
		boolQuery = e.buildQuery(filters)
	}
	searchService := e.es.Scroll().Index(index).Size(pageSize)
	if len(scrollID) > 0 {
		searchService.ScrollId(scrollID)
	}
	searchResult, err := searchService.Query(boolQuery).Scroll("5m").Do(ctx)
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			e.loggeer.Infof("Document not found: %v", err)
			return nil, nil
		case elastic.IsTimeout(err):
			e.loggeer.Errorf("Timeout retrieving document: %v", err)
		case elastic.IsConnErr(err):
			e.loggeer.Errorf("Connection problem: %v", err)
		}
		return nil, err
	}

	return searchResult, nil
}

func (e *Elastic) buildQuery(filters []*Filters) (filersQuery *elastic.BoolQuery) {
	query := elastic.NewBoolQuery()
	for i := range filters {
		if filters[i].Qu != "" || filters[i].T != "" {
			switch filters[i].Wildcard {
			case "=":
				query = query.Must(elastic.NewTermQuery(filters[i].Qu, filters[i].T))
			case "in":
				tmpInter := make([]interface{}, 0)
				tsString, ok := filters[i].T.([]string)
				if ok {
					for _, v := range tsString {
						tmpInter = append(tmpInter, v)
					}
				}
				tsInt, ok := filters[i].T.([]int)
				if ok {
					for _, v := range tsInt {
						tmpInter = append(tmpInter, v)
					}
				}
				query = query.Must(elastic.NewTermsQuery(filters[i].Qu, tmpInter...))
			case "range":
				ts := filters[i].T.([]string)
				query = query.Must(elastic.NewRangeQuery(filters[i].Qu).Gte(ts[0]).Lte(ts[1]))
			case "!=":
				query = query.MustNot(elastic.NewTermQuery(filters[i].Qu, filters[i].T))
			case ">=":
				query = query.Must(elastic.NewRangeQuery(filters[i].Qu).Gte(filters[i].T))
			case "<=":
				query = query.Must(elastic.NewRangeQuery(filters[i].Qu).Lte(filters[i].T))
			default:
				query = query.Must(elastic.NewTermQuery(filters[i].Qu, filters[i].T))
			}
		}
	}
	return query
}

func (e *Elastic) RemoveScrollID(ctx context.Context, scrollID string) error {
	_, err := e.es.ClearScroll(scrollID).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

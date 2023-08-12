package slice

import (
	"gosearch/pkg/crawler"
)

type Slice struct {
	docs []crawler.Document
}

func New() *Slice {
	s := Slice{}
	return &s
}

// StoreDocs добавляет новые документы.
func (s *Slice) StoreDocs(docs []crawler.Document) error {
	s.docs = append(s.docs, docs...)
	return nil
}

func (s *Slice) Insert(doc crawler.Document) {
	s.docs = append(s.docs, doc)
}

func (s *Slice) Search(id int) crawler.Document {
	for _, doc := range s.docs {
		if doc.ID == id {
			return doc
		}
	}
	return crawler.Document{}
}

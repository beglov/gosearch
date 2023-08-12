package engine

import (
	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
	"gosearch/pkg/storage"
)

// Engine - поисковый движок.
// Его задача - обслуживание поисковых запросов.
// функциональность:
// - обработка поискового запроса;
// - поиск документов в индексе;
// - запрос документов из хранилища;
// - добавление документов в хранилище и индекс;
// - очистка хранилища и индекса;
// - возврат посиковой выдачи.

// Service - поисковый движок.
type Service struct {
	Index   index.Interface
	Storage storage.Interface
}

// New - конструктор.
func New(index index.Interface, storage storage.Interface) *Service {
	s := Service{
		Index:   index,
		Storage: storage,
	}
	return &s
}

// Search ищет документы, соответствующие поисковому запросу.
func (s *Service) Search(query string) []crawler.Document {
	if query == "" {
		return nil
	}
	ids := s.Index.Search(query)
	docs := s.Storage.Docs(ids)
	return docs
}

// Clear очищает индекс и хранилище
func (s *Service) Clear() {
	s.Index.Clear()
	s.Storage.Clear()
}

// Add добавляет документы в хранилице и индексирует их
func (s *Service) Add(docs []crawler.Document) error {
	s.Index.Add(docs)
	err := s.Storage.StoreDocs(docs)
	return err
}

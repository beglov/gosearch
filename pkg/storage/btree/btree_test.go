package btree

import (
	"fmt"
	"gosearch/pkg/crawler"
	"testing"
)

func BenchmarkTree_StoreDocs(b *testing.B) {
	tests := []struct {
		name  string
		count int
	}{
		{
			name:  "Тест №1",
			count: 1,
		},
		{
			name:  "Тест №2",
			count: 10,
		},
		{
			name:  "Тест №3",
			count: 100,
		},
	}
	for _, tt := range tests {
		t := New()
		docs := documents(tt.count)

		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				t.StoreDocs(docs)
			}
		})
	}
}

func BenchmarkTree_Search(b *testing.B) {
	tests := []struct {
		name  string
		count int
	}{
		{
			name:  "Тест №1",
			count: 1,
		},
		{
			name:  "Тест №2",
			count: 10,
		},
		{
			name:  "Тест №3",
			count: 100,
		},
	}
	for _, tt := range tests {
		t := New()
		docs := documents(tt.count)
		t.StoreDocs(docs)

		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				doc := t.Search(i)
				_ = doc
			}
		})
	}
}

func documents(count int) []crawler.Document {
	var docs []crawler.Document
	for i := 0; i <= count; i++ {
		docs = append(docs, crawler.Document{
			ID:    i,
			URL:   fmt.Sprintf("https://example%d.com", i),
			Title: fmt.Sprintf("Example %d Site", i),
		})
	}
	return docs
}

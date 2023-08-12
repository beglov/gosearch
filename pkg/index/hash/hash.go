package hash

import (
	"gosearch/pkg/crawler"
	"strings"
)

// Index - индекс на основе хэш-таблицы.
type Index struct {
	Data map[string][]int
}

// New - конструктор.
func New() *Index {
	var ind Index
	ind.Data = make(map[string][]int)
	return &ind
}

// Add добавляет данные из переданных документов в индекс.
//
// Сначала происходит выделение лексем как ключей словаря из данных документа.
// Потом проверяется наличие номера документа в занчении словаря для лексемы.
// Если номер документа не найден, то он добавляется в значение словаря.
func (index *Index) Add(docs []crawler.Document) {
	for _, doc := range docs {
		for _, token := range tokens(doc.Title) {
			if !exists(index.Data[token], doc.ID) {
				index.Data[token] = append(index.Data[token], doc.ID)
			}
		}
	}
}

// Search возвращает номера документов, где встречается данная лексема.
func (index *Index) Search(token string) []int {
	return index.Data[strings.ToLower(token)]
}

// Clear очищает индекс
func (index *Index) Clear() {
	index.Data = make(map[string][]int)
}

// Разделение строки на лексемы.
func tokens(s string) []string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

// Проверка наличия элемента в массиве.
func exists(ids []int, item int) bool {
	for _, id := range ids {
		if id == item {
			return true
		}
	}
	return false
}

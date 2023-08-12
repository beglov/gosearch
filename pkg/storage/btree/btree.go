package btree

import (
	"gosearch/pkg/crawler"
	"sync"
)

// Tree - Двоичное дерево поиска
type Tree struct {
	Mux  *sync.Mutex
	Root *Element
}

// Element - элемент дерева
type Element struct {
	Left, Right *Element
	Value       crawler.Document
}

// New - конструктор.
func New() *Tree {
	db := Tree{
		Mux: new(sync.Mutex),
	}
	return &db
}

// Clear очищает хранилище
func (t *Tree) Clear() {
	t.Root = nil
}

// StoreDocs добавляет новые документы.
func (t *Tree) StoreDocs(docs []crawler.Document) error {
	for _, doc := range docs {
		t.Insert(doc)
	}
	return nil
}

// Docs возвращает документы по их номерам.
func (t *Tree) Docs(ids []int) []crawler.Document {
	var result []crawler.Document
	t.Mux.Lock()
	defer t.Mux.Unlock()
	for _, id := range ids {
		s := t.Search(id)
		result = append(result, s)
	}
	return result
}

// Insert - вставка элемента в дерево
func (t *Tree) Insert(doc crawler.Document) {
	e := &Element{Value: doc}
	if t.Root == nil {
		t.Root = e
		return
	}
	insert(t.Root, e)
}

// inset рекурсивно вставляет элемент в нужный уровень дерева.
func insert(node, new *Element) {
	if new.Value.ID < node.Value.ID {
		if node.Left == nil {
			node.Left = new
			return
		}
		insert(node.Left, new)
	}
	if new.Value.ID >= node.Value.ID {
		if node.Right == nil {
			node.Right = new
			return
		}
		insert(node.Right, new)
	}
}

// Search - поиск значения в дереве, выдаёт документ если найдено, иначе nil
func (t *Tree) Search(x int) crawler.Document {
	return search(t.Root, x)
}

func search(el *Element, x int) crawler.Document {
	if el == nil {
		return crawler.Document{}
	}
	if el.Value.ID == x {
		return el.Value
	}
	if el.Value.ID < x {
		return search(el.Right, x)
	}
	return search(el.Left, x)
}

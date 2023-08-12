package hash

import (
	"gosearch/pkg/crawler"
	"os"
	"reflect"
	"testing"
)

var ind *Index

func TestMain(m *testing.M) {
	ind = New()
	docs := []crawler.Document{
		{
			ID:    10,
			Title: "Два Слова",
		},
		{
			ID:    20,
			Title: "And Another Three",
		},
		{
			ID:    30,
			Title: "Three Tokens More",
		},
	}
	ind.Add(docs)
	os.Exit(m.Run())
}

func TestIndex_Add(t *testing.T) {
	ind2 := New()
	docs := []crawler.Document{
		{
			ID:    40,
			Title: "Go Java Perl Perl",
		},
	}
	ind2.Add(docs)
	got := len(ind2.Data)
	want := 3
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}

func TestIndex_Search(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  []int
	}{
		{
			name:  "Тест №1",
			token: "ДВА",
			want:  []int{10},
		},
		{
			name:  "Тест №2",
			token: "THree",
			want:  []int{20, 30},
		},
		{
			name:  "Тест №3",
			token: "NotAToken",
			want:  nil,
		},
		{
			name:  "Тест №4",
			token: "three tokens",
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ind.Search(tt.token); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("получили %v, ожидалось %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Clear(t *testing.T) {
	ind.Clear()
	got := len(ind.Data)
	want := 0
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}

func Test_tokens(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "Тест №1",
			s:    "qwe rty",
			want: []string{"qwe", "rty"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tokens(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exists(t *testing.T) {
	tests := []struct {
		name string
		ids  []int
		item int
		want bool
	}{
		{
			name: "Тест №1",
			ids:  []int{10, 20, 30},
			item: 10,
			want: true,
		},
		{
			name: "Тест №2",
			ids:  []int{10, 20, 30},
			item: 11,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exists(tt.ids, tt.item); got != tt.want {
				t.Errorf("exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

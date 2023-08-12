package spider

import (
	"golang.org/x/net/html"
	"reflect"
	"testing"
)

func Test_pageTitle(t *testing.T) {
	want := "Заголовок"
	n := &html.Node{
		FirstChild: &html.Node{
			Type: html.ElementNode,
			Data: "title",
			FirstChild: &html.Node{
				Data: want,
			},
		},
	}
	if got := pageTitle(n); got != want {
		t.Errorf("pageTitle() = %v, want %v", got, want)
	}
}

func Test_pageLinks(t *testing.T) {
	n := &html.Node{
		Type: html.ElementNode,
		Data: "a",
		Attr: []html.Attribute{
			{
				Key: "href",
				Val: "www.ya.ru",
			},
			{
				Key: "href",
				Val: "www.yandex.ru",
			},
			{
				Key: "href",
				Val: "www.ya.ru",
			},
		},
		FirstChild: &html.Node{
			Type: html.ElementNode,
			Data: "a",
			Attr: []html.Attribute{
				{
					Key: "href",
					Val: "www.ya.ru",
				},
				{
					Key: "href",
					Val: "www.google.ru",
				},
			},
		},
	}
	want := []string{
		"www.ya.ru",
		"www.yandex.ru",
		"www.google.ru",
	}
	if got := pageLinks([]string{}, n); !reflect.DeepEqual(got, want) {
		t.Errorf("pageTitle() = %v, want %v", got, want)
	}
}

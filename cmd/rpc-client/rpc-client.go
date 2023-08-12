package main

import (
	"bufio"
	"fmt"
	"gosearch/pkg/crawler"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	// создание клиента RPC
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	fmt.Println("Подключение установлено. Введите поисковую фразу:")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите слово для поиска или exit для выхода: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			break
		}

		var results []crawler.Document
		err = client.Call("Service.Search", text, &results)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range results {
			fmt.Printf("%s - %s\n", v.URL, v.Title)
		}
	}
}

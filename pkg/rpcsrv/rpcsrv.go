package rpcsrv

import (
	"net"
	"net/rpc"

	"gosearch/pkg/crawler"
	"gosearch/pkg/engine"
)

// Service - сетевая служба
type Service struct {
	addr   string
	engine *engine.Service
}

// New - конструктор.
func New(addr string, engine *engine.Service) *Service {
	s := Service{
		addr:   addr,
		engine: engine,
	}
	return &s
}

// Start запускает сетевую службу
func (s *Service) Start() error {
	err := rpc.Register(s)
	if err != nil {
		return err
	}
	// регистрация сетевой службы RPC-сервера
	listener, err := net.Listen("tcp4", s.addr)
	if err != nil {
		return err
	}

	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go rpc.ServeConn(conn)
	}
}

// Search ищет документы, соответствующие поисковому запросу.
func (s *Service) Search(query string, result *[]crawler.Document) error {
	*result = s.engine.Search(query)
	return nil
}

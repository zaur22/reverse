package serv

import(
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/config"
	"syscall"
	"time"
)

type Server struct {
	s          *http.Server
	stopped    chan bool
	osStopSigs chan os.Signal
}

func NewServer(router http.Handler) *Server {
	var server = Server{
		s: &http.Server{
			Addr:           config.GetString(config.ServerAddr),
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}

	server.osStopSigs = make(chan os.Signal, 1)
	server.stopped = make(chan bool, 1)
	signal.Notify(server.osStopSigs, os.Interrupt, syscall.SIGTERM)

	return &server
}

func (s *Server) Serve() {

	go func() {
		log.Printf("starting server at %s", s.s.Addr)
		if err := s.s.ListenAndServe(); err != nil {
			log.Fatalf("Server error: %v", err.Error())
		}
	}()

	go func() {
		<-s.osStopSigs
		log.Println()
		log.Println("Stopping service...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		s.s.Shutdown(ctx)
		s.stopped <- true
	}()

	<-s.stopped
	log.Printf("Done\n")
	os.Exit(0)
}
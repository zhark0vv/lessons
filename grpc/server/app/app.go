package app

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"lessons/grpc/server/controller"
	educationapi "lessons/grpc/server/gen/go/zhark0vv/grpc/education/api"
)

type App struct {
	grpcServer *grpc.Server
	httpServer *http.Server
	lis        net.Listener
}

func Init(ctx context.Context, c *controller.Controller) (*App, error) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	educationapi.RegisterEducationAPIServer(grpcServer, c)

	l, err := net.Listen("tcp", ":8085")
	if err != nil {
		return nil, err
	}

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = educationapi.RegisterEducationAPIHandlerFromEndpoint(
		ctx,
		gwmux,
		":8085",
		opts,
	)
	if err != nil {
		return nil, err
	}

	httpServer := &http.Server{
		Addr:    ":8086",
		Handler: gwmux,
	}

	return &App{
		grpcServer: grpcServer,
		httpServer: httpServer,
		lis:        l,
	}, nil
}

func (a *App) Run() error {
	// Канал для отслеживания сигналов остановки
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Запускаем gRPC сервер в отдельной горутине
	go func() {
		log.Println("Starting gRPC server on :8085")
		if err := a.grpcServer.Serve(a.lis); err != nil {
			log.Printf("gRPC server error: %v", err)
		}
	}()

	// Запускаем HTTP сервер в отдельной горутине
	go func() {
		log.Println("Starting HTTP server on :8086")
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	// Ожидание сигнала остановки
	<-stopChan
	log.Println("Shutting down servers...")

	// Останавливаем gRPC сервер
	a.grpcServer.GracefulStop()
	log.Println("gRPC server stopped gracefully")

	// Останавливаем HTTP сервер с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	} else {
		log.Println("HTTP server stopped gracefully")
	}

	return nil
}

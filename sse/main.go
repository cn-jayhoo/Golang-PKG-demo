package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
	"sse/service"
	"sse/service_api"
	"strings"
)

func sseHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/v1/events") {
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		// 检查是否是事件流端点
		next.ServeHTTP(w, r)
	})
}

type SSEMarshaler struct {
	runtime.JSONPb
}

func (S SSEMarshaler) Marshal(v interface{}) ([]byte, error) {
	// 先使用JSONPb进行序列化
	jsonData, err := S.JSONPb.Marshal(v)
	if err != nil {
		return nil, err
	}

	// 包装为SSE格式
	sseData := fmt.Sprintf("data: %s\n", jsonData)
	return []byte(sseData), nil
}

func (S SSEMarshaler) ContentType(v interface{}) string {
	return "text/event-stream"
}

func main() {
	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()
	service_api.RegisterEventServiceServer(grpcServer, service.NewPerson())
	go func() {
		lis, _ := net.Listen("tcp", ":50051")
		grpcServer.Serve(lis)
	}()

	// 创建 gRPC Gateway mux
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			"text/event-stream",
			&SSEMarshaler{
				JSONPb: runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						EmitUnpopulated: true,
					},
				},
			},
		))

	// 注册 gRPC Gateway 端点
	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := service_api.RegisterEventServiceHandlerFromEndpoint(ctx, gwmux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	// 启动 HTTP 服务器
	mux := http.NewServeMux()
	mux.Handle("/", sseHeadersMiddleware(gwmux))

	log.Println("Starting HTTP server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	grpc2 "github.com/ricejson/apollo-backend/controller/grpc"
	http2 "github.com/ricejson/apollo-backend/controller/http"
	"github.com/ricejson/apollo-backend/models/mongodb"
	_ "github.com/ricejson/apollo-backend/models/mongodb"
	"github.com/ricejson/apollo-backend/repository"
	"github.com/ricejson/apollo-backend/repository/dao"
	"github.com/ricejson/apollo-backend/service/toggle"
	proto2 "github.com/ricejson/apollo-idl-go/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	col := mongodb.GetClient().Database("apollo").Collection("toggles")
	toggleDAO := dao.NewMongoToggleDAO(col)
	toggleRepository := repository.NewDefaultToggleRepository(toggleDAO)
	toggleService := toggle.NewDefaultToggleService(toggleRepository)
	go func() {
		// 测试grpc
		s := grpc.NewServer()
		proto2.RegisterRPCToggleServiceServer(s, grpc2.NewGRPCToggleServerImpl(toggleService))
		listen, err := net.Listen("tcp", ":8992")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s.Serve(listen)
	}()
	server := gin.Default()
	// 使用cors中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许所有源访问
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	config.AllowCredentials = true // 允许携带cookie

	server.Use(cors.New(config))

	toggleController := http2.NewToggleController(toggleService)
	toggleController.RegisterServices(server)
	if err := server.Run(":8991"); err != nil {
		panic(err)
	}
}

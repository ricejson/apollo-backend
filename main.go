package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ricejson/apollo-backend/controller/grpc/proto"
	http2 "github.com/ricejson/apollo-backend/controller/http"
	"github.com/ricejson/apollo-backend/models/mongodb"
	_ "github.com/ricejson/apollo-backend/models/mongodb"
	"github.com/ricejson/apollo-backend/repository"
	"github.com/ricejson/apollo-backend/repository/dao"
	"github.com/ricejson/apollo-backend/service/toggle"
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
		proto.RegisterRPCToggleServiceServer(s, proto.NewGRPCToggleServerImpl(toggleService))
		listen, err := net.Listen("tcp", ":8992")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s.Serve(listen)
	}()
	server := gin.Default()

	toggleController := http2.NewToggleController(toggleService)
	toggleController.RegisterServices(server)
	if err := server.Run(":8991"); err != nil {
		panic(err)
	}
}

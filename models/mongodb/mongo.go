package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/event"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

var client *mongo.Client

func GetClient() *mongo.Client {
	return client
}

func init() {
	// 1. 准备命令监视器
	monitor := &event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			// 输出查询命令
			log.Println(startedEvent.Command)
		},
		Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {},
		Failed:    func(ctx context.Context, errEvent *event.CommandFailedEvent) {},
	}
	// 2. 建立连接
	opts := options.Client().ApplyURI("mongodb://root:37zl3M4X42V3z3Ei@dbconn.sealoshzh.site:34949/?directConnection=true").SetMonitor(monitor)
	c, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
	}
	client = c
}

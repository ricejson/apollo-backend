package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// 定义数据结构（与之前相同）
type Rule struct {
	ID        string `json:"id" bson:"id"`
	Attribute string `json:"attribute" bson:"attribute"`
	Operator  string `json:"operator" bson:"operator"`
	Value     string `json:"value" bson:"value"`
}

type Audience struct {
	ID    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Rules []Rule `json:"rules" bson:"rules"`
}

type Toggle struct {
	ID          string     `json:"id" bson:"id"`
	Name        string     `json:"name" bson:"name"`
	Key         string     `json:"key" bson:"key"`
	Description string     `json:"description" bson:"description"`
	Status      string     `json:"status" bson:"status"`
	CreatedAt   string     `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string     `json:"updatedAt" bson:"updatedAt"`
	Audiences   []Audience `json:"audiences" bson:"audiences"`
}

func main() {
	// 请修改以下参数为你的实际MongoDB认证信息
	username := "root"      // 替换为你的用户名
	password := "example"   // 替换为你的密码
	databaseName := "admin" // 通常是admin，也可能是其他认证数据库

	// MongoDB连接字符串（带认证）
	connectionString := fmt.Sprintf("mongodb://%s:%s@localhost:27017/%s",
		username, password, databaseName)

	// 设置连接选项
	clientOptions := options.Client().
		ApplyURI(connectionString).
		SetAuth(options.Credential{
			Username:   username,
			Password:   password,
			AuthSource: databaseName, // 认证数据库
		})

	// 连接到MongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("连接MongoDB失败:", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal("断开连接失败:", err)
		}
	}()

	// 测试连接
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Ping失败:", err)
	}
	fmt.Println("成功连接到MongoDB!")

	// 获取数据库和集合
	database := client.Database("apollo")
	collection := database.Collection("toggles")

	// JSON数据
	jsonData := `{
		"id": "tg_1",
		"name": "智能推荐算法 v2",
		"key": "smart_recommender_v2",
		"description": "基于用户行为路径的新型推荐引擎，旨在提高点击率。",
		"status": "enabled",
		"createdAt": "2024-03-20",
		"updatedAt": "2026-01-11",
		"audiences": [
			{
				"id": "aud_1",
				"name": "北京/上海核心测试用户",
				"rules": [
					{
						"id": "r1",
						"attribute": "city",
						"operator": "in",
						"value": "Beijing, Shanghai"
					},
					{
						"id": "r2",
						"attribute": "traffic",
						"operator": "lt",
						"value": "20"
					},
					{
						"id": "kgt543dcg",
						"attribute": "user_id",
						"operator": "equals",
						"value": "1"
					}
				]
			}
		]
	}`

	// 解析JSON数据
	var toggle Toggle
	err = json.Unmarshal([]byte(jsonData), &toggle)
	if err != nil {
		log.Fatal("解析JSON失败:", err)
	}

	// 检查是否已存在相同ID的文档
	filter := bson.M{"id": toggle.ID}
	var existingDoc bson.M
	err = collection.FindOne(context.Background(), filter).Decode(&existingDoc)
	if err == nil {
		// 文档已存在，询问用户是否要更新
		fmt.Printf("ID为 '%s' 的文档已存在。是否要更新？(y/n): ", toggle.ID)
		var response string
		fmt.Scanln(&response)

		if response == "y" || response == "Y" {
			// 更新文档
			update := bson.M{"$set": toggle}
			result, err := collection.UpdateOne(context.Background(), filter, update)
			if err != nil {
				log.Fatal("更新文档失败:", err)
			}
			fmt.Printf("文档更新成功! 匹配%d个文档，修改%d个文档\n",
				result.MatchedCount, result.ModifiedCount)
		} else {
			fmt.Println("操作已取消。")
			return
		}
	} else if err == mongo.ErrNoDocuments {
		// 文档不存在，插入新文档
		result, err := collection.InsertOne(context.Background(), toggle)
		if err != nil {
			log.Fatal("插入文档失败:", err)
		}
		fmt.Printf("文档插入成功! 文档ID: %v\n", result.InsertedID)
	} else {
		log.Fatal("查询文档失败:", err)
	}
}

// model/gamedb/register_user.go
package gamedb

import (
	"context"
	"fmt"
	"log"
	"time"

	"yourbot/Events/Client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func init() {
	dbClient = client.Mongodb() // MongoDB 클라이언트 초기화
	if dbClient == nil {
		log.Fatal("MongoDB 클라이언트 초기화에 실패했습니다.")
	}
}

func Accessiondb(userID string, money int, joinTime time.Time) {
	if dbClient == nil {
		log.Println("MongoDB 클라이언트가 초기화되지 않았습니다.")
		return
	}

	collection := dbClient.Database("gamebot").Collection("gameusers")

	// 업데이트 또는 삽입을 위한 필터 및 데이터 설정
	filter := bson.M{"userID": userID}
	update := bson.M{
		"$set": bson.M{
			"userID":   userID,
			"money":    money,
			"joinTime": joinTime,
		},
	}

	// Upsert 옵션 설정 (데이터가 없을 경우 삽입)
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Println("오류 발생: 데이터를 저장할 수 없습니다.", err)
		return
	}

	fmt.Printf("%s 데이터가 성공적으로 저장되었습니다.", userID)
}
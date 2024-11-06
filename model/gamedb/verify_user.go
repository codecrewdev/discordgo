package gamedb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// VerifyUser 함수는 사용자가 이미 가입했는지 확인합니다.
func VerifyUser(userID string) (bool, error) {
	if dbClient == nil {
		log.Println("MongoDB 클라이언트가 초기화되지 않았습니다.")
		return false, fmt.Errorf("MongoDB 클라이언트가 초기화되지 않았습니다.")
	}

	collection := dbClient.Database("gamebot").Collection("gameusers")

	// 사용자 ID로 필터링
	filter := bson.M{"userID": userID}
	var result bson.M

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// 사용자가 존재하지 않는 경우
			return false, nil
		}
		log.Println("오류 발생: 사용자 데이터 검색 중 오류가 발생했습니다.", err)
		return false, err
	}

	// 사용자가 이미 가입한 경우
	return true, nil
}

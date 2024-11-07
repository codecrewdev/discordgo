package gamedb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// VerifyUser 함수는 사용자가 이미 가입했는지 확인합니다.
func IsUserRegistered(userID string) (bool, error) {
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

func VerifyUser(userID string) (bool, int, time.Time, error) {
	if dbClient == nil {
		return false, 0, time.Time{}, fmt.Errorf("MongoDB 클라이언트가 초기화되지 않았습니다.")
	}

	collection := dbClient.Database("gamebot").Collection("gameusers")
	
	// 유저 정보 확인
	filter := bson.M{"userID": userID}
	var result struct {
		Money    int       `bson:"money"`
		JoinTime time.Time `bson:"joinTime"`
	}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, 0, time.Time{}, nil // 유저가 존재하지 않음
		}
		return false, 0, time.Time{}, fmt.Errorf("사용자 정보를 가져오는 동안 오류가 발생했습니다: %v", err)
	}

	return true, result.Money, result.JoinTime, nil
}

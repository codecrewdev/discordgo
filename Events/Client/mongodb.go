package client

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB 연결 함수
func Mongodb() *mongo.Client {

	// .env 파일 로드
    err := godotenv.Load()
    if err != nil {
        fmt.Println(".env 파일 로딩 오류")
        return nil
    }

    // 환경 변수에서 토큰 가져오기
    mongoURI := os.Getenv("mongo")
    if mongoURI == "" {
        fmt.Println(".env 파일에 토큰이 설정되지 않았습니다.")
        return nil
    }

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 연결 확인
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB에 성공적으로 연결되었습니다!")
	return client
}

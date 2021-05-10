package cash

import (
	redisDb "aapanavyapar-service-smssender/configurations/redis"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strings"
	"time"
)

type RedisDataBase struct {
	Cash *redis.Client
}

func NewRedisClient() *RedisDataBase {
	return &RedisDataBase{Cash: redisDb.InitRedis()}
}

func (dataBase *RedisDataBase) InitSmsStream(ctx context.Context) error {
	err := dataBase.CreateSmsStream(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "already exist") {
			fmt.Println("Sms Stream Already Exist")
			return nil
		}
		return err
	}
	return nil
}

func (dataBase *RedisDataBase) CreateSmsStream(ctx context.Context) error {
	return dataBase.Cash.XGroupCreateMkStream(ctx, os.Getenv("REDIS_STREAM_MSG_NAME"), os.Getenv("REDIS_STREAM_MSG_GROUP"), "$").Err()

}

func (dataBase *RedisDataBase) AckToSmsStream(ctx context.Context, val string) error {
	err := dataBase.Cash.XAck(ctx, os.Getenv("REDIS_STREAM_MSG_NAME"), os.Getenv("REDIS_STREAM_MSG_GROUP"), val).Err()
	return err
}

func (dataBase *RedisDataBase) DelFromSmsStream(ctx context.Context, val string) error {
	err := dataBase.Cash.XDel(ctx, os.Getenv("REDIS_STREAM_MSG_NAME"), val).Err()
	return err
}

func (dataBase *RedisDataBase) ReadFromSmsStream(ctx context.Context, count int64, timeout time.Duration, myKeyId, workerName string) *redis.XStreamSliceCmd {

	readGroup := dataBase.Cash.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    os.Getenv("REDIS_STREAM_MSG_GROUP"),
		Consumer: workerName,
		Streams:  []string{os.Getenv("REDIS_STREAM_MSG_NAME"), myKeyId},
		Count:    count,   // No Of Data To Retrieve
		Block:    timeout, //TimeOut
		NoAck:    false,
	})
	return readGroup

}

func (dataBase *RedisDataBase) AddToSmsStream(ctx context.Context, mobileNo, messageId, message string) error {

	err := dataBase.Cash.XAdd(ctx, &redis.XAddArgs{
		Stream:       os.Getenv("REDIS_STREAM_MSG_NAME"),
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           messageId,
		Values:       []string{"mobileNo", mobileNo, "message", message},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

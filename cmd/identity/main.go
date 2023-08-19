package main

import (
	"Identity/cmd/config"
	"Identity/internal/casbin"
	"Identity/internal/controller"
	authapiv1 "Identity/internal/proto/identity/authapi/v1"
	userapiv1 "Identity/internal/proto/identity/userapi/v1"
	"Identity/internal/redis"
	"Identity/internal/repository/sql"
	service "Identity/internal/service/impl"
	"Identity/internal/store"
	"Identity/pkg/jwt"
	"Identity/pkg/limiter"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"
	"google.golang.org/grpc"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

func run() error {
	conf, err := config.Load()
	if err != nil {
		return err
	}

	db, err := store.NewMSSQLStorage(conf)
	if err != nil {
		return err
	}

	userRepo := sql.NewUserRepo(db)
	sessionRepo := sql.NewSessionRepo(db)
	trackRepo := sql.NewTrackRepo(db)

	cb, err := casbin.NewEnforcer(conf)
	if err != nil {
		return err
	}
	node, _ := snowflake.NewNode(1)
	rc := redis.NewRedisClient(conf)
	jt := jwt.NewJwtHandler(conf.RSAPair)
	il := limiter.NewIPLimiter(5 * time.Minute)

	userSvc := service.NewUserService(userRepo, cb, node)
	sessionSvc := service.NewSessionService(sessionRepo, rc)
	authSvc := service.NewAuthService(userRepo, sessionRepo, trackRepo, jt, rc)

	uss := controller.NewUserServiceServer(userSvc, sessionSvc, il, cb, jt)
	ass := controller.NewAuthServiceServer(authSvc, sessionSvc, il, jt)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.HttpPort))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	userapiv1.RegisterUserServiceServer(server, uss)
	authapiv1.RegisterAuthServiceServer(server, ass)

	return server.Serve(lis)
}

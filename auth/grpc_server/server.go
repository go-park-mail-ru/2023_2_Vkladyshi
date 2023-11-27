package server

import (
	"context"
	"flag"
	"log/slog"

	pb "auth/auth"
	"auth/configs"
	"auth/repository/session"
)

var (
	port = flag.Int("port", 50051, "The server port")
	lg = &slog.Logger{}
	cfg_sessions, _ = configs.ReadSessionRedisConfig()
	sessions, _ = session.GetSessionRepo(*cfg_sessions, lg)
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) Auth_Check(ctx context.Context, in *pb.Auth_Check_Request) (*pb.Auth_Check_Reply, error) {
	sid := in.GetSid()
	status, _ := sessions.CheckActiveSession(ctx, sid, lg)
	return &pb.Auth_Check_Reply{Status: status}, nil
}
package delivery_auth_grpc

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"log/slog"
	"net"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/authorization/repository/profile"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/authorization/repository/session"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"

	pb "github.com/go-park-mail-ru/2023_2_Vkladyshi/authorization/proto"
)

type Config struct {
	Port           string `yaml:"port"`
	ConnectionType string `yaml:"connection_type"`
}

type server struct {
	pb.UnimplementedAuthorizationServer
	userRepo    *profile.RepoPostgre
	sessionRepo *session.SessionRepo
	lg          *slog.Logger
}

func (s *server) GetId(ctx context.Context, req *pb.FindIdRequest) (*pb.FindIdResponse, error) {
	login, err := s.sessionRepo.GetUserLogin(ctx, req.Sid, s.lg)
    id, err := s.userRepo.GetUserProfileId(login)
    if err != nil {
        return nil, err
    }
    return &pb.FindIdResponse{
        Value: id,
    }, nil
}

func (s *server) GetIdsAndPaths(ctx context.Context, req *pb.IdsAndPathsListRequest) (*pb.IdsAndPathsResponse, error) {
    ids, paths, err := s.userRepo.GetIdsAndPaths()
    if err != nil {
        return nil, err
    }
    return &pb.IdsAndPathsResponse{
        Ids:   ids,
        Paths: paths,
    }, nil
}


func (s *server) GetAuthorizationStatus(ctx context.Context, req *pb.AuthorizationCheckRequest) (*pb.AuthorizationCheckResponse, error) {
    status, err := s.sessionRepo.CheckActiveSession(ctx, req.Sid, s.lg)
    if err != nil {
        return nil, err
    }
    return &pb.AuthorizationCheckResponse{
        Status: status,
    }, nil
}


func ListenAndServeGrpc(l *slog.Logger) {
	filename := flag.String("config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	configData, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}

	lis, err := net.Listen(config.ConnectionType, ":"+config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthorizationServer(s, &server{
		lg: l,
	})
	log.Printf("Server started on %s port %s", config.ConnectionType, config.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

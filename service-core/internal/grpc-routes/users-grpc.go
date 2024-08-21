package grpcroutes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JoaoRafa19/plataforma-ead-service-core/domain/entities"
	"github.com/JoaoRafa19/plataforma-ead-service-core/pb"
)

type Server struct {
	pb.UserServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	res := &pb.CreateUserResponse{}

	user, err := entities.NewUser(req.GetName(), req.GetEmail(), req.GetPassword(), req.GetPhone())
	if err != nil {
		res.Status = 400
		res.Message = "usuário inválido"
		return nil, err
	}

	jsonData, _ := json.Marshal(user)

	fmt.Println(string(jsonData))

	res.Message = "Usuário criado com sucesso"
	res.Status = 201
	
	return res, nil

}
package service

import (
	"context"
	pb "helloworld/api/piano/v1"
	"helloworld/internal/data"
)

type PianoService struct {
	pb.UnimplementedPianoServer
	dao *data.Data
}

func NewPianoService(d *data.Data) *PianoService {
	return &PianoService{dao: d}
}

func (s *PianoService) CreatePiano(ctx context.Context, req *pb.CreatePianoRequest) (*pb.CreatePianoReply, error) {

	return &pb.CreatePianoReply{
		Code: 1,
		Msg:  "success",
	}, nil
}
func (s *PianoService) UpdatePiano(ctx context.Context, req *pb.UpdatePianoRequest) (*pb.UpdatePianoReply, error) {
	return &pb.UpdatePianoReply{}, nil
}
func (s *PianoService) DeletePiano(ctx context.Context, req *pb.DeletePianoRequest) (*pb.DeletePianoReply, error) {
	return &pb.DeletePianoReply{}, nil
}
func (s *PianoService) GetPiano(ctx context.Context, req *pb.GetPianoRequest) (*pb.GetPianoReply, error) {
	return &pb.GetPianoReply{}, nil
}
func (s *PianoService) ListPiano(ctx context.Context, req *pb.ListPianoRequest) (*pb.ListPianoReply, error) {
	return &pb.ListPianoReply{}, nil
}

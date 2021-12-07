package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	pb "helloworld/api/piano/v1"
	"helloworld/common/jwt"
	"helloworld/internal/dao"
	model "helloworld/internal/model/jwt"
	"helloworld/internal/model/piano"
	"time"
)

type PianoService struct {
	pb.UnimplementedPianoServer
	dao *dao.Dao
}

func NewPianoService(d *dao.Dao) *PianoService {
	return &PianoService{dao: d}
}

func (s *PianoService) CreatePiano(ctx context.Context, req *pb.CreatePianoRequest) (*pb.CreatePianoReply, error) {
	if req.Name == "" {
		return nil, errors.New("no name")
	}
	create := &piano.Create{
		Name:      req.Name,
		Status:    0,
		Type:      1,
		StartTime: time.Now().Unix(),
		StopTime:  time.Now().Unix() + 14400,
		NickName:  "baby",
	}

	err := s.dao.CreateList(ctx, create)
	if err != nil {
		return &pb.CreatePianoReply{
			Code: 0,
			Msg:  "CreateList fail",
		}, err
	}
	return &pb.CreatePianoReply{
		Code: 1,
		Msg:  "success",
	}, nil
}

func (s *PianoService) RegisterUser(ctx context.Context, req *pb.CreateRegisterRequest) (*pb.CommonResponse, error) {
	if req != nil {
		//查找用户
		cnt, err := s.dao.CheckUser(&model.RegisterReq{
			Name:     req.Name,
			Phone:    req.Phone,
			Email:    req.Email,
			Password: req.Password,
		})
		if cnt != 0 {
			return &pb.CommonResponse{
				Code: 0,
				Msg:  "用户已存在" + err.Error(),
			}, nil
		} else {
			password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
			if err != nil {
				return &pb.CommonResponse{
					Code: 0,
					Msg:  "内部错误" + err.Error(),
				}, nil
			}
			err = s.dao.AddUser(&model.UserCreate{
				UserName: req.Name,
				Phone:    req.Phone,
				Email:    req.Email,
				Password: string(password),
			})
			if err != nil {
				return &pb.CommonResponse{
					Code: 0,
					Msg:  "新增错误" + err.Error(),
				}, nil
			}
		}
	} else {
		return &pb.CommonResponse{
			Code: 0,
			Msg:  "请传入参数",
		}, nil
	}
	return &pb.CommonResponse{
		Code: 1,
		Msg:  "成功",
	}, nil
}

//todo 生成的token可以存到redis里，之后每次登陆只要取redis就行
func (s *PianoService) LoginIn(ctx context.Context, req *pb.CreateLogInRequest) (*pb.CommonResponse, error) {
	if req != nil {
		// 登陆逻辑校验(查库，验证用户是否存在以及登陆信息是否正确)
		user, err := s.dao.LoginCheck(&model.LoginReq{
			UserName: req.Name,
			//Password: req.Password,
		})
		// 验证通过后为该次请求生成token
		if user != nil {
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
			if err != nil {
				return &pb.CommonResponse{
					Code: 0,
					Msg:  "密码错误" + err.Error(),
				}, nil
			}
			data, err := jwt.GenerateToken(*user)
			if err != nil {
				return &pb.CommonResponse{
					Code: 0,
					Msg:  err.Error(),
				}, nil
			}
			return &pb.CommonResponse{
				Code: 1,
				Msg:  data.Token,
			}, nil
		} else {
			return &pb.CommonResponse{
				Code: 0,
				Msg:  "验证失败" + err.Error(),
			}, nil
		}

	} else {
		return &pb.CommonResponse{
			Code: 0,
			Msg:  "请传入参数",
		}, nil
	}
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

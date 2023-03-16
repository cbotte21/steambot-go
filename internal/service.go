package internal

import (
	"context"
	"errors"
	"github.com/cbotte21/chess-go/internal/game"
	"github.com/cbotte21/chess-go/internal/game/position"
	"github.com/cbotte21/chess-go/pb"
	"github.com/cbotte21/microservice-common/pkg/jwtParser"
)

type Chess struct {
	JwtSecret       *jwtParser.JwtSecret
	InstanceManager *game.Instances
	pb.UnimplementedChessServiceServer
}

func NewChess(secret *jwtParser.JwtSecret, instanceManager *game.Instances) Chess {
	return Chess{JwtSecret: secret, InstanceManager: instanceManager}
}

func (chess *Chess) Create(ctx context.Context, createRequest *pb.CreateRequest) (*pb.CreateResponse, error) {
	p1 := createRequest.Player1.GetXId()
	p2 := createRequest.Player2.GetXId()
	var err error = nil

	if chess.InstanceManager.InGame(p1) || chess.InstanceManager.InGame(p2) {
		err = chess.InstanceManager.CreateGame(p1, p2, false) //TODO: All games unranked
	} else {
		err = errors.New("player(s) are already in a game")
	}

	return &pb.CreateResponse{Status: false}, err
}

func (chess *Chess) Move(ctx context.Context, moveRequest *pb.MoveRequest) (*pb.Bool, error) {
	player, err := chess.JwtSecret.Redeem(moveRequest.GetJwt().GetJwt())
	p1 := position.NewPosition(int(moveRequest.GetInitial().GetX()), int(moveRequest.GetInitial().GetY()))
	p2 := position.NewPosition(int(moveRequest.GetFinal().GetX()), int(moveRequest.GetFinal().GetY()))

	if err == nil {
		pgame, err := chess.InstanceManager.GetGame(player.Id)
		if err == nil {
			err = pgame.Move(player.Id, p1, p2)
			if err == nil {
				return &pb.Bool{Status: true}, err
			}
		}
	}
	return &pb.Bool{Status: false}, err
}

func (chess *Chess) Update(jwt *pb.Jwt, stream pb.ChessService_UpdateServer) error {
	player, err := chess.JwtSecret.Redeem(jwt.GetJwt())
	if err != nil {
		return err
	}
	pgame, err := chess.InstanceManager.GetGame(player.Id)
	if err != nil {
		return err
	}
	sent := false //bool to only send once

	for stream.Context().Err() == nil && err == nil { //While client is connected
		if pgame.IsTurn(player.Id) && !sent {
			err = stream.Send(&pb.BoardStatus{
				Turn:     false,
				Opponent: nil,
				Board:    pgame.GetBoard(),
			})
			sent = true
		} else {
			sent = false
		}
	}

	if err != nil {
		return err
	}
	return stream.Context().Err()
}

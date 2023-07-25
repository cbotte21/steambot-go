package internal

import (
	"context"
	"github.com/cbotte21/chess-go/internal/game/templates"
	"github.com/cbotte21/chess-go/pb"
	"github.com/cbotte21/chess-go/schema"
	"github.com/cbotte21/microservice-common/pkg/jwtParser"
)

type Chess struct {
	JwtSecret *jwtParser.JwtSecret
	pb.UnimplementedChessServiceServer
}

func NewChess(secret *jwtParser.JwtSecret) Chess {
	return Chess{JwtSecret: secret}
}

// Create is an internal function, should only be called by queue. Once per game.
func (chess *Chess) Create(ctx context.Context, createRequest *pb.CreateRequest) (*pb.CreateResponse, error) {
	p1 := createRequest.Player1.GetXId()
	p2 := createRequest.Player2.GetXId()
	var err error = nil

	//Check if in game

	//Get initial game data
	game := schema.CachedGame{
		PlayerOne: p1,
		PlayerTwo: p2,
		Board:     templates.GetTemplate("Default"),
		Ranked:    false,
		Turn:      false,
	}

	//Append new game to redis

	return &pb.CreateResponse{Status: err == nil}, err
}

func (chess *Chess) Play(ctx context.Context, moveRequest *pb.MoveRequest) (*pb.Bool, error) {
	/*
		Should be a stream!
		1) Link requesting user's _id to a game.
		2) On move, verify and Publish.
		3) Subscribe to game id, sent opponent move to user and save relevant data.
	*/

	return &pb.Bool{Status: false}, nil
}

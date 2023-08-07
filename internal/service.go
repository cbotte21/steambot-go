package internal

import (
	"context"
	"errors"
	"github.com/cbotte21/chess-go/internal/game"
	"github.com/cbotte21/chess-go/internal/game/position"
	"github.com/cbotte21/chess-go/internal/game/templates"
	"github.com/cbotte21/chess-go/pb"
	"github.com/cbotte21/chess-go/schema"
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/jwtParser"
	"io"
)

type Chess struct {
	JwtRedeemer    *jwtParser.JwtSecret
	GameCache      *datastore.RedisClient[schema.CachedGame]
	SvcRecordCache *datastore.RedisClient[schema.SVCRecord]
	pb.UnimplementedChessServiceServer
}

func NewChess(jwtRedeemer *jwtParser.JwtSecret,
	cachedGame *datastore.RedisClient[schema.CachedGame],
	svcRecordCache *datastore.RedisClient[schema.SVCRecord]) Chess {
	return Chess{JwtRedeemer: jwtRedeemer, GameCache: cachedGame, SvcRecordCache: svcRecordCache}
}

// Create is an internal function, should only be called by queue. Once per game.
func (chess *Chess) Create(ctx context.Context, createRequest *pb.CreateRequest) (*pb.CreateResponse, error) {
	gameInstance := game.NewGame(templates.Default)

	//Initial game data
	cachedGame := schema.CachedGame{
		White:  createRequest.White.GetXId(),
		Black:  createRequest.Black.GetXId(),
		Ranked: createRequest.GetRanked(),
		Turn:   false, // ExportUpdate will set to true!
	}
	gameInstance.ExportUpdate(cachedGame)

	// Create game
	err := chess.GameCache.Create(cachedGame)
	if err != nil {
		return &pb.CreateResponse{Status: false}, err
	}

	// Create service records
	err = chess.SvcRecordCache.Create(schema.SVCRecord{Player: createRequest.White.GetXId()})
	if err != nil {
		_ = chess.GameCache.Delete(schema.CachedGame{White: createRequest.White.GetXId()})
		return &pb.CreateResponse{Status: false}, err
	}
	err = chess.SvcRecordCache.Create(schema.SVCRecord{Player: createRequest.Black.GetXId()})
	if err != nil {
		_ = chess.SvcRecordCache.Delete(schema.SVCRecord{Player: createRequest.White.GetXId()})
		_ = chess.GameCache.Delete(schema.CachedGame{White: createRequest.White.GetXId()})
		return &pb.CreateResponse{Status: false}, err
	}

	return &pb.CreateResponse{Status: true}, nil
}

func (chess *Chess) Play(ctx context.Context, stream pb.ChessService_MoveServer) (*pb.MoveResponse, error) {
	key := ""        // Game id
	var isWhite bool // Player is on team white

	for { // Handle move requests
		moveRequest, err := stream.Recv()
		if err == io.EOF {
			return nil, errors.New("stream terminated")
		}
		if err != nil {
			return nil, errors.New("internal server error")
		}

		// First time, sync game from cache, register move handler
		if key == "" {
			key, isWhite, err = chess.sync(moveRequest)
			// Move handler
			go func() {
				sub := chess.GameCache.Subscribe(key)
				for range sub.Channel() {
					break
				}
			}()
		}

		// Attempt move
		if moveRequest.Initial != nil && moveRequest.Final != nil {
			chess.move(moveRequest, isWhite, key, &stream)
		}
	}
}

func (chess *Chess) sync(moveRequest *pb.MoveRequest) (string, bool, error) {
	jwtClaim, err := chess.JwtRedeemer.Redeem(moveRequest.Jwt.GetJwt())
	if err != nil {
		return "", false, err // Player is not logged in
	}

	gameKey, err := chess.SvcRecordCache.Find(schema.SVCRecord{Player: jwtClaim.Id})
	if err != nil {
		return "", false, err // Player is not in game
	}
	key := gameKey.Game

	cachedGame, err := chess.GameCache.Find(schema.CachedGame{White: key})

	return key, cachedGame.White == gameKey.Player, err
}

func (chess *Chess) move(moveRequest *pb.MoveRequest, isWhite bool, key string, stream *pb.ChessService_MoveServer) {
	initial := position.Position{X: int(moveRequest.Initial.GetX()), Y: int(moveRequest.Initial.GetY())}
	final := position.Position{X: int(moveRequest.Final.GetX()), Y: int(moveRequest.Final.GetY())}

	cachedGame, err := chess.GameCache.Find(schema.CachedGame{White: key})

	var whitesTurn bool = cachedGame.Turn
	isTurn := whitesTurn && isWhite || !whitesTurn && !isWhite

	gameInstance := game.LoadGame(&cachedGame)
	if !isTurn {
		err = errors.New("invalid turn sequence")
	} else {
		err = gameInstance.Move(initial, final, isWhite)
	}

	if err == nil { // Valid move
		// Push to redis
		err = chess.GameCache.Create(gameInstance.ExportUpdate(cachedGame))

		if err == nil { // Publish move to pub-sub
			err = chess.GameCache.Publish(key, "")
			if err == nil {
				return // Successful move
			}
			// Revert changes to keep server in sync with clients
			_ = chess.GameCache.Create(cachedGame)
		}
	}

	// Invalid move, notify err
	_ = (*stream).Send(&pb.MoveResponse{
		Turn:  isTurn,
		State: gameInstance,
	})
}

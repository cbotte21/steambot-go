[string]$CHESS_DIR = $Env:CHESS_DIR
[string]$PROJ_DIR = $Env:PROJ_DIR

Start-Process -FilePath "protoc" -ArgumentList "--go_out=$CHESS_DIR/pb --go_opt=paths=source_relative --go-grpc_out=$CHESS_DIR/pb --go-grpc_opt=paths=source_relative --proto_path=$PROJ_DIR/proto/ chess.proto"

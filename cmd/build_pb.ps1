[string]$PROJ_DIR = $Env:PROJ_DIR
[string]$PROJ_NAME = $Env:PROJ_NAME

Start-Process -FilePath "protoc" -ArgumentList "--go_out=$PROJ_DIR/pb --go_opt=paths=source_relative --go-grpc_out=$PROJ_DIR/pb --go-grpc_opt=paths=source_relative --proto_path=$PROJ_DIR/proto/ $PROJ_NAME.proto"

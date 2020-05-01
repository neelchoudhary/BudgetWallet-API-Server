protoc services/auth/auth.proto -I=. --go_opt=paths=source_relative --go_out=plugins=grpc:.
protoc services/shared/shared.proto -I=. --go_opt=paths=source_relative --go_out=plugins=grpc:.
protoc services/plaidfinances/plaidFinances.proto -I=. --go_opt=paths=source_relative --go_out=plugins=grpc:.
protoc services/userfinances/userFinances.proto -I=. --go_opt=paths=source_relative --go_out=plugins=grpc:. 
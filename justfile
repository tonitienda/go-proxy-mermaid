start-servers:
    cd proxy && PORT=8080 SERVICE_NAME="S1" NEXT_SERVICE=http://localhost:8081 go run cmd/main.go &
    cd proxy && PORT=8081 SERVICE_NAME="S2" NEXT_SERVICE=http://localhost:8082 go run cmd/main.go &
    cd proxy && PORT=8082 SERVICE_NAME="S3" NEXT_SERVICE=http://localhost:8083 go run cmd/main.go &
    cd proxy && PORT=8083 SERVICE_NAME="S4" NEXT_SERVICE=http://localhost:8084 go run cmd/main.go &
    cd proxy && PORT=8084 SERVICE_NAME="S5" go run cmd/main.go &
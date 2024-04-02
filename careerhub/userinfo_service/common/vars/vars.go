package vars

import (
	"fmt"
	"os"
	"strconv"
)

type DBUser struct {
	Username string
	Password string
}

type Vars struct {
	MongoUri        string
	DbName          string
	DBUser          *DBUser
	RestApiGrpcPort int
}

type ErrNotExistedVar struct {
	VarName string
}

func NotExistedVar(varName string) *ErrNotExistedVar {
	return &ErrNotExistedVar{VarName: varName}
}

func (e *ErrNotExistedVar) Error() string {
	return fmt.Sprintf("%s is not existed", e.VarName)
}

func Variables() (*Vars, error) {
	mongoUri, err := getFromEnv("MONGO_URI")
	if err != nil {
		return nil, err
	}

	dbUsername := getFromEnvPtr("DB_USERNAME")
	dbPassword := getFromEnvPtr("DB_PASSWORD")

	var dbUser *DBUser
	if dbUsername != nil && dbPassword != nil {
		dbUser = &DBUser{
			Username: *dbUsername,
			Password: *dbPassword,
		}
	}

	dbName, err := getFromEnv("DB_NAME")
	if err != nil {
		return nil, err
	}

	restApiGrpcPort, err := getFromEnv("RESTAPI_GRPC_PORT")
	if err != nil {
		return nil, err
	}

	restApiGrpcPortInt, err := strconv.ParseInt(restApiGrpcPort, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("RESTAPI_GRPC_PORT is not integer.\tREST_API_PORT: %s", restApiGrpcPort)
	}

	return &Vars{
		MongoUri:        mongoUri,
		DBUser:          dbUser,
		DbName:          dbName,
		RestApiGrpcPort: int(restApiGrpcPortInt),
	}, nil
}

func getFromEnv(envVar string) (string, error) {
	ev := os.Getenv(envVar)

	if ev == "" {
		return "", fmt.Errorf("%s is not existed", envVar)
	}

	return ev, nil
}

func getFromEnvPtr(envVar string) *string {
	ev := os.Getenv(envVar)

	if ev == "" {
		return nil
	}

	return &ev
}

package main

import (
	"context"

	"github.com/jae2274/careerhub-userinfo-service/careerhub/userinfo_service/app"
)

func main() {
	app.Run(context.Background())
}

package main

import (
	_ "be-golang-chapter-46-implem/docs"
	"be-golang-chapter-46-implem/infra"
	"be-golang-chapter-46-implem/router"
)

// @title shipping API
// @version 1.0
// @description this is a shipping api
// @termOfService
// @contact.name lumoshive academy
// @contact.url https://shipping.com/
// @contact.email lumoshice.academy@gmail.com
// @host localhost:8080
// @schemes http
// @basePath /v1/api
// @securityDefinitions.apikey ApiKeyAuthToken
// @in header
// @name token
// @securityDefinitions.apikey ApiKeyAuthIDKey
// @in header
// @name ID-KEY
func main() {
	ctx, err := infra.NewContext()
	if err != nil {
		ctx.Log.Panic("error:") // error panic
	}

	router.SetupReouter(ctx, ctx.Middleware)
	// commit form feature a
}

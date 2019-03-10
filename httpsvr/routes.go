package httpsvr

import ()

func initializeRoutes() {

	// Group user related routes together
	userRoutes := router.Group("/api")
	{

		userRoutes.GET("/listUser", listUser)

	}
}

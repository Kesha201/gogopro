package main

import (
	"os"

	routes "jwt-go-token/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// API-1
	router.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	// API-2
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	// API-3
    	router.GET("/api-3", func(c *gin.Context) {
    		c.JSON(200, gin.H{"success": "Access granted for api-3"})
    	})
    // API-4
    	router.GET("/api-4", func(c *gin.Context) {
    		c.JSON(200, gin.H{"success": "Access granted for api-4"})
    	})
    // API-5
        router.GET("/api-5", func(c *gin.Context) {
            c.JSON(200, gin.H{"success": "Access granted for api-5"})
        	})
   // API-6
            	        	router.GET("/api-6", func(c *gin.Context) {
                    		c.JSON(200, gin.H{"success": "Access granted for api-6"})
                    	})
   // API-7
                                    	        	router.GET("/api-7", func(c *gin.Context) {
                                            		c.JSON(200, gin.H{"success": "Access granted for api-7"})
                                            	})
	router.Run(":" + port)
}

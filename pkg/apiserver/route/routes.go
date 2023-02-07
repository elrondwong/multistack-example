package route

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "time"
)

var InstallRouteFuncs []func(*gin.Engine, *gin.RouterGroup)

func InstallRoutes(r *gin.Engine) {
    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    r.Use(gin.Recovery())
    rootGroup := r.Group("/multistack/apis/v1", cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "OPTIONS", "POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders:     []string{"X-CSRF-Token", "Authorization"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
    for _, f := range InstallRouteFuncs {
        f(r, rootGroup)
    }
}

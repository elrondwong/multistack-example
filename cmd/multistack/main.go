package main

import (
    "github.com/elrondwong/multistack-example/pkg/apiserver/route"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()
    route.InstallRoutes(r)
    serverBindAddr := "0.0.0.0:8800"

    log.Printf("Run server at %s", serverBindAddr)

    //go func() {
        if err := r.Run(serverBindAddr); err != nil {
            log.Printf("Failed to run server: %v", err)
            return
        }
    //}()


}

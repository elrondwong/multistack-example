package tencents3

import (
    "github.com/elrondwong/multistack-example/pkg/apiserver/controller"
    "github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine, rootGroup *gin.RouterGroup) {
    qs3Controller := controller.NewQclouds3Controller()
    rootGroup.POST("/qs3bucket" , qs3Controller.CreateBucket)
}


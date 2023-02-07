package controller

import (
    "fmt"
    "github.com/elrondwong/multistack-example/pkg/apiserver/domain/qs3"
    "github.com/gin-gonic/gin"
)

type Qclouds3Controller struct {
    qs3Service *qs3.Service
}

func NewQclouds3Controller() *Qclouds3Controller {
    return &Qclouds3Controller{
        qs3Service: qs3.NewQs3Service(),
    }
}

func (qc *Qclouds3Controller) CreateBucket(c *gin.Context) {
    type name struct {
        Name string `json:"name"`
    }
    var bucketName name
    if err := c.ShouldBindJSON(&bucketName); err != nil {
        fmt.Printf("Error binding")
        return
    }

    bucket, err := qc.qs3Service.CreateBucket(bucketName.Name)
    if  err != nil {
        fmt.Printf("bucket create error: %v", err)
        return
    }

    c.JSON(200, fmt.Sprintf("bucket name: %v", bucket))

}
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"

	"github.com/bokjo/grpc-kube-test/proto"
)

func main() {
	conn, err := grpc.Dial("nzd-service:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}

	nzdClient := nzd.NewNZDServiceClient(conn)

	r := gin.Default()

	// Endpoint handler
	r.GET("/nzd/:a/:b", func(c *gin.Context) {

		a, err := strconv.ParseUint(c.Param("a"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
			return
		}

		b, err := strconv.ParseUint(c.Param("b"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter B"})
			return
		}

		serviceReq := &nzd.NZDRequest{A: a, B: b}

		if result, err := nzdClient.Compute(c, serviceReq); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(result.Result),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	// Run the server...
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run the API server: %v", err)
	}
}

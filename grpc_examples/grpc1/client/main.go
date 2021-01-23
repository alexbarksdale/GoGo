package main

import (
	"fmt"
	"grpc_tut/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	g := gin.Default()

	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param A"})
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param B"})
		}

		req := &proto.Request{A: int64(a), B: int64(b)}

		if resp, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(resp.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"ereror": err.Error(),
			})
		}
	})

	g.GET("/multiply/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param A"})
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param B"})
		}

		req := &proto.Request{A: int64(a), B: int64(b)}

		if resp, err := client.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(resp.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"ereror": err.Error(),
			})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
}

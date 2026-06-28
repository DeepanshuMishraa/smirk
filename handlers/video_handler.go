package handlers

import (
	"errors"
	"net/http"

	"github.com/DeepanshuMishraa/vid-processing-go.git/models"
	"github.com/DeepanshuMishraa/vid-processing-go.git/repository"
	"github.com/DeepanshuMishraa/vid-processing-go.git/types"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateVideoHandler(conn *amqp.Connection, db *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateVideoRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		if req.VideoID == "" || req.Title == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "video_id and title are required"})
			return
		}

		video := models.Video{
			ID:          req.VideoID,
			Title:       req.Title,
			OriginalURL: req.OriginalUrl,
		}

		if err := repository.CreateVideo(conn, db, video); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, types.CreateVideoResponse{
			VideoID: video.ID,
			Status:  string(models.UPLOADED),
		})
	}
}

func GetAllVideosHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		videos, err := repository.GetAllVideos(db)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, videos)
	}
}

func GetVideoByIdHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "video id is required"})
			return
		}

		video, err := repository.GetVideoById(db, id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "video not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, video)
	}
}

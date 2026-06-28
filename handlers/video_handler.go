package handlers

import (
	"errors"
	"net/http"

	"github.com/DeepanshuMishraa/vid-processing-go.git/models"
	"github.com/DeepanshuMishraa/vid-processing-go.git/repository"
	"github.com/DeepanshuMishraa/vid-processing-go.git/services"
	"github.com/DeepanshuMishraa/vid-processing-go.git/types"
	"github.com/DeepanshuMishraa/vid-processing-go.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateVideoHandler(conn *amqp.Connection, db *pgxpool.Pool, r2Svc *types.R2Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		videoID := ctx.PostForm("video_id")
		title := ctx.PostForm("title")

		if title == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
			return
		}

		if videoID == "" {
			videoID = uuid.New().String()
		}

		file, header, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
			return
		}
		defer file.Close()

		tmpFile, err := utils.CreateTempFile(file, header.Filename)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer utils.RemoveFile(tmpFile.Name())

		video := models.Video{
			ID:    videoID,
			Title: title,
		}

		if err := services.CreateVideo(conn, db, r2Svc, video, tmpFile.Name()); err != nil {
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



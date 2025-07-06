package routers

import (
	"errors"
	"net/http"
	"server/db"
	"server/models"
	"server/schemas"
	"server/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterPostRoutes(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		userID := c.QueryParam("user_id") // or get from context/session
		var posts []models.Post
		if err := db.DB.Find(&posts).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		responses := make([]schemas.PostResponse, len(posts))
		for i, p := range posts {
			var vote models.Vote
			err := db.DB.Where("post_id = ? AND user_id = ?", p.ID, userID).First(&vote).Error
			responses[i] = schemas.PostResponse{Post: p}
			if err == nil {
				if vote.IsUp {
					responses[i].Upvoted = true
				} else {
					responses[i].Downvoted = true
				}
			}
		}

		return c.JSON(http.StatusOK, echo.Map{
			"ok":            true,
			"posts":         responses,
			"rows_affected": int64(len(posts)),
		})
	})

	g.POST("", func(c echo.Context) error {
		var body schemas.CreatePostReq
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
		}

		if err := utils.ValidateStruct(body); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		post := models.Post{Content: body.Content}
		result := db.DB.Create(&post)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": result.Error})
		}

		return c.JSON(http.StatusOK, echo.Map{"ok": true, "post": post, "rows_affected": result.RowsAffected, "error": result.Error})	
	})

	g.POST("/vote", func(c echo.Context) error {
		var body schemas.CreateVoteReq
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		if err := utils.ValidateStruct(body); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		vote := models.Vote{
			PostID: uint(body.PostId),
			UserID: body.UserId,
			IsUp:   *body.IsUp,
		}

		var existing models.Vote
		err := db.DB.Where("post_id = ? AND user_id = ?", body.PostId, body.UserId).First(&existing).Error
		if err == nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "vote already exists"})
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "db error"})
		}

		result := db.DB.Create(&vote)

		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": result.Error})
		}

		var post models.Post
		if err := db.DB.First(&post, body.PostId).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "post not found"})
		}

		if *body.IsUp {
			db.DB.Model(&post).UpdateColumn("upvotes", gorm.Expr("upvotes + ?", 1))
		} else {
			db.DB.Model(&post).UpdateColumn("downvotes", gorm.Expr("downvotes + ?", 1))
		}

		return c.JSON(http.StatusOK, echo.Map{"ok": true, "rows_affected": result.RowsAffected, "error": result.Error})	
	})

	g.DELETE("/vote", func(c echo.Context) error {
		var body schemas.DeleteVoteReq
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
		}

		if err := utils.ValidateStruct(body); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		var vote models.Vote
		if err := db.DB.Where("post_id = ? AND user_id = ?", body.PostId, body.UserId).First(&vote).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		result := db.DB.Unscoped().Delete(&vote)

		if vote.IsUp {
			db.DB.Model(&models.Post{}).Where("id = ?", vote.PostID).UpdateColumn("upvotes", gorm.Expr("upvotes - ?", 1))
		} else {
			db.DB.Model(&models.Post{}).Where("id = ?", vote.PostID).UpdateColumn("downvotes", gorm.Expr("downvotes - ?", 1))
		}

		return c.JSON(http.StatusOK, echo.Map{"ok": true, "rows_affected": result.RowsAffected, "error": result.Error})	
	})
}
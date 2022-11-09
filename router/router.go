package router

import (
	"myapp/dto"
	"myapp/middleware"
	"myapp/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WebRouter(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware(), middleware.AuthMiddleware())

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "berhasil",
		})
	})

	router.POST("/login", func(ctx *gin.Context) {
		var input dto.LoginUser

		err := ctx.ShouldBind(&input)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		token, err := service.LoginUser(ctx, input)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})

	router.POST("/new-user", func(ctx *gin.Context) {
		var input dto.NewUser

		err := ctx.BindJSON(&input)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		user, err := service.CreateUser(ctx, input)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, user)
	})

	router.GET("/get-all-user", func(ctx *gin.Context) {
		users, _ := service.GetAllUser(ctx)

		ctx.JSON(http.StatusOK, users)
	})

	router.GET("/user/:userID", func(ctx *gin.Context) {
		userID := ctx.Param("userID")

		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			panic(err)
		}

		user, err := service.GetUserById(ctx, userIDInt)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "user not found",
			})
			return
		}

		ctx.JSON(http.StatusOK, user)
	})

	authRouter := router.Use(middleware.IsUser())

	authRouter.DELETE("/delete-user", func(ctx *gin.Context) {
		var (
			userID = middleware.AuthCtx(ctx.Request.Context()).ID
		)

		resp, err := service.Deleteuser(ctx, userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, resp)

	})

	authRouter.PUT("/update-user", func(ctx *gin.Context) {
		var (
			userID = middleware.AuthCtx(ctx.Request.Context()).ID
		)

		var input dto.Updateuser

		err := ctx.ShouldBind(&input)
		if err != nil {
			panic(err)
		}

		resp, err := service.Updateuser(ctx, userID, input)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, resp)
	})

	router.POST("/new-post", func(ctx *gin.Context) {
		var input dto.NewPost

		err := ctx.ShouldBind(&input)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		_, err = service.GetUserById(ctx, input.UserId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "user not found",
			})
			return
		}

		post, err := service.CreatePost(ctx, input)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, post)
	})

	router.GET("/get-all-post", func(ctx *gin.Context) {
		posts, _ := service.GetAllPost(ctx)

		ctx.JSON(http.StatusOK, posts)
	})

	router.GET("/get-post/:postID", func(ctx *gin.Context) {
		postID := ctx.Param("postID")
		postIDInt, err := strconv.Atoi(postID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		}

		post, err := service.GetPostById(ctx, postIDInt)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, post)
	})

	authRouter.DELETE("/delete/post/:postID", func(ctx *gin.Context) {
		var (
			userID = middleware.AuthCtx(ctx.Request.Context()).ID
		)

		postID := ctx.Param("postID")
		postIDInt, err := strconv.Atoi(postID)
		if err != nil {
			panic(err)
		}

		resp, err := service.DeletePost(ctx, postIDInt, userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, resp)
	})

	authRouter.PUT("/update-post/:postID", func(ctx *gin.Context) {
		var (
			userID = middleware.AuthCtx(ctx.Request.Context()).ID
		)
		postID := ctx.Param("postID")
		postIDInt, err := strconv.Atoi(postID)
		if err != nil {
			panic(err)
		}

		var input dto.Updatepost

		//err = ctx.ShouldBind(&input)

		resp, err := service.UpdatePost(ctx, postIDInt, input, userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, resp)

	})

	router.POST("/new-product", func(ctx *gin.Context) {
		var input dto.NewProduct

		err := ctx.BindJSON(&input)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		product, err := service.CreateProduct(ctx, input)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, product)
	})

	router.GET("/get-all-product", func(ctx *gin.Context) {
		products, _ := service.GetAllProduct(ctx)

		ctx.JSON(http.StatusOK, products)
	})
}

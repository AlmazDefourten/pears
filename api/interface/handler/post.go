package handler

import (
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models/post_models"
	"github.com/AlmazDefourten/goapp/models/requests_models"
	"github.com/golobby/container/v3"
	"github.com/kataras/iris/v12"
)

// PostHandler - handler for handle requests with post/s info
type PostHandler struct {
}

func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

// List takes array of posts
//
//	@Summary		List of posts
//	@Description	takes array of posts
//	@Accept			json
//	@Produce		json
//	@Failure		401	{object}	requests_models.Response
//	@Success		200	{object}	[]post_models.Post
//	@Router			/post/list [get]
//  @Security 		JWTToken
func (postHandler *PostHandler) List(ctx iris.Context) {
	var postService post_models.IPostService
	err := container.Resolve(&postService)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		err = ctx.JSON(requests_models.Response{Ok: false, Message: requests_models.StandardAnswerOnError})
		if err != nil {
			logger_instance.GlobalLogger.Error(err)
		}
		return
	}
	data, err := postService.ListPosts()
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		err = ctx.JSON(requests_models.Response{Ok: false, Message: requests_models.StandardAnswerOnError})
		if err != nil {
			logger_instance.GlobalLogger.Error(err)
		}
		return
	}
	err = ctx.JSON(data)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
}

// Get post by identificator
//
//	@Summary		Get a post
//	@Description	get post info
//	@Accept			json
//	@Produce		json
//  @Param        	id   	path      int  	true  "post ID"
//	@Failure		401	{object}	requests_models.Response
//	@Success		200	{object}	requests_models.Response
//	@Router			/post/get/{id} [get]
//  @Security 		JWTToken
func (postHandler *PostHandler) Get(ctx iris.Context) {
	var postService post_models.IPostService
	err := container.Resolve(&postService)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		err = ctx.JSON(requests_models.Response{Ok: false, Message: requests_models.StandardAnswerOnError})
		if err != nil {
			logger_instance.GlobalLogger.Error(err)
		}
		return
	}

	id, err := ctx.URLParamInt("id")
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		err = ctx.JSON(requests_models.Response{Ok: false, Message: requests_models.StandardAnswerOnError})
		return
	}
	data, err := postService.GetPost(id)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		err = ctx.JSON(requests_models.Response{Ok: false, Message: requests_models.StandardAnswerOnError})
		if err != nil {
			logger_instance.GlobalLogger.Error(err)
		}
		return
	}
	err = ctx.JSON(data)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
}

// Create ShowAccount godoc
//
//	@Summary		Create new post
//	@Description	creating a new post with body parameters
//	@Accept			json
//	@Produce		json
//	@Param			body		body		post_models.Post		true	"request body with info about post"
//	@Failure		401	{object}	requests_models.Response
//	@Success		200	{object}	requests_models.Response
//	@Router			/post/create [post]
//  @Security 		JWTToken
func (postHandler *PostHandler) Create(ctx iris.Context) {
	var postService post_models.IPostService
	err := container.Resolve(&postService)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		err = ctx.JSON(requests_models.Response{Ok: false, Message: requests_models.StandardAnswerOnError})
		if err != nil {
			logger_instance.GlobalLogger.Error(err)
		}
		return
	}

	var post post_models.Post
	err = ctx.ReadJSON(&post)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		err = ctx.JSON(requests_models.Response{Ok: false, Message: requests_models.StandardAnswerOnError})
		return
	}
	err = postService.CreatePost(post)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		err = ctx.JSON(requests_models.Response{Ok: false, Message: requests_models.StandardAnswerOnError})
		if err != nil {
			logger_instance.GlobalLogger.Error(err)
		}
		return
	}
	err = ctx.JSON(requests_models.Response{Ok: true, Message: "Пост успешно создан"})
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
}

package api

import (
	"github.com/gin-gonic/gin"
	"leonzhangxf-api/db"
	"leonzhangxf-api/util"
	"log"
	"strconv"
)

func InitArticle() {
	Engine.GET("/open/v1/articles", GetArticles)
	Engine.GET("/open/v1/articles/:id", GetArticle)
}

func GetArticles(context *gin.Context) {
	articles := db.QueryArticles()
	var articleOpenDtoList []ArticleOpenDto
	if len(articles) == 0 {
		context.JSON(404, "Not found")
		return
	}
	for _, article := range articles {
		dto := ConvertToArticleOpenDto(&article)
		articleOpenDtoList = append(articleOpenDtoList, *dto)
	}
	context.JSON(200, &articleOpenDtoList)
}

func GetArticle(ctx *gin.Context) {
	idParam := ctx.Param("id")
	articleId, err := strconv.ParseInt(idParam, 10, 64)
	if nil != err {
		log.Printf("The id param is %v", idParam)
		ctx.JSON(400, "Bad article id param")
		return
	}
	article := db.GetArticleById(articleId)
	if article.Id == 0 {
		ctx.JSON(404, "Not found")
		return
	}
	dto := ConvertToArticleOpenDto(article)
	ctx.JSON(200, &dto)
}

func ConvertToArticleOpenDto(article *db.Article) *ArticleOpenDto {
	articleDto := &ArticleOpenDto{}
	articleDto.Id = article.Id
	articleDto.Title = article.Title
	articleDto.PreviewImg = article.PreviewImg
	articleDto.Desc = article.Desc
	articleDto.Content = article.Content
	log.Printf("The gmt publish is %v", article.GmtPublish)
	if nil != article.GmtPublish {
		articleDto.GmtPublish = &util.JsonTime{Time: *article.GmtPublish}
	}
	articleDto.GmtCreate = &util.JsonTime{Time: *article.GmtCreate}
	articleDto.GmtModified = &util.JsonTime{Time: *article.GmtModified}
	return articleDto
}

type ArticleOpenDto struct {
	Id          int64
	Title       string
	PreviewImg  string
	Desc        string
	Content     string
	GmtPublish  *util.JsonTime `json:",omitempty"`
	GmtCreate   *util.JsonTime
	GmtModified *util.JsonTime
}

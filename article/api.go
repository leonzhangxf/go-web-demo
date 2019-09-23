package article

import (
	"github.com/gin-gonic/gin"
	"leonzhangxf-api/util"
	"net/http"
	"strconv"
)

type Api struct {
}

// @Summary Query published articles
// @Description Query published articles
// @Produce  json
// @Success 200 {array} article.OpenDto
// @Failure 404 {string} Not found
// @Router /open/v1/articles [get]
func (api *Api) GetOpenArticles(context *gin.Context) {
	articles := QueryPublishedArticles()
	var articleOpenDtoList []OpenDto
	if len(articles) == 0 {
		context.JSON(http.StatusNotFound, "Not found")
		return
	}
	for _, article := range articles {
		dto := ConvertToOpenDto(&article)
		articleOpenDtoList = append(articleOpenDtoList, *dto)
	}
	context.JSON(http.StatusOK, &articleOpenDtoList)
}

// @Summary Get published article by id
// @Description Get published article by id
// @Produce json
// @Param id path int true "Article Id"
// @Success 200 {object} article.OpenDto
// @Failure 400 {string} Bad article id param
// @Failure 404 {string} Not found
// @Router /open/v1/articles/{id} [get]
func (api *Api) GetOpenArticleById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	articleId, err := strconv.ParseInt(idParam, 10, 64)
	if nil != err {
		util.Log.Printf("The id param is %v", idParam)
		ctx.JSON(http.StatusBadRequest, "Bad article id param")
		return
	}
	article := GetPublishedArticleById(articleId)
	if article.Id == 0 {
		ctx.JSON(http.StatusNotFound, "Not found")
		return
	}
	dto := ConvertToOpenDto(article)
	ctx.JSON(http.StatusOK, &dto)
}

// @Summary Query articles
// @Description Query articles, include no published articles
// @Produce  json
// @Success 200 {array} article.Dto
// @Failure 404 {string} Not found
// @Router /api/v1/articles [get]
func (api *Api) GetArticles(context *gin.Context) {
	articles := QueryArticles()
	var articleDtoList []Dto
	if len(articles) == 0 {
		context.JSON(http.StatusNotFound, "Not found")
		return
	}
	for _, article := range articles {
		dto := ConvertToArticleDto(&article)
		articleDtoList = append(articleDtoList, *dto)
	}
	context.JSON(http.StatusOK, &articleDtoList)
}

func ConvertToArticleDto(article *Article) *Dto {
	articleDto := &Dto{}
	articleDto.Id = article.Id
	articleDto.Title = article.Title
	articleDto.PreviewImg = article.PreviewImg
	articleDto.Desc = article.Desc
	articleDto.Content = article.Content
	articleDto.Status = article.Status
	if nil != article.GmtPublish {
		articleDto.GmtPublish = &util.JsonTime{Time: *article.GmtPublish}
	}
	articleDto.GmtCreate = &util.JsonTime{Time: *article.GmtCreate}
	articleDto.GmtModified = &util.JsonTime{Time: *article.GmtModified}
	return articleDto
}

type Dto struct {
	Id          int64          `json:"id" example:"1"`
	Title       string         `json:"title" example:"文章标题"`
	PreviewImg  string         `json:"previewImg" example:"文章预览图片"`
	Desc        string         `json:"desc" example:"文章描述"`
	Content     string         `json:"content" example:"文章内容"`
	Status      int8           `json:"status" example:"0"`
	GmtPublish  *util.JsonTime `json:"gmtPublish,omitempty" swaggertype:"string" example:"2006-01-02 15:04:05"`
	GmtCreate   *util.JsonTime `json:"gmtCreate" swaggertype:"string" example:"2006-01-02 15:04:05"`
	GmtModified *util.JsonTime `json:"gmtModified" swaggertype:"string" example:"2006-01-02 15:04:05"`
}

func ConvertToOpenDto(article *Article) *OpenDto {
	articleDto := &OpenDto{}
	articleDto.Id = article.Id
	articleDto.Title = article.Title
	articleDto.PreviewImg = article.PreviewImg
	articleDto.Desc = article.Desc
	articleDto.Content = article.Content
	if nil != article.GmtPublish {
		articleDto.GmtPublish = &util.JsonTime{Time: *article.GmtPublish}
	}
	articleDto.GmtCreate = &util.JsonTime{Time: *article.GmtCreate}
	articleDto.GmtModified = &util.JsonTime{Time: *article.GmtModified}
	return articleDto
}

type OpenDto struct {
	Id          int64          `json:"id" example:"1"`
	Title       string         `json:"title" example:"文章标题"`
	PreviewImg  string         `json:"previewImg" example:"文章预览图片"`
	Desc        string         `json:"desc" example:"文章描述"`
	Content     string         `json:"content" example:"文章内容"`
	GmtPublish  *util.JsonTime `json:"gmtPublish,omitempty" swaggertype:"string" example:"2006-01-02 15:04:05"`
	GmtCreate   *util.JsonTime `json:"gmtCreate" swaggertype:"string" example:"2006-01-02 15:04:05"`
	GmtModified *util.JsonTime `json:"gmtModified" swaggertype:"string" example:"2006-01-02 15:04:05"`
}

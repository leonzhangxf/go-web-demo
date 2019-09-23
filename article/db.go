package article

import (
	"leonzhangxf-api/config"
	"leonzhangxf-api/util"
	"time"
)

func QueryArticles() []Article {
	var articles []Article
	sql := "select * from article"
	err := config.Db.Select(&articles, sql)
	if nil != err {
		util.Log.Warningln("QueryArticles err.", err)
	}
	return articles
}

func QueryPublishedArticles() []Article {
	var articles []Article
	sql := "select * from article where status = 1"
	err := config.Db.Select(&articles, sql)
	if nil != err {
		util.Log.Warningln("QueryPublishedArticles err.", err)
	}
	return articles
}

func GetPublishedArticleById(articleId int64) *Article {
	var article Article
	sql := "select * from article where status = 1 and id = ?"
	err := config.Db.QueryRowx(sql, articleId).StructScan(&article)
	if nil != err {
		util.Log.Warningln("GetPublishedArticleById err.", err)
	}
	return &article
}

type Article struct {
	Id          int64      `db:"id"`
	Title       string     `db:"title"`
	PreviewImg  string     `db:"preview_img"`
	Desc        string     `db:"desc"`
	Content     string     `db:"content"`
	Status      int8       `db:"status"`
	GmtPublish  *time.Time `db:"gmt_publish"`
	GmtCreate   *time.Time `db:"gmt_create"`
	GmtModified *time.Time `db:"gmt_modified"`
}

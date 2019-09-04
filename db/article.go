package db

import (
	"log"
	"time"
)

func QueryArticles() []Article {
	var articles []Article
	sql := "select * from article"
	err := Db.Select(&articles, sql)
	if nil != err {
		log.Printf("QueryArticles err. %v", err)
	}
	return articles
}

func GetArticleById(articleId int64) *Article {
	var article Article
	sql := "select * from article where id = ?"
	err := Db.QueryRowx(sql, articleId).StructScan(&article)
	if nil != err {
		log.Printf("GetArticleById err. %v", err)
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

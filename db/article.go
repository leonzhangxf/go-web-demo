package db

import "time"

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

package mysql

import (
	"fmt"
	"hellonil/models"
	"hellonil/pkg/snowflake"
)

func InsertVideo(video *models.Videos) (err error) {
	sqlStr := `insert into videos(id,author_id,play_url,cover_url,favorite_count,comment_count,title,create_time) VALUES (?,?,?,?,?,?,?,?)`
	id := snowflake.GenID()
	_, err = db.Exec(sqlStr, id, video.AuthorID, video.PlayUrl, video.CoverUrl, video.FavoriteCount, video.CommentCount, video.Title, video.CreateTime)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

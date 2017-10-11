package models

import (
	"time"
)

type post struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Link    string    `json:"href"`
}

type params struct {
	title, content, link, start, end string
}

func NewParams(title, content, link, start, end string) *params {
	p := new(params)
	p.title = `%`
	p.content = `%`
	p.link = `%`
	p.start = `1970-01-01`
	p.end = `3000-01-01`
	if title != "" {
		p.title = title
	}
	if content != "" {
		p.content = content
	}
	if link != "" {
		p.link = link
	}
	if start != "1010-01-01" && start != "" {
		p.start = start
	}
	if end != "1010-01-01" && end != "" {
		p.end = end
	}
	return p
}

func AllPosts(params params) ([]*post, error) {
	rows, err := db.Query(
		`SELECT title, summary, datetime, href_original 
		FROM "tableNewsRus" 
		WHERE 
			title LIKE $1 AND 
			summary LIKE $2 AND 
			href_original LIKE $3 AND 
			datetime >= $4 AND
			datetime <= $5`,
		params.title, params.content, params.link, params.start, params.end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]*post, 0)
	for rows.Next() {
		post := new(post)
		err := rows.Scan(&post.Title, &post.Content, &post.Created, &post.Link)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPost(id int) (*post, error) {
	post := new(post)
	db.QueryRow(
		`SELECT title, summary, datetime, href_original FROM "tableNewsRus" 
		WHERE id = $1`, id).
		Scan(&post.Title, &post.Content, &post.Created, &post.Link)
	return post, nil
}

package repository

import (
	m "blog/model"
	"database/sql"
	"strconv"
	"time"
)

var (
	rows *sql.Rows
)

// GetPosts : to get post data from table based on parameters
func GetPosts(db *sql.DB, params []*m.Param) (posts []*m.Post, err error) {
	var values []any

	selectQ := "select p.id, p.title, p.content, p.published_date "
	fromQ := "from post p "
	whereQ := "where 1=1 "

	for i, v := range params {
		if v.Key == "title" {
			whereQ += "and p.title ilike $" + strconv.Itoa(i+1) + " "
			v.Value = "%" + v.Value + "%"
			goto valuesInput
		}

		if v.Key == "content" {
			whereQ += "and p.content ilike $" + strconv.Itoa(i+1) + " "
			v.Value = "%" + v.Value + "%"
			goto valuesInput
		}

		if v.Key == "status" {
			whereQ += "and p.status = $" + strconv.Itoa(i+1) + " "
			goto valuesInput
		}

		if v.Key == "published_date" {
			whereQ += "and date(p.published_date) = $" + strconv.Itoa(i+1) + " "
			goto valuesInput
		}

		if v.Key == "tag" {
			fromQ += "join post_to_tag ptt on p.id = ptt.post_id join tag t on ptt.tag_id = t.id "
			whereQ += "and lower(t.label) = lower($" + strconv.Itoa(i+1) + ") "
			goto valuesInput
		}

	valuesInput:
		values = append(values, v.Value)
	}

	query := selectQ + fromQ + whereQ + " order by p.id"
	rows, err = db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := new(m.Post)
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.PublishedDate)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// GetPostById : to get post data from table filtered by id
func GetPostById(db *sql.DB, id string) (*m.Post, error) {
	var (
		err           error
		publishedDate sql.NullTime
		updatedDate   sql.NullTime
	)

	post := new(m.Post)
	query := "select id, title, content, status, published_date, created_date, updated_date from post where id = $1"
	status := 2

	err = db.QueryRow(query, id).Scan(&post.Id, &post.Title, &post.Content, &status, &publishedDate, &post.CreatedDate, &updatedDate)
	if err != nil {
		return nil, err
	}

	if publishedDate.Valid {
		post.PublishedDate = &publishedDate.Time
	}

	if status == 1 {
		post.Status = "draft"
	} else {
		post.Status = "publish"
	}

	query = "select t.id, t.label from post_to_tag ptt join tag t on ptt.tag_id = t.id where ptt.post_id = $1"
	rows, err = db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tag := new(m.Tag)
		err = rows.Scan(&tag.Id, &tag.Label)
		if err != nil {
			return nil, err
		}

		post.Tags = append(post.Tags, tag)
	}

	return post, nil
}

// AddPost : to add new post data into table
func AddPost(db *sql.DB, post *m.Post) error {
	var (
		tx    *sql.Tx
		err   error
		query string
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query = "insert into post (title, content) values ($1, $2) returning id"
	err = tx.QueryRow(query, post.Title, post.Content).Scan(&post.Id)
	if err != nil {
		return err
	}

	for _, v := range post.Tags {
		query = "insert into post_to_tag (post_id, tag_id) values ($1, $2)"
		_, err = tx.Exec(query, post.Id, v.Id)
		if err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}

// ModifyPost : to change existing post data in data
func ModifyPost(db *sql.DB, post *m.Post) error {
	var (
		tx    *sql.Tx
		err   error
		query string
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query = "update post set title = $1, content = $2, updated_date = $3 where id = $4"
	_, err = tx.Exec(query, post.Title, post.Content, time.Now(), post.Id)
	if err != nil {
		return err
	}

	query = "delete from post_to_tag where post_id = $1"
	_, err = tx.Exec(query, post.Id)
	if err != nil {
		return err
	}

	for _, v := range post.Tags {
		query = "insert into post_to_tag (post_id, tag_id) values ($1, $2)"
		_, err = tx.Exec(query, post.Id, v.Id)
		if err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}

// RemovePost : to remove post data from table
func RemovePost(db *sql.DB, id string) error {
	var (
		tx    *sql.Tx
		err   error
		query string
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query = "delete from post_to_tag where post_id = $1"
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	query = "delete from post where id = $1"
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

// PublishPost : to change status post data from draft into publish in table
func PublishPost(db *sql.DB, id string) error {
	query := "update post set status = $1, published_date = $2 where id = $3"
	_, err := db.Exec(query, 2, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

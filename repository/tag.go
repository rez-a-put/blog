package repository

import (
	m "blog/model"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

// GetTagById : to get tag data from table filtered by paramaters
func GetTags(db *sql.DB, params []*m.Param) (tags []*m.Tag, err error) {
	var values []any

	selectQ := "select t.id, t.label "
	fromQ := "from tag t "
	whereQ := "where 1=1 "

	for i, v := range params {
		if v.Key == "tag" {
			whereQ += "and t.label ilike $" + strconv.Itoa(i+1) + " "
			v.Value = "%" + v.Value + "%"
			goto valuesInput
		}

	valuesInput:
		values = append(values, v.Value)
	}

	query := selectQ + fromQ + whereQ + " order by t.id"
	fmt.Println(query, values)
	rows, err = db.Query(query, values...)
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

		tags = append(tags, tag)
	}

	return tags, nil
}

// GetTagByLabel : to get tag data filtered by label from table
func GetTagByLabel(db *sql.DB, label string) (*m.Tag, error) {
	var err error

	tag := new(m.Tag)
	query := "select id, label from tag where label = $1"

	err = db.QueryRow(query, label).Scan(&tag.Id, &tag.Label)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

// GetTagById : to get tag data filtered by id from table
func GetTagById(db *sql.DB, id string) (*m.Tag, error) {
	var err error

	tag := new(m.Tag)
	query := "select id, label from tag where id = $1"

	err = db.QueryRow(query, id).Scan(&tag.Id, &tag.Label)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

// AddTag : to add new tag data into table
func AddTag(db *sql.DB, tag *m.Tag) error {
	query := "insert into tag (label) values ($1)"
	_, err := db.Exec(query, tag.Label)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// ModifyTag : to change existing tag data in data
func ModifyTag(db *sql.DB, tag *m.Tag) error {
	query := "update tag set label = $1, updated_date = $2 where id = $3"
	_, err := db.Exec(query, tag.Label, time.Now(), tag.Id)
	if err != nil {
		return err
	}

	return nil
}

// RemoveTag : to remove tag data from table
func RemoveTag(db *sql.DB, id string) error {
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

	query = "delete from post_to_tag where tag_id = $1"
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	query = "delete from tag where id = $1"
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

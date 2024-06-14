package controller

import (
	m "blog/model"
	r "blog/repository"
	u "blog/utils"
	"errors"
	"net/http"
	"strconv"
	"time"
)

// GetPost : to get multiple data of post from database
func (h *BaseHandler) GetPosts(title, content, status, publishedDate, tag string) (retDatas []*m.Post, statusCode int, err error) {
	var params []*m.Param

	if title != "" {
		params = append(params, &m.Param{Key: "title", Value: title})
	}

	if content != "" {
		params = append(params, &m.Param{Key: "content", Value: content})
	}

	if status != "" {
		params = append(params, &m.Param{Key: "status", Value: status})
	}

	if publishedDate != "" {
		_, err = time.Parse("2006-01-02", publishedDate)
		if err != nil {
			err = errors.New("date format should be yyyy-mm-dd")
			return nil, http.StatusBadRequest, err
		}

		params = append(params, &m.Param{Key: "published_date", Value: publishedDate})
	}

	if tag != "" {
		params = append(params, &m.Param{Key: "tag", Value: tag})
	}

	retDatas, err = r.GetPosts(h.db, params)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return retDatas, http.StatusOK, nil
}

// GetPost : to get data of post from database
func (h *BaseHandler) GetPost(id string) (retData *m.Post, statusCode int, err error) {
	retData, err = r.GetPostById(h.db, id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return retData, http.StatusOK, nil
}

// AddPost : to add new data of post into database
func (h *BaseHandler) AddPost(reqData *m.ReqDataPost) (statusCode int, err error) {
	post := &m.Post{
		Title:   reqData.Title,
		Content: reqData.Content,
	}

	for _, v := range reqData.Tags {
		tag, err := r.GetTagByLabel(h.db, v)
		if err != nil {
			err = errors.New(u.DataNotExist("tag " + v))
			return http.StatusBadRequest, err
		}

		post.Tags = append(post.Tags, &m.Tag{Id: tag.Id})
	}

	err = r.AddPost(h.db, post)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusCreated, nil
}

// ModifyPost : to change existing post data in database
func (h *BaseHandler) ModifyPost(reqData *m.ReqDataPost) (statusCode int, err error) {
	var id int

	id, err = strconv.Atoi(reqData.Id)
	if err != nil {
		return http.StatusBadRequest, err
	}

	post := &m.Post{
		Id:      int64(id),
		Title:   reqData.Title,
		Content: reqData.Content,
	}

	for _, v := range reqData.Tags {
		tag, err := r.GetTagByLabel(h.db, v)
		if err != nil {
			err = errors.New(u.DataNotExist("tag " + v))
			return http.StatusBadRequest, err
		}

		post.Tags = append(post.Tags, &m.Tag{Id: tag.Id})
	}

	err = r.ModifyPost(h.db, post)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

// RemovePost : to remove post data from database
func (h *BaseHandler) RemovePost(id string) (statusCode int, err error) {
	err = r.RemovePost(h.db, id)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

// PublishPost : to change status post data from draft into publish in database
func (h *BaseHandler) PublishPost(id string) (statusCode int, err error) {
	err = r.PublishPost(h.db, id)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

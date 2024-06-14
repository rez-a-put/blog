package controller

import (
	m "blog/model"
	r "blog/repository"
	"fmt"
	"net/http"
	"strconv"
)

// GetTag : to get multiple data of tag from database
func (h *BaseHandler) GetTags(tag string) (retDatas []*m.Tag, statusCode int, err error) {
	var params []*m.Param

	if tag != "" {
		params = append(params, &m.Param{Key: "tag", Value: tag})
	}

	retDatas, err = r.GetTags(h.db, params)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return retDatas, http.StatusOK, nil
}

// GetTag : to get data of tag from database
func (h *BaseHandler) GetTagById(id string) (retData *m.Tag, statusCode int, err error) {
	retData, err = r.GetTagById(h.db, id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return retData, http.StatusOK, nil
}

// GetTagByLabel : to get data of tag from database based on label filter
func (h *BaseHandler) GetTagByLabel(label string) (retData *m.Tag, statusCode int, err error) {
	retData, err = r.GetTagByLabel(h.db, label)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return retData, http.StatusOK, nil
}

// AddTag : to add new data of tag into database
func (h *BaseHandler) AddTag(reqData *m.ReqDataTag) (statusCode int, err error) {
	tag := &m.Tag{
		Label: reqData.Label,
	}
	fmt.Println(tag)
	err = r.AddTag(h.db, tag)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusCreated, nil
}

// ModifyTag : to change existing tag data in database
func (h *BaseHandler) ModifyTag(reqData *m.ReqDataTag) (statusCode int, err error) {
	var id int

	id, err = strconv.Atoi(reqData.Id)
	if err != nil {
		return http.StatusBadRequest, err
	}

	tag := &m.Tag{
		Id:    int64(id),
		Label: reqData.Label,
	}

	err = r.ModifyTag(h.db, tag)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

// RemoveTag : to remove tag data from database
func (h *BaseHandler) RemoveTag(id string) (statusCode int, err error) {
	err = r.RemoveTag(h.db, id)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

package ads

import (
	"aura-test/repository/items"
	"aura-test/repository/user"
	"context"
	"errors"
)

type (
	taskOfList struct {
		req     *ReqOfList
		res     []*ResOfList
		storage *storageOfList
	}

	ReqOfList struct {
		Username string `form:"username" binding:"required" example:"test"`
		ItemType string `form:"item" binding:"required" example:"tools"`
	}

	ResOfList struct {
		ItemID   int    `json:"itemid"`
		ItemName string `json:"item_name"`
		Category string `json:"category"`
	}

	storageOfList struct {
	}
)

func List(ctx context.Context, in *ReqOfList) ([]*ResOfList, error) {
	task := newTaskOfList(in)
	if err := task.exec(ctx); err != nil {
		return nil, err
	}

	return task.res, nil
}

func newTaskOfList(in *ReqOfList) *taskOfList {
	return &taskOfList{
		req:     in,
		storage: &storageOfList{},
	}
}

func (t *taskOfList) exec(ctx context.Context) error {
	if err := t.validate(ctx); err != nil {
		return err
	}

	if err := t.getItemList(ctx); err != nil {
		return err
	}

	return nil
}

func (t *taskOfList) validate(ctx context.Context) error {
	userInfo, err := user.GetUserInfo(ctx, t.req.Username)
	if err != nil {
		return errors.New("get user info failed")
	}
	if userInfo == nil {
		return errors.New("user not exist")
	}

	return nil
}

func (t *taskOfList) getItemList(ctx context.Context) error {
	list, err := items.GetList(ctx, &items.ReqOfList{})
	if err != nil {
		return errors.New("get item list failed")
	}
	if len(list) == 0 {
		t.res = make([]*ResOfList, 0)
		return nil
	}

	for _, v := range list {
		row := &ResOfList{
			ItemID:   v.ID,
			ItemName: v.Name,
			Category: v.Category,
		}
		t.res = append(t.res, row)
	}

	return nil
}

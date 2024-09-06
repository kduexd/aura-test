package items

import (
	dbconn "aura-test/pkg/db"
	"aura-test/pkg/log"
	"context"
	"strings"

	"go.uber.org/zap"
)

type (
	ReqOfList struct {
		UserName string
		ItemType string
	}

	ResOfList struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Category string `json:"category"`
	}
)

func GetList(ctx context.Context, req *ReqOfList) ([]*ResOfList, error) {
	db := dbconn.CreateConnection()

	sql := `
		SELECT
		    id,
		    name,
		    category
		FROM items
	`

	var params []interface{}
	var cond []string
	if req.UserName != "" {
		cond = append(cond, "AND user_name = ? ")
		params = append(params, req.UserName)
	}

	if req.ItemType != "" {
		cond = append(cond, "AND type = ? ")
		params = append(params, req.ItemType)
	}

	if len(cond) > 0 {
		cond[0] = strings.Replace(cond[0], "AND", "WHERE", -1)
	}

	res := make([]*ResOfList, 0)
	err := db.SelectContext(ctx, &res, sql, params...)
	if err != nil {
		log.WithFields(
			zap.Any("Error", err),
			zap.String("SQL", sql),
			zap.Any("Params", params),
		).Error()
		return nil, err
	}

	return res, nil
}

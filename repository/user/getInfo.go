package user

import (
	dbconn "aura-test/pkg/db"
	"aura-test/pkg/log"
	"context"
	sqls "database/sql"
	"time"

	"go.uber.org/zap"
)

type ResOfInfo struct {
	UUID     string    `db:"uuid"`
	Name     string    `db:"name"`
	Password string    `db:"password"`
	Created  time.Time `db:"created"`
}

func GetUserInfo(ctx context.Context, name string) (*ResOfInfo, error) {
	db := dbconn.CreateConnection()

	sql := `
		SELECT
		    uuid,
		    name,
		    password,
		    created
		FROM users
		WHERE name = ?
	`

	var params []interface{}
	params = append(params, name)

	res := new(ResOfInfo)
	err := db.GetContext(ctx, res, sql, params...)
	if err != nil {
		if err == sqls.ErrNoRows {
			return nil, nil
		}

		log.WithFields(
			zap.Any("Error", err),
			zap.String("SQL", sql),
			zap.Any("Params", params),
		).Error()
		return nil, err
	}

	return res, nil
}

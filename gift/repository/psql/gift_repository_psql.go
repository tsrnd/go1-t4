package psql

import (
	"database/sql"

	model "github.com/goweb4/gift"
)

type giftRepository struct {
	DB *sql.DB
}

func (gR *giftRepository) Create(fromUserID, toUserID, productID int64, message string) (int64, error) {
	const query = `
		INSERT INTO gifts (
			from_user_id,
			to_user_id,
			product_id,
			message
		) values (
			$1,
			$2,
			$3,
			$4
		) returning id;
	`
	var id int64
	err := gR.DB.QueryRow(query, fromUserID, toUserID, productID, message).Scan(&id)
	return id, err
}

func (gR *giftRepository) Update(id, productID int64, message string) error {
	const query = `
		UPDATE gifts SET product_id = $1, message = $2 WHERE id = $3;
	`
	_, err := gR.DB.Exec(query, productID, message, id)
	return err
}

func (gR *giftRepository) Delete(id int64) error {
	const query = `
		DELETE FROM gifts WHERE id = $1;
	`
	_, err := gR.DB.Exec(query, id)
	return err
}

func (gR *giftRepository) GetByID(id int64) (*model.Gift, error) {
	gift := model.Gift{}
	const query = `
		SELECT
			id,
			from_user_id,
			to_user_id,
			product_id,
			message
		FROM
			gifts
		WHERE
			id = $1;
	`
	err := gR.DB.
		QueryRow(query, id).
		Scan(&gift.ID, &gift.FromUserID, &gift.ToUserID, &gift.ProductID, &gift.Message)
	return &gift, err
}

func (gR *giftRepository) GetByFromUserID(id int64) ([]*model.Gift, error) {
	var gifts []*model.Gift
	const query = `
		SELECT
			id,
			from_user_id,
			to_user_id,
			product_id,
			message
		FROM
			gifts
		WHERE
			from_user_id = $1;
	`
	rows, err := gR.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var gift model.Gift
		err = rows.Scan(&gift.ID, &gift.FromUserID, &gift.ToUserID, &gift.ProductID, &gift.Message)
		if err != nil {
			return gifts, err
		}
		gifts = append(gifts, &gift)
	}
	return gifts, err
}

func (gR *giftRepository) Fetch(offset, limit int64) ([]*model.Gift, error) {
	var gifts []*model.Gift
	const query = `
		SELECT
			id,
			from_user_id,
			to_user_id,
			group_id,
			message
		FROM
			gifts
		LIMIT $1 OFFSET $2;
	`
	rows, err := gR.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var gift model.Gift
		err = rows.Scan(&gift.ID, &gift.FromUserID, &gift.ToUserID, &gift.ProductID, &gift.Message)
		if err != nil {
			return gifts, err
		}
		gifts = append(gifts, &gift)
	}
	return gifts, err
}

func NewGiftRepository(db *sql.DB) *giftRepository {
	return &giftRepository{DB: db}
}

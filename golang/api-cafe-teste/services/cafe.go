package services

import (
	"context"
	"time"
)

type Coffe struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Roast     string    `json:"roast"`
	Image     string    `json:"image"`
	Region    string    `json:"region"`
	Price     float32   `json:"price"`
	GrindUnit int16     `json:"grind_unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func (c *Coffe) GetAllCoffess() ([]*Coffe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `select id, name, image, roast, region, priece, grind_unit, created_at, updated_at from coffes`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var coffess []*Coffe
	for rows.Next() {
		var coffe Coffe
		err := rows.Scan(
			&coffe.ID,
			&coffe.Name,
			&coffe.Image,
			&coffe.Roast,
			&coffe.Region,
			&coffe.Price,
			&coffe.GrindUnit,
			&coffe.CreatedAt,
			&coffe.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		coffess = append(coffess, &coffe)
	}
	return coffess, nil
}

func (c *Coffe) GetCoffeById(id string) (*Coffe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `
		select * from coffes where id = $1
	`
	var coffe Coffe

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&coffe.ID,
		&coffe.Name,
		&coffe.Image,
		&coffe.Roast,
		&coffe.Region,
		&coffe.Price,
		&coffe.GrindUnit,
		&coffe.CreatedAt,
		&coffe.UpdateAt,
	)
	if err != nil {
		return nil, err
	}
	return &coffe, nil
}

func (c *Coffe) CreateCoffe(coffe Coffe) (*Coffe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `
		insert into coffes(name, image, region, roast, price, grind_unit, created_at, updated_at)
		values($1, $2, $3, $4, $5, $6, $7, $8) returning *
	`
	_, err := db.ExecContext(
		ctx, query, coffe.Name, coffe.Image, coffe.Region, coffe.Roast, coffe.Price, coffe.GrindUnit, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	return &coffe, nil
}

func (c *Coffe) UpdateCoffe(id string, body Coffe) (*Coffe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `
	update coffes 
	set 
		name = $1
		image = $2, 
		roast = $3 
		price = $4
		grind_uint = $5
		region = $6
		updated_at = $7
	where id = $8
	`

	_, err := db.ExecContext(ctx, query, body.Name, body.Image, body.Roast, body.Price, body.GrindUnit, body.Region, body.UpdateAt, id)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (c *Coffe) DeleteCoffe(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `delete from coffes where id = $1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

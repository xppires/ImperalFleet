package repository

import (
	"app/internal/interfaces"
	"app/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type SpacecraftRepositoryMysql struct {
	conn interfaces.DBStore
}

func NewSpacecraftRepositoryMysql(conn interfaces.DBStore) *SpacecraftRepositoryMysql {
	return &SpacecraftRepositoryMysql{
		conn: conn,
	}
}

// Create an entry for a new spaceship.
func (r *SpacecraftRepositoryMysql) Create(ctx context.Context, craft *models.SpacecraftRequest) (int64, error) {
	qSpaceship := `INSERT INTO spaceships (name, class, status, image, crew, value) VALUES (?, ?, ?, ?, ?, ?)`
	qArmaments := `INSERT INTO armaments (spaceship_id,title, qty) VALUES (?, ?, ?)`

	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("spacecraft_repo: begin tx: %w", err)
	}

	insertArmament, err := tx.Prepare(qArmaments)
	if err != nil {
		return 0, fmt.Errorf("spacecraft_repo: prepare stmt: %w", err)
	}

	result, err := tx.ExecContext(ctx, qSpaceship, craft.Name, craft.Class, craft.Status, craft.Image, craft.Crew, craft.Value)
	if err != nil {
		err2 := tx.Rollback()
		return 0, errors.Join(fmt.Errorf("spacecraft_repo: insert spacecraft: %w", err), err2)
	}
	lastID, _ := result.LastInsertId()

	for _, armament := range craft.Armament {
		_, err := insertArmament.ExecContext(ctx, lastID, armament.Title, armament.Quantity)
		if err != nil {
			err2 := tx.Rollback()
			return 0, errors.Join(fmt.Errorf("spacecraft_repo: insert armament: %w", err), err2)
		}
	}

	if tx.Commit() != nil {
		err2 := tx.Rollback()
		return 0, errors.Join(fmt.Errorf("spacecraft_repo: insert armament: %w", err), err2)
	}
	return lastID, nil
}
func (r *SpacecraftRepositoryMysql) Update(ctx context.Context, id string, craft *models.SpacecraftRequest) error {
	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("spacecraft_repo: begin tx: %w", err)
	}

	q := "UPDATE spaceships SET name = ?, class = ?, status = ?, image = ?, crew = ?, value = ? WHERE id = ?"
	updateQuery, err := tx.Prepare(q)
	if err != nil {
		return fmt.Errorf("spacecraft_repo: prepare stmt: %w", err)
	}
	res, err := updateQuery.ExecContext(ctx, craft.Name, craft.Class, craft.Status, craft.Image, craft.Crew, craft.Value, id)
	if err != nil {
		return fmt.Errorf("spacecraft_repo: update spacecraft: %w", err)
	}
	rAff, _ := res.RowsAffected()
	if rAff == 0 {
		return errNotFound
	}

	if tx.Commit() != nil {
		err2 := tx.Rollback()
		return fmt.Errorf("spacecraft_repo: insert armament: %w", err2)
	}
	return nil
}
func (r *SpacecraftRepositoryMysql) Delete(ctx context.Context, id int) error {
	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("spacecraft_repo: begin tx: %w", err)
	}

	q := "DELETE FROM spaceships WHERE id = ?"
	deleteQuery, err := tx.Prepare(q)
	if err != nil {
		return fmt.Errorf("spacecraft_repo: prepare stmt: %w", err)
	}

	res, err := deleteQuery.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("spacecraft_repo: delete spacecraft: %w", err)
	}
	rAff, _ := res.RowsAffected()
	if rAff == 0 {
		return errNotFound
	}

	if tx.Commit() != nil {
		err2 := tx.Rollback()
		return fmt.Errorf("spacecraft_repo: insert armament: %w", err2)
	}
	return nil
}

func (r *SpacecraftRepositoryMysql) GetByID(ctx context.Context, id int, _ *string) (models.Spacecraft, error) {
	qSpaceship := "SELECT id, name, class, status, image, crew, value FROM spaceships WHERE id = ?"
	qArmaments := "SELECT id, spaceship_id, title, qty FROM armaments WHERE spaceship_id = ? "

	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return models.Spacecraft{}, fmt.Errorf("spacecraft_repo: begin tx: %w", err)
	}

	var craft models.Spacecraft
	row := tx.QueryRowContext(ctx, qSpaceship, id)
	if err := row.Scan(&craft.ID, &craft.Name, &craft.Class, &craft.Status, &craft.Image, &craft.Crew, &craft.Value); err != nil {
		err2 := tx.Rollback()
		return models.Spacecraft{}, fmt.Errorf("spacecraft_repo: retrieve spacecraft: %w", err2)
	}

	armRows, err := tx.QueryContext(ctx, qArmaments, craft.ID)
	if err != nil {
		err2 := tx.Rollback()
		return models.Spacecraft{}, errors.Join(fmt.Errorf("spacecraft_repo: retrieve armaments: %w", err), err2)
	}

	armaments := make([]models.Armament, 0)
	for armRows.Next() {
		var armament models.Armament
		if err := armRows.Scan(&armament.ID, &armament.CraftID, &armament.Title, &armament.Quantity); err != nil {
			err2 := tx.Rollback()
			return models.Spacecraft{}, errors.Join(fmt.Errorf("spacecraft_repo: retrieve armaments: %w", err), err2)
		}
		armaments = append(armaments, armament)
	}
	craft.Armament = armaments

	return craft, nil
}

func (r *SpacecraftRepositoryMysql) Get(ctx context.Context, _ *map[string][]string) ([]models.Spacecraft, error) {
	// list := []models.Spacecraft{}
	// return list, fmt.Errorf("not implemented")

	// q := `
	// SELECT id, name, class, status , image, crew, value FROM spaceships
	// WHERE (LOWER(name) = LOWER($1)) OR $1 = ''
	// AND (LOWER(class) = LOWER($2)) OR $2 = ''
	// AND (LOWER(status) = LOWER($3)) OR $3 = ''
	// ORDER BY id
	// `
	q := `
	SELECT id, name, class, status , image, crew, value FROM spaceships 
	ORDER BY id
	`
	// var name, class ,status string = "", "", ""

	// name := filters.Get("name")
	// class := filters.Get("class")
	// status := filters.Get("status")

	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("spacecraft_repo: begin tx: %w", err)
	}

	selectArmaments, err := tx.Prepare("SELECT id, spaceship_id, title, qty FROM armaments WHERE spaceship_id = ?")
	if err != nil {
		return nil, fmt.Errorf("spacecraft_repo: preparing armaments stmt: %w", err)
	}
	defer func(selectArmaments *sql.Stmt) {
		err := selectArmaments.Close()
		if err != nil {

		}
	}(selectArmaments)
	rows, err := tx.QueryContext(ctx, q)
	if err != nil {
		err2 := tx.Rollback()
		return nil, errors.Join(fmt.Errorf("spacecraft_repo: retrieve spacecrafts spacecrafts: %w", err), err2)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	spaceshipsMap := make(map[int]*models.Spacecraft, 0)
	for rows.Next() {
		var spacecraft models.Spacecraft
		if err := rows.Scan(&spacecraft.ID, &spacecraft.Name, &spacecraft.Class, &spacecraft.Status, &spacecraft.Image, &spacecraft.Crew, &spacecraft.Value); err != nil {
			err2 := tx.Rollback()
			return nil, errors.Join(fmt.Errorf("spacecraft_repo: retrieve spacecrafts: %w", err), err2)
		}
		spaceshipsMap[len(spaceshipsMap)] = &spacecraft

	}
	for craftIdx, spacecraft := range spaceshipsMap {
		armaments := make([]models.Armament, 0)
		armRows, err := selectArmaments.Query(spacecraft.ID)
		if err != nil {
			err2 := tx.Rollback()
			return nil, errors.Join(fmt.Errorf("spacecraft_repo: retrieve armaments: %w", err), err2)
		}

		for armRows.Next() {
			var armament models.Armament
			if err := armRows.Scan(&armament.ID, &armament.CraftID, &armament.Title, &armament.Quantity); err != nil {
				return nil, fmt.Errorf("spacecraft_repo: retrieve armaments: %w", err)
			}
			armaments = append(armaments, armament)
			spaceshipsMap[craftIdx].Armament = append(spaceshipsMap[craftIdx].Armament, armament)
		}
		clear(armaments)

	}
	defer func(selectArmaments *sql.Stmt) {
		err := selectArmaments.Close()
		if err != nil {

		}
	}(selectArmaments)
	returnedSpacecrafts := make([]models.Spacecraft, 0, len(spaceshipsMap))
	for _, spacecraft := range spaceshipsMap {
		returnedSpacecrafts = append(returnedSpacecrafts, *spacecraft)
	}
	return returnedSpacecrafts, nil
}

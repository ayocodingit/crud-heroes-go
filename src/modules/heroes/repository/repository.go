package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/entity"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{db}
}

var querySelect = "select heroes.id as id, name, ability, role, defend, attack, armor  from heroes LEFT JOIN attribute ON heroes.id = attribute.hero_id"

func (r repository) FindAll(req entity.QueryFindAll) []entity.Hero {
	log.Println(req)
	var heroes []entity.Hero
	db := r.db

	results, err := db.Query(querySelect)

	if err != nil {
		log.Fatal(err.Error())
	}

	for results.Next() {
		var hero entity.Hero
		err = results.Scan(&hero.Id, &hero.Name, &hero.Ability, &hero.Role, &hero.Attribute.Defend, &hero.Attribute.Attack, &hero.Attribute.Armor)
		if err != nil {
			log.Fatal(err.Error())
		}
		heroes = append(heroes, hero)
	}

	return heroes
}

func (r repository) FindById(id int) (hero entity.Hero, err error) {
	db := r.db

	err = db.QueryRow(querySelect+" where heroes.id = ?", id).Scan(&hero.Id, &hero.Name, &hero.Ability, &hero.Role, &hero.Attribute.Defend, &hero.Attribute.Attack, &hero.Attribute.Armor)

	return
}

func (r repository) Delete(id int) {
	db := r.db

	db.QueryRow("delete from heroes where id = ?", id)

}

func (r repository) Store(hero *entity.Hero) (err error) {
	db := r.db

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	results, err := tx.ExecContext(ctx, "insert into heroes set name = ?, ability = ?, role = ?", hero.Name, hero.Ability, hero.Role)
	if err != nil {
		tx.Rollback()
		return
	}

	id, err := results.LastInsertId()
	hero.Id = int(id)

	if err != nil {
		tx.Rollback()
		return
	}
	_, err = tx.ExecContext(ctx, "insert into attribute set attack = ?, defend = ?, armor = ?, hero_id = ?", hero.Attribute.Attack, hero.Attribute.Defend, hero.Attribute.Armor, hero.Id)

	if err != nil {
		tx.Rollback()
		return
	}

	if err == nil {
		tx.Commit()
	}

	return
}

func (r repository) Update(hero entity.Hero) (err error) {
	db := r.db

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = tx.ExecContext(ctx, "update heroes set name = ?, ability = ?, role = ? where id = ?", hero.Name, hero.Ability, hero.Role, hero.Id)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.ExecContext(ctx, "update attribute set attack = ?, defend = ?, armor = ?, hero_id = ? where hero_id = ?", hero.Attribute.Attack, hero.Attribute.Defend, hero.Attribute.Armor, hero.Id, hero.Id)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}

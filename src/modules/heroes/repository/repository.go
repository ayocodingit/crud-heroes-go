package repository

import (
	"context"
	"log"

	"github.com/ayocodingit/crud-heroes-go/src/config"
	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/entity"
)

func FindAll() []entity.Hero {
	var heroes []entity.Hero
	db := config.LoadDB()

	results, err := db.Query("select heroes.id as id, name, ability, role, defend, attack, armor  from heroes LEFT JOIN attribute ON heroes.id = attribute.hero_id")

	defer db.Close()

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

func FindById(id int) (hero entity.Hero, err error) {
	db := config.LoadDB()

	err = db.QueryRow("select heroes.id as id, name, ability, role, defend, attack, armor  from heroes LEFT JOIN attribute ON heroes.id = attribute.hero_id where heroes.id = ?", id).Scan(&hero.Id, &hero.Name, &hero.Ability, &hero.Role, &hero.Attribute.Defend, &hero.Attribute.Attack, &hero.Attribute.Armor)

	defer db.Close()

	return
}

func Delete(id int) {
	db := config.LoadDB()

	db.QueryRow("delete from heroes where id = ?", id)

	defer db.Close()
}

func Store(hero *entity.Hero) (err error) {
	db := config.LoadDB()
	defer db.Close()

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

func Update(hero *entity.Hero) (err error) {
	db := config.LoadDB()
	defer db.Close()

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

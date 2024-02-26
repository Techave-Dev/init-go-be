package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/techave-dev/init-go-be/internal/repo"
	"github.com/techave-dev/init-go-be/tools"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
	cfg, err := tools.LoadConfig()
	if err != nil {
		logrus.Fatal("cannot load config")
	}

	pool, err := pgxpool.New(context.Background(), cfg.PostgresURL)
	if err != nil {
		logrus.Fatal("cannot connect to db")
	}

	queries := repo.New(pool)

	abilities := []repo.Ability{}
	for _, abilityEnum := range repo.AllAbilityEnumValues() {
		ability, err := queries.UpsertAbility(context.Background(), repo.UpsertAbilityParams{
			Name: abilityEnum,
			Desc: pgtype.Text{Valid: false},
		})

		if err != nil {
			panic(err.Error())
		}

		abilities = append(abilities, ability)
	}

	fmt.Printf("abilities: %v\n", abilities)

	roles := []repo.Role{}
	for _, roleEnum := range repo.AllRoleEnumValues() {
		role, err := queries.UpsertRole(context.Background(), repo.UpsertRoleParams{
			Name: roleEnum,
			Desc: pgtype.Text{Valid: false},
		})

		if err != nil {
			panic(err.Error())
		}

		roles = append(roles, role)
	}

	fmt.Printf("roles: %v\n", roles)

	roleAbilityMap := map[repo.RoleEnum][]repo.AbilityEnum{
		repo.RoleEnumAdmin: {},
		repo.RoleEnumUser:  {},
	}

	for role, v := range roleAbilityMap {
		for _, roleAbilities := range v {
			roleAbility, err := queries.UpsertRoleAbilityByName(context.Background(), repo.UpsertRoleAbilityByNameParams{
				Role:    role,
				Ability: roleAbilities,
			})

			if err != nil {
				panic(err.Error())
			}

			fmt.Printf("roleAbility: %v\n", roleAbility)
		}
	}

	logrus.Info("db seeded successfully")
}

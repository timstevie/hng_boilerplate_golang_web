package migrations

import "github.com/hngprojects/hng_boilerplate_golang_web/internal/models"

func AuthMigrationModels() []interface{} {
	return []interface{}{
		models.Organisation{},
		models.Profile{},
		models.Product{},
		models.User{},
		models.WaitlistUser{},
		models.NewsLetter{},
	} // an array of db models, example: User{}
}

func AlterColumnModels() []AlterColumn {
	return []AlterColumn{}
}

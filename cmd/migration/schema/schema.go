package schema

import (
	"log"

	model "github.com/Robihamanto/produktif/internal"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// Schema struct represent schema for migration
type Schema struct {
	db         *gorm.DB
	migrator   *gormigrate.Gormigrate
	migrations []*gormigrate.Migration
}

// New create Schema instance
func New(db *gorm.DB) *Schema {
	migrations := initMigrations()
	m := gormigrate.New(db, &gormigrate.Options{
		TableName:      "migrations",
		IDColumnName:   "id",
		IDColumnSize:   255,
		UseTransaction: true,
	}, migrations)

	return &Schema{db, m, migrations}
}

func initMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "20200324125355",
			Migrate: func(tx *gorm.DB) error {
				var (
					userTable = "CREATE TABLE `users` " +
						"(" +
						"`id` int(10) unsigned NOT NULL AUTO_INCREMENT," +
						"`created_at` timestamp NULL DEFAULT NULL," +
						"`updated_at` timestamp NULL DEFAULT NULL," +
						"`deleted_at` timestamp NULL DEFAULT NULL," +
						"`username` varchar(255) NOT NULL," +
						"`password` varchar(255) NOT NULL," +
						"`email` varchar(255) NOT NULL," +
						"`full_name` varchar(255) NOT NULL," +
						"PRIMARY KEY (`id`)," +
						"UNIQUE KEY `username` (`username`)," +
						"UNIQUE KEY `email` (`email`)" +
						") ENGINE=InnoDB DEFAULT CHARSET=utf8;"
				)

				if err := tx.Exec(userTable).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.DropTableIfExists(&model.User{}).Error; err != nil {
					return err
				}
				return nil
			},
		},
	}
}

// MyMigration represent a migration history points
type MyMigration struct {
	ID string `gorm:"column:id;size:255;not null;primary_key"`
}

// func TableName() string {

// }

// AutoMigrate do an AutoMigration from exsisting models
func (s *Schema) autoMigrate() error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.AutoMigrate(
		&model.User{},
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	if !tx.HasTable("migrations") {
		if err := tx.CreateTable(&MyMigration{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, migration := range s.migrations {
		var temp MyMigration
		if err := tx.Model(&MyMigration{}).Where("id = ?", migration.ID).Find(&temp).Error; err != nil && err != gorm.ErrRecordNotFound {
			tx.Rollback()
			return err
		}
		if temp.ID == "" {
			temp = MyMigration{migration.ID}
			if err := tx.Model(&MyMigration{}).Create(temp).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// Migrate migrates all registered migrations
func (s *Schema) Migrate() error {
	if s.db.HasTable("migrations") {
		var migrationCount int
		if err := s.db.Model(&MyMigration{}).Count(&migrationCount).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if migrationCount > 0 {
			log.Print("Running manual migration")
			return s.migrator.Migrate()
		}
	}

	return s.autoMigrate()
}

// ReverseLast reverse / undo last migration
func (s *Schema) ReverseLast() error {
	return s.migrator.RollbackLast()
}

// ReverseN reverse / undo N migration from current
func (s *Schema) ReverseN(n int) error {
	for ; n > 0; n-- {
		if err := s.migrator.RollbackLast(); err != nil {
			return err
		}
	}
	return nil
}

// ReverseAll reverse all migrations
func (s *Schema) ReverseAll() error {
	// list all migrations in DB
	var dbMigs int
	if err := s.db.
		Model(&MyMigration{}).
		Count(&dbMigs).
		Error; err != nil {
		return err
	}
	for i := dbMigs - 1; i >= 0; i-- {
		if err := s.migrator.RollbackMigration(s.migrations[i]); err != nil {
			return err
		}
	}
	return nil
}

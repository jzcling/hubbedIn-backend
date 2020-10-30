package migrations

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating candidates table")
		_, err := db.Exec(`
            create table if not exists candidates (
                id bigserial not null primary key,
                first_name varchar(50) not null,
                last_name varchar(50) not null,
                email varchar(256) not null,
                contact_number varchar(80) not null,
                gender varchar(20),
                nationality varchar(50),
                residence_city varchar(50),
                expected_salary_currency varchar(3),
                expected_salary int,
                linked_in_url varchar(256),
                scm_url varchar(256),
                education_level varchar(50),
                birthday timestamp,
                notice_period int,
                created_at timestamp,
                updated_at timestamp,
                deleted_at timestamp
            );
        `)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping candidates table")
		_, err := db.Exec(`drop table if exists candidates`)
		return err
	})
}

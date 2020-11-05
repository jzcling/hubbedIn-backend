package main

import (
	"fmt"

	migrations "github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating tables")
		_, err := db.Exec(`
            create table if not exists candidates (
                id bigserial not null primary key,
                first_name text not null,
                last_name text not null,
                email text not null,
                contact_number text not null,
                gender text,
                nationality text,
                residence_city text,
                expected_salary_currency text,
                expected_salary int,
                linked_in_url text,
                scm_url text,
                education_level text,
                birthday timestamp,
                notice_period int,
                created_at timestamptz,
                updated_at timestamptz,
                deleted_at timestamptz
            );

            create index on candidates (id, deleted_at);
            create index on candidates (email, deleted_at);

            create table if not exists skills (
                id bigserial not null primary key,
                name text not null
            );

            create index on skills (name);

            create table if not exists users_skills (
                id bigserial not null primary key,
                candidate_id bigint not null,
                skill_id bigint not null,
                created_at timestamptz,
                updated_at timestamptz,
                deleted_at timestamptz,
                constraint fk_candidates foreign key(candidate_id) references candidates(id),
                constraint fk_skills foreign key(skill_id) references skills(id)
            );

            create index on users_skills (candidate_id, skill_id, deleted_at);

            create table if not exists institutions (
                id bigserial not null primary key,
                country text,
                name text not null
            );

            create index on institutions (name);

            create table if not exists courses (
                id bigserial not null primary key,
                institution_id bigint,
                level text,
                name text not null,
                constraint fk_institutions foreign key(institution_id) references institutions(id)
            );

            create index on courses (name);

            create table if not exists academic_histories (
                id bigserial not null primary key,
                candidate_id bigint not null,
                institution_id bigint not null,
                course_id bigint not null,
                year_obtained bigint,
                created_at timestamptz,
                updated_at timestamptz,
                deleted_at timestamptz,
                constraint fk_candidates foreign key(candidate_id) references candidates(id),
                constraint fk_institutions foreign key(institution_id) references institutions(id),
                constraint fk_courses foreign key(course_id) references courses(id)
            );

            create index on academic_histories (candidate_id, institution_id, course_id, deleted_at);

            create table if not exists companies (
                id bigserial not null primary key,
                name text not null
            );

            create index on companies (name);

            create table if not exists departments (
                id bigserial not null primary key,
                company_id bigint,
                name text not null,
                constraint fk_companies foreign key(company_id) references companies(id)
            );

            create index on departments (name);

            create table if not exists job_histories (
                id bigserial not null primary key,
                candidate_id bigint not null,
                company_id bigint not null,
                department_id bigint,
                country text not null,
                city text,
                title text not null,
                start_date timestamptz not null,
                end_date timestamptz,
                salary_currency text,
                salary bigint,
                description text,
                created_at timestamptz,
                updated_at timestamptz,
                deleted_at timestamptz
            );

            create index on job_histories (candidate_id, company_id, department_id, deleted_at);
            create index on job_histories (candidate_id, title, deleted_at);
        `)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping tables")
		_, err := db.Exec(`
            drop table if exists job_histories;
            drop table if exists departments;
            drop table if exists companies;
            drop table if exists academic_histories;
            drop table if exists courses;
            drop table if exists institutions;
            drop table if exists users_skills;
            drop table if exists skills;
            drop table if exists candidates;
        `)
		return err
	})
}

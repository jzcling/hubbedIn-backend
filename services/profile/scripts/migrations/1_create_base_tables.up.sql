create table if not exists candidates (
    id bigserial not null primary key,
    auth_id text not null unique,
    first_name text,
    last_name text,
    email text not null unique,
    contact_number text unique,
    picture text,
    gender text,
    nationality text,
    residence_city text,
    expected_salary_currency text,
    expected_salary int,
    linked_in_url text,
    scm_url text,
    website_url text,
    education_level text,
    summary text,
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
    name text not null unique
);

create index on skills (name);

create table if not exists users_skills (
    id bigserial not null primary key,
    candidate_id bigint not null,
    skill_id bigint not null,
    created_at timestamptz,
    updated_at timestamptz,
    constraint fk_candidates foreign key(candidate_id) references candidates(id) on delete cascade on update cascade,
    constraint fk_skills foreign key(skill_id) references skills(id) on delete cascade on update cascade
);

create index on users_skills (candidate_id, skill_id);

create table if not exists institutions (
    id bigserial not null primary key,
    country text,
    name text not null unique
);

create index on institutions (name);

create table if not exists courses (
    id bigserial not null primary key,
    level text,
    name text not null unique
);

create index on courses (name);

create table if not exists courses_institutions (
    id bigserial not null primary key,
    course_id bigint not null,
    institution_id bigint not null,
    constraint fk_courses foreign key(course_id) references courses(id) on delete cascade on update cascade,
    constraint fk_institutions foreign key(institution_id) references institutions(id) on delete cascade on update cascade
);

create index on courses_institutions (course_id, institution_id);

create table if not exists academic_histories (
    id bigserial not null primary key,
    candidate_id bigint not null,
    institution_id bigint not null,
    course_id bigint not null,
    year_obtained bigint,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    constraint fk_candidates foreign key(candidate_id) references candidates(id) on delete cascade on update cascade,
    constraint fk_institutions foreign key(institution_id) references institutions(id) on delete cascade on update cascade,
    constraint fk_courses foreign key(course_id) references courses(id) on delete cascade on update cascade
);

create index on academic_histories (candidate_id, institution_id, course_id, deleted_at);

create table if not exists companies (
    id bigserial not null primary key,
    name text not null unique
);

create index on companies (name);

create table if not exists departments (
    id bigserial not null primary key,
    name text not null unique
);

create index on departments (name);

create table if not exists companies_departments (
    id bigserial not null primary key,
    company_id bigint not null,
    department_id bigint not null,
    constraint fk_companies foreign key(company_id) references companies(id) on delete cascade on update cascade,
    constraint fk_departments foreign key(department_id) references departments(id) on delete cascade on update cascade
);

create index on companies_departments (company_id, department_id);

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
    deleted_at timestamptz,
    constraint fk_candidates foreign key(candidate_id) references candidates(id) on delete cascade on update cascade,
    constraint fk_companies foreign key(company_id) references companies(id) on delete cascade on update cascade,
    constraint fk_departments foreign key(department_id) references departments(id) on delete cascade on update cascade
);

create index on job_histories (candidate_id, company_id, department_id, deleted_at);
create index on job_histories (candidate_id, title, deleted_at);
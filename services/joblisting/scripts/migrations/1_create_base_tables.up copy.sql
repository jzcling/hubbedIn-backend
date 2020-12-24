create table if not exists companies (
    id bigserial not null primary key,
	name text not null unique,
	logo_url text not null,
	size bigint
);

create table if not exists industries (
    id bigserial not null primary key,
    name text not null unique
);

create table if not exists companies_industries (
    id bigserial not null primary key,
    company_id bigint not null,
	industry_id bigint not null,
	constraint fk_companies foreign key(company_id) references companies(id) on delete cascade on update cascade,
    constraint fk_industries foreign key(industry_id) references industries(id) on delete cascade on update cascade
);

create table if not exists job_functions (
    id bigserial not null primary key,
    name text not null unique
);

create table if not exists key_persons (
    id bigserial not null primary key,
	company_id bigint not null,
    name text not null,
	contact_number text,
	email text,
	job_title text,
	updated_at timestamptz,
	constraint fk_companies foreign key(company_id) references companies(id) on delete cascade on update cascade
);

create table if not exists job_platforms (
    id bigserial not null primary key,
    name text not null,
	base_url text not null
);

create table if not exists job_posts (
    id bigserial not null primary key,
	company_id bigint not null,
	hr_contact_id bigint,
	hiring_manager_id bigint,
	job_platform_id bigint,
	skill_id bigint[],
	title text not null,
	description text not null,
	seniority_level text,
	years_experience bigint,
	employment_type text,
	function_id bigint,
	industry_id bigint,
	location text,
	remote boolean,
	salary_currency text,
	min_salary bigint,
	max_salary bigint,
	created_at timestamptz,
	updated_at timestamptz,
	start_at timestamptz,
	expire_at timestamptz,
    constraint fk_companies foreign key(company_id) references companies(id) on delete cascade on update cascade,
    constraint fk_job_functions foreign key(function_id) references job_functions(id) on delete cascade on update cascade,
    constraint fk_industries foreign key(industry_id) references industries(id) on delete cascade on update cascade,
    constraint fk_key_persons foreign key(hr_contact_id) references key_persons(id) on delete cascade on update cascade,
    constraint fk_key_persons_2 foreign key(hiring_manager_id) references key_persons(id) on delete cascade on update cascade,
    constraint fk_job_platforms foreign key(job_platform_id) references job_platforms(id) on delete cascade on update cascade
);

create index on job_posts (title, seniority_level, location);
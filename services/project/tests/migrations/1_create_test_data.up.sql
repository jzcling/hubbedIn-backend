create table if not exists projects (
    id bigserial not null primary key,
    name text not null,
    repo_url text not null unique,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz
);

create index on projects (id, deleted_at);
create index on projects (repo_url, deleted_at);

create table if not exists candidates_projects (
    id bigserial not null primary key,
    candidate_id bigint not null,
    project_id bigint not null,
    constraint fk_projects foreign key(project_id) references projects(id) on delete cascade on update cascade
);

create index on candidates_projects (candidate_id, project_id);

create table if not exists ratings (
    id bigserial not null primary key,
    project_id bigint not null,
    reliability_rating integer,
    maintainability_rating integer,
    security_rating integer,
    security_review_rating integer,
    coverage real,
    duplications real,
    lines bigint,
    created_at timestamptz,
    constraint fk_projects foreign key(project_id) references projects(id) on delete cascade on update cascade
);

insert into projects (name, repo_url, created_at, updated_at)
    values ('testname', 'testrepo', now(), now());
insert into projects (name, repo_url, created_at, updated_at)
    values ('testname2', 'testrepo2', now(), now());

insert into candidates_projects (candidate_id, project_id)
    values (1, 1);
insert into candidates_projects (candidate_id, project_id)
    values (1, 2);

insert into ratings (project_id, reliability_rating, maintainability_rating, security_rating, security_review_rating, coverage, duplications, lines, created_at)
    values (1, 1, 1, 1, 1, 1.0, 1.0, 1, now());
insert into ratings (project_id, reliability_rating, maintainability_rating, security_rating, security_review_rating, coverage, duplications, lines, created_at)
    values (1, 2, 2, 2, 2, 2.0, 2.0, 2, now());
insert into ratings (project_id, reliability_rating, maintainability_rating, security_rating, security_review_rating, coverage, duplications, lines, created_at)
    values (2, 1, 1, 1, 1, 1.0, 1.0, 1, now());
insert into ratings (project_id, reliability_rating, maintainability_rating, security_rating, security_review_rating, coverage, duplications, lines, created_at)
    values (2, 2, 2, 2, 2, 2.0, 2.0, 2, now());
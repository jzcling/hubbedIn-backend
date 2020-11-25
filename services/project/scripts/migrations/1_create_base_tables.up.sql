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
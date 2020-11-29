create table if not exists joblistings (
    id bigserial not null primary key,
    name text not null,
    repo_url text not null unique,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz
);

create index on joblistings (id, deleted_at);
create index on joblistings (repo_url, deleted_at);

create table if not exists candidates_joblistings (
    id bigserial not null primary key,
    candidate_id bigint not null,
    joblisting_id bigint not null,
    constraint fk_joblistings foreign key(joblisting_id) references joblistings(id) on delete cascade on update cascade
);

create index on candidates_joblistings (candidate_id, joblisting_id);

create table if not exists ratings (
    id bigserial not null primary key,
    joblisting_id bigint not null,
    reliability_rating integer,
    maintainability_rating integer,
    security_rating integer,
    security_review_rating integer,
    coverage real,
    duplications real,
    lines bigint,
    created_at timestamptz,
    constraint fk_joblistings foreign key(joblisting_id) references joblistings(id) on delete cascade on update cascade
);
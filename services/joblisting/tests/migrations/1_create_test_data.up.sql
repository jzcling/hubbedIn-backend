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

insert into joblistings (name, repo_url, created_at, updated_at)
    values ('testname', 'testrepo', now(), now());
insert into joblistings (name, repo_url, created_at, updated_at)
    values ('testname2', 'testrepo2', now(), now());

insert into candidates_joblistings (candidate_id, joblisting_id)
    values (1, 1);
insert into candidates_joblistings (candidate_id, joblisting_id)
    values (1, 2);

insert into ratings (joblisting_id, reliability_rating, maintainability_rating, security_rating, security_review_rating, coverage, duplications, lines, created_at)
    values (1, 1, 1, 1, 1, 1.0, 1.0, 1, now());
insert into ratings (joblisting_id, reliability_rating, maintainability_rating, security_rating, security_review_rating, coverage, duplications, lines, created_at)
    values (1, 2, 2, 2, 2, 2.0, 2.0, 2, now());
insert into ratings (joblisting_id, reliability_rating, maintainability_rating, security_rating, security_review_rating, coverage, duplications, lines, created_at)
    values (2, 1, 1, 1, 1, 1.0, 1.0, 1, now());
insert into ratings (joblisting_id, reliability_rating, maintainability_rating, security_rating, security_review_rating, coverage, duplications, lines, created_at)
    values (2, 2, 2, 2, 2, 2.0, 2.0, 2, now());
create table if not exists users (
    id bigserial not null primary key,
    auth_id text not null unique,
    first_name text,
    last_name text,
    email text not null unique,
    contact_number text unique,
    picture text,
    gender text,
    roles text[],
    candidate_id bigint,
    job_company_id bigint,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz
);

create index on users (id, deleted_at);
create index on users (email, deleted_at);

insert into users (id, auth_id, first_name, last_name, email, contact_number, picture, gender, roles, candidate_id, created_at, updated_at, deleted_at)
select id, auth_id, first_name, last_name, email, contact_number, picture, gender, array['Candidate'], id, created_at, updated_at, deleted_at
from candidates;

alter table candidates
drop column auth_id,
drop column first_name,
drop column last_name,
drop column email,
drop column contact_number,
drop column picture,
drop column gender;
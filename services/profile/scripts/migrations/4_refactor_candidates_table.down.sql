alter table candidates
add column auth_id text,
add column first_name text,
add column last_name text,
add column email text,
add column contact_number text,
add column picture text,
add column gender text;

update candidates c 
set auth_id = u.auth_id, 
    first_name = u.first_name, 
    last_name = u.last_name,
    email = u.email,
    contact_number = u.contact_number,
    picture = u.picture,
    gender = u.gender 
from (select * from users) u
where c.id = u.candidate_id;

drop table if exists users;
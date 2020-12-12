create table if not exists assessments (
    id bigserial not null primary key,
	name text not null unique,
	description text,
	notes text,
	image_url text,
	difficulty text,
	time_allowed bigint,
	type text,
	randomise boolean,
	num_questions bigint
);

create index on assessments (name);

create table if not exists assessment_attempts (
    id bigserial not null primary key,
	assessment_id bigint not null,
	candidate_id bigint not null,
	status text not null,
	started_at timestamptz,
	completed_at timestamptz,
	score bigint,
    constraint fk_assessments foreign key(assessment_id) references assessments(id) on delete cascade on update cascade
);

create index on assessment_attempts (assessment_id, candidate_id);

create table if not exists questions (
    id bigserial not null primary key,
    created_by bigint,
	type text not null,
	text text,
	image_url text,
	options text[],
	answer bigint,
    tags text[]
);

create table if not exists attempts_questions (
	id bigserial not null primary key,
	attempt_id bigint not null,
	question_id bigint not null,
	candidate_id bigint not null,
	selection bigint,
	text text,
	score bigint,
	time_taken bigint,
	created_at timestamptz,
	updated_at timestamptz,
    constraint fk_assessment_attempts foreign key(attempt_id) references assessment_attempts(id) on delete cascade on update cascade,
    constraint fk_questions foreign key(question_id) references questions(id) on delete cascade on update cascade
)

create index on attempts_questions (attempt_id, question_id, candidate_id);

create table if not exists tags (
    id bigserial not null primary key,
    name text not null unique
);

create index on tags (name);

create table if not exists questions_tags (
    id bigserial not null primary key,
    question_id bigint not null,
    tag_id bigint not null,
    constraint fk_questions foreign key(question_id) references questions(id) on delete cascade on update cascade,
    constraint fk_tags foreign key(tag_id) references tags(id) on delete cascade on update cascade
);

create index on questions_tags (question_id, tag_id);

create table if not exists assessments_questions (
    id bigserial not null primary key,
    assessment_id bigint not null,
	question_id bigint not null,
    constraint fk_assessments foreign key(assessment_id) references assessments(id) on delete cascade on update cascade,
    constraint fk_questions foreign key(question_id) references questions(id) on delete cascade on update cascade
);

create index on assessments_questions (assessment_id, question_id);
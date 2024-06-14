create table "post" (
	"id" serial primary key,
	"title" varchar(100),
	"content" text,
	"status" int default 1,
	"published_date" timestamp,
	"created_date" timestamp default current_timestamp,
	"updated_date" timestamp
);

comment on column "post"."status" is 'state of post data';

create table "tag" (
	"id" serial primary key,
	"label" varchar(100) unique,
	"created_date" timestamp default current_timestamp,
	"updated_date" timestamp
);

create table "post_to_tag" (
	"id" serial primary key,
	"post_id" int,
	"tag_id" int,
	unique("post_id", "tag_id"),
	foreign key ("post_id") references "post"("id"),
	foreign key ("tag_id") references "tag"("id")
);
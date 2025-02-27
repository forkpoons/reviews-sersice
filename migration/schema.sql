CREATE TYPE reviews_status AS ENUM ('created','published','deleted');

CREATE TABLE "reviews"
(
    "id"          uuid PRIMARY KEY,
    "rtype"       varchar,
    "created_at"  timestamptz,
    "updated_at"  timestamptz,
    "product_id"  uuid,
    "user_id"     uuid    NOT NULL,
    "review_text" varchar NOT NULL,
    "media"       varchar NOT NULL,
    "rate"        int default 0,
    "status"      reviews_status
);

CREATE TABLE "answers"
(
    "id"          uuid PRIMARY KEY,
    "answer_text" varchar NOT NULL,
    "user_id"     uuid    NOT NULL,
    "created_at"  timestamptz,
    "updated_at"  timestamptz,
    "question_id" uuid,
    "status"      reviews_status
);
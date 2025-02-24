CREATE TYPE reviews_status AS ENUM ('created','published','deleted');

CREATE TABLE "reviews"
(
    "id"          uuid PRIMARY KEY,
    "type"        varchar,
    "created_at"  timestamptz,
    "updated_at"  timestamptz,
    "product_id"  uuid,
    "user_id"     uuid    NOT NULL,
    "review_text" varchar NOT NULL,
    "media"       jsonb   NOT NULL,
    "rate"        int,
    "status"      reviews_status
);


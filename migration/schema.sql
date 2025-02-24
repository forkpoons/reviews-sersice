CREATE TABLE "reviews"
(
    "id"          uuid PRIMARY KEY,
    "user_id"     uuid      NOT NULL,
    "review_text" varchar   NOT NULL,
    "media"       varchar[] NOT NULL,
    "product_id"  uuid,
    "rate"        int
);

CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "customer_name" varchar,
  "customer_cont_no" varchar,
  "customer_address" text,
  "total_buy" bigserial,
  "creator_id" bigserial,
  "created_at" date NOT NULL DEFAULT 'now()'
);

CREATE TABLE "sells" (
  "id" bigserial PRIMARY KEY,
  "customer_id" bigserial,
  "product_id" varchar,
  "purse_price" bigserial,
  "quantity" bigserial,
  "total_price" bigserial,
  "created_at" date NOT NULL DEFAULT 'now()'
);

ALTER TABLE "sells" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");
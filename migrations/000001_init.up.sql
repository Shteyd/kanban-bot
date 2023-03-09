CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "customer" (
    "id" UUID NOT NULL DEFAULT uuid_generate_v4(),
    "username" VARCHAR(255) NOT NULL,

    PRIMARY KEY ("id")
);

CREATE TABLE "board" (
    "id" UUID NOT NULL DEFAULT uuid_generate_v4(),
    "article" VARCHAR(5) NOT NULL,
    "description" TEXT,
    "customer_id" UUID NOT NULL,

    PRIMARY KEY ("id"),
    FOREIGN KEY ("customer_id") REFERENCES "customer" ("id")
);

CREATE TABLE "task" (
    "id" UUID NOT NULL DEFAULT uuid_generate_v4(),
    "position" INTEGER NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "board_id" UUID NOT NULL,

    PRIMARY KEY ("id"),
    FOREIGN KEY ("board_id") REFERENCES "board" ("id")
);


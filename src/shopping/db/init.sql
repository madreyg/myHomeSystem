CREATE TABLE "product" (
  "id"          SERIAL       NOT NULL,
  "name"        VARCHAR(127) NOT NULL,
  "photo"       VARCHAR(255) NOT NULL DEFAULT '',
  "category_id" INT          NOT NULL,
  CONSTRAINT products_pk PRIMARY KEY ("id")
) WITH (OIDS = FALSE
);

CREATE TABLE "category" (
  "id"           SERIAL       NOT NULL,
  "name"         VARCHAR(127) NOT NULL,
  "directory_id" INT          NOT NULL,
  "photo"        VARCHAR(255) NOT NULL DEFAULT '',
  CONSTRAINT category_pk PRIMARY KEY ("id")
) WITH (OIDS = FALSE
);

CREATE TABLE "directory" (
  "id"    SERIAL       NOT NULL,
  "name"  VARCHAR(127) NOT NULL,
  "photo" VARCHAR(255) NOT NULL DEFAULT '',
  CONSTRAINT directory_pk PRIMARY KEY ("id")
) WITH (OIDS = FALSE
);

CREATE TABLE "shoplist_product" (
  "id"          SERIAL                      NOT NULL,
  "name"        VARCHAR(127)                NOT NULL,
  "date"        TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT now(),
  "cost"        DECIMAL                     NOT NULL DEFAULT 0,
  "result"      BOOLEAN                     NOT NULL DEFAULT FALSE,
  "shoplist_id" INT                         NOT NULL,
  "product_id"  INT                         NOT NULL,
  "count"       INT                         NOT NULL,
  CONSTRAINT shoplist_product_pk PRIMARY KEY ("id")
) WITH (OIDS = FALSE
);


CREATE TABLE "shoplist" (
  "id"         SERIAL                      NOT NULL,
  "name"       VARCHAR(127)                NOT NULL,
  "date"       TIMESTAMP(2) WITH TIME ZONE NOT NULL DEFAULT now(),
  CONSTRAINT shoplist_pk PRIMARY KEY ("id")
) WITH (OIDS = FALSE);



ALTER TABLE "product"
  ADD CONSTRAINT "product_fk0" FOREIGN KEY ("category_id") REFERENCES "category" ("id");
ALTER TABLE "category"
  ADD CONSTRAINT "category_fk0" FOREIGN KEY ("directory_id") REFERENCES "directory" ("id");
ALTER TABLE "shoplist_product"
  ADD CONSTRAINT "shoplist_product_fk0" FOREIGN KEY ("shoplist_id") REFERENCES "shoplist" ("id");
ALTER TABLE "shoplist_product"
  ADD CONSTRAINT "shoplist_product_fk1" FOREIGN KEY ("product_id") REFERENCES "product" ("id");



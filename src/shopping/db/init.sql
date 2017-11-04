CREATE TABLE "products" (
	"id" serial NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"count" integer NOT NULL DEFAULT '0',
	"photo" VARCHAR(255) NOT NULL DEFAULT '0',
	"shop_list_id" integer(255) NOT NULL DEFAULT '0',
	"category_id" integer(255) NOT NULL DEFAULT '0',
	CONSTRAINT products_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "shopping_list" (
	"id" serial NOT NULL,
	"cost" DECIMAL NOT NULL,
	"group_id" serial NOT NULL,
	CONSTRAINT shopping_list_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "category" (
	"id" serial NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	CONSTRAINT category_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



ALTER TABLE "products" ADD CONSTRAINT "products_fk0" FOREIGN KEY ("shop_list_id") REFERENCES "shopping_list"("id");
ALTER TABLE "products" ADD CONSTRAINT "products_fk1" FOREIGN KEY ("category_id") REFERENCES "category"("id");




CREATE TABLE IF NOT EXISTS "users" (
	"id" serial NOT NULL UNIQUE,
	"first_name" varchar(255) NULL,
	"last_name" varchar(255) NULL,
	"email" varchar(255) UNIQUE NULL,
	"phone" varchar(11) NULL UNIQUE,
	"username" varchar(255) NOT NULL UNIQUE,
	"password" varchar(255),
	"role" varchar(255) NOT NULL,
	"created_questions_count" bigint NOT NULL,
	"solved_questions_count" bigint NOT NULL,
	"submissions_count" bigint NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "questions" (
	"id" serial NOT NULL UNIQUE,
	"user_id" bigint NOT NULL,
	"publish_date" timestamp with time zone,
	"status" varchar(255) NOT NULL,
	"title" varchar(255) NOT NULL,
	"body" varchar(10000) NOT NULL,
	"time_limit" bigint NOT NULL,
	"memory_limit" bigint NOT NULL,
	"input_url" varchar(255) NOT NULL,
	"deadline" timestamp with time zone NOT NULL,
	"output_url" varchar(255) NOT NULL,
	"submissions_count" bigint NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "submissions" (
	"id" serial NOT NULL UNIQUE,
	"question_id" bigint NOT NULL,
	"user_id" bigint NOT NULL,
	"submit_url" varchar(255) NOT NULL,
	"status" bigint NOT NULL,
	"is_final" boolean NOT NULL DEFAULT '1',
	"submit_time" timestamp with time zone NOT NULL,
	"runner_started_at" TIMESTAMP NULL,
	"try_count" INT DEFAULT 0,
	PRIMARY KEY ("id")
);


ALTER TABLE "questions" ADD CONSTRAINT "questions_fk1" FOREIGN KEY ("user_id") REFERENCES "users"("id");
ALTER TABLE "submissions" ADD CONSTRAINT "submissions_fk1" FOREIGN KEY ("question_id") REFERENCES "questions"("id");

ALTER TABLE "submissions" ADD CONSTRAINT "submissions_fk2" FOREIGN KEY ("user_id") REFERENCES "users"("id");
CREATE INDEX "idx_submissions_status_runner_started_at" ON "submissions" ("status", "runner_started_at");

INSERT INTO "users" ("first_name", "last_name", "email", "phone", "username", "password", "role", "created_questions_count", "solved_questions_count", "submissions_count") VALUES ('admin', 'admin', 'admin@admin.com', '09121112233', 'admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin', 0, 0, 0);

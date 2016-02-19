-- Database init
CREATE USER jose WITH UNENCRYPTED PASSWORD 'jose';
CREATE DATABASE "jose";
GRANT ALL ON DATABASE "jose" TO "jose";

-- Switch to the audiences db as the audiences user.
\connect "jose";
set role "jose";
--
-- Table structure for table "follow"
--

CREATE TABLE "follow" (
  "id" SERIAL PRIMARY KEY,
  "userid" BIGINT NOT NULL,
  "username" varchar(100) NOT NULL,
  "status" text,
  "followdate" TIMESTAMP WITH TIME ZONE NOT NULL,
  "unfollowdate" TIMESTAMP WITH TIME ZONE DEFAULT NULL,
  "lastaction" TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX ON "follow" ("userid");

--
-- Table structure for table "tweet"
--

CREATE TABLE "tweet" (
  "id" SERIAL PRIMARY KEY,
  "content" text NOT NULL,
  "date" TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX ON "tweet" ("content");

--
-- Table structure for table "reply"
--

CREATE TABLE "reply" (
  "id" SERIAL PRIMARY KEY,
  "userid" BIGINT NOT NULL,
  "username" VARCHAR(100) NOT NULL,
  "tweetid" BIGINT NOT NULL,
  "status" TEXT NOT NULL,
  "answer" TEXT NOT NULL,
  "replydate" TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX ON "reply" ("tweetid");

--
-- Table structure for table "favorite"
--

CREATE TABLE "favorite" (
  "id" SERIAL PRIMARY KEY,
  "userid" BIGINT NOT NULL,
  "username" VARCHAR(100) NOT NULL,
  "tweetid" BIGINT NOT NULL,
  "status" TEXT NOT NULL,
  "favdate" TIMESTAMP WITH TIME ZONE NOT NULL,
  "unfavdate" TIMESTAMP WITH TIME ZONE NULL,
  "lastaction" TIMESTAMP WITH TIME ZONE NOT NULL
);

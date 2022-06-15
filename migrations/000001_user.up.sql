/*
For the purpose of this project I dont create segment table and users table is enough
*/
CREATE TABLE "users" (
    mobile      VARCHAR(15) UNIQUE NOT NULL PRIMARY KEY,
    credit      DECIMAL default 0
)

CREATE TABLE "code"(
    code VARCHAR(50) UNIQUE NOT NULL PRIMARY KEY,
    credit  DECIMAL DEFAULT 1000000,
    consumer_count  int default 1000
)


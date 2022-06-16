CREATE TABLE "users" (
    mobile      VARCHAR(15) UNIQUE NOT NULL PRIMARY KEY,
    credit      DECIMAL DEFAULT 0,
    recevied_charge BOOLEAN DEFAULT FALSE
);

CREATE TABLE "codes"(
    code VARCHAR(50) UNIQUE NOT NULL PRIMARY KEY,
    credit  DECIMAL DEFAULT 1000000,
    consumer_count  int default 1000
);


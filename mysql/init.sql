DROP DATABASE IF EXISTS lambda_rdb;
CREATE DATABASE lambda_rdb;

DROP TABLE IF EXISTS lambda_rdb.sample_table;
CREATE TABLE lambda_rdb.sample_table (
  id VARCHAR(11) PRIMARY KEY,
  title TEXT NOT NULL,
  created_at DATETIME,
  deleted_at DATETIME
);

INSERT INTO lambda_rdb.sample_table VALUES (
  '111-111-111',
  'This is a sample record',
  NOW(6),
  NOW(6)
)

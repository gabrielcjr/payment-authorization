CREATE TABLE IF NOT EXISTS transactions (
  id                UUID, 
  authorization_id  UUID,
  client_id         UUID,
  status            varchar(15), 
  value             float, 
  approved_at       timestamp, 
  denied_at         timestamp, 
  created_at        timestamp, 
  updated_at        timestamp, 
  deleted_at        timestamp,
  PRIMARY KEY (id)
);

ALTER TABLE transactions
RENAME COLUMN transaction_id TO client_id;

ALTER TABLE transactions
ADD COLUMN client_id UUID;

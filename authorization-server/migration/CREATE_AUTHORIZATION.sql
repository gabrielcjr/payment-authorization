CREATE TABLE IF NOT EXISTS authorizations (
  id                UUID, 
  client_id         UUID, 
  transaction_id    UUID, 
  status      varchar(15), 
  value       float, 
  approved_at timestamp, 
  denied_at   timestamp, 
  created_at  timestamp, 
  updated_at  timestamp, 
  deleted_at  timestamp,
  PRIMARY KEY (id)
);

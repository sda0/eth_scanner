CREATE ROLE ethscanner WITH LOGIN;
ALTER ROLE ethscanner WITH PASSWORD 'password';

CREATE TABLE accounts (
  "number" bytea PRIMARY KEY,
  balance NUMERIC NOT NULL,
  lastBlock BIGINT NOT NULL
);

CREATE TABLE transactions (
  hash TEXT PRIMARY KEY,
  blocknumber BIGINT NOT NULL,
  "from" bytea NOT NULL,
  "to" bytea NOT NULL,
  "value" NUMERIC NOT NULL,
  timestamp INT NOT NULL,
  transactionIndex BIGINT NOT NULL,
  "input" TEXT,
  v TEXT,
  r TEXT,
  s TEXT
);


CREATE INDEX idx_transactions_from
  ON transactions("from");

CREATE INDEX idx_transactions_to
  ON transactions("to");

CREATE INDEX idx_transactions_blockNumber
  ON transactions("blocknumber");


GRANT INSERT, SELECT, UPDATE ON TABLE transactions TO ethscanner;

CREATE OR REPLACE FUNCTION update_accounts_table()
  RETURNS trigger
LANGUAGE plpgsql
AS $$
BEGIN
  INSERT INTO accounts("number", balance, lastBlock) VALUES(NEW."from", -NEW.value, NEW.blocknumber) ON CONFLICT ("number") DO UPDATE SET balance = accounts.balance - NEW.value, lastBlock = GREATEST(accounts.lastBlock, NEW.blocknumber);
  IF NEW."to" <> ''  THEN
    INSERT INTO accounts("number", balance, lastBlock) VALUES(NEW."to", NEW.value, NEW.blocknumber) ON CONFLICT ("number") DO UPDATE SET balance = accounts.balance + NEW.value, lastBlock = GREATEST(accounts.lastBlock, NEW.blocknumber);
  END IF;

  RETURN NEW;
END;
$$;



CREATE TRIGGER update_account
  AFTER INSERT OR UPDATE
  ON transactions
  FOR EACH ROW
EXECUTE PROCEDURE update_accounts_table();
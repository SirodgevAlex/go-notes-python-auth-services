CREATE OR REPLACE FUNCTION update_last_password_change()
RETURNS TRIGGER AS $$
BEGIN
    NEW.last_password_change = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_last_password_change_trigger
BEFORE UPDATE ON users
FOR EACH ROW
WHEN (NEW.password <> OLD.password)
EXECUTE FUNCTION update_last_password_change();

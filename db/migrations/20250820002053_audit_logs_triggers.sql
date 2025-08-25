-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE OR REPLACE FUNCTION audit_trigger_function()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        INSERT INTO audit_logs(table_name, operation, user_id, old_data, new_data)
        VALUES (TG_TABLE_NAME, TG_OP, current_setting('app.user_id', true), row_to_json(OLD)::jsonb, row_to_json(NEW)::jsonb);

        RETURN OLD;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO audit_logs(table_name, operation, user_id, old_data, new_data)
        VALUES (TG_TABLE_NAME, TG_OP, current_setting('app.user_id', true), row_to_json(OLD)::jsonb, row_to_json(NEW)::jsonb);
        RETURN NEW;
    ELSIF (TG_OP = 'INSERT') THEN
        INSERT INTO audit_logs(table_name, operation, user_id, new_data)
        VALUES (TG_TABLE_NAME, TG_OP, current_setting('app.user_id', true), row_to_json(NEW)::jsonb);
        RETURN NEW;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP FUNCTION IF EXISTS audit_trigger_function() CASCADE;
-- +goose StatementEnd

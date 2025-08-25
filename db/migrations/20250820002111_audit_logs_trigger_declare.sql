-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
DO $$
DECLARE
    t RECORD;
BEGIN
    FOR t IN
        SELECT tablename FROM pg_tables 
        WHERE schemaname = 'public' AND tablename != 'audit_logs'
    LOOP
        EXECUTE format('
            CREATE TRIGGER %I_audit_trigger
            AFTER INSERT OR UPDATE OR DELETE ON %I
            FOR EACH ROW EXECUTE FUNCTION audit_trigger_function();
        ', t.tablename, t.tablename);
    END LOOP;
END;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DO $$
DECLARE
    t RECORD;
BEGIN
    FOR t IN
        SELECT tablename FROM pg_tables 
        WHERE schemaname = 'public' AND tablename != 'audit_logs'
    LOOP
        EXECUTE format('DROP TRIGGER IF EXISTS %I_audit_trigger ON %I;', t.tablename, t.tablename);
    END LOOP;
END;
$$;
-- +goose StatementEnd

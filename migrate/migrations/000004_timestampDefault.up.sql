ALTER TABLE todo.tasks ALTER COLUMN created_at SET DEFAULT NOW();
ALTER TABLE todo.users ALTER COLUMN created_at SET DEFAULT NOW();
ALTER TABLE todo.groups ALTER COLUMN created_at SET DEFAULT NOW();
ALTER TABLE todo.history ALTER COLUMN created_at SET DEFAULT NOW();

CREATE OR REPLACE FUNCTION todo.update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

DROP TRIGGER IF EXISTS update_tasks_row_timestamp_default ON todo.tasks;
CREATE TRIGGER update_tasks_row_timestamp_default
BEFORE UPDATE ON todo.tasks
FOR EACH ROW
EXECUTE PROCEDURE todo.update_updated_at_column();

DROP TRIGGER IF EXISTS update_users_row_timestamp_default ON todo.users;
CREATE TRIGGER update_users_row_timestamp_default
BEFORE UPDATE ON todo.users
FOR EACH ROW
EXECUTE PROCEDURE todo.update_updated_at_column();

DROP TRIGGER IF EXISTS update_groups_row_timestamp_default ON todo.groups;
CREATE TRIGGER update_groups_row_timestamp_default
BEFORE UPDATE ON todo.groups
FOR EACH ROW
EXECUTE PROCEDURE todo.update_updated_at_column();
DROP TRIGGER IF EXISTS update_tasks_row_timestamp_default ON todo.tasks;
DROP TRIGGER IF EXISTS update_users_row_timestamp_default ON todo.users;
DROP TRIGGER IF EXISTS update_groups_row_timestamp_default ON todo.groups;
DROP FUNCTION IF EXISTS todo.update_updated_at_column();

ALTER TABLE todo.tasks ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE todo.users ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE todo.groups ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE todo.history ALTER COLUMN created_at DROP DEFAULT;
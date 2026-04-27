DROP TYPE IF EXISTS todo.EnumTaskStatus CASCADE;
ALTER TABLE todo.history DROP CONSTRAINT IF EXISTS fk_history_prev;
DROP TABLE todo.history;
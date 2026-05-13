ALTER TABLE todo.tasks DROP COLUMN IF EXISTS deleted_at;

ALTER TYPE todo.EnumTaskStatus RENAME TO EnumTaskStatusOld;
CREATE TYPE todo.EnumTaskStatus AS ENUM ('in_progress', 'done');
ALTER TABLE todo.history ALTER COLUMN status SET DATA TYPE todo.EnumTaskStatus USING status::text::todo.EnumTaskStatus;
DROP TYPE IF EXISTS todo.EnumTaskStatusOld;
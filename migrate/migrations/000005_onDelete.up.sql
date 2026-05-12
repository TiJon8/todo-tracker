ALTER TABLE todo.tasks ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ;

ALTER TYPE todo.EnumTaskStatus ADD VALUE IF NOT EXISTS 'deleted' AFTER 'done';
ALTER TYPE todo.EnumTaskStatus ADD VALUE IF NOT EXISTS 'patched' AFTER 'in_progress';

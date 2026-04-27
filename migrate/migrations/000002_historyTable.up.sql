CREATE TYPE todo.EnumTaskStatus AS ENUM ('in_progress', 'done');

CREATE TABLE IF NOT EXISTS todo.history (
	id UUID PRIMARY KEY DEFAULT uuidv4(),
	row_version INTEGER NOT NULL DEFAULT 1,

	task_id UUID NOT NULL REFERENCES todo.tasks(id),
	status todo.EnumTaskStatus NOT NULL, 
	updated_by UUID NOT NULL REFERENCES todo.users(id),
	previous UUID,

	title VARCHAR(100) NOT NULL,
	description text,

	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ
);

ALTER TABLE todo.history ADD CONSTRAINT fk_history_prev FOREIGN KEY (previous) REFERENCES todo.history(id);
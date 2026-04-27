CREATE SCHEMA todo;

CREATE TABLE IF NOT EXISTS todo.users(
	id UUID PRIMARY KEY DEFAULT uuidv4(),
	row_version INTEGER NOT NULL DEFAULT 1,

	name VARCHAR(32),
	phone VARCHAR(16) CHECK (
		phone ~ '^\+[0-9]{10,16}$'
	),

	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS todo.groups(
	id UUID PRIMARY KEY DEFAULT uuidv4(),
	row_version INTEGER NOT NULL DEFAULT 1,

	name VARCHAR(32) NOT NULL,
	owner UUID NOT NULL REFERENCES todo.users(id),

	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS todo.tasks(
	id UUID PRIMARY KEY DEFAULT uuidv4(),
	row_version INTEGER NOT NULL DEFAULT 1,

	title VARCHAR(100) NOT NULL,
	description text,
	completed BOOLEAN NOT NULL DEFAULT FALSE,
	completed_at TIMESTAMPTZ,

	user_id UUID NOT NULL REFERENCES todo.users(id),
	group_id UUID REFERENCES todo.groups(id),

	created_at TIMESTAMPTZ,
	updated_at TIMESTAMPTZ,

	CHECK (
		(completed=FALSE AND completed_at IS NULL)
		OR
		(completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
	)
);

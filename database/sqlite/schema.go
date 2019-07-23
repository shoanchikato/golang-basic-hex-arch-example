package sqlite

// Schema for the database
var Schema = `CREATE TABLE IF NOT EXISTS toys (
	id TEXT NOT NULL PRIMARY KEY, 
	name TEXT,
	owner TEXT,
	createdAt DATE,
	updatedAt DATE,
	deletedAt DATE
);`
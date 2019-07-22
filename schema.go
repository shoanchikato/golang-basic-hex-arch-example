package main

stmt := `CREATE TABLE IF NOT EXISTS toys (
				id TEXT NOT NULL PRIMARY KEY, 
				name TEXT,
				owner TEXT,
				createdAt DATE,
				updatedAt DATE,
				deletedAt DATE
			);`
CREATE TABLE IF NOT EXISTS "WorldDirectory" (
	`world_id`	BLOB NOT NULL,
	`name`	TEXT NOT NULL UNIQUE,
	`hostport`	TEXT NOT NULL,

  /* optional */
  `creation_time_usec` INTEGER,
  `last_hearbeat_usec` INTEGER,

	PRIMARY KEY(`world_id`)
);

INSERT INTO "WorldDirectory" 
VALUES ("123-456-789", "world1", "localhost:8080"),
       ("123-456-123", "world2", "localhost:8080");


world
  id
  name
  hostport
  creation_time_usec
  last_hearbeat_usec


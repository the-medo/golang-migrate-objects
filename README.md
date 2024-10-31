# golang-migrate-objects

## Overview

`golang-migrate-objects` is a Go-based tool that extends `golang-migrate` to manage database objects. It allows for migrations of complex database objects by applying structured steps. It includes automatic generation of files for object creation and dropping, helping maintain consistency across versions.


## Main features
- Supports both up and down migrations with database objects.
- Can automatically generate a "sum file" that merges the latest versions of DB objects.
- Utilizes a configuration structure to define paths and database settings.
- Supports command-line arguments to control migration actions and file management.


## Project structure
- **main.go**: Initializes configuration and manages command-line interface.
- **/migrator/migrator.go**: Contains core migration logic for managing DB objects.
- **/migrator/types.go**: Defines data structures used in the migrator.

---

## Configuration Structure
Config Struct (in /migrator/types.go)

| Field                | Type     | Description                        |
|----------------------|----------|------------------------------------|
| DB                   | *sql.DB  | Database connection instance.      |
| PriorityLpad         | int      | Padding length for priority fields.|
| VersionLpad          | int      | Padding length for version fields. |
| MigrationFilesPath   | string   | Path to SQL migration files.       |
| DbObjectPath         | string   | Path to DB object files.           |
| CreateObjectsFilename| string   | Filename for creating merged sum files.|
| DropObjectsFilename  | string   | Filename for object drop scripts.  |


## Usage instructions
1. Set Up Migration Paths
    - `mpath`: Path to migration files.
    - `obj_path`: Path to DB object files.
    - `db_source`: Database connection string.
2. Execute commands
    - Run commands with `go run main.go` with flags to define migration behavior:
    `go run main.go -mpath=<migration_path> -obj_path=<object_path> -db_source=<db_source>`


## Available Command-Line Flags

| Flag             | Type     | Description                        |
|------------------|----------|------------------------------------|
| `-mpath`	        | string	| 	Path to SQL migration files. Required	| 
| `-obj_path`	     | 	string	| 	Path to DB object files. Required	|
| `-db_source`	    | 	string	| 	Database connection string. Required	|
| `-co_filename`	  | 	string	| 	Filename for generating the sum file. Required	|
| `-do_filename`		 | string	| 	Filename for object drop scripts. Required	| 
| `-sumfile`		     | bool	| 	If set, creates a sum file of all object versions.	| 
| `-up`            | 	bool	| 	Runs migrations in the up direction.	| 
| `-down`	         | 	bool		| Runs migrations in the down direction.	| 
| `-step`	         | 	int	| 	Number of steps to migrate. If 0, runs all migrations.	|


## Core Components

### Migrator Struct (in /migrator/migrator.go)

| Field   | Type     | Description                                     |
|---------|----------|-------------------------------------------------|
| Migrate | *migrate.Migrate  | Instance for migration operations.              |
| Config  | *Config      | Configuration settings from the `Config` struct.  |

### Methods in Migrator
- GetObjectList: Returns a list of DB objects based on the directory structure.
- GetObjectsForStep: Retrieves DB objects for a specific migration step, considering direction.
- CreateObjectsForStep: Creates objects required for a specific step.
- CreateObjectsFile: Merges latest versions of DB objects into a single file.
- DropObjects: Executes a script to drop DB objects.
- RunFile: Executes SQL file at the specified path.
- RunAll: Migrates the database to the highest available version.

### Additional Helper Functions
- `parseDir` and `parseVersionFiles`: Parses directories and files to extract object versions.
- `orderDirEntries`: Orders directory entries by name.

## Example usage
1. Run All Migrations Up
```sh
go run main.go -mpath="path/to/migrations" -obj_path="path/to/objects" -db_source="postgres://user:pass@host/db" -up
```

2. Generate a Sum File
```sh
go run main.go -mpath="path/to/migrations" -obj_path="path/to/objects" -db_source="postgres://user:pass@host/db" -sumfile
```

## Error Handling
The package includes several error checks, such as:
- **ErrDifferentPriorityLength**: Ensures all priority values are of equal length.
- **ErrInvalidPriority**: Ensures priority values are valid integers.
- **ErrDifferentVersionLength**: Checks if all version numbers have consistent lengths.
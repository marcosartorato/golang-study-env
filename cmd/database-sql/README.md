# Go `database/sql` (curated)

The [`database/sql`](https://pkg.go.dev/database/sql) package provides a DB-agnostic API.  
This demo uses the pure-Go SQLite driver (`modernc.org/sqlite`) and an in-memory database.

## Common Functions

- Connect: `sql.Open`, `(*DB).PingContext`
- Execute: `(*DB).ExecContext` (with `?` placeholders)
- Query single/many: `(*DB).QueryRowContext`, `(*DB).QueryContext`, `(*Rows).Next`, `(*Rows).Scan`
- Prepared statements: `(*DB).PrepareContext`, `(*Stmt).ExecContext`
- Transactions: `(*DB).BeginTx`, `(*Tx).ExecContext`, `(*Tx).Commit`

```go
func Open(driverName, dataSourceName string) (*DB, error)
func (db *DB) PingContext(ctx context.Context) error
func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
```

`Open` opens a database specified by its database driver name and a driver-specific data source name.
`PingContext` verifies a connection to the database is still alive, establishing a connection if necessary.
`ExecContext` executes a query without returning any rows. The args are for any placeholder parameters in the query.

```go
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row
func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
func (rs *Rows) Next() bool
func (rs *Rows) Scan(dest ...any) error
```

`QueryRowContext` executes a query that is expected to return at most one row.
`QueryContext` executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
`Next` prepares the next result row for reading with the `Rows.Scan` method. It returns true on success, or false if there is no next result row or an error happened while preparing it. `Rows.Err` should be consulted to distinguish between the two cases.

```go
func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error)
func (s *Stmt) ExecContext(ctx context.Context, args ...any) (Result, error)
```

PrepareContext creates a prepared statement for later queries or executions. Multiple queries or executions may be run concurrently from the returned statement. The caller must call the statement's `*Stmt.Close` method when the statement is no longer needed.
`ExecContext` executes a prepared statement with the given arguments and returns a Result summarizing the effect of the statement.

```go
func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)
func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
func (tx *Tx) Commit() error
```

`BeginTx` starts a transaction.
`ExecContext` executes a query without returning any rows. The args are for any placeholder parameters in the query.
`Commit` commits the transaction.

---

[Go Back](../../README.md)
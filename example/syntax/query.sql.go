// Code generated by pggen. DO NOT EDIT.

package syntax

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Querier is a typesafe Go interface backed by SQL queries.
//
// Methods ending with Batch enqueue a query to run later in a pgx.Batch. After
// calling SendBatch on pgx.Conn, pgxpool.Pool, or pgx.Tx, use the Scan methods
// to parse the results.
type Querier interface {
	// Query to test escaping in generated Go.
	Backtick(ctx context.Context) (string, error)
	// BacktickBatch enqueues a Backtick query into batch to be executed
	// later by the batch.
	BacktickBatch(batch *pgx.Batch)
	// BacktickScan scans the result of an executed BacktickBatch query.
	BacktickScan(results pgx.BatchResults) (string, error)

	// Query to test escaping in generated Go.
	BacktickQuoteBacktick(ctx context.Context) (string, error)
	// BacktickQuoteBacktickBatch enqueues a BacktickQuoteBacktick query into batch to be executed
	// later by the batch.
	BacktickQuoteBacktickBatch(batch *pgx.Batch)
	// BacktickQuoteBacktickScan scans the result of an executed BacktickQuoteBacktickBatch query.
	BacktickQuoteBacktickScan(results pgx.BatchResults) (string, error)

	// Query to test escaping in generated Go.
	BacktickNewline(ctx context.Context) (string, error)
	// BacktickNewlineBatch enqueues a BacktickNewline query into batch to be executed
	// later by the batch.
	BacktickNewlineBatch(batch *pgx.Batch)
	// BacktickNewlineScan scans the result of an executed BacktickNewlineBatch query.
	BacktickNewlineScan(results pgx.BatchResults) (string, error)

	// Query to test escaping in generated Go.
	BacktickDoubleQuote(ctx context.Context) (string, error)
	// BacktickDoubleQuoteBatch enqueues a BacktickDoubleQuote query into batch to be executed
	// later by the batch.
	BacktickDoubleQuoteBatch(batch *pgx.Batch)
	// BacktickDoubleQuoteScan scans the result of an executed BacktickDoubleQuoteBatch query.
	BacktickDoubleQuoteScan(results pgx.BatchResults) (string, error)

	// Query to test escaping in generated Go.
	BacktickBackslashN(ctx context.Context) (string, error)
	// BacktickBackslashNBatch enqueues a BacktickBackslashN query into batch to be executed
	// later by the batch.
	BacktickBackslashNBatch(batch *pgx.Batch)
	// BacktickBackslashNScan scans the result of an executed BacktickBackslashNBatch query.
	BacktickBackslashNScan(results pgx.BatchResults) (string, error)

	// Illegal names.
	IllegalNameSymbols(ctx context.Context, helloWorld string) (IllegalNameSymbolsRow, error)
	// IllegalNameSymbolsBatch enqueues a IllegalNameSymbols query into batch to be executed
	// later by the batch.
	IllegalNameSymbolsBatch(batch *pgx.Batch, helloWorld string)
	// IllegalNameSymbolsScan scans the result of an executed IllegalNameSymbolsBatch query.
	IllegalNameSymbolsScan(results pgx.BatchResults) (IllegalNameSymbolsRow, error)

	// Enum named 123.
	BadEnumName(ctx context.Context) (UnnamedEnum123, error)
	// BadEnumNameBatch enqueues a BadEnumName query into batch to be executed
	// later by the batch.
	BadEnumNameBatch(batch *pgx.Batch)
	// BadEnumNameScan scans the result of an executed BadEnumNameBatch query.
	BadEnumNameScan(results pgx.BatchResults) (UnnamedEnum123, error)

	GoKeyword(ctx context.Context, go_ string) (string, error)
	// GoKeywordBatch enqueues a GoKeyword query into batch to be executed
	// later by the batch.
	GoKeywordBatch(batch *pgx.Batch, go_ string)
	// GoKeywordScan scans the result of an executed GoKeywordBatch query.
	GoKeywordScan(results pgx.BatchResults) (string, error)
}

type DBQuerier struct {
	conn genericConn
}

var _ Querier = &DBQuerier{}

// genericConn is a connection to a Postgres database. This is usually backed by
// *pgx.Conn, pgx.Tx, or *pgxpool.Pool.
type genericConn interface {
	// Query executes sql with args. If there is an error the returned Rows will
	// be returned in an error state. So it is allowed to ignore the error
	// returned from Query and handle it in Rows.
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)

	// QueryRow is a convenience wrapper over Query. Any error that occurs while
	// querying is deferred until calling Scan on the returned Row. That Row will
	// error with pgx.ErrNoRows if no rows are returned.
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row

	// Exec executes sql. sql can be either a prepared statement name or an SQL
	// string. arguments should be referenced positionally from the sql string
	// as $1, $2, etc.
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

// NewQuerier creates a DBQuerier that implements Querier. conn is typically
// *pgx.Conn, pgx.Tx, or *pgxpool.Pool.
func NewQuerier(conn genericConn) *DBQuerier {
	return &DBQuerier{
		conn: conn,
	}
}

// WithTx creates a new DBQuerier that uses the transaction to run all queries.
func (q *DBQuerier) WithTx(tx pgx.Tx) (*DBQuerier, error) {
	return &DBQuerier{conn: tx}, nil
}

// ignoredOID means we don't know or care about the OID for a type. This is okay
// because pgx only uses the OID to encode values and lookup a decoder. We only
// use ignoredOID for decoding and we always specify a concrete decoder for scan
// methods.
const ignoredOID = 0

// UnnamedEnum123 represents the Postgres enum "123".
type UnnamedEnum123 string

const (
	UnnamedEnum123InconvertibleEnumName UnnamedEnum123 = "inconvertible_enum_name"
	UnnamedEnum123UnnamedLabel1         UnnamedEnum123 = ""
	UnnamedEnum123UnnamedLabel2111      UnnamedEnum123 = "111"
	UnnamedEnum123UnnamedLabel3         UnnamedEnum123 = "!!"
)

func (u UnnamedEnum123) String() string { return string(u) }

const backtickSQL = "SELECT '`';"

// Backtick implements Querier.Backtick.
func (q *DBQuerier) Backtick(ctx context.Context) (string, error) {
	row := q.conn.QueryRow(ctx, backtickSQL)
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query Backtick: %w", err)
	}
	return item, nil
}

// BacktickBatch implements Querier.BacktickBatch.
func (q *DBQuerier) BacktickBatch(batch *pgx.Batch) {
	batch.Queue(backtickSQL)
}

// BacktickScan implements Querier.BacktickScan.
func (q *DBQuerier) BacktickScan(results pgx.BatchResults) (string, error) {
	row := results.QueryRow()
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan BacktickBatch row: %w", err)
	}
	return item, nil
}

const backtickQuoteBacktickSQL = "SELECT '`\"`';"

// BacktickQuoteBacktick implements Querier.BacktickQuoteBacktick.
func (q *DBQuerier) BacktickQuoteBacktick(ctx context.Context) (string, error) {
	row := q.conn.QueryRow(ctx, backtickQuoteBacktickSQL)
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query BacktickQuoteBacktick: %w", err)
	}
	return item, nil
}

// BacktickQuoteBacktickBatch implements Querier.BacktickQuoteBacktickBatch.
func (q *DBQuerier) BacktickQuoteBacktickBatch(batch *pgx.Batch) {
	batch.Queue(backtickQuoteBacktickSQL)
}

// BacktickQuoteBacktickScan implements Querier.BacktickQuoteBacktickScan.
func (q *DBQuerier) BacktickQuoteBacktickScan(results pgx.BatchResults) (string, error) {
	row := results.QueryRow()
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan BacktickQuoteBacktickBatch row: %w", err)
	}
	return item, nil
}

const backtickNewlineSQL = "SELECT '`\n';"

// BacktickNewline implements Querier.BacktickNewline.
func (q *DBQuerier) BacktickNewline(ctx context.Context) (string, error) {
	row := q.conn.QueryRow(ctx, backtickNewlineSQL)
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query BacktickNewline: %w", err)
	}
	return item, nil
}

// BacktickNewlineBatch implements Querier.BacktickNewlineBatch.
func (q *DBQuerier) BacktickNewlineBatch(batch *pgx.Batch) {
	batch.Queue(backtickNewlineSQL)
}

// BacktickNewlineScan implements Querier.BacktickNewlineScan.
func (q *DBQuerier) BacktickNewlineScan(results pgx.BatchResults) (string, error) {
	row := results.QueryRow()
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan BacktickNewlineBatch row: %w", err)
	}
	return item, nil
}

const backtickDoubleQuoteSQL = "SELECT '`\"';"

// BacktickDoubleQuote implements Querier.BacktickDoubleQuote.
func (q *DBQuerier) BacktickDoubleQuote(ctx context.Context) (string, error) {
	row := q.conn.QueryRow(ctx, backtickDoubleQuoteSQL)
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query BacktickDoubleQuote: %w", err)
	}
	return item, nil
}

// BacktickDoubleQuoteBatch implements Querier.BacktickDoubleQuoteBatch.
func (q *DBQuerier) BacktickDoubleQuoteBatch(batch *pgx.Batch) {
	batch.Queue(backtickDoubleQuoteSQL)
}

// BacktickDoubleQuoteScan implements Querier.BacktickDoubleQuoteScan.
func (q *DBQuerier) BacktickDoubleQuoteScan(results pgx.BatchResults) (string, error) {
	row := results.QueryRow()
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan BacktickDoubleQuoteBatch row: %w", err)
	}
	return item, nil
}

const backtickBackslashNSQL = "SELECT '`\\n';"

// BacktickBackslashN implements Querier.BacktickBackslashN.
func (q *DBQuerier) BacktickBackslashN(ctx context.Context) (string, error) {
	row := q.conn.QueryRow(ctx, backtickBackslashNSQL)
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query BacktickBackslashN: %w", err)
	}
	return item, nil
}

// BacktickBackslashNBatch implements Querier.BacktickBackslashNBatch.
func (q *DBQuerier) BacktickBackslashNBatch(batch *pgx.Batch) {
	batch.Queue(backtickBackslashNSQL)
}

// BacktickBackslashNScan implements Querier.BacktickBackslashNScan.
func (q *DBQuerier) BacktickBackslashNScan(results pgx.BatchResults) (string, error) {
	row := results.QueryRow()
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan BacktickBackslashNBatch row: %w", err)
	}
	return item, nil
}

const illegalNameSymbolsSQL = "SELECT '`\\n' as \"$\", $1 as \"foo.bar!@#$%&*()\"\"--+\";"

type IllegalNameSymbolsRow struct {
	UnnamedColumn0 string `json:"$"`
	FooBar         string `json:"foo.bar!@#$%&*()\"--+"`
}

// IllegalNameSymbols implements Querier.IllegalNameSymbols.
func (q *DBQuerier) IllegalNameSymbols(ctx context.Context, helloWorld string) (IllegalNameSymbolsRow, error) {
	row := q.conn.QueryRow(ctx, illegalNameSymbolsSQL, helloWorld)
	var item IllegalNameSymbolsRow
	if err := row.Scan(&item.UnnamedColumn0, &item.FooBar); err != nil {
		return item, fmt.Errorf("query IllegalNameSymbols: %w", err)
	}
	return item, nil
}

// IllegalNameSymbolsBatch implements Querier.IllegalNameSymbolsBatch.
func (q *DBQuerier) IllegalNameSymbolsBatch(batch *pgx.Batch, helloWorld string) {
	batch.Queue(illegalNameSymbolsSQL, helloWorld)
}

// IllegalNameSymbolsScan implements Querier.IllegalNameSymbolsScan.
func (q *DBQuerier) IllegalNameSymbolsScan(results pgx.BatchResults) (IllegalNameSymbolsRow, error) {
	row := results.QueryRow()
	var item IllegalNameSymbolsRow
	if err := row.Scan(&item.UnnamedColumn0, &item.FooBar); err != nil {
		return item, fmt.Errorf("scan IllegalNameSymbolsBatch row: %w", err)
	}
	return item, nil
}

const badEnumNameSQL = `SELECT 'inconvertible_enum_name'::"123";`

// BadEnumName implements Querier.BadEnumName.
func (q *DBQuerier) BadEnumName(ctx context.Context) (UnnamedEnum123, error) {
	row := q.conn.QueryRow(ctx, badEnumNameSQL)
	var item UnnamedEnum123
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query BadEnumName: %w", err)
	}
	return item, nil
}

// BadEnumNameBatch implements Querier.BadEnumNameBatch.
func (q *DBQuerier) BadEnumNameBatch(batch *pgx.Batch) {
	batch.Queue(badEnumNameSQL)
}

// BadEnumNameScan implements Querier.BadEnumNameScan.
func (q *DBQuerier) BadEnumNameScan(results pgx.BatchResults) (UnnamedEnum123, error) {
	row := results.QueryRow()
	var item UnnamedEnum123
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan BadEnumNameBatch row: %w", err)
	}
	return item, nil
}

const goKeywordSQL = `SELECT $1::text;`

// GoKeyword implements Querier.GoKeyword.
func (q *DBQuerier) GoKeyword(ctx context.Context, go_ string) (string, error) {
	row := q.conn.QueryRow(ctx, goKeywordSQL, go_)
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query GoKeyword: %w", err)
	}
	return item, nil
}

// GoKeywordBatch implements Querier.GoKeywordBatch.
func (q *DBQuerier) GoKeywordBatch(batch *pgx.Batch, go_ string) {
	batch.Queue(goKeywordSQL, go_)
}

// GoKeywordScan implements Querier.GoKeywordScan.
func (q *DBQuerier) GoKeywordScan(results pgx.BatchResults) (string, error) {
	row := results.QueryRow()
	var item string
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan GoKeywordBatch row: %w", err)
	}
	return item, nil
}

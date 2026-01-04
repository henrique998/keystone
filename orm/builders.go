package orm

type QueryBuilder interface {
	Select(columns ...string) QueryBuilder
	OrderBy(column string, direction string) QueryBuilder
	Limit(limit int) QueryBuilder
	Offset(offset int) QueryBuilder
	Exec(dest interface{}) error
	WithDeleted() QueryBuilder
	OnlyDeleted() QueryBuilder
	Include(opt WithOption) QueryBuilder
	Require(opt WithOption) QueryBuilder
	IncludeFrom(opt WithOption) QueryBuilder
}

type UpdateBuilder interface {
	Set(column string, value interface{}) UpdateBuilder
	SetMap(values map[string]interface{}) UpdateBuilder
	Exec() error
}

type DeleteBuilder interface {
	Exec() error
	ExecReturningCount() (int64, error)
	AllowUnsafe() DeleteBuilder
	Force() DeleteBuilder
}

type DeleteBatchBuilder interface {
	AllowUnsafe() DeleteBatchBuilder
	HardDelete() DeleteBatchBuilder
	Exec() error
}

type ColumnBuilder interface {
	PrimaryKey() ColumnBuilder
	Unique() ColumnBuilder
	NotNull() ColumnBuilder
	Check(expr string) ColumnBuilder
	DefaultNow() ColumnBuilder
	DefaultUUID() ColumnBuilder
	AutoIncrement() ColumnBuilder
	References(table, column string) ColumnBuilder
	OnDelete(action string) ColumnBuilder
	OnUpdate(action string) ColumnBuilder
	Default(value string) ColumnBuilder
	BelongsTo(name string, refTable string) ColumnBuilder
}

type TableBuilder interface {
	Char(name string, size int) ColumnBuilder
	Varchar(name string, size int) ColumnBuilder
	Text(name string) ColumnBuilder
	SmallInt(name string) ColumnBuilder
	Int(name string) ColumnBuilder
	BigInt(name string) ColumnBuilder
	SmallSerial(name string) ColumnBuilder
	Serial(name string) ColumnBuilder
	BigSerial(name string) ColumnBuilder
	Numeric(name string, precision ...int) ColumnBuilder
	Decimal(name string, precision, scale int) ColumnBuilder
	Real(name string) ColumnBuilder
	Double(name string) ColumnBuilder
	Float(name string, precision ...int) ColumnBuilder
	Bool(name string) ColumnBuilder
	Date(name string) ColumnBuilder
	Timestamp(name string) ColumnBuilder
	Time(name string) ColumnBuilder
	Timestamptz(name string) ColumnBuilder
	TimeWithTimeZone(name string) ColumnBuilder
	JSON(name string) ColumnBuilder
	JSONB(name string) ColumnBuilder
	UUID(name string) ColumnBuilder
	Bytea(name string) ColumnBuilder
}

package database

// ================================================== FILTERS =================================================

type filter struct {
	field string
	op    string
	value any
	table string
}

type filterColumn[T any] struct {
	table string
	name  string
}

type compoundFilter struct {
	op      string
	filters []filter
}

// ================================================== TABLE ===================================================

type relationBelongsTo struct {
	name       string
	columnName string
	refTable   string
}

type tableDefinition struct {
	name          string
	columns       map[string]*columnInfo
	softDeleteCol string

	belongsTo []*relationBelongsTo
}

type tableBuilder struct {
	tableName string
	columns   map[string]*columnInfo

	belongsTo []*relationBelongsTo
}

// ================================================== COLUMN ===================================================

type condition struct {
	sql  string
	args []any
}

type column[T any] struct {
	tableName  string
	columnName string
	sqlType    string
}

// ================================================== BUILDERS ==================================================

type foreignKeyDef struct {
	RefTable  string `json:"ref_table"`
	RefColumn string `json:"ref_column"`
}

type columnInfo struct {
	name        string
	sqlType     string
	constraints []string
	ForeignKey  *foreignKeyDef `json:"foreign_key,omitempty"`
}

type columnBuilder struct {
	col *columnInfo
	tb  *tableBuilder
}

type queryBuilder struct {
	db             db
	table          string
	tableDef       *tableDefinition
	columns        []string
	whereClause    string
	args           []any
	orderBy        string
	limit          int
	offset         int
	includeDeleted bool
	onlyDeleted    bool
	joins          []joinDef
	hasJoins       bool
}

type joinDef struct {
	joinType    string
	localTable  string // original table ex: posts
	localColumn string // original column ex: user_id
	refTable    string // joined table ex: users
	refColumn   string // joined column ex: id
}

type deleteBuilder struct {
	db          db
	table       string
	filters     []any
	allowUnsafe bool
	hardDelete  bool
}

type deleteBatchBuilder struct {
	db          db
	table       string
	filters     []any
	allowUnsafe bool
	hardDelete  bool
}

type updateBuilder struct {
	db       db
	table    string
	tableDef *tableDefinition
	sets     map[string]any
	filters  []any
}

package models

// Generated with gopkg.in/reform.v1. Do not edit by hand.

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/xaionaro/reform"
)

type PipelineScope struct {
	item *Pipeline

	db           reform.ReformDBTX
	where        [][]interface{}
	order        []string
	groupBy      []string
	limit        int
	tableQuery   *string
	fieldsFilter []string
	appendTail   string

	loggingEnabled bool
	loggingAuthor  *string
	loggingComment string
}
type PipelineFilter Pipeline

type PipelineLogRow struct {
	Pipeline
	LogAuthor  *string
	LogAction  string
	LogDate    time.Time
	LogComment string
}

// Schema returns a schema name in SQL database ("").
type pipelineTableTypeType struct {
	s reform.StructInfo
	z []interface{}
}

func (v pipelineTableTypeType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("pipelines").
func (v pipelineTableTypeType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v pipelineTableTypeType) Columns() []string {
	return []string{"id", "project_name", "tag_name", "created_at", "approved_at"}
}

// NewStruct makes a new struct for that view or table.
func (v pipelineTableTypeType) NewStruct() reform.Struct {
	return new(Pipeline)
}

// NewRecord makes a new record for that table.
func (v *pipelineTableTypeType) NewRecord() reform.Record {
	return new(Pipeline)
}

func (v *pipelineTableTypeType) NewScope() *PipelineScope {
	return &PipelineScope{item: &Pipeline{}}
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *pipelineTableTypeType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

func (v pipelineTableTypeType) CreateTableIfNotExists(db *reform.DB) (bool, error) {
	if db == nil {
		db = defaultDB_Pipeline
	}
	return db.CreateTableIfNotExists(v.s)
}

func (v pipelineTableTypeType) StructInfo() reform.StructInfo {
	return v.s
}

// PipelineTable represents pipelines view or table in SQL database.
var PipelineTable = &pipelineTableTypeType{
	s: reform.StructInfo{Type: "Pipeline", SQLSchema: "", SQLName: "pipelines", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "ProjectName", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "project_name", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "TagName", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "tag_name", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "CreatedAt", IsPK: false, IsUnique: false, HasIndex: false, Type: "*extime.Time", Column: "created_at", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "ApprovedAt", IsPK: false, IsUnique: false, HasIndex: false, Type: "*extime.Time", Column: "approved_at", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(Pipeline).Values(),
}

type pipelineTableTypeType_log struct {
	s reform.StructInfo
	z []interface{}
}

func (v *pipelineTableTypeType_log) Schema() string {
	return v.s.SQLSchema
}

func (v *pipelineTableTypeType_log) Name() string {
	return v.s.SQLName
}

func (v *pipelineTableTypeType_log) Columns() []string {
	return []string{"id", "project_name", "tag_name", "created_at", "approved_at", "log_author", "log_action", "log_date", "log_comment"}
}

func (v *pipelineTableTypeType_log) NewStruct() reform.Struct {
	return new(Pipeline)
}

func (v *pipelineTableTypeType_log) NewRecord() reform.Record {
	return new(Pipeline)
}

func (v *pipelineTableTypeType_log) NewScope() *PipelineScope {
	return &PipelineScope{item: &Pipeline{}}
}

func (v *pipelineTableTypeType_log) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

var PipelineTableLogRow = &pipelineTableTypeType_log{
	s: reform.StructInfo{Type: "Pipeline", SQLSchema: "", SQLName: "pipelines_log", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "ProjectName", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "project_name", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "TagName", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "tag_name", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "CreatedAt", IsPK: false, IsUnique: false, HasIndex: false, Type: "*extime.Time", Column: "created_at", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "ApprovedAt", IsPK: false, IsUnique: false, HasIndex: false, Type: "*extime.Time", Column: "approved_at", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAuthor", IsPK: false, IsUnique: false, HasIndex: false, Type: "*string", Column: "log_author", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAction", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_action", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogDate", IsPK: false, IsUnique: false, HasIndex: false, Type: "time.Time", Column: "log_date", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogComment", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_comment", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(PipelineLogRow).Values(),
}

func (s pipelineTableTypeType) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	case "ProjectName":
		return "project_name"
	case "TagName":
		return "tag_name"
	case "CreatedAt":
		return "created_at"
	case "ApprovedAt":
		return "approved_at"
	}
	return ""
}

func (s pipelineTableTypeType_log) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	case "ProjectName":
		return "project_name"
	case "TagName":
		return "tag_name"
	case "CreatedAt":
		return "created_at"
	case "ApprovedAt":
		return "approved_at"
	case "LogAuthor":
		return "log_author"
	case "LogAction":
		return "log_action"
	case "LogDate":
		return "log_date"
	case "LogComment":
		return "log_comment"
	}
	return ""
}

func (s *Pipeline) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
	if len(fieldNames) == 0 {
		return s.Pointers()
	}

	for _, fieldName := range fieldNames {
		fieldPointer := s.FieldPointerByName(fieldName)
		if fieldPointer == nil {
			panic("Invalid field name:" + fieldName)
		}
		fieldPointers = append(fieldPointers, fieldPointer)
	}

	return
}

func (s *PipelineLogRow) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
	if len(fieldNames) == 0 {
		return s.Pointers()
	}

	for _, fieldName := range fieldNames {
		fieldPointer := s.FieldPointerByName(fieldName)
		if fieldPointer == nil {
			panic("Invalid field name:" + fieldName)
		}
		fieldPointers = append(fieldPointers, fieldPointer)
	}

	return
}

func (s *Pipeline) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	case "ProjectName":
		return &s.ProjectName
	case "TagName":
		return &s.TagName
	case "CreatedAt":
		return &s.CreatedAt
	case "ApprovedAt":
		return &s.ApprovedAt
	}

	return nil
}

func (s *PipelineLogRow) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	case "ProjectName":
		return &s.ProjectName
	case "TagName":
		return &s.TagName
	case "CreatedAt":
		return &s.CreatedAt
	case "ApprovedAt":
		return &s.ApprovedAt
	case "LogAuthor":
		return &s.LogAuthor
	case "LogAction":
		return &s.LogAction
	case "LogDate":
		return &s.LogDate
	case "LogComment":
		return &s.LogComment
	}

	return nil
}

// String returns a string representation of this struct or record.
func (s Pipeline) String() string {
	res := make([]string, 5)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "ProjectName: " + reform.Inspect(s.ProjectName, true)
	res[2] = "TagName: " + reform.Inspect(s.TagName, true)
	res[3] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[4] = "ApprovedAt: " + reform.Inspect(s.ApprovedAt, true)
	return strings.Join(res, ", ")
}
func (s PipelineLogRow) String() string {
	res := make([]string, 9)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "ProjectName: " + reform.Inspect(s.ProjectName, true)
	res[2] = "TagName: " + reform.Inspect(s.TagName, true)
	res[3] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[4] = "ApprovedAt: " + reform.Inspect(s.ApprovedAt, true)
	res[5] = "LogAuthor: " + reform.Inspect(s.LogAuthor, true)
	res[6] = "LogAction: " + reform.Inspect(s.LogAction, true)
	res[7] = "LogDate: " + reform.Inspect(s.LogDate, true)
	res[8] = "LogComment: " + reform.Inspect(s.LogComment, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Pipeline) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.ProjectName,
		s.TagName,
		s.CreatedAt,
		s.ApprovedAt,
	}
}
func (s *PipelineLogRow) Values() []interface{} {
	return append(s.Pipeline.Values(), []interface{}{
		s.LogAuthor,
		s.LogAction,
		s.LogDate,
		s.LogComment,
	}...)
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Pipeline) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.ProjectName,
		&s.TagName,
		&s.CreatedAt,
		&s.ApprovedAt,
	}
}
func (s *PipelineLogRow) Pointers() []interface{} {
	return append(s.Pipeline.Pointers(), []interface{}{
		&s.LogAuthor,
		&s.LogAction,
		&s.LogDate,
		&s.LogComment,
	}...)
}

// View returns View object for that struct.
func (s Pipeline) View() reform.View {
	return PipelineTable
}
func (s PipelineScope) View() reform.View {
	return s.item.View()
}
func (s PipelineLogRow) View() reform.View {
	return PipelineTableLogRow
}

// Generate a scope for object
func (s Pipeline) Scope() *PipelineScope {
	return &PipelineScope{item: &s, db: defaultDB_Pipeline}
}
func (s *Pipeline) PtrScope() *PipelineScope {
	return &PipelineScope{item: s, db: defaultDB_Pipeline}
}

// Sets DB to do queries
func (s Pipeline) DB(db reform.ReformDBTX) (scope *PipelineScope) { return s.Scope().DB(db) }
func (s *PipelineScope) DB(db reform.ReformDBTX) *PipelineScope {
	if db != nil {
		s.db = db
	}
	afterDBer, ok := interface{}(s).(reform.AfterDBer)
	if ok {
		afterDBer.AfterDB()
	}
	return s
}

// Gets DB
func (s Pipeline) GetDB() (db *reform.DB) { return s.Scope().GetDB() }
func (s PipelineScope) GetDB() *reform.DB {
	return s.db.(*reform.DB)
}

func (s Pipeline) StartTransaction() (*reform.TX, error) { return s.Scope().StartTransaction() }
func (s PipelineScope) StartTransaction() (*reform.TX, error) {
	return s.db.(*reform.DB).Begin()
}

// Sets default DB (to do not call the scope.DB() method every time)
func (s *Pipeline) SetDefaultDB(db *reform.DB) (err error) {
	defaultDB_Pipeline = db
	return nil
}

// Compiles SQL tail for defined limit scope
// TODO: should be compiled via dialects
func (s *PipelineScope) getLimitTail() (tail string, args []interface{}, err error) {
	if s.limit <= 0 {
		return
	}

	tail = fmt.Sprintf("%v", s.limit)
	return
}

// Compiles SQL tail for defined group scope
// TODO: should be compiled via dialects
func (s *PipelineScope) getGroupTail() (tail string, args []interface{}, err error) {
	tail = strings.Join(s.groupBy, ", ")

	return
}

// Compiles SQL tail for defined order scope
// TODO: should be compiled via dialects
func (s *PipelineScope) getOrderTail() (tail string, args []interface{}, err error) {
	var fieldName string
	var orderStringParts []string

	for idx, orderStr := range s.order {
		switch idx % 2 {
		case 0:
			fieldName = orderStr
		case 1:
			orderDirection := orderStr

			orderStringParts = append(orderStringParts, s.db.EscapeTableName(fieldName)+" "+orderDirection)
		}
	}

	tail = strings.Join(orderStringParts, ", ")

	return
}

// Compiles SQL tail for defined filter
// TODO: should be compiled via dialects
func (s *PipelineScope) getWhereTailForFilter(filter PipelineFilter) (tail string, whereTailArgs []interface{}, err error) {
	return s.db.GetWhereTailForFilter(Pipeline(filter), nil, "", false)
}

// parseQuerierArgs considers different ways of defning the tail (using scope properties or/and in_args)
func (s PipelineScope) parseWhereTailComponent(in_args []interface{}, placeholderCounter *int) (tail string, args []interface{}, err error) {
	if len(in_args) > 0 {
		switch arg := in_args[0].(type) {
		case int:
			tail = "id " + s.db.OperatorAndPlaceholderOfValueForSQL(in_args[0], *placeholderCounter)
			*placeholderCounter++
			args = s.db.ValueForSQL(in_args[0])
		case string:
			tailWords := s.db.SplitConditionByPlaceholders(arg)

			if len(tailWords)-1 != len(in_args[1:]) {
				panic(fmt.Errorf("The pattern doesn't fit for passed arguments (wrong number of question marks?): len(tailWords)-1 != len(in_args[1:]): <%v> <%v>", arg, in_args[1:]))
			}

			for idx, rawNewArgs := range in_args[1:] {
				newArgs := s.db.ValueForSQL(rawNewArgs)
				newTailWords := []string{}
				for range newArgs {
					newTailWords = append(newTailWords, s.db.GetDialect().Placeholder(*placeholderCounter))
					*placeholderCounter++
				}
				tail += tailWords[idx] + strings.Join(newTailWords, ",")
				args = append(args, newArgs...)
			}
			tail += tailWords[len(in_args[1:])]

			return
		case *Pipeline:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case *PipelineFilter:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case Pipeline:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(PipelineFilter(arg))
		case PipelineFilter:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(arg)
		default:
			err = fmt.Errorf("Invalid first element of \"in_args\" (%T). It should be a string or PipelineFilter.", arg)
			return
		}
	}

	return
}

// Compiles SQL tail for defined filter
// TODO: should be compiled via dialects
func (s *PipelineScope) getWhereTail() (tail string, whereTailArgs []interface{}, err error) {
	var whereTailStringParts []string

	if len(s.where) == 0 {
		return
	}

	placeholderCounter := 0

	for _, whereComponent := range s.where {
		var whereTailStringPart string
		var whereTailArgsPart []interface{}

		whereTailStringPart, whereTailArgsPart, err = s.parseWhereTailComponent(whereComponent, &placeholderCounter)
		if err != nil {
			return
		}

		if len(whereTailStringPart) > 0 {
			whereTailStringParts = append(whereTailStringParts, whereTailStringPart)
		}
		whereTailArgs = append(whereTailArgs, whereTailArgsPart...)
	}

	if len(whereTailStringParts) == 0 {
		return
	}

	tail = "(" + strings.Join(whereTailStringParts, ") AND (") + ")"

	return
}

func (s Pipeline) Where(requiredArg interface{}, args ...interface{}) (scope *PipelineScope) {
	return s.Scope().Where(requiredArg, args...)
}
func (s PipelineScope) Where(requiredArg interface{}, in_args ...interface{}) *PipelineScope {
	s.where = append(s.where, append([]interface{}{requiredArg}, in_args...))
	return &s
}
func (s PipelineScope) SetWhere(where [][]interface{}) *PipelineScope {
	s.where = where
	return &s
}
func (s PipelineScope) GetWhere() [][]interface{} {
	return s.where
}

// Sets all scope-related parameters to be equal as in passed scope (as an argument)
func (s PipelineScope) SetScope(anotherScope reform.Scope) *PipelineScope {
	s.where = anotherScope.GetWhere()
	s.order = anotherScope.GetOrder()
	s.groupBy = anotherScope.GetGroup()
	s.limit = anotherScope.GetLimit()
	s.db = anotherScope.GetDB()

	return &s
}
func (s PipelineScope) ISetScope(anotherScope reform.Scope) reform.Scope {
	return s.ISetScope(anotherScope)
}

// Compiles SQL tail for defined db/where/order/limit scope
// TODO: should be compiled via dialects
func (s *PipelineScope) getTail() (tail string, args []interface{}, err error) {
	whereTailString, whereTailArgs, err := s.getWhereTail()

	if err != nil {
		return
	}
	groupTailString, groupTailArgs, err := s.getGroupTail()
	if err != nil {
		return
	}
	orderTailString, orderTailArgs, err := s.getOrderTail()
	if err != nil {
		return
	}
	limitTailString, _, err := s.getLimitTail()
	if err != nil {
		return
	}

	args = append(whereTailArgs, append(groupTailArgs, orderTailArgs...)...)

	if len(whereTailString) > 0 {
		whereTailString = " WHERE " + whereTailString + " "
	}

	if len(groupTailString) > 0 {
		groupTailString = " GROUP BY " + groupTailString + " "
	}

	if len(orderTailString) > 0 {
		orderTailString = " ORDER BY " + orderTailString + " "
	}

	if len(limitTailString) > 0 {
		limitTailString = " LIMIT " + limitTailString + " "
	}

	tail = whereTailString + groupTailString + orderTailString + limitTailString

	if len(s.appendTail) > 0 {
		tail += " " + s.appendTail
	}

	return

}

// SelectRows is a simple wrapper to get raw "sql.Rows"
func (s Pipeline) SelectRows(query string, args ...interface{}) (rows *sql.Rows, err error) {
	return s.Scope().SelectRows(query, args...)
}
func (s *PipelineScope) SelectRows(query string, queryArgs ...interface{}) (rows *sql.Rows, err error) {
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	return s.db.Query("SELECT "+query+" FROM `"+PipelineTable.Name()+"` "+tail, append(queryArgs, args...)...)
}

func (s *PipelineScope) callStructMethod(str *Pipeline, methodName string) error {
	if method := reflect.ValueOf(str).MethodByName(methodName); method.IsValid() {
		switch f := method.Interface().(type) {
		case func():
			f()

		case func(reform.ReformDBTX):
			f(s.db)

		case func(*PipelineScope):
			f(s)

		case func(interface{}): // For compatibility with other ORMs
			f(s.db)

		case func() error:
			return f()

		case func(reform.ReformDBTX) error:
			return f(s.db)

		case func(*PipelineScope) error:
			return f(s)

		case func(interface{}) error: // For compatibility with other ORMS
			return f(s.db)

		default:
			panic("Unknown type of method: \"" + methodName + "\"")
		}
	}
	return nil
}

func (s PipelineScope) checkDb() {
	if s.db == nil {
		panic("s.db == nil")
	}
}

// Select is a handy wrapper for SelectRows() and NextRow(): it makes a query and collects the result into a slice
func (s Pipeline) Select(args ...interface{}) (result []Pipeline, err error) {
	return s.Scope().Select(args...)
}
func (s PipelineScope) Select(args ...interface{}) (result []Pipeline, err error) {
	s.checkDb()

	if len(args) > 0 {
		s = *s.Where(args[0], args[1:]...)
	}
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	rows, err := s.db.FlexSelectRows(PipelineTable, s.tableQuery, s.fieldsFilter, tail, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		item := Pipeline{}
		err = rows.Scan(item.FieldPointersByNames(s.fieldsFilter)...)
		if err != nil {
			return
		}

		s.callStructMethod(&item, "AfterFind")

		result = append(result, item)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}
func (s Pipeline) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().Select(args...)
}
func (s PipelineScope) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Select(args...)
}

// "First" a method to select and return only one record.
func (s Pipeline) First(args ...interface{}) (result Pipeline, err error) {
	return s.Scope().First(args...)
}
func (s PipelineScope) First(args ...interface{}) (result Pipeline, err error) {
	s.checkDb()

	if len(args) > 0 {
		s = *s.Where(args[0], args[1:]...)
	}
	tail, args, err := s.Limit(1).getTail()
	if err != nil {
		return
	}

	err = s.db.FlexSelectOneTo(&result, s.tableQuery, s.fieldsFilter, tail, args...)

	return
}
func (s Pipeline) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().First(args...)
}
func (s PipelineScope) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.First(args...)
}

// Sets "GROUP BY".
func (s Pipeline) Group(args ...interface{}) (scope *PipelineScope) { return s.Scope().Group(args...) }
func (s PipelineScope) Group(argsI ...interface{}) *PipelineScope {
	for _, argI := range argsI {
		s.groupBy = append(s.groupBy, argI.(string))
	}

	return &s
}
func (s PipelineScope) SetGroup(groupBy []string) *PipelineScope {
	s.groupBy = groupBy
	return &s
}
func (s PipelineScope) GetGroup() []string {
	return s.groupBy
}

// Sets a table query. For example SetTableQuery("table1 JOIN table2 USING(key)")
func (s Pipeline) SetTableQuery(query string) (scope *PipelineScope) {
	return s.Scope().SetTableQuery(query)
}
func (s PipelineScope) SetTableQuery(query string) *PipelineScope {
	if query == "" {
		s.tableQuery = nil
	} else {
		s.tableQuery = &query
	}
	return &s
}
func (s PipelineScope) GetTableQuery() string {
	if s.tableQuery != nil {
		return *s.tableQuery
	}
	return s.db.QualifiedView(s.View())
}

// Sets which structure fields should be queried while Select()/First(). For example SetFields("StructField1", "StructIdField", "StructCommentsField"). Could be used just to speed up a query.
// It's not recommended to use this function!
func (s Pipeline) SetQueryFieldsByNames(fields ...string) (scope *PipelineScope) {
	return s.Scope().SetQueryFieldsByNames(fields...)
}
func (s PipelineScope) SetQueryFieldsByNames(fields ...string) *PipelineScope {
	s.fieldsFilter = fields
	return &s
}
func (s PipelineScope) GetQueryFields() []string {
	return s.fieldsFilter
}

// Sets order. Arguments should be passed by pairs column-{ASC,DESC}. For example Order("id", "ASC", "value" "DESC")
func (s Pipeline) Order(args ...interface{}) (scope *PipelineScope) { return s.Scope().Order(args...) }
func (s PipelineScope) Order(argsI ...interface{}) *PipelineScope {
	switch len(argsI) {
	case 0:
	case 1:
		arg := argsI[0].(string)
		args0 := strings.Split(arg, ",")
		var args []string
		for _, arg0 := range args0 {
			args = append(args, strings.Split(arg0, ":")...)
		}
		s.order = args
	default:
		var args []string
		for _, argI := range argsI {
			args = append(args, argI.(string))
		}
		s.order = args
	}

	return &s
}
func (s PipelineScope) SetOrder(order []string) *PipelineScope {
	s.order = order
	return &s
}
func (s PipelineScope) GetOrder() []string {
	return s.order
}

func (s Pipeline) SetSQLAppend(appendTail string) (scope *PipelineScope) {
	return s.Scope().SetSQLAppend(appendTail)
}
func (s PipelineScope) SetSQLAppend(appendTail string) *PipelineScope {
	s.appendTail = appendTail
	return &s
}

// Sets limit.
func (s Pipeline) Limit(limit int) (scope *PipelineScope) { return s.Scope().Limit(limit) }
func (s *PipelineScope) Limit(limit int) *PipelineScope {
	s.limit = limit
	return s
}

// Gets limit
func (s PipelineScope) GetLimit() int {
	return s.limit
}

// "Reload" reloads record using Primary Key
func (s *PipelineFilter) Reload(db *reform.DB) error { return (*Pipeline)(s).Reload(db) }
func (s *Pipeline) Reload(db *reform.DB) (err error) {
	return db.FindByPrimaryKeyTo(s, s.PKValue())
}

// Create and Insert inserts new record to DB
func (s *Pipeline) Create() (err error) { return s.PtrScope().Create() }
func (s *PipelineScope) Create() (err error) {
	return s.Insert()
}
func (s *Pipeline) Insert() (err error) { return s.PtrScope().Insert() }
func (s *PipelineScope) Insert() (err error) {
	s.checkDb()
	err = s.db.Insert(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Replace "REPLACE INTO" new record to DB
func (s *Pipeline) Replace() (err error) { return s.PtrScope().Replace() }
func (s *PipelineScope) Replace() (err error) {
	s.checkDb()
	err = s.db.Replace(s.item)
	if err == nil {
		s.doLog("REPLACE")
	}
	return err
}

// Save inserts new record to DB is PK is zero and updates existing record if PK is not zero
func (s *Pipeline) Save() (err error) { return s.PtrScope().Save() }
func (s *PipelineScope) Save() (err error) {
	s.checkDb()
	err = s.db.Save(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Update updates existing record in DB
func (s Pipeline) Update() (err error) { return s.Scope().Update() }
func (s *PipelineScope) Update() (err error) {
	s.checkDb()
	err = s.db.Update(s.item)
	if err == nil {
		s.doLog("UPDATE")
	}
	return err
}

// Delete deletes existing record in DB
func (s Pipeline) Delete() (err error) { return s.Scope().Delete() }
func (s *PipelineScope) Delete() (err error) {
	s.checkDb()
	err = s.db.Delete(s.item)
	if err == nil {
		s.doLog("DELETE")
	}
	return err
}

func (s *PipelineScope) doLog(requestType string) {
	if !s.loggingEnabled {
		return
	}

	var logRow PipelineLogRow
	logRow.Pipeline = *s.item
	logRow.LogAuthor = s.loggingAuthor
	logRow.LogAction = requestType
	logRow.LogDate = time.Now()
	logRow.LogComment = s.loggingComment

	s.db.Insert(&logRow)
}

// Enables logging to table "pipelines_log". This table should has the same schema, except:
// - Unique/Primary keys should be removed
// - Should be added next fields: "log_author" (nullable string), "log_date" (timestamp), "log_action" (enum("INSERT", "UPDATE", "DELETE")), "log_comment" (string)
func (s *Pipeline) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *PipelineScope) {
	return s.Scope().Log(enableLogging, author, commentFormat, commentArgs...)
}
func (s *PipelineScope) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *PipelineScope) {
	s.loggingEnabled = enableLogging
	s.loggingAuthor = author
	s.loggingComment = fmt.Sprintf(commentFormat, commentArgs...)

	return s
}

// Table returns Table object for that record.
func (s Pipeline) Table() reform.Table {
	return PipelineTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s Pipeline) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s Pipeline) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s Pipeline) HasPK() bool {
	return s.Id != PipelineTable.z[PipelineTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *PipelineFilter) SetPK(pk interface{}) { (*Pipeline)(s).SetPK(pk) }
func (s *Pipeline) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.Id = int(i64)
	} else {
		s.Id = pk.(int)
	}
}

var (
	// check interfaces
	_ reform.View   = PipelineTable
	_ reform.Struct = (*Pipeline)(nil)
	_ reform.Table  = PipelineTable
	_ reform.Record = (*Pipeline)(nil)
	_ fmt.Stringer  = (*Pipeline)(nil)

	// querier
	PipelineSQL        = Pipeline{} // Should be read only
	defaultDB_Pipeline *reform.DB
)

func init() {
	//parse.AssertUpToDate(&PipelineTable.s, new(Pipeline)) // Temporary disabled (doesn't work with arbitary types like "type sliceString []string")
}

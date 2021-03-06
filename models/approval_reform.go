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

type ApprovalScope struct {
	item *Approval

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
type ApprovalFilter Approval

type ApprovalLogRow struct {
	Approval
	LogAuthor  *string
	LogAction  string
	LogDate    time.Time
	LogComment string
}

// Schema returns a schema name in SQL database ("").
type approvalTableTypeType struct {
	s reform.StructInfo
	z []interface{}
}

func (v approvalTableTypeType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("approvals").
func (v approvalTableTypeType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v approvalTableTypeType) Columns() []string {
	return []string{"id", "username", "pipeline_id", "created_at"}
}

// NewStruct makes a new struct for that view or table.
func (v approvalTableTypeType) NewStruct() reform.Struct {
	return new(Approval)
}

// NewRecord makes a new record for that table.
func (v *approvalTableTypeType) NewRecord() reform.Record {
	return new(Approval)
}

func (v *approvalTableTypeType) NewScope() *ApprovalScope {
	return &ApprovalScope{item: &Approval{}}
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *approvalTableTypeType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

func (v approvalTableTypeType) CreateTableIfNotExists(db *reform.DB) (bool, error) {
	if db == nil {
		db = defaultDB_Approval
	}
	return db.CreateTableIfNotExists(v.s)
}

func (v approvalTableTypeType) StructInfo() reform.StructInfo {
	return v.s
}

// ApprovalTable represents approvals view or table in SQL database.
var ApprovalTable = &approvalTableTypeType{
	s: reform.StructInfo{Type: "Approval", SQLSchema: "", SQLName: "approvals", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Username", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "username", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "PipelineId", IsPK: false, IsUnique: false, HasIndex: false, Type: "int", Column: "pipeline_id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "CreatedAt", IsPK: false, IsUnique: false, HasIndex: false, Type: "*extime.Time", Column: "created_at", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(Approval).Values(),
}

type approvalTableTypeType_log struct {
	s reform.StructInfo
	z []interface{}
}

func (v *approvalTableTypeType_log) Schema() string {
	return v.s.SQLSchema
}

func (v *approvalTableTypeType_log) Name() string {
	return v.s.SQLName
}

func (v *approvalTableTypeType_log) Columns() []string {
	return []string{"id", "username", "pipeline_id", "created_at", "log_author", "log_action", "log_date", "log_comment"}
}

func (v *approvalTableTypeType_log) NewStruct() reform.Struct {
	return new(Approval)
}

func (v *approvalTableTypeType_log) NewRecord() reform.Record {
	return new(Approval)
}

func (v *approvalTableTypeType_log) NewScope() *ApprovalScope {
	return &ApprovalScope{item: &Approval{}}
}

func (v *approvalTableTypeType_log) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

var ApprovalTableLogRow = &approvalTableTypeType_log{
	s: reform.StructInfo{Type: "Approval", SQLSchema: "", SQLName: "approvals_log", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Username", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "username", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "PipelineId", IsPK: false, IsUnique: false, HasIndex: false, Type: "int", Column: "pipeline_id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "CreatedAt", IsPK: false, IsUnique: false, HasIndex: false, Type: "*extime.Time", Column: "created_at", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAuthor", IsPK: false, IsUnique: false, HasIndex: false, Type: "*string", Column: "log_author", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAction", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_action", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogDate", IsPK: false, IsUnique: false, HasIndex: false, Type: "time.Time", Column: "log_date", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogComment", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_comment", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(ApprovalLogRow).Values(),
}

func (s approvalTableTypeType) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	case "Username":
		return "username"
	case "PipelineId":
		return "pipeline_id"
	case "CreatedAt":
		return "created_at"
	}
	return ""
}

func (s approvalTableTypeType_log) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	case "Username":
		return "username"
	case "PipelineId":
		return "pipeline_id"
	case "CreatedAt":
		return "created_at"
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

func (s *Approval) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
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

func (s *ApprovalLogRow) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
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

func (s *Approval) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	case "Username":
		return &s.Username
	case "PipelineId":
		return &s.PipelineId
	case "CreatedAt":
		return &s.CreatedAt
	}

	return nil
}

func (s *ApprovalLogRow) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	case "Username":
		return &s.Username
	case "PipelineId":
		return &s.PipelineId
	case "CreatedAt":
		return &s.CreatedAt
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
func (s Approval) String() string {
	res := make([]string, 4)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "Username: " + reform.Inspect(s.Username, true)
	res[2] = "PipelineId: " + reform.Inspect(s.PipelineId, true)
	res[3] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	return strings.Join(res, ", ")
}
func (s ApprovalLogRow) String() string {
	res := make([]string, 8)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "Username: " + reform.Inspect(s.Username, true)
	res[2] = "PipelineId: " + reform.Inspect(s.PipelineId, true)
	res[3] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[4] = "LogAuthor: " + reform.Inspect(s.LogAuthor, true)
	res[5] = "LogAction: " + reform.Inspect(s.LogAction, true)
	res[6] = "LogDate: " + reform.Inspect(s.LogDate, true)
	res[7] = "LogComment: " + reform.Inspect(s.LogComment, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Approval) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.Username,
		s.PipelineId,
		s.CreatedAt,
	}
}
func (s *ApprovalLogRow) Values() []interface{} {
	return append(s.Approval.Values(), []interface{}{
		s.LogAuthor,
		s.LogAction,
		s.LogDate,
		s.LogComment,
	}...)
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Approval) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.Username,
		&s.PipelineId,
		&s.CreatedAt,
	}
}
func (s *ApprovalLogRow) Pointers() []interface{} {
	return append(s.Approval.Pointers(), []interface{}{
		&s.LogAuthor,
		&s.LogAction,
		&s.LogDate,
		&s.LogComment,
	}...)
}

// View returns View object for that struct.
func (s Approval) View() reform.View {
	return ApprovalTable
}
func (s ApprovalScope) View() reform.View {
	return s.item.View()
}
func (s ApprovalLogRow) View() reform.View {
	return ApprovalTableLogRow
}

// Generate a scope for object
func (s Approval) Scope() *ApprovalScope {
	return &ApprovalScope{item: &s, db: defaultDB_Approval}
}
func (s *Approval) PtrScope() *ApprovalScope {
	return &ApprovalScope{item: s, db: defaultDB_Approval}
}

// Sets DB to do queries
func (s Approval) DB(db reform.ReformDBTX) (scope *ApprovalScope) { return s.Scope().DB(db) }
func (s *ApprovalScope) DB(db reform.ReformDBTX) *ApprovalScope {
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
func (s Approval) GetDB() (db *reform.DB) { return s.Scope().GetDB() }
func (s ApprovalScope) GetDB() *reform.DB {
	return s.db.(*reform.DB)
}

func (s Approval) StartTransaction() (*reform.TX, error) { return s.Scope().StartTransaction() }
func (s ApprovalScope) StartTransaction() (*reform.TX, error) {
	return s.db.(*reform.DB).Begin()
}

// Sets default DB (to do not call the scope.DB() method every time)
func (s *Approval) SetDefaultDB(db *reform.DB) (err error) {
	defaultDB_Approval = db
	return nil
}

// Compiles SQL tail for defined limit scope
// TODO: should be compiled via dialects
func (s *ApprovalScope) getLimitTail() (tail string, args []interface{}, err error) {
	if s.limit <= 0 {
		return
	}

	tail = fmt.Sprintf("%v", s.limit)
	return
}

// Compiles SQL tail for defined group scope
// TODO: should be compiled via dialects
func (s *ApprovalScope) getGroupTail() (tail string, args []interface{}, err error) {
	tail = strings.Join(s.groupBy, ", ")

	return
}

// Compiles SQL tail for defined order scope
// TODO: should be compiled via dialects
func (s *ApprovalScope) getOrderTail() (tail string, args []interface{}, err error) {
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
func (s *ApprovalScope) getWhereTailForFilter(filter ApprovalFilter) (tail string, whereTailArgs []interface{}, err error) {
	return s.db.GetWhereTailForFilter(Approval(filter), nil, "", false)
}

// parseQuerierArgs considers different ways of defning the tail (using scope properties or/and in_args)
func (s ApprovalScope) parseWhereTailComponent(in_args []interface{}, placeholderCounter *int) (tail string, args []interface{}, err error) {
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
		case *Approval:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case *ApprovalFilter:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case Approval:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(ApprovalFilter(arg))
		case ApprovalFilter:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(arg)
		default:
			err = fmt.Errorf("Invalid first element of \"in_args\" (%T). It should be a string or ApprovalFilter.", arg)
			return
		}
	}

	return
}

// Compiles SQL tail for defined filter
// TODO: should be compiled via dialects
func (s *ApprovalScope) getWhereTail() (tail string, whereTailArgs []interface{}, err error) {
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

func (s Approval) Where(requiredArg interface{}, args ...interface{}) (scope *ApprovalScope) {
	return s.Scope().Where(requiredArg, args...)
}
func (s ApprovalScope) Where(requiredArg interface{}, in_args ...interface{}) *ApprovalScope {
	s.where = append(s.where, append([]interface{}{requiredArg}, in_args...))
	return &s
}
func (s ApprovalScope) SetWhere(where [][]interface{}) *ApprovalScope {
	s.where = where
	return &s
}
func (s ApprovalScope) GetWhere() [][]interface{} {
	return s.where
}

// Sets all scope-related parameters to be equal as in passed scope (as an argument)
func (s ApprovalScope) SetScope(anotherScope reform.Scope) *ApprovalScope {
	s.where = anotherScope.GetWhere()
	s.order = anotherScope.GetOrder()
	s.groupBy = anotherScope.GetGroup()
	s.limit = anotherScope.GetLimit()
	s.db = anotherScope.GetDB()

	return &s
}
func (s ApprovalScope) ISetScope(anotherScope reform.Scope) reform.Scope {
	return s.ISetScope(anotherScope)
}

// Compiles SQL tail for defined db/where/order/limit scope
// TODO: should be compiled via dialects
func (s *ApprovalScope) getTail() (tail string, args []interface{}, err error) {
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
func (s Approval) SelectRows(query string, args ...interface{}) (rows *sql.Rows, err error) {
	return s.Scope().SelectRows(query, args...)
}
func (s *ApprovalScope) SelectRows(query string, queryArgs ...interface{}) (rows *sql.Rows, err error) {
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	return s.db.Query("SELECT "+query+" FROM `"+ApprovalTable.Name()+"` "+tail, append(queryArgs, args...)...)
}

func (s *ApprovalScope) callStructMethod(str *Approval, methodName string) error {
	if method := reflect.ValueOf(str).MethodByName(methodName); method.IsValid() {
		switch f := method.Interface().(type) {
		case func():
			f()

		case func(reform.ReformDBTX):
			f(s.db)

		case func(*ApprovalScope):
			f(s)

		case func(interface{}): // For compatibility with other ORMs
			f(s.db)

		case func() error:
			return f()

		case func(reform.ReformDBTX) error:
			return f(s.db)

		case func(*ApprovalScope) error:
			return f(s)

		case func(interface{}) error: // For compatibility with other ORMS
			return f(s.db)

		default:
			panic("Unknown type of method: \"" + methodName + "\"")
		}
	}
	return nil
}

func (s ApprovalScope) checkDb() {
	if s.db == nil {
		panic("s.db == nil")
	}
}

// Select is a handy wrapper for SelectRows() and NextRow(): it makes a query and collects the result into a slice
func (s Approval) Select(args ...interface{}) (result []Approval, err error) {
	return s.Scope().Select(args...)
}
func (s ApprovalScope) Select(args ...interface{}) (result []Approval, err error) {
	s.checkDb()

	if len(args) > 0 {
		s = *s.Where(args[0], args[1:]...)
	}
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	rows, err := s.db.FlexSelectRows(ApprovalTable, s.tableQuery, s.fieldsFilter, tail, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		item := Approval{}
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
func (s Approval) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().Select(args...)
}
func (s ApprovalScope) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Select(args...)
}

// "First" a method to select and return only one record.
func (s Approval) First(args ...interface{}) (result Approval, err error) {
	return s.Scope().First(args...)
}
func (s ApprovalScope) First(args ...interface{}) (result Approval, err error) {
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
func (s Approval) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().First(args...)
}
func (s ApprovalScope) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.First(args...)
}

// Sets "GROUP BY".
func (s Approval) Group(args ...interface{}) (scope *ApprovalScope) { return s.Scope().Group(args...) }
func (s ApprovalScope) Group(argsI ...interface{}) *ApprovalScope {
	for _, argI := range argsI {
		s.groupBy = append(s.groupBy, argI.(string))
	}

	return &s
}
func (s ApprovalScope) SetGroup(groupBy []string) *ApprovalScope {
	s.groupBy = groupBy
	return &s
}
func (s ApprovalScope) GetGroup() []string {
	return s.groupBy
}

// Sets a table query. For example SetTableQuery("table1 JOIN table2 USING(key)")
func (s Approval) SetTableQuery(query string) (scope *ApprovalScope) {
	return s.Scope().SetTableQuery(query)
}
func (s ApprovalScope) SetTableQuery(query string) *ApprovalScope {
	if query == "" {
		s.tableQuery = nil
	} else {
		s.tableQuery = &query
	}
	return &s
}
func (s ApprovalScope) GetTableQuery() string {
	if s.tableQuery != nil {
		return *s.tableQuery
	}
	return s.db.QualifiedView(s.View())
}

// Sets which structure fields should be queried while Select()/First(). For example SetFields("StructField1", "StructIdField", "StructCommentsField"). Could be used just to speed up a query.
// It's not recommended to use this function!
func (s Approval) SetQueryFieldsByNames(fields ...string) (scope *ApprovalScope) {
	return s.Scope().SetQueryFieldsByNames(fields...)
}
func (s ApprovalScope) SetQueryFieldsByNames(fields ...string) *ApprovalScope {
	s.fieldsFilter = fields
	return &s
}
func (s ApprovalScope) GetQueryFields() []string {
	return s.fieldsFilter
}

// Sets order. Arguments should be passed by pairs column-{ASC,DESC}. For example Order("id", "ASC", "value" "DESC")
func (s Approval) Order(args ...interface{}) (scope *ApprovalScope) { return s.Scope().Order(args...) }
func (s ApprovalScope) Order(argsI ...interface{}) *ApprovalScope {
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
func (s ApprovalScope) SetOrder(order []string) *ApprovalScope {
	s.order = order
	return &s
}
func (s ApprovalScope) GetOrder() []string {
	return s.order
}

func (s Approval) SetSQLAppend(appendTail string) (scope *ApprovalScope) {
	return s.Scope().SetSQLAppend(appendTail)
}
func (s ApprovalScope) SetSQLAppend(appendTail string) *ApprovalScope {
	s.appendTail = appendTail
	return &s
}

// Sets limit.
func (s Approval) Limit(limit int) (scope *ApprovalScope) { return s.Scope().Limit(limit) }
func (s *ApprovalScope) Limit(limit int) *ApprovalScope {
	s.limit = limit
	return s
}

// Gets limit
func (s ApprovalScope) GetLimit() int {
	return s.limit
}

// "Reload" reloads record using Primary Key
func (s *ApprovalFilter) Reload(db *reform.DB) error { return (*Approval)(s).Reload(db) }
func (s *Approval) Reload(db *reform.DB) (err error) {
	return db.FindByPrimaryKeyTo(s, s.PKValue())
}

// Create and Insert inserts new record to DB
func (s *Approval) Create() (err error) { return s.PtrScope().Create() }
func (s *ApprovalScope) Create() (err error) {
	return s.Insert()
}
func (s *Approval) Insert() (err error) { return s.PtrScope().Insert() }
func (s *ApprovalScope) Insert() (err error) {
	s.checkDb()
	err = s.db.Insert(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Replace "REPLACE INTO" new record to DB
func (s *Approval) Replace() (err error) { return s.PtrScope().Replace() }
func (s *ApprovalScope) Replace() (err error) {
	s.checkDb()
	err = s.db.Replace(s.item)
	if err == nil {
		s.doLog("REPLACE")
	}
	return err
}

// Save inserts new record to DB is PK is zero and updates existing record if PK is not zero
func (s *Approval) Save() (err error) { return s.PtrScope().Save() }
func (s *ApprovalScope) Save() (err error) {
	s.checkDb()
	err = s.db.Save(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Update updates existing record in DB
func (s Approval) Update() (err error) { return s.Scope().Update() }
func (s *ApprovalScope) Update() (err error) {
	s.checkDb()
	err = s.db.Update(s.item)
	if err == nil {
		s.doLog("UPDATE")
	}
	return err
}

// Delete deletes existing record in DB
func (s Approval) Delete() (err error) { return s.Scope().Delete() }
func (s *ApprovalScope) Delete() (err error) {
	s.checkDb()
	err = s.db.Delete(s.item)
	if err == nil {
		s.doLog("DELETE")
	}
	return err
}

func (s *ApprovalScope) doLog(requestType string) {
	if !s.loggingEnabled {
		return
	}

	var logRow ApprovalLogRow
	logRow.Approval = *s.item
	logRow.LogAuthor = s.loggingAuthor
	logRow.LogAction = requestType
	logRow.LogDate = time.Now()
	logRow.LogComment = s.loggingComment

	s.db.Insert(&logRow)
}

// Enables logging to table "approvals_log". This table should has the same schema, except:
// - Unique/Primary keys should be removed
// - Should be added next fields: "log_author" (nullable string), "log_date" (timestamp), "log_action" (enum("INSERT", "UPDATE", "DELETE")), "log_comment" (string)
func (s *Approval) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *ApprovalScope) {
	return s.Scope().Log(enableLogging, author, commentFormat, commentArgs...)
}
func (s *ApprovalScope) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *ApprovalScope) {
	s.loggingEnabled = enableLogging
	s.loggingAuthor = author
	s.loggingComment = fmt.Sprintf(commentFormat, commentArgs...)

	return s
}

// Table returns Table object for that record.
func (s Approval) Table() reform.Table {
	return ApprovalTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s Approval) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s Approval) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s Approval) HasPK() bool {
	return s.Id != ApprovalTable.z[ApprovalTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *ApprovalFilter) SetPK(pk interface{}) { (*Approval)(s).SetPK(pk) }
func (s *Approval) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.Id = int(i64)
	} else {
		s.Id = pk.(int)
	}
}

var (
	// check interfaces
	_ reform.View   = ApprovalTable
	_ reform.Struct = (*Approval)(nil)
	_ reform.Table  = ApprovalTable
	_ reform.Record = (*Approval)(nil)
	_ fmt.Stringer  = (*Approval)(nil)

	// querier
	ApprovalSQL        = Approval{} // Should be read only
	defaultDB_Approval *reform.DB
)

func init() {
	//parse.AssertUpToDate(&ApprovalTable.s, new(Approval)) // Temporary disabled (doesn't work with arbitary types like "type sliceString []string")
}

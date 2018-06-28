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

type UserProfileScope struct {
	item *UserProfile

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
type UserProfileFilter UserProfile

type UserProfileLogRow struct {
	UserProfile
	LogAuthor  *string
	LogAction  string
	LogDate    time.Time
	LogComment string
}

// Schema returns a schema name in SQL database ("").
type userProfileTableTypeType struct {
	s reform.StructInfo
	z []interface{}
}

func (v userProfileTableTypeType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("user_profiles").
func (v userProfileTableTypeType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v userProfileTableTypeType) Columns() []string {
	return []string{"id", "username", "sms_to"}
}

// NewStruct makes a new struct for that view or table.
func (v userProfileTableTypeType) NewStruct() reform.Struct {
	return new(UserProfile)
}

// NewRecord makes a new record for that table.
func (v *userProfileTableTypeType) NewRecord() reform.Record {
	return new(UserProfile)
}

func (v *userProfileTableTypeType) NewScope() *UserProfileScope {
	return &UserProfileScope{item: &UserProfile{}}
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *userProfileTableTypeType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

func (v userProfileTableTypeType) CreateTableIfNotExists(db *reform.DB) (bool, error) {
	if db == nil {
		db = defaultDB_UserProfile
	}
	return db.CreateTableIfNotExists(v.s)
}

func (v userProfileTableTypeType) StructInfo() reform.StructInfo {
	return v.s
}

// UserProfileTable represents user_profiles view or table in SQL database.
var UserProfileTable = &userProfileTableTypeType{
	s: reform.StructInfo{Type: "UserProfile", SQLSchema: "", SQLName: "user_profiles", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Username", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "username", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "SMSTo", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "sms_to", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(UserProfile).Values(),
}

type userProfileTableTypeType_log struct {
	s reform.StructInfo
	z []interface{}
}

func (v *userProfileTableTypeType_log) Schema() string {
	return v.s.SQLSchema
}

func (v *userProfileTableTypeType_log) Name() string {
	return v.s.SQLName
}

func (v *userProfileTableTypeType_log) Columns() []string {
	return []string{"id", "username", "sms_to", "log_author", "log_action", "log_date", "log_comment"}
}

func (v *userProfileTableTypeType_log) NewStruct() reform.Struct {
	return new(UserProfile)
}

func (v *userProfileTableTypeType_log) NewRecord() reform.Record {
	return new(UserProfile)
}

func (v *userProfileTableTypeType_log) NewScope() *UserProfileScope {
	return &UserProfileScope{item: &UserProfile{}}
}

func (v *userProfileTableTypeType_log) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

var UserProfileTableLogRow = &userProfileTableTypeType_log{
	s: reform.StructInfo{Type: "UserProfile", SQLSchema: "", SQLName: "user_profiles_log", Fields: []reform.FieldInfo{{Name: "Id", IsPK: true, IsUnique: false, HasIndex: false, Type: "int", Column: "id", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "Username", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "username", FieldsPath: []reform.FieldInfo{}, SQLSize: 255, Embedded: "", StructFile: ""}, {Name: "SMSTo", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "sms_to", FieldsPath: []reform.FieldInfo{}, SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAuthor", IsPK: false, IsUnique: false, HasIndex: false, Type: "*string", Column: "log_author", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogAction", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_action", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogDate", IsPK: false, IsUnique: false, HasIndex: false, Type: "time.Time", Column: "log_date", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}, {Name: "LogComment", IsPK: false, IsUnique: false, HasIndex: false, Type: "string", Column: "log_comment", FieldsPath: []reform.FieldInfo(nil), SQLSize: 0, Embedded: "", StructFile: ""}}, PKFieldIndex: 0, ImitateGorm: false, SkipMethodOrder: false},
	z: new(UserProfileLogRow).Values(),
}

func (s userProfileTableTypeType) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	case "Username":
		return "username"
	case "SMSTo":
		return "sms_to"
	}
	return ""
}

func (s userProfileTableTypeType_log) ColumnNameByFieldName(fieldName string) string {
	switch fieldName {
	case "Id":
		return "id"
	case "Username":
		return "username"
	case "SMSTo":
		return "sms_to"
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

func (s *UserProfile) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
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

func (s *UserProfileLogRow) FieldPointersByNames(fieldNames []string) (fieldPointers []interface{}) {
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

func (s *UserProfile) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	case "Username":
		return &s.Username
	case "SMSTo":
		return &s.SMSTo
	}

	return nil
}

func (s *UserProfileLogRow) FieldPointerByName(fieldName string) interface{} {
	switch fieldName {
	case "Id":
		return &s.Id
	case "Username":
		return &s.Username
	case "SMSTo":
		return &s.SMSTo
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
func (s UserProfile) String() string {
	res := make([]string, 3)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "Username: " + reform.Inspect(s.Username, true)
	res[2] = "SMSTo: " + reform.Inspect(s.SMSTo, true)
	return strings.Join(res, ", ")
}
func (s UserProfileLogRow) String() string {
	res := make([]string, 7)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "Username: " + reform.Inspect(s.Username, true)
	res[2] = "SMSTo: " + reform.Inspect(s.SMSTo, true)
	res[3] = "LogAuthor: " + reform.Inspect(s.LogAuthor, true)
	res[4] = "LogAction: " + reform.Inspect(s.LogAction, true)
	res[5] = "LogDate: " + reform.Inspect(s.LogDate, true)
	res[6] = "LogComment: " + reform.Inspect(s.LogComment, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *UserProfile) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.Username,
		s.SMSTo,
	}
}
func (s *UserProfileLogRow) Values() []interface{} {
	return append(s.UserProfile.Values(), []interface{}{
		s.LogAuthor,
		s.LogAction,
		s.LogDate,
		s.LogComment,
	}...)
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *UserProfile) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.Username,
		&s.SMSTo,
	}
}
func (s *UserProfileLogRow) Pointers() []interface{} {
	return append(s.UserProfile.Pointers(), []interface{}{
		&s.LogAuthor,
		&s.LogAction,
		&s.LogDate,
		&s.LogComment,
	}...)
}

// View returns View object for that struct.
func (s UserProfile) View() reform.View {
	return UserProfileTable
}
func (s UserProfileScope) View() reform.View {
	return s.item.View()
}
func (s UserProfileLogRow) View() reform.View {
	return UserProfileTableLogRow
}

// Generate a scope for object
func (s UserProfile) Scope() *UserProfileScope {
	return &UserProfileScope{item: &s, db: defaultDB_UserProfile}
}
func (s *UserProfile) PtrScope() *UserProfileScope {
	return &UserProfileScope{item: s, db: defaultDB_UserProfile}
}

// Sets DB to do queries
func (s UserProfile) DB(db reform.ReformDBTX) (scope *UserProfileScope) { return s.Scope().DB(db) }
func (s *UserProfileScope) DB(db reform.ReformDBTX) *UserProfileScope {
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
func (s UserProfile) GetDB() (db *reform.DB) { return s.Scope().GetDB() }
func (s UserProfileScope) GetDB() *reform.DB {
	return s.db.(*reform.DB)
}

func (s UserProfile) StartTransaction() (*reform.TX, error) { return s.Scope().StartTransaction() }
func (s UserProfileScope) StartTransaction() (*reform.TX, error) {
	return s.db.(*reform.DB).Begin()
}

// Sets default DB (to do not call the scope.DB() method every time)
func (s *UserProfile) SetDefaultDB(db *reform.DB) (err error) {
	defaultDB_UserProfile = db
	return nil
}

// Compiles SQL tail for defined limit scope
// TODO: should be compiled via dialects
func (s *UserProfileScope) getLimitTail() (tail string, args []interface{}, err error) {
	if s.limit <= 0 {
		return
	}

	tail = fmt.Sprintf("%v", s.limit)
	return
}

// Compiles SQL tail for defined group scope
// TODO: should be compiled via dialects
func (s *UserProfileScope) getGroupTail() (tail string, args []interface{}, err error) {
	tail = strings.Join(s.groupBy, ", ")

	return
}

// Compiles SQL tail for defined order scope
// TODO: should be compiled via dialects
func (s *UserProfileScope) getOrderTail() (tail string, args []interface{}, err error) {
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
func (s *UserProfileScope) getWhereTailForFilter(filter UserProfileFilter) (tail string, whereTailArgs []interface{}, err error) {
	return s.db.GetWhereTailForFilter(UserProfile(filter), nil, "", false)
}

// parseQuerierArgs considers different ways of defning the tail (using scope properties or/and in_args)
func (s UserProfileScope) parseWhereTailComponent(in_args []interface{}, placeholderCounter *int) (tail string, args []interface{}, err error) {
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
		case *UserProfile:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case *UserProfileFilter:
			in_args[0] = *arg
			return s.parseWhereTailComponent(in_args, placeholderCounter)
		case UserProfile:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(UserProfileFilter(arg))
		case UserProfileFilter:
			if len(in_args) > 1 {
				s = *s.Where(in_args[1], in_args[2:]...)
			}
			tail, args, err = s.getWhereTailForFilter(arg)
		default:
			err = fmt.Errorf("Invalid first element of \"in_args\" (%T). It should be a string or UserProfileFilter.", arg)
			return
		}
	}

	return
}

// Compiles SQL tail for defined filter
// TODO: should be compiled via dialects
func (s *UserProfileScope) getWhereTail() (tail string, whereTailArgs []interface{}, err error) {
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

func (s UserProfile) Where(requiredArg interface{}, args ...interface{}) (scope *UserProfileScope) {
	return s.Scope().Where(requiredArg, args...)
}
func (s UserProfileScope) Where(requiredArg interface{}, in_args ...interface{}) *UserProfileScope {
	s.where = append(s.where, append([]interface{}{requiredArg}, in_args...))
	return &s
}
func (s UserProfileScope) SetWhere(where [][]interface{}) *UserProfileScope {
	s.where = where
	return &s
}
func (s UserProfileScope) GetWhere() [][]interface{} {
	return s.where
}

// Sets all scope-related parameters to be equal as in passed scope (as an argument)
func (s UserProfileScope) SetScope(anotherScope reform.Scope) *UserProfileScope {
	s.where = anotherScope.GetWhere()
	s.order = anotherScope.GetOrder()
	s.groupBy = anotherScope.GetGroup()
	s.limit = anotherScope.GetLimit()
	s.db = anotherScope.GetDB()

	return &s
}
func (s UserProfileScope) ISetScope(anotherScope reform.Scope) reform.Scope {
	return s.ISetScope(anotherScope)
}

// Compiles SQL tail for defined db/where/order/limit scope
// TODO: should be compiled via dialects
func (s *UserProfileScope) getTail() (tail string, args []interface{}, err error) {
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
func (s UserProfile) SelectRows(query string, args ...interface{}) (rows *sql.Rows, err error) {
	return s.Scope().SelectRows(query, args...)
}
func (s *UserProfileScope) SelectRows(query string, queryArgs ...interface{}) (rows *sql.Rows, err error) {
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	return s.db.Query("SELECT "+query+" FROM `"+UserProfileTable.Name()+"` "+tail, append(queryArgs, args...)...)
}

func (s *UserProfileScope) callStructMethod(str *UserProfile, methodName string) error {
	if method := reflect.ValueOf(str).MethodByName(methodName); method.IsValid() {
		switch f := method.Interface().(type) {
		case func():
			f()

		case func(reform.ReformDBTX):
			f(s.db)

		case func(*UserProfileScope):
			f(s)

		case func(interface{}): // For compatibility with other ORMs
			f(s.db)

		case func() error:
			return f()

		case func(reform.ReformDBTX) error:
			return f(s.db)

		case func(*UserProfileScope) error:
			return f(s)

		case func(interface{}) error: // For compatibility with other ORMS
			return f(s.db)

		default:
			panic("Unknown type of method: \"" + methodName + "\"")
		}
	}
	return nil
}

func (s UserProfileScope) checkDb() {
	if s.db == nil {
		panic("s.db == nil")
	}
}

// Select is a handy wrapper for SelectRows() and NextRow(): it makes a query and collects the result into a slice
func (s UserProfile) Select(args ...interface{}) (result []UserProfile, err error) {
	return s.Scope().Select(args...)
}
func (s UserProfileScope) Select(args ...interface{}) (result []UserProfile, err error) {
	s.checkDb()

	if len(args) > 0 {
		s = *s.Where(args[0], args[1:]...)
	}
	tail, args, err := s.getTail()
	if err != nil {
		return
	}

	rows, err := s.db.FlexSelectRows(UserProfileTable, s.tableQuery, s.fieldsFilter, tail, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		item := UserProfile{}
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
func (s UserProfile) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().Select(args...)
}
func (s UserProfileScope) SelectI(args ...interface{}) (result interface{}, err error) {
	return s.Select(args...)
}

// "First" a method to select and return only one record.
func (s UserProfile) First(args ...interface{}) (result UserProfile, err error) {
	return s.Scope().First(args...)
}
func (s UserProfileScope) First(args ...interface{}) (result UserProfile, err error) {
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
func (s UserProfile) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.Scope().First(args...)
}
func (s UserProfileScope) FirstI(args ...interface{}) (result interface{}, err error) {
	return s.First(args...)
}

// Sets "GROUP BY".
func (s UserProfile) Group(args ...interface{}) (scope *UserProfileScope) {
	return s.Scope().Group(args...)
}
func (s UserProfileScope) Group(argsI ...interface{}) *UserProfileScope {
	for _, argI := range argsI {
		s.groupBy = append(s.groupBy, argI.(string))
	}

	return &s
}
func (s UserProfileScope) SetGroup(groupBy []string) *UserProfileScope {
	s.groupBy = groupBy
	return &s
}
func (s UserProfileScope) GetGroup() []string {
	return s.groupBy
}

// Sets a table query. For example SetTableQuery("table1 JOIN table2 USING(key)")
func (s UserProfile) SetTableQuery(query string) (scope *UserProfileScope) {
	return s.Scope().SetTableQuery(query)
}
func (s UserProfileScope) SetTableQuery(query string) *UserProfileScope {
	if query == "" {
		s.tableQuery = nil
	} else {
		s.tableQuery = &query
	}
	return &s
}
func (s UserProfileScope) GetTableQuery() string {
	if s.tableQuery != nil {
		return *s.tableQuery
	}
	return s.db.QualifiedView(s.View())
}

// Sets which structure fields should be queried while Select()/First(). For example SetFields("StructField1", "StructIdField", "StructCommentsField"). Could be used just to speed up a query.
// It's not recommended to use this function!
func (s UserProfile) SetQueryFieldsByNames(fields ...string) (scope *UserProfileScope) {
	return s.Scope().SetQueryFieldsByNames(fields...)
}
func (s UserProfileScope) SetQueryFieldsByNames(fields ...string) *UserProfileScope {
	s.fieldsFilter = fields
	return &s
}
func (s UserProfileScope) GetQueryFields() []string {
	return s.fieldsFilter
}

// Sets order. Arguments should be passed by pairs column-{ASC,DESC}. For example Order("id", "ASC", "value" "DESC")
func (s UserProfile) Order(args ...interface{}) (scope *UserProfileScope) {
	return s.Scope().Order(args...)
}
func (s UserProfileScope) Order(argsI ...interface{}) *UserProfileScope {
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
func (s UserProfileScope) SetOrder(order []string) *UserProfileScope {
	s.order = order
	return &s
}
func (s UserProfileScope) GetOrder() []string {
	return s.order
}

func (s UserProfile) SetSQLAppend(appendTail string) (scope *UserProfileScope) {
	return s.Scope().SetSQLAppend(appendTail)
}
func (s UserProfileScope) SetSQLAppend(appendTail string) *UserProfileScope {
	s.appendTail = appendTail
	return &s
}

// Sets limit.
func (s UserProfile) Limit(limit int) (scope *UserProfileScope) { return s.Scope().Limit(limit) }
func (s *UserProfileScope) Limit(limit int) *UserProfileScope {
	s.limit = limit
	return s
}

// Gets limit
func (s UserProfileScope) GetLimit() int {
	return s.limit
}

// "Reload" reloads record using Primary Key
func (s *UserProfileFilter) Reload(db *reform.DB) error { return (*UserProfile)(s).Reload(db) }
func (s *UserProfile) Reload(db *reform.DB) (err error) {
	return db.FindByPrimaryKeyTo(s, s.PKValue())
}

// Create and Insert inserts new record to DB
func (s *UserProfile) Create() (err error) { return s.PtrScope().Create() }
func (s *UserProfileScope) Create() (err error) {
	return s.Insert()
}
func (s *UserProfile) Insert() (err error) { return s.PtrScope().Insert() }
func (s *UserProfileScope) Insert() (err error) {
	s.checkDb()
	err = s.db.Insert(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Replace "REPLACE INTO" new record to DB
func (s *UserProfile) Replace() (err error) { return s.PtrScope().Replace() }
func (s *UserProfileScope) Replace() (err error) {
	s.checkDb()
	err = s.db.Replace(s.item)
	if err == nil {
		s.doLog("REPLACE")
	}
	return err
}

// Save inserts new record to DB is PK is zero and updates existing record if PK is not zero
func (s *UserProfile) Save() (err error) { return s.PtrScope().Save() }
func (s *UserProfileScope) Save() (err error) {
	s.checkDb()
	err = s.db.Save(s.item)
	if err == nil {
		s.doLog("INSERT")
	}
	return err
}

// Update updates existing record in DB
func (s UserProfile) Update() (err error) { return s.Scope().Update() }
func (s *UserProfileScope) Update() (err error) {
	s.checkDb()
	err = s.db.Update(s.item)
	if err == nil {
		s.doLog("UPDATE")
	}
	return err
}

// Delete deletes existing record in DB
func (s UserProfile) Delete() (err error) { return s.Scope().Delete() }
func (s *UserProfileScope) Delete() (err error) {
	s.checkDb()
	err = s.db.Delete(s.item)
	if err == nil {
		s.doLog("DELETE")
	}
	return err
}

func (s *UserProfileScope) doLog(requestType string) {
	if !s.loggingEnabled {
		return
	}

	var logRow UserProfileLogRow
	logRow.UserProfile = *s.item
	logRow.LogAuthor = s.loggingAuthor
	logRow.LogAction = requestType
	logRow.LogDate = time.Now()
	logRow.LogComment = s.loggingComment

	s.db.Insert(&logRow)
}

// Enables logging to table "user_profiles_log". This table should has the same schema, except:
// - Unique/Primary keys should be removed
// - Should be added next fields: "log_author" (nullable string), "log_date" (timestamp), "log_action" (enum("INSERT", "UPDATE", "DELETE")), "log_comment" (string)
func (s *UserProfile) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *UserProfileScope) {
	return s.Scope().Log(enableLogging, author, commentFormat, commentArgs...)
}
func (s *UserProfileScope) Log(enableLogging bool, author *string, commentFormat string, commentArgs ...interface{}) (scope *UserProfileScope) {
	s.loggingEnabled = enableLogging
	s.loggingAuthor = author
	s.loggingComment = fmt.Sprintf(commentFormat, commentArgs...)

	return s
}

// Table returns Table object for that record.
func (s UserProfile) Table() reform.Table {
	return UserProfileTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s UserProfile) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s UserProfile) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s UserProfile) HasPK() bool {
	return s.Id != UserProfileTable.z[UserProfileTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *UserProfileFilter) SetPK(pk interface{}) { (*UserProfile)(s).SetPK(pk) }
func (s *UserProfile) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.Id = int(i64)
	} else {
		s.Id = pk.(int)
	}
}

var (
	// check interfaces
	_ reform.View   = UserProfileTable
	_ reform.Struct = (*UserProfile)(nil)
	_ reform.Table  = UserProfileTable
	_ reform.Record = (*UserProfile)(nil)
	_ fmt.Stringer  = (*UserProfile)(nil)

	// querier
	UserProfileSQL        = UserProfile{} // Should be read only
	defaultDB_UserProfile *reform.DB
)

func init() {
	//parse.AssertUpToDate(&UserProfileTable.s, new(UserProfile)) // Temporary disabled (doesn't work with arbitary types like "type sliceString []string")
}

package token

// Reserved
const (
	ALL               = "ALL"
	ANALYSE           = "ANALYSE"
	ANALYZE           = "ANALYZE"
	AND               = "AND"
	ANY               = "ANY"
	ARRAY             = "ARRAY"
	AS                = "AS"
	ASC               = "ASC"
	ASYMMETRIC        = "ASYMMETRIC"
	AUTHORIZATION     = "AUTHORIZATION"
	BINARY            = "BINARY"
	BOTH              = "BOTH"
	CASE              = "CASE"
	CAST              = "CAST"
	CHECK             = "CHECK"
	COLLATE           = "COLLATE"
	COLLATION         = "COLLATION"
	COLUMN            = "COLUMN"
	CONCURRENTLY      = "CONCURRENTLY"
	CONSTRAINT        = "CONSTRAINT"
	CREATE            = "CREATE"
	CROSS             = "CROSS"
	CURRENT_CATALOG   = "CURRENT_CATALOG"
	CURRENT_DATE      = "CURRENT_DATE"
	CURRENT_ROLE      = "CURRENT_ROLE"
	CURRENT_SCHEMA    = "CURRENT_SCHEMA"
	CURRENT_TIME      = "CURRENT_TIME"
	CURRENT_TIMESTAMP = "CURRENT_TIMESTAMP"
	CURRENT_USER      = "CURRENT_USER"
	DEFAULT           = "DEFAULT"
	DEFERRABLE        = "DEFERRABLE"
	DESC              = "DESC"
	DISTINCT          = "DISTINCT"
	DO                = "DO"
	ELSE              = "ELSE"
	END               = "END"
	EXCEPT            = "EXCEPT"
	FALSE             = "FALSE"
	FETCH             = "FETCH"
	FOR               = "FOR"
	FOREIGN           = "FOREIGN"
	FREEZE            = "FREEZE"
	FROM              = "FROM"
	FULL              = "FULL"
	GRANT             = "GRANT"
	GROUP             = "GROUP"
	HAVING            = "HAVING"
	ILIKE             = "ILIKE"
	IN                = "IN"
	INITIALLY         = "INITIALLY"
	INNER             = "INNER"
	INTERSECT         = "INTERSECT"
	INTO              = "INTO"
	IS                = "IS"
	ISNULL            = "ISNULL"
	JOIN              = "JOIN"
	LATERAL           = "LATERAL"
	LEADING           = "LEADING"
	LEFT              = "LEFT"
	LIKE              = "LIKE"
	LIMIT             = "LIMIT"
	LOCALTIME         = "LOCALTIME"
	LOCALTIMESTAMP    = "LOCALTIMESTAMP"
	NATURAL           = "NATURAL"
	NOT               = "NOT"
	NOTNULL           = "NOTNULL"
	NULL              = "NULL"
	OFFSET            = "OFFSET"
	ON                = "ON"
	ONLY              = "ONLY"
	OR                = "OR"
	ORDER             = "ORDER"
	OUTER             = "OUTER"
	OVERLAPS          = "OVERLAPS"
	PLACING           = "PLACING"
	PRIMARY           = "PRIMARY"
	REFERENCES        = "REFERENCES"
	RETURNING         = "RETURNING"
	RIGHT             = "RIGHT"
	SELECT            = "SELECT"
	SESSION_USER      = "SESSION_USER"
	SIMILAR           = "SIMILAR"
	SOME              = "SOME"
	SYMMETRIC         = "SYMMETRIC"
	TABLE             = "TABLE"
	TABLESAMPLE       = "TABLESAMPLE"
	THEN              = "THEN"
	TO                = "TO"
	TRAILING          = "TRAILING"
	TRUE              = "TRUE"
	UNION             = "UNION"
	UNIQUE            = "UNIQUE"
	USER              = "USER"
	USING             = "USING"
	VARIADIC          = "VARIADIC"
	VERBOSE           = "VERBOSE"
	WHEN              = "WHEN"
	WHERE             = "WHERE"
	WINDOW            = "WINDOW"
	WITH              = "WITH"
)

// Non-reserved
const (
	ABORT           = "ABORT"
	ABSOLUTE        = "ABSOLUTE"
	ACCESS          = "ACCESS"
	ACTION          = "ACTION"
	ADD             = "ADD"
	ADMIN           = "ADMIN"
	AFTER           = "AFTER"
	AGGREGATE       = "AGGREGATE"
	ALSO            = "ALSO"
	ALTER           = "ALTER"
	ALWAYS          = "ALWAYS"
	ASENSITIVE      = "ASENSITIVE"
	ASSERTION       = "ASSERTION"
	ASSIGNMENT      = "ASSIGNMENT"
	AT              = "AT"
	ATOMIC          = "ATOMIC"
	ATTACH          = "ATTACH"
	ATTRIBUTE       = "ATTRIBUTE"
	BACKWARD        = "BACKWARD"
	BEFORE          = "BEFORE"
	BEGIN           = "BEGIN"
	BETWEEN         = "BETWEEN"
	BIGINT          = "BIGINT"
	BIT             = "BIT"
	BOOLEAN         = "BOOLEAN"
	BREADTH         = "BREADTH"
	BY              = "BY"
	CACHE           = "CACHE"
	CALL            = "CALL"
	CALLED          = "CALLED"
	CASCADE         = "CASCADE"
	CASCADED        = "CASCADED"
	CATALOG         = "CATALOG"
	CHAIN           = "CHAIN"
	CHAR            = "CHAR"
	CHARACTER       = "CHARACTER"
	CHARACTERISTICS = "CHARACTERISTICS"
	CHECKPOINT      = "CHECKPOINT"
	CLASS           = "CLASS"
	CLOSE           = "CLOSE"
	CLUSTER         = "CLUSTER"
	COALESCE        = "COALESCE"
	COLUMNS         = "COLUMNS"
	COMMENT         = "COMMENT"
	COMMENTS        = "COMMENTS"
	COMMIT          = "COMMIT"
	COMMITTED       = "COMMITTED"
	COMPRESSION     = "COMPRESSION"
	CONFIGURATION   = "CONFIGURATION"
	CONFLICT        = "CONFLICT"
	CONNECTION      = "CONNECTION"
	CONSTRAINTS     = "CONSTRAINTS"
	CONTENT         = "CONTENT"
	CONTINUE        = "CONTINUE"
	CONVERSION      = "CONVERSION"
	COPY            = "COPY"
	COST            = "COST"
	CSV             = "CSV"
	CUBE            = "CUBE"
	CURRENT         = "CURRENT"
	CURSOR          = "CURSOR"
	CYCLE           = "CYCLE"
	DATA            = "DATA"
	DATABASE        = "DATABASE"
	DAY             = "DAY"
	DEALLOCATE      = "DEALLOCATE"
	DEC             = "DEC"
	DECIMAL         = "DECIMAL"
	DECLARE         = "DECLARE"
	DEFAULTS        = "DEFAULTS"
	DEFERRED        = "DEFERRED"
	DEFINER         = "DEFINER"
	DELETE          = "DELETE"
	DELIMITER       = "DELIMITER"
	DELIMITERS      = "DELIMITERS"
	DEPENDS         = "DEPENDS"
	DEPTH           = "DEPTH"
	DETACH          = "DETACH"
	DICTIONARY      = "DICTIONARY"
	DISABLE         = "DISABLE"
	DISCARD         = "DISCARD"
	DOCUMENT        = "DOCUMENT"
	DOMAIN          = "DOMAIN"
	DOUBLE          = "DOUBLE"
	DROP            = "DROP"
	EACH            = "EACH"
	ENABLE          = "ENABLE"
	ENCODING        = "ENCODING"
	ENCRYPTED       = "ENCRYPTED"
	ENUM            = "ENUM"
	ESCAPE          = "ESCAPE"
	EVENT           = "EVENT"
	EXCLUDE         = "EXCLUDE"
	EXCLUDING       = "EXCLUDING"
	EXCLUSIVE       = "EXCLUSIVE"
	EXECUTE         = "EXECUTE"
	EXISTS          = "EXISTS"
	EXPLAIN         = "EXPLAIN"
	EXPRESSION      = "EXPRESSION"
	EXTENSION       = "EXTENSION"
	EXTERNAL        = "EXTERNAL"
	EXTRACT         = "EXTRACT"
	FAMILY          = "FAMILY"
	FILTER          = "FILTER"
	FINALIZE        = "FINALIZE"
	FIRST           = "FIRST"
	FLOAT           = "FLOAT"
	FOLLOWING       = "FOLLOWING"
	FORCE           = "FORCE"
	FORWARD         = "FORWARD"
	FUNCTION        = "FUNCTION"
	FUNCTIONS       = "FUNCTIONS"
	GENERATED       = "GENERATED"
	GLOBAL          = "GLOBAL"
	GRANTED         = "GRANTED"
	GREATEST        = "GREATEST"
	GROUPING        = "GROUPING"
	GROUPS          = "GROUPS"
	HANDLER         = "HANDLER"
	HEADER          = "HEADER"
	HOLD            = "HOLD"
	HOUR            = "HOUR"
	IDENTITY        = "IDENTITY"
	IF              = "IF"
	IMMEDIATE       = "IMMEDIATE"
	IMMUTABLE       = "IMMUTABLE"
	IMPLICIT        = "IMPLICIT"
	IMPORT          = "IMPORT"
	INCLUDE         = "INCLUDE"
	INCLUDING       = "INCLUDING"
	INCREMENT       = "INCREMENT"
	INDEX           = "INDEX"
	INDEXES         = "INDEXES"
	INHERIT         = "INHERIT"
	INHERITS        = "INHERITS"
	INLINE          = "INLINE"
	INOUT           = "INOUT"
	INPUT           = "INPUT"
	INSENSITIVE     = "INSENSITIVE"
	INSERT          = "INSERT"
	INSTEAD         = "INSTEAD"
	INT             = "INT"
	INTEGER         = "INTEGER"
	INTERVAL        = "INTERVAL"
	INVOKER         = "INVOKER"
	ISOLATION       = "ISOLATION"
	KEY             = "KEY"
	LABEL           = "LABEL"
	LANGUAGE        = "LANGUAGE"
	LARGE           = "LARGE"
	LAST            = "LAST"
	LEAKPROOF       = "LEAKPROOF"
	LEAST           = "LEAST"
	LEVEL           = "LEVEL"
	LISTEN          = "LISTEN"
	LOAD            = "LOAD"
	LOCAL           = "LOCAL"
	LOCATION        = "LOCATION"
	LOCK            = "LOCK"
	LOCKED          = "LOCKED"
	LOGGED          = "LOGGED"
	MAPPING         = "MAPPING"
	MATCH           = "MATCH"
	MATCHED         = "MATCHED"
	MATERIALIZED    = "MATERIALIZED"
	MAXVALUE        = "MAXVALUE"
	MERGE           = "MERGE"
	METHOD          = "METHOD"
	MINUTE          = "MINUTE"
	MINVALUE        = "MINVALUE"
	MODE            = "MODE"
	MONTH           = "MONTH"
	MOVE            = "MOVE"
	NAME            = "NAME"
	NAMES           = "NAMES"
	NATIONAL        = "NATIONAL"
	NCHAR           = "NCHAR"
	NEW             = "NEW"
	NEXT            = "NEXT"
	NFC             = "NFC"
	NFD             = "NFD"
	NFKC            = "NFKC"
	NFKD            = "NFKD"
	NO              = "NO"
	NONE            = "NONE"
	NORMALIZE       = "NORMALIZE"
	NORMALIZED      = "NORMALIZED"
	NOTHING         = "NOTHING"
	NOTIFY          = "NOTIFY"
	NOWAIT          = "NOWAIT"
	NULLIF          = "NULLIF"
	NULLS           = "NULLS"
	NUMERIC         = "NUMERIC"
	OBJECT          = "OBJECT"
	OF              = "OF"
	OFF             = "OFF"
	OIDS            = "OIDS"
	OLD             = "OLD"
	OPERATOR        = "OPERATOR"
	OPTION          = "OPTION"
	OPTIONS         = "OPTIONS"
	ORDINALITY      = "ORDINALITY"
	OTHERS          = "OTHERS"
	OUT             = "OUT"
	OVER            = "OVER"
	OVERLAY         = "OVERLAY"
	OVERRIDING      = "OVERRIDING"
	OWNED           = "OWNED"
	OWNER           = "OWNER"
	PARALLEL        = "PARALLEL"
	PARAMETER       = "PARAMETER"
	PARSER          = "PARSER"
	PARTIAL         = "PARTIAL"
	PARTITION       = "PARTITION"
	PASSING         = "PASSING"
	PASSWORD        = "PASSWORD"
	PLANS           = "PLANS"
	POLICY          = "POLICY"
	POSITION        = "POSITION"
	PRECEDING       = "PRECEDING"
	PRECISION       = "PRECISION"
	PREPARE         = "PREPARE"
	PREPARED        = "PREPARED"
	PRESERVE        = "PRESERVE"
	PRIOR           = "PRIOR"
	PRIVILEGES      = "PRIVILEGES"
	PROCEDURAL      = "PROCEDURAL"
	PROCEDURE       = "PROCEDURE"
	PROCEDURES      = "PROCEDURES"
	PROGRAM         = "PROGRAM"
	PUBLICATION     = "PUBLICATION"
	QUOTE           = "QUOTE"
	RANGE           = "RANGE"
	READ            = "READ"
	REAL            = "REAL"
	REASSIGN        = "REASSIGN"
	RECHECK         = "RECHECK"
	RECURSIVE       = "RECURSIVE"
	REF             = "REF"
	REFERENCING     = "REFERENCING"
	REFRESH         = "REFRESH"
	REINDEX         = "REINDEX"
	RELATIVE        = "RELATIVE"
	RELEASE         = "RELEASE"
	RENAME          = "RENAME"
	REPEATABLE      = "REPEATABLE"
	REPLACE         = "REPLACE"
	REPLICA         = "REPLICA"
	RESET           = "RESET"
	RESTART         = "RESTART"
	RESTRICT        = "RESTRICT"
	RETURN          = "RETURN"
	RETURNS         = "RETURNS"
	REVOKE          = "REVOKE"
	ROLE            = "ROLE"
	ROLLBACK        = "ROLLBACK"
	ROLLUP          = "ROLLUP"
	ROUTINE         = "ROUTINE"
	ROUTINES        = "ROUTINES"
	ROW             = "ROW"
	ROWS            = "ROWS"
	RULE            = "RULE"
	SAVEPOINT       = "SAVEPOINT"
	SCHEMA          = "SCHEMA"
	SCHEMAS         = "SCHEMAS"
	SCROLL          = "SCROLL"
	SEARCH          = "SEARCH"
	SECOND          = "SECOND"
	SECURITY        = "SECURITY"
	SEQUENCE        = "SEQUENCE"
	SEQUENCES       = "SEQUENCES"
	SERIALIZABLE    = "SERIALIZABLE"
	SERVER          = "SERVER"
	SESSION         = "SESSION"
	SET             = "SET"
	SETOF           = "SETOF"
	SETS            = "SETS"
	SHARE           = "SHARE"
	SHOW            = "SHOW"
	SIMPLE          = "SIMPLE"
	SKIP            = "SKIP"
	SMALLINT        = "SMALLINT"
	SNAPSHOT        = "SNAPSHOT"
	SQL             = "SQL"
	STABLE          = "STABLE"
	STANDALONE      = "STANDALONE"
	START           = "START"
	STATEMENT       = "STATEMENT"
	STATISTICS      = "STATISTICS"
	STDIN           = "STDIN"
	STDOUT          = "STDOUT"
	STORAGE         = "STORAGE"
	STORED          = "STORED"
	STRICT          = "STRICT"
	STRIP           = "STRIP"
	SUBSCRIPTION    = "SUBSCRIPTION"
	SUBSTRING       = "SUBSTRING"
	SUPPORT         = "SUPPORT"
	SYSID           = "SYSID"
	SYSTEM          = "SYSTEM"
	TABLES          = "TABLES"
	TABLESPACE      = "TABLESPACE"
	TEMP            = "TEMP"
	TEMPLATE        = "TEMPLATE"
	TEMPORARY       = "TEMPORARY"
	TEXT            = "TEXT"
	TIES            = "TIES"
	TIME            = "TIME"
	TIMESTAMP       = "TIMESTAMP"
	TRANSACTION     = "TRANSACTION"
	TRANSFORM       = "TRANSFORM"
	TREAT           = "TREAT"
	TRIGGER         = "TRIGGER"
	TRIM            = "TRIM"
	TRUNCATE        = "TRUNCATE"
	TRUSTED         = "TRUSTED"
	TYPE            = "TYPE"
	TYPES           = "TYPES"
	UESCAPE         = "UESCAPE"
	UNBOUNDED       = "UNBOUNDED"
	UNCOMMITTED     = "UNCOMMITTED"
	UNENCRYPTED     = "UNENCRYPTED"
	UNKNOWN         = "UNKNOWN"
	UNLISTEN        = "UNLISTEN"
	UNLOGGED        = "UNLOGGED"
	UNTIL           = "UNTIL"
	UPDATE          = "UPDATE"
	VACUUM          = "VACUUM"
	VALID           = "VALID"
	VALIDATE        = "VALIDATE"
	VALIDATOR       = "VALIDATOR"
	VALUE           = "VALUE"
	VALUES          = "VALUES"
	VARCHAR         = "VARCHAR"
	VARYING         = "VARYING"
	VERSION         = "VERSION"
	VIEW            = "VIEW"
	VIEWS           = "VIEWS"
	VOLATILE        = "VOLATILE"
	WHITESPACE      = "WHITESPACE"
	WITHIN          = "WITHIN"
	WITHOUT         = "WITHOUT"
	WORK            = "WORK"
	WRAPPER         = "WRAPPER"
	WRITE           = "WRITE"
	XML             = "XML"
	XMLATTRIBUTES   = "XMLATTRIBUTES"
	XMLCONCAT       = "XMLCONCAT"
	XMLELEMENT      = "XMLELEMENT"
	XMLEXISTS       = "XMLEXISTS"
	XMLFOREST       = "XMLFOREST"
	XMLNAMESPACES   = "XMLNAMESPACES"
	XMLPARSE        = "XMLPARSE"
	XMLPI           = "XMLPI"
	XMLROOT         = "XMLROOT"
	XMLSERIALIZE    = "XMLSERIALIZE"
	XMLTABLE        = "XMLTABLE"
	YEAR            = "YEAR"
	YES             = "YES"
	ZONE            = "ZONE"
)

var KEYWORDS = map[string]TokenType{"ALL": ALL,
	"ANALYSE":           ANALYSE,
	"ANALYZE":           ANALYZE,
	"AND":               AND,
	"ANY":               ANY,
	"ARRAY":             ARRAY,
	"AS":                AS,
	"ASC":               ASC,
	"ASYMMETRIC":        ASYMMETRIC,
	"AUTHORIZATION":     AUTHORIZATION,
	"BINARY":            BINARY,
	"BOTH":              BOTH,
	"CASE":              CASE,
	"CAST":              CAST,
	"CHECK":             CHECK,
	"COLLATE":           COLLATE,
	"COLLATION":         COLLATION,
	"COLUMN":            COLUMN,
	"CONCURRENTLY":      CONCURRENTLY,
	"CONSTRAINT":        CONSTRAINT,
	"CREATE":            CREATE,
	"CROSS":             CROSS,
	"CURRENT_CATALOG":   CURRENT_CATALOG,
	"CURRENT_DATE":      CURRENT_DATE,
	"CURRENT_ROLE":      CURRENT_ROLE,
	"CURRENT_SCHEMA":    CURRENT_SCHEMA,
	"CURRENT_TIME":      CURRENT_TIME,
	"CURRENT_TIMESTAMP": CURRENT_TIMESTAMP,
	"CURRENT_USER":      CURRENT_USER,
	"DEFAULT":           DEFAULT,
	"DEFERRABLE":        DEFERRABLE,
	"DESC":              DESC,
	"DISTINCT":          DISTINCT,
	"DO":                DO,
	"ELSE":              ELSE,
	"END":               END,
	"EXCEPT":            EXCEPT,
	"FALSE":             FALSE,
	"FETCH":             FETCH,
	"FOR":               FOR,
	"FOREIGN":           FOREIGN,
	"FREEZE":            FREEZE,
	"FROM":              FROM,
	"FULL":              FULL,
	"GRANT":             GRANT,
	"GROUP":             GROUP,
	"HAVING":            HAVING,
	"ILIKE":             ILIKE,
	"IN":                IN,
	"INITIALLY":         INITIALLY,
	"INNER":             INNER,
	"INTERSECT":         INTERSECT,
	"INTO":              INTO,
	"IS":                IS,
	"ISNULL":            ISNULL,
	"JOIN":              JOIN,
	"LATERAL":           LATERAL,
	"LEADING":           LEADING,
	"LEFT":              LEFT,
	"LIKE":              LIKE,
	"LIMIT":             LIMIT,
	"LOCALTIME":         LOCALTIME,
	"LOCALTIMESTAMP":    LOCALTIMESTAMP,
	"NATURAL":           NATURAL,
	"NOT":               NOT,
	"NOTNULL":           NOTNULL,
	"NULL":              NULL,
	"OFFSET":            OFFSET,
	"ON":                ON,
	"ONLY":              ONLY,
	"OR":                OR,
	"ORDER":             ORDER,
	"OUTER":             OUTER,
	"OVERLAPS":          OVERLAPS,
	"PLACING":           PLACING,
	"PRIMARY":           PRIMARY,
	"REFERENCES":        REFERENCES,
	"RETURNING":         RETURNING,
	"RIGHT":             RIGHT,
	"SELECT":            SELECT,
	"SESSION_USER":      SESSION_USER,
	"SIMILAR":           SIMILAR,
	"SOME":              SOME,
	"SYMMETRIC":         SYMMETRIC,
	"TABLE":             TABLE,
	"TABLESAMPLE":       TABLESAMPLE,
	"THEN":              THEN,
	"TO":                TO,
	"TRAILING":          TRAILING,
	"TRUE":              TRUE,
	"UNION":             UNION,
	"UNIQUE":            UNIQUE,
	"USER":              USER,
	"USING":             USING,
	"VARIADIC":          VARIADIC,
	"VERBOSE":           VERBOSE,
	"WHEN":              WHEN,
	"WHERE":             WHERE,
	"WINDOW":            WINDOW,
	"WITH":              WITH}

var NRKEYWORDS = map[string]TokenType{"ABORT": ABORT,
	"ABSOLUTE":        ABSOLUTE,
	"ACCESS":          ACCESS,
	"ACTION":          ACTION,
	"ADD":             ADD,
	"ADMIN":           ADMIN,
	"AFTER":           AFTER,
	"AGGREGATE":       AGGREGATE,
	"ALSO":            ALSO,
	"ALTER":           ALTER,
	"ALWAYS":          ALWAYS,
	"ASENSITIVE":      ASENSITIVE,
	"ASSERTION":       ASSERTION,
	"ASSIGNMENT":      ASSIGNMENT,
	"AT":              AT,
	"ATOMIC":          ATOMIC,
	"ATTACH":          ATTACH,
	"ATTRIBUTE":       ATTRIBUTE,
	"BACKWARD":        BACKWARD,
	"BEFORE":          BEFORE,
	"BEGIN":           BEGIN,
	"BETWEEN":         BETWEEN,
	"BIGINT":          BIGINT,
	"BIT":             BIT,
	"BOOLEAN":         BOOLEAN,
	"BREADTH":         BREADTH,
	"BY":              BY,
	"CACHE":           CACHE,
	"CALL":            CALL,
	"CALLED":          CALLED,
	"CASCADE":         CASCADE,
	"CASCADED":        CASCADED,
	"CATALOG":         CATALOG,
	"CHAIN":           CHAIN,
	"CHAR":            CHAR,
	"CHARACTER":       CHARACTER,
	"CHARACTERISTICS": CHARACTERISTICS,
	"CHECKPOINT":      CHECKPOINT,
	"CLASS":           CLASS,
	"CLOSE":           CLOSE,
	"CLUSTER":         CLUSTER,
	"COALESCE":        COALESCE,
	"COLUMNS":         COLUMNS,
	"COMMENT":         COMMENT,
	"COMMENTS":        COMMENTS,
	"COMMIT":          COMMIT,
	"COMMITTED":       COMMITTED,
	"COMPRESSION":     COMPRESSION,
	"CONFIGURATION":   CONFIGURATION,
	"CONFLICT":        CONFLICT,
	"CONNECTION":      CONNECTION,
	"CONSTRAINTS":     CONSTRAINTS,
	"CONTENT":         CONTENT,
	"CONTINUE":        CONTINUE,
	"CONVERSION":      CONVERSION,
	"COPY":            COPY,
	"COST":            COST,
	"CSV":             CSV,
	"CUBE":            CUBE,
	"CURRENT":         CURRENT,
	"CURSOR":          CURSOR,
	"CYCLE":           CYCLE,
	"DATA":            DATA,
	"DATABASE":        DATABASE,
	"DAY":             DAY,
	"DEALLOCATE":      DEALLOCATE,
	"DEC":             DEC,
	"DECIMAL":         DECIMAL,
	"DECLARE":         DECLARE,
	"DEFAULTS":        DEFAULTS,
	"DEFERRED":        DEFERRED,
	"DEFINER":         DEFINER,
	"DELETE":          DELETE,
	"DELIMITER":       DELIMITER,
	"DELIMITERS":      DELIMITERS,
	"DEPENDS":         DEPENDS,
	"DEPTH":           DEPTH,
	"DETACH":          DETACH,
	"DICTIONARY":      DICTIONARY,
	"DISABLE":         DISABLE,
	"DISCARD":         DISCARD,
	"DOCUMENT":        DOCUMENT,
	"DOMAIN":          DOMAIN,
	"DOUBLE":          DOUBLE,
	"DROP":            DROP,
	"EACH":            EACH,
	"ENABLE":          ENABLE,
	"ENCODING":        ENCODING,
	"ENCRYPTED":       ENCRYPTED,
	"ENUM":            ENUM,
	"ESCAPE":          ESCAPE,
	"EVENT":           EVENT,
	"EXCLUDE":         EXCLUDE,
	"EXCLUDING":       EXCLUDING,
	"EXCLUSIVE":       EXCLUSIVE,
	"EXECUTE":         EXECUTE,
	"EXISTS":          EXISTS,
	"EXPLAIN":         EXPLAIN,
	"EXPRESSION":      EXPRESSION,
	"EXTENSION":       EXTENSION,
	"EXTERNAL":        EXTERNAL,
	"EXTRACT":         EXTRACT,
	"FAMILY":          FAMILY,
	"FILTER":          FILTER,
	"FINALIZE":        FINALIZE,
	"FIRST":           FIRST,
	"FLOAT":           FLOAT,
	"FOLLOWING":       FOLLOWING,
	"FORCE":           FORCE,
	"FORWARD":         FORWARD,
	"FUNCTION":        FUNCTION,
	"FUNCTIONS":       FUNCTIONS,
	"GENERATED":       GENERATED,
	"GLOBAL":          GLOBAL,
	"GRANTED":         GRANTED,
	"GREATEST":        GREATEST,
	"GROUPING":        GROUPING,
	"GROUPS":          GROUPS,
	"HANDLER":         HANDLER,
	"HEADER":          HEADER,
	"HOLD":            HOLD,
	"HOUR":            HOUR,
	"IDENTITY":        IDENTITY,
	"IF":              IF,
	"IMMEDIATE":       IMMEDIATE,
	"IMMUTABLE":       IMMUTABLE,
	"IMPLICIT":        IMPLICIT,
	"IMPORT":          IMPORT,
	"INCLUDE":         INCLUDE,
	"INCLUDING":       INCLUDING,
	"INCREMENT":       INCREMENT,
	"INDEX":           INDEX,
	"INDEXES":         INDEXES,
	"INHERIT":         INHERIT,
	"INHERITS":        INHERITS,
	"INLINE":          INLINE,
	"INOUT":           INOUT,
	"INPUT":           INPUT,
	"INSENSITIVE":     INSENSITIVE,
	"INSERT":          INSERT,
	"INSTEAD":         INSTEAD,
	"INT":             INT,
	"INTEGER":         INTEGER,
	"INTERVAL":        INTERVAL,
	"INVOKER":         INVOKER,
	"ISOLATION":       ISOLATION,
	"KEY":             KEY,
	"LABEL":           LABEL,
	"LANGUAGE":        LANGUAGE,
	"LARGE":           LARGE,
	"LAST":            LAST,
	"LEAKPROOF":       LEAKPROOF,
	"LEAST":           LEAST,
	"LEVEL":           LEVEL,
	"LISTEN":          LISTEN,
	"LOAD":            LOAD,
	"LOCAL":           LOCAL,
	"LOCATION":        LOCATION,
	"LOCK":            LOCK,
	"LOCKED":          LOCKED,
	"LOGGED":          LOGGED,
	"MAPPING":         MAPPING,
	"MATCH":           MATCH,
	"MATCHED":         MATCHED,
	"MATERIALIZED":    MATERIALIZED,
	"MAXVALUE":        MAXVALUE,
	"MERGE":           MERGE,
	"METHOD":          METHOD,
	"MINUTE":          MINUTE,
	"MINVALUE":        MINVALUE,
	"MODE":            MODE,
	"MONTH":           MONTH,
	"MOVE":            MOVE,
	"NAME":            NAME,
	"NAMES":           NAMES,
	"NATIONAL":        NATIONAL,
	"NCHAR":           NCHAR,
	"NEW":             NEW,
	"NEXT":            NEXT,
	"NFC":             NFC,
	"NFD":             NFD,
	"NFKC":            NFKC,
	"NFKD":            NFKD,
	"NO":              NO,
	"NONE":            NONE,
	"NORMALIZE":       NORMALIZE,
	"NORMALIZED":      NORMALIZED,
	"NOTHING":         NOTHING,
	"NOTIFY":          NOTIFY,
	"NOWAIT":          NOWAIT,
	"NULLIF":          NULLIF,
	"NULLS":           NULLS,
	"NUMERIC":         NUMERIC,
	"OBJECT":          OBJECT,
	"OF":              OF,
	"OFF":             OFF,
	"OIDS":            OIDS,
	"OLD":             OLD,
	"OPERATOR":        OPERATOR,
	"OPTION":          OPTION,
	"OPTIONS":         OPTIONS,
	"ORDINALITY":      ORDINALITY,
	"OTHERS":          OTHERS,
	"OUT":             OUT,
	"OVER":            OVER,
	"OVERLAY":         OVERLAY,
	"OVERRIDING":      OVERRIDING,
	"OWNED":           OWNED,
	"OWNER":           OWNER,
	"PARALLEL":        PARALLEL,
	"PARAMETER":       PARAMETER,
	"PARSER":          PARSER,
	"PARTIAL":         PARTIAL,
	"PARTITION":       PARTITION,
	"PASSING":         PASSING,
	"PASSWORD":        PASSWORD,
	"PLANS":           PLANS,
	"POLICY":          POLICY,
	"POSITION":        POSITION,
	"PRECEDING":       PRECEDING,
	"PRECISION":       PRECISION,
	"PREPARE":         PREPARE,
	"PREPARED":        PREPARED,
	"PRESERVE":        PRESERVE,
	"PRIOR":           PRIOR,
	"PRIVILEGES":      PRIVILEGES,
	"PROCEDURAL":      PROCEDURAL,
	"PROCEDURE":       PROCEDURE,
	"PROCEDURES":      PROCEDURES,
	"PROGRAM":         PROGRAM,
	"PUBLICATION":     PUBLICATION,
	"QUOTE":           QUOTE,
	"RANGE":           RANGE,
	"READ":            READ,
	"REAL":            REAL,
	"REASSIGN":        REASSIGN,
	"RECHECK":         RECHECK,
	"RECURSIVE":       RECURSIVE,
	"REF":             REF,
	"REFERENCING":     REFERENCING,
	"REFRESH":         REFRESH,
	"REINDEX":         REINDEX,
	"RELATIVE":        RELATIVE,
	"RELEASE":         RELEASE,
	"RENAME":          RENAME,
	"REPEATABLE":      REPEATABLE,
	"REPLACE":         REPLACE,
	"REPLICA":         REPLICA,
	"RESET":           RESET,
	"RESTART":         RESTART,
	"RESTRICT":        RESTRICT,
	"RETURN":          RETURN,
	"RETURNS":         RETURNS,
	"REVOKE":          REVOKE,
	"ROLE":            ROLE,
	"ROLLBACK":        ROLLBACK,
	"ROLLUP":          ROLLUP,
	"ROUTINE":         ROUTINE,
	"ROUTINES":        ROUTINES,
	"ROW":             ROW,
	"ROWS":            ROWS,
	"RULE":            RULE,
	"SAVEPOINT":       SAVEPOINT,
	"SCHEMA":          SCHEMA,
	"SCHEMAS":         SCHEMAS,
	"SCROLL":          SCROLL,
	"SEARCH":          SEARCH,
	"SECOND":          SECOND,
	"SECURITY":        SECURITY,
	"SEQUENCE":        SEQUENCE,
	"SEQUENCES":       SEQUENCES,
	"SERIALIZABLE":    SERIALIZABLE,
	"SERVER":          SERVER,
	"SESSION":         SESSION,
	"SET":             SET,
	"SETOF":           SETOF,
	"SETS":            SETS,
	"SHARE":           SHARE,
	"SHOW":            SHOW,
	"SIMPLE":          SIMPLE,
	"SKIP":            SKIP,
	"SMALLINT":        SMALLINT,
	"SNAPSHOT":        SNAPSHOT,
	"SQL":             SQL,
	"STABLE":          STABLE,
	"STANDALONE":      STANDALONE,
	"START":           START,
	"STATEMENT":       STATEMENT,
	"STATISTICS":      STATISTICS,
	"STDIN":           STDIN,
	"STDOUT":          STDOUT,
	"STORAGE":         STORAGE,
	"STORED":          STORED,
	"STRICT":          STRICT,
	"STRIP":           STRIP,
	"SUBSCRIPTION":    SUBSCRIPTION,
	"SUBSTRING":       SUBSTRING,
	"SUPPORT":         SUPPORT,
	"SYSID":           SYSID,
	"SYSTEM":          SYSTEM,
	"TABLES":          TABLES,
	"TABLESPACE":      TABLESPACE,
	"TEMP":            TEMP,
	"TEMPLATE":        TEMPLATE,
	"TEMPORARY":       TEMPORARY,
	"TEXT":            TEXT,
	"TIES":            TIES,
	"TIME":            TIME,
	"TIMESTAMP":       TIMESTAMP,
	"TRANSACTION":     TRANSACTION,
	"TRANSFORM":       TRANSFORM,
	"TREAT":           TREAT,
	"TRIGGER":         TRIGGER,
	"TRIM":            TRIM,
	"TRUNCATE":        TRUNCATE,
	"TRUSTED":         TRUSTED,
	"TYPE":            TYPE,
	"TYPES":           TYPES,
	"UESCAPE":         UESCAPE,
	"UNBOUNDED":       UNBOUNDED,
	"UNCOMMITTED":     UNCOMMITTED,
	"UNENCRYPTED":     UNENCRYPTED,
	"UNKNOWN":         UNKNOWN,
	"UNLISTEN":        UNLISTEN,
	"UNLOGGED":        UNLOGGED,
	"UNTIL":           UNTIL,
	"UPDATE":          UPDATE,
	"VACUUM":          VACUUM,
	"VALID":           VALID,
	"VALIDATE":        VALIDATE,
	"VALIDATOR":       VALIDATOR,
	"VALUE":           VALUE,
	"VALUES":          VALUES,
	"VARCHAR":         VARCHAR,
	"VARYING":         VARYING,
	"VERSION":         VERSION,
	"VIEW":            VIEW,
	"VIEWS":           VIEWS,
	"VOLATILE":        VOLATILE,
	"WHITESPACE":      WHITESPACE,
	"WITHIN":          WITHIN,
	"WITHOUT":         WITHOUT,
	"WORK":            WORK,
	"WRAPPER":         WRAPPER,
	"WRITE":           WRITE,
	"XML":             XML,
	"XMLATTRIBUTES":   XMLATTRIBUTES,
	"XMLCONCAT":       XMLCONCAT,
	"XMLELEMENT":      XMLELEMENT,
	"XMLEXISTS":       XMLEXISTS,
	"XMLFOREST":       XMLFOREST,
	"XMLNAMESPACES":   XMLNAMESPACES,
	"XMLPARSE":        XMLPARSE,
	"XMLPI":           XMLPI,
	"XMLROOT":         XMLROOT,
	"XMLSERIALIZE":    XMLSERIALIZE,
	"XMLTABLE":        XMLTABLE,
	"YEAR":            YEAR,
	"YES":             YES,
	"ZONE":            ZONE}

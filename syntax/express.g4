grammar express;

parse: express+;

express
    // 数据类型
    : INTEGER				 									                                                        # Integer
    | FLOAT                                                                                                             # Float
    | TEXT			 											                                                        # Text

    // 判断类
    | IF LLIMIT  boolExpress COMMA express COMMA express  RLIMIT 	                                                    # If
    | IFS LLIMIT  subBoolExpress ( COMMA subBoolExpress)*  RLIMIT 	                                                    # Ifs
    | OR LLIMIT  boolExpress ( COMMA  boolExpress)*  RLIMIT 		                                                    # TrueOr
    | AND LLIMIT  boolExpress ( COMMA  boolExpress)*  RLIMIT 		                                                    # TrueAnd

    // 字符串类
    | CONCATSTRING  LLIMIT  express ( COMMA express)*  RLIMIT                                                           # ConcatString
    | FINDSTRING LLIMIT express COMMA express ( COMMA express )* RLIMIT                                                 # FindString

    // 数学表达式
    | MAX LLIMIT  express( COMMA  express)*  RLIMIT 							                                        # Max
    | MIN LLIMIT  express( COMMA  express)*  RLIMIT                                                                     # Min
    | SUM LLIMIT  express( COMMA  express)*  RLIMIT                                                                     # Sum
    | AVERAGE LLIMIT  express( COMMA  express)*  RLIMIT                                                                 # Average
    | ROUND LLIMIT express  COMMA  express  RLIMIT                                                                      # Round

    // 数学运算
    | express op=DIV express                                                                                            # Div       // 除法
    | express op=MUL express                                                                                            # Mul       // 乘法
    | express op=MOD express                                                                                            # Mod       // 求余数
    | express op=SUB express                                                                                            # Sub       // 减法
    | express op=ADD express                                                                                            # Add       // 加法
    | LLIMIT '-'express RLIMIT                                                                                          # Minus     // 负数处理
    | LLIMIT express RLIMIT                                                                                             # Plus      // 正数处理

    // 日期类
    | TODAY LLIMIT RLIMIT                                                                                               # Today
    | DATE LLIMIT express COMMA express COMMA express RLIMIT                                                            # Date
    | DATEDIF LLIMIT  express  COMMA  express  COMMA op=TEXT RLIMIT                                                     # DateDiff
    | DATEADD LLIMIT  express  COMMA  express  COMMA  op=TEXT  RLIMIT                                                   # DateAdd

    // 扩展函数
    | Helloworld LLIMIT  express  COMMA  express  RLIMIT                                                                # Helloworld
    ;


boolExpress
	: express op=COMPARATOR express									                                                    # Comparator
	| OR LLIMIT  boolExpress ( COMMA  boolExpress)*  RLIMIT 		                                                    # Or
    | AND LLIMIT  boolExpress ( COMMA  boolExpress)*  RLIMIT 		                                                    # And
    ;

subBoolExpress
    : boolExpress  COMMA  express                                                                                       # SubIfs
    ;


// 辅助类
WHITE_SPACE  :[ \t\r\n]+ -> skip ;

fragment ESC : '\\"' | '\\\\';

fragment
ExponentPart
    : [eE] [+-]? Digit+
    ;

fragment
Digit
    : [0-9]
    ;

// 数据类型
INTEGER:    Digit+;
TEXT
    : '#' (ESC|.)*? '#'
    | '\'' (ESC|.)*? '\''
    | '’' (ESC|.)*? '‘'
    | '"' (ESC|.)*? '"'
    | '“' (ESC|.)*? '“'
    ;

FLOAT
    : Digit + '.' Digit * ExponentPart?
    | '.' Digit + ExponentPart?
    | Digit + ExponentPart
    ;

COMMA
    : ','
    |  '，' ;

LLIMIT: '(' |  '（' ;
RLIMIT: ')' |  '）' ;

// 判断控制
IF						:'IF';
IFS                     :'IFS';
ELSE					:'ELSE';

// 布尔运算
AND						:'AND';
OR						:'OR';
COMPARATOR				:'>'|'＞'|'<'|'＜'|'='|'＝'|'≠'|'!='|'>='|'≥'|'＞＝'|'＜＝'|'≤'|'<=';


// 字符串处理
CONCATSTRING            :'CONCATSTR';
FINDSTRING              :'FINDSTR';


// 基础数学
MAX						:'MAX';
MIN                     :'MIN';
SUM                     :'SUM';
AVERAGE                 :'AVERAGE';
ROUND                   :'ROUND';

ADD                     :'+';
SUB                     :'-';
MUL                     :'×'|'*';
DIV                     :'÷'| '/';
MOD                     :'%';

// 日期
DoDays                  : 'D';
DoMonths                : 'M';
DoYears                 : 'Y';

TODAY                   :'TODAY';
DATE                    :'DATE';
DATEDIF                 :'DATEDIF';
DATEADD                 :'DATEADD';


// 扩展函数
Helloworld          :'HELLO_WORLD';








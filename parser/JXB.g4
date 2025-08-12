grammar JXB;

document
  : statement* EOF
  ;
statement
  : assignment
  | block
  | insert
  ;
assignment
  : KEYWORD STRING* NL
  ;
block
  : BEGIN STRING NL statement* END STRING NL 
  ;
insert
  : INSERT KEYWORD NL
  | INSERT STRING NL
  ;

BEGIN
  : 'begin'
  ;
END
  : 'end'
  ;
INSERT
  : 'insert'
  ;
KEYWORD
  : 'cv'
  | 'papersize'
  | 'name'
  | 'selfie'
  | 'address'
  | 'phone'
  | 'email'
  | 'website'
  | 'title'
  ;

fragment ALPHA
  : [A-Za-z]
  ;
fragment DIGIT
  : [0-9]
  ;
fragment SYM
  : [-.,+_:/?#@]
  ;
STRING
  : (ALPHA | DIGIT | SYM)+
  ;

WS
  : [ \t]+ -> skip
  ;

NL
  : ('\r' | '\n')+
  ;

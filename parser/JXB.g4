grammar JXB;

document
  : statement* EOF
  ;
statement
  : assignment
  | block
  | insert
  | maroto
  ;
assignment
  : KEYWORD STRING* NL
  ;
block
  : BEGIN STRING NL assignment* END STRING NL 
  ;
insert
  : INSERT (KEYWORD | STRING) ALIGNMENT? NL
  ;
maroto
  : ROW marcol* NL
  ;
marcol
  : COL insert*
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
  | 'name'
  | 'selfie'
  | 'address'
  | 'phone'
  | 'email'
  | 'website'
  | 'title'
  ;
ALIGNMENT
  : 'left'
  | 'right'
  | 'center'
  | 'bottom'
  | 'top'
  ;
ROW
  : 'row'
  ;
COL
  : 'col'
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

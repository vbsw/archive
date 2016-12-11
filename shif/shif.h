#ifndef  SHIF_H
#define  SHIF_H

#include <stdlib.h>

#define SHIF_EROR_WRONG_NODE_NAME 1
#define SHIF_EROR_WRONG_INDENT_DEPTH 2

#define SHIF_TRIM_NOTHING 0
#define SHIF_TRIM_BLANKS_LEFT 1
#define SHIF_TRIM_BLANKS_RIGHT 2
#define SHIF_TRIM_TABS_LEFT 4
#define SHIF_TRIM_TABS_RIGHT 8
#define SHIF_TRIM_ALL 15

extern int
SHiF_parse
(
	char *const in_cstring,
	void *const inout_data,
	int ( *const in_newNodeFct )
	(
		void *data,
		int indentDepth,
		char *nodeNameBegin,
		char *nodeNameEnd,
		char *nodeValueBegin,
		char *nodeValueEnd
	),
	int ( *const in_errorFct )
	(
		void *data,
		int errorCode,
		int errorLineNum
	),
	const int in_trimOption
);

#endif /* SHIF_H */

// This file is modified from source code from the FBC project.

#ifndef FBTYPES_H_
#define FBTYPES_H_

#include <stdio.h>


/*********************************************************************
 This file contains the mapping between IEC 61131 and C data types.
 Since C data types are non-portable, this file MUST be customized to
 match the data width assumed by the C compiler used for compiling the
 generated code. This enables the generated code to be easily ported
 to other platforms.
*********************************************************************/

// bool stuff
#ifndef __cplusplus
typedef char bool;
#endif // __cplusplus

#ifndef true
#define true 1
#endif

#ifndef false
#define false 0
#endif

#ifndef TRUE
#define TRUE true
#endif

#ifndef FALSE
#define FALSE false
#endif

#define STRING_LENGTH 32

// Bit strings
typedef bool BOOL;
typedef unsigned char BYTE;
typedef unsigned short WORD;
typedef unsigned int DWORD;
typedef unsigned long long LWORD;

// Integers
typedef char SINT;
typedef unsigned char USINT;
typedef short INT;
typedef unsigned short UINT;
typedef int DINT;
typedef unsigned int UDINT;
typedef long long LINT;
typedef unsigned long long ULINT;

// Reals
typedef float REAL;
typedef double LREAL;

// Duration
typedef long long TIME;

// Strings
typedef char FBstring[STRING_LENGTH];
typedef FBstring STRING;
typedef FBstring WSTRING;

#endif // FBTYPES_H_

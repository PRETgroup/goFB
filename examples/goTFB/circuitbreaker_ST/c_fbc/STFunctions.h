/*
Type conversion functions
..._TO_... Conversion between integers
BOOL_TO BOOL -->Type X
TO_BOOL Type X --> BOOL
TIME_TO / TIME_OF_DAY TIME / TIME_OF_DAY --> Type X
DATE_TO / DT_TO DATE / DATE_AND_TIME --> Type X
STRING_TO STRING --> Type X
TRUNC REAL --> INT
Numeric functions
*/

#include "fbtypes.h"

/**************************************** TO SINT ****************/ 
SINT UINT_TO_SINT(UINT var);
SINT INT_TO_SINT(INT var);
SINT DINT_TO_SINT(DINT var);
SINT LINT_TO_SINT(LINT var);
/**************************************** TO INT ****************/ 
INT SINT_TO_INT(SINT var);
INT UINT_TO_INT(UINT var);
INT DINT_TO_INT(DINT var);
INT LINT_TO_INT(LINT var);
INT REAL_TO_INT(REAL var);
/**************************************** TO UINT ****************/
UINT INT_TO_UINT(INT var);
UINT LINT_TO_UINT(LINT var);
UINT DINT_TO_UINT(DINT var);
UINT SINT_TO_UINT(SINT var);
/**************************************** TO LINT ****************/
LINT SINT_TO_LINT(SINT var);
LINT INT_TO_LINT(INT var);
LINT UINT_TO_LINT(UINT var);
LINT DINT_TO_LINT(DINT var);
/**************************************** TO DINT ****************/
DINT LINT_TO_DINT(LINT var);
DINT INT_TO_DINT(INT var);
DINT SINT_TO_DINT(SINT var);
DINT UINT_TO_DINT(UINT var);
int REAL_TO_DINT(float var);
/**************************************** TO UDINT ***************/
UDINT DINT_TO_UDINT(DINT var);
/**************************************** TO REAL ***************/
REAL INT_TO_REAL(INT var);
REAL DINT_TO_REAL(DINT var);


/* Other Functions */
INT ABS(INT var);
REAL SQRT(REAL var);
REAL COS(REAL var);
REAL SIN(REAL var);
REAL ATAN(REAL var);

#include <math.h> // for math functions

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
// Using 
// http://msdn.microsoft.com/en-us/library/dd607323(VS.85).aspx
// for ref


/**************************************** TO SINT ****************/
char UINT_TO_SINT(unsigned short var)
{
	char newvar = (char)var;
	if( var > 127 )
		newvar = 127;
		
	return newvar;
}

char DINT_TO_SINT(int var)
{
	char newvar;
	if( var > 127 )
		newvar = 127;
	else if( var < -127 )
		newvar = -127;
	else 
		newvar = (char)var;
		
	return newvar;
}

char LINT_TO_SINT(long var)
{
	char newvar;
	if( var >= 127 )
		newvar = 127;
	else if( var < -127 )
		newvar = -127;
	else 
		newvar = (char)var;
	return newvar;
}
/**************************************** TO INT ****************/
short UINT_TO_INT(unsigned short var)
{
	short newvar = (short)var;
	if( var > 32767 )
		newvar = 32767;
		
	return newvar;
}

short DINT_TO_INT(int var)
{
	short newvar;
	if( var > 32767 )
		newvar = 32767;
	else if( var < -32767 )
		newvar = -32767;
	else 
		newvar = (short)var;
		
	return newvar;
}

short LINT_TO_INT(long var)
{
	short newvar;
	if( var >= 32766 )
		newvar = 32766;
	else if( var < -32767 )
		newvar = -32767;
	else 
		newvar = (short)var;
	return newvar;
}

short REAL_TO_INT(float var)
{
	short newvar = (short)var;
	return newvar;
}
/**************************************** TO UINT ****************/
unsigned short LINT_TO_UINT(long var)
{
	unsigned short newvar;
	if( var >= 255 )
		newvar = 255;
	else if( var < 0 )
		newvar = 0;
	else 
		newvar = (unsigned short)var;
	return newvar;
}

unsigned short DINT_TO_UINT(int var)
{
	unsigned short newvar;
	if( var >= 255 )
		newvar = 255;
	else if( var < 0 )
		newvar = 0;
	else 
		newvar = (unsigned short)var;
	return newvar;
}

unsigned short SINT_TO_UINT(short var)
{
	unsigned short newvar;
	if( var <= 0 )
		newvar = 0;
	else 
		newvar = (unsigned short)var;
		
	return newvar;
}



/**************************************** TO LINT ****************/

long SINT_TO_LINT(short var)
{
	long newvar;
	newvar = (long)var;
	return newvar;
}

long UINT_TO_LINT(unsigned short var)
{
	long newvar;
	newvar = (long)var;
	return newvar;
}

long DINT_TO_LINT(int var)
{
	long newvar;
	newvar = (long)var;
	return newvar;
}

/**************************************** TO DINT ****************/
int INT_TO_DINT(short var)
{
	int newvar;
	newvar = (int)var;
	return newvar;
}

int LINT_TO_DINT(long var)
{
	int newvar;
	if( var >= 65535 )
		newvar = 65535;
	else 
		newvar = (int)var;
	return newvar;
}

int SINT_TO_DINT(short var)
{
	int newvar;
	newvar = (int)var;
	return newvar;
}

int UINT_TO_DINT(unsigned short var)
{
	int newvar;
	newvar = (int)var;
	return newvar;
}

int REAL_TO_DINT(float var)
{
	int newvar;
	newvar = (int)var;
	return newvar;
}

/**************************************** TO UDINT ***************/
unsigned int DINT_TO_UDINT(int var)
{
	unsigned int newvar;
	if( var < 0 )
		newvar = 0;
	else 
		newvar = (unsigned int)var;
	return newvar;

}

/**************************************** TO REAL ***************/
float INT_TO_REAL(short var)
{
	float newvar = (float)var;
	return newvar;
}

float DINT_TO_REAL(int var)
{
	float newvar = (float)var;
	return newvar;
}



/* ************************ Other Functions ***************/
short ABS(short var)
{
	short newvar = var;
	if( newvar < 0 )
		newvar *= -1;
	return newvar;
}

float SQRT(float var)
{
	return (float)sqrt(var);
}

float COS(float var)
{
	return (float)cos(var);
}

float SIN(float var)
{
	return (float)sin(var);
}

float ATAN(float var)
{
	return (float)atan(var);
}


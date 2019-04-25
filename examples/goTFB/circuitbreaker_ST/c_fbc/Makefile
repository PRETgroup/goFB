CC = gcc
CFLAGS = -Wall -O2
INCLUDES = 
LDFLAGS = 
LIBS = -lm

TARGET = BfbIDMTCurve
C_SOURCES =  \
            STFunctions.c  \
            BfbIDMTCurve.c 
RM = rm -f

.PHONY: all clean cleanall

all: $(TARGET)

$(TARGET): $(C_SOURCES) Makefile
	$(CC) $(CFLAGS) $(INCLUDES) $(C_SOURCES) $(LDFLAGS) $(LIBS) -o $(TARGET)

clean:
	$(RM) *.o

cleanall:
	$(RM) *.o
	$(RM) $(TARGET)
	$(RM) $(TARGET).exe

TCRESTHOMEDIR?=$(HOME)/t-crest
TCRESTINCLUDEDIR?=$(TCRESTHOMEDIR)/patmos/c
TCRESTLINKDIR?=$(TCRESTHOMEDIR)/patmos/tmp
TCRESTNOCINIT?=$(TCRESTHOMEDIR)/patmos/c/cmp/nocinit.c
FBINCLUDEDIR?=./include

PROJECT?=bottlingplant
#PROJECTINCLUDE?=./$(PROJECT)/include

BUILDDIR?=../tmp
LIBNOC=$(BUILDDIR)/libnoc.a
LIBMP=$(BUILDDIR)/libmp.a
LIBCORETHREAD=$(BUILDDIR)/libcorethread.a
LIBETH=$(BUILDDIR)/libeth.a
LIBELF=$(BUILDDIR)/libelf.a

#(not needed?) -I /home/hammond/t-crest/patmos/c/libelf/ 
# libiec61499/bottlingplant.*.c 
all:
%: ; patmos-clang \
	-target patmos-unknown-unknown-elf \
	-O2 \
	-I $(TCRESTINCLUDEDIR) \
	-I $(FBINCLUDEDIR) \
	-I ./$(PROJECT)/c \
	-mpatmos-disable-vliw \
	-mpatmos-method-cache-size=0x1000 \
	-mpatmos-stack-base=0x200000 \
	-mpatmos-shadow-stack-base=0x1f8000 \
	-Xgold --defsym -Xgold __heap_end=0x1f0000 \
	-Xgold -T \
	-Xgold $(TCRESTHOMEDIR)/patmos/hardware/spm_ram.t \
	-o /home/hammond/t-crest/patmos/tmp/$(PROJECT).elf \
	./$(PROJECT)/c/*.c \
	$(TCRESTNOCINIT) \
	-L $(TCRESTLINKDIR) \
	-lmp \
	-lnoc \
	-lcorethread \
	-leth \
	-lelf

#single:
#%: ; patmos-clang \
#	-target patmos-unknown-unknown-elf \
#	-O2 \
#	-I $(TCRESTINCLUDEDIR) \
#	-I /home/hammond/t-crest/patmos/c/libelf/ \
#	-mpatmos-disable-vliw \
#	-mpatmos-method-cache-size=0x1000 \
#	-mpatmos-stack-base=0x200000 \
#	-mpatmos-shadow-stack-base=0x1f8000 \
#	-Xgold --defsym -Xgold __heap_end=0x1f0000 \
#	-Xgold -T \
#	-Xgold $(TCRESTHOMEDIR)/patmos/hardware/spm_ram.t \
#	-o /home/hammond/t-crest/patmos/tmp/$(PROJECT).elf \
#	 echo.c cmp/nocinit.c -L/home/hammond/t-crest/patmos/tmp -lmp -lnoc -lcorethread -leth -lelf

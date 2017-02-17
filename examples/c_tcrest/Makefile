TCRESTHOMEDIR?=$(HOME)/t-crest
TCRESTINCLUDEDIR?=$(TCRESTHOMEDIR)/patmos/c
TCRESTLINKDIR?=$(TCRESTHOMEDIR)/patmos/tmp
TCRESTNOCINIT?=$(TCRESTHOMEDIR)/patmos/c/cmp/nocinit.c
FBINCLUDEDIR?=./include

PROJECT?=bottlingplant
FUNCTION?=main

BUILDDIR?=../tmp
LIBNOC=$(BUILDDIR)/libnoc.a
LIBMP=$(BUILDDIR)/libmp.a
LIBCORETHREAD=$(BUILDDIR)/libcorethread.a
LIBETH=$(BUILDDIR)/libeth.a
LIBELF=$(BUILDDIR)/libelf.a

#Delay to global memory per request in cycles
T_DELAY_CYCLES?=83
#Global memory transfer time per burst in cycles
G_DELAY_CYCLES?=0
#Transfer size (burst size) of the global memory in bytes
B_SIZE?=16


$(TCRESTHOMEDIR)/patmos/tmp/$(PROJECT).elf: ./$(PROJECT)/c/*.c ./$(PROJECT)/c/*.h
	patmos-clang \
	-target patmos-unknown-unknown-elf \
	-O2 \
	-I $(TCRESTINCLUDEDIR) \
	-I $(TCRESTINCLUDEDIR)/libelf/ \
	-I $(FBINCLUDEDIR) \
	-I ./$(PROJECT)/c \
	-mpatmos-disable-vliw \
	-mpatmos-method-cache-size=0x1000 \
	-mpatmos-stack-base=0x200000 \
	-mpatmos-shadow-stack-base=0x1f8000 \
	-Xgold --defsym -Xgold __heap_end=0x1f0000 \
	-Xgold -T \
	-Xgold $(TCRESTHOMEDIR)/patmos/hardware/spm_ram.t \
	-o $(TCRESTHOMEDIR)/patmos/tmp/$(PROJECT).elf \
	./$(PROJECT)/c/*.c \
	$(TCRESTNOCINIT) \
	-L $(TCRESTLINKDIR) \
	-lmp \
	-lnoc \
	-lcorethread \
	-leth \
	-lelf \
	-mserialize=$(TCRESTHOMEDIR)/patmos/tmp/$(PROJECT).pml

build: $(TCRESTHOMEDIR)/patmos/tmp/$(PROJECT).elf
	
wcet-config: 
	platin pml-config \
	--target patmos-unknown-unknown-elf \
	-o $(TCRESTHOMEDIR)/patmos/tmp/$(PROJECT)config.pml \
	--tdelay $(T_DELAY_CYCLES) \
	--gtime $(G_DELAY_CYCLES) \
	--bsize $(B_SIZE)

wcet: build wcet-config
	platin wcet \
	-i $(TCRESTHOMEDIR)/patmos/tmp/$(PROJECT).pml \
	--enable-wca \
	--disable-ait \
	--stats \
	--binary $(TCRESTHOMEDIR)/patmos/tmp/$(PROJECT).elf \
	-e $(FUNCTION) \
	-o $(PROJECT).wca \
	--report $(PROJECT)_wcet_report.txt \
	-i $(TCRESTHOMEDIR)/patmos/tmp/$(PROJECT)config.pml

c:	
	cd $(PROJECT); bash fbt_to_c.sh
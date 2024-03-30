CLANG   ?= clang
LLC     ?= llc
OPT     ?= opt
DIS     ?= llvm-dis
RM      ?= rm
MV      ?= mv
CC      ?= gcc
MAKE    ?= make
OBJDUMP ?= llvm-objdump

HEADER := -I./headers         \
         -I./kernel

CFLAGS_KV :=					\
	-g -O2 -Wall                            \
	-emit-llvm -target bpf                  \
	-D__KERNEL__                            \
	-D__TARGET_ARCH_x86_64                  \
	-D__NATIVE_EBPF__						\
	-D__x86_64__

all: connect.o

connect.o: connect.c
	$(CLANG) $(CFLAGS_KV) $(HEADER) -O2 -c $< -o -  | \
	$(LLC) -march=bpf -filetype=obj -o connect.o
	$(MV) connect.o ../output/

clean:
	$(RM) -rf ../output/connect.o
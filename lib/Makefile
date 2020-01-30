CC=gcc
FLAGS=-std=gnu99 -I. -I./sha3
CFLAGS=-fPIC -pedantic -Wall -Wextra -march=native
LDFLAGS=-shared

TARGET=libx16rv2.so
SOURCES = x16rv2.c $(shell echo sha3/*.c)
HEADERS = x16rv2.h $(shell echo sha3/*.h)
#OBJECTS = $(SOURCES:.c=.o)
OBJECTS = \
    x16rv2.o \
    ./sha3/fugue.o \
    ./sha3/aes_helper.o \
    ./sha3/blake.o \
    ./sha3/bmw.o \
    ./sha3/cubehash.o \
    ./sha3/echo.o \
    ./sha3/groestl.o \
    ./sha3/jh.o \
    ./sha3/keccak.o \
    ./sha3/luffa.o \
    ./sha3/shavite.o \
    ./sha3/simd.o \
    ./sha3/skein.o \
    ./sha3/whirlpool.o \
    ./sha3/shabal.o \
    ./sha3/hamsi.o \
    ./sha3/sha2big.o \
    ./sha3/tiger.o \

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) $(LDFLAGS) -o $(TARGET) $(OBJECTS)

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

clean:
	rm *.o sha3/*.o *.so

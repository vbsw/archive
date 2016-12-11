OBJECTS = iqueue.o
CC      = gcc
CFLAGS  = -Wall

libcds.a: $(OBJECTS)
	ar rs libcds.a $(OBJECTS);

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@



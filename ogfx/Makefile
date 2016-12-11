OBJECTS = window.o
CC      = gcc
CFLAGS  = -Wall

libogfx.a: $(OBJECTS)
	ar rs libogfx.a $(OBJECTS);

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	-rm $(OBJECTS) libogfx.a

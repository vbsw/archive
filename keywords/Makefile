HEADERS =
SOURCES = main.c
OUTPUT = keywords

keywords: $(HEADERS) $(SOURCES)
	gcc -O3 -Wall -g $(SOURCES) -o $(OUTPUT) `pkg-config --cflags --libs gtk+-2.0`

HEADERS = gui.h tui.h args.h settings.h strings.h output.h parse.h input.h buffer.h clock.h
SOURCES = main.c gui.c tui.c args.c settings.c strings.c output.c parse.c input.c buffer.c clock.c
OUTPUT = ched

ched: $(HEADERS) $(SOURCES)
	gcc -O3 -Wall -g $(SOURCES) -o $(OUTPUT) `pkg-config --cflags --libs gtk+-3.0`

all: window label button application layout notebook

window: window.c
	gcc -O3 -Wall -g window.c -o window `pkg-config --cflags --libs gtk+-3.0`

label: label.c
	gcc -O3 -Wall -g label.c -o label `pkg-config --cflags --libs gtk+-3.0`

button: button.c
	gcc -O3 -Wall -g button.c -o button `pkg-config --cflags --libs gtk+-3.0`

application: application.c
	gcc -O3 -Wall -g application.c -o application `pkg-config --cflags --libs gtk+-3.0`

layout: layout.c
	gcc -O3 -Wall -g layout.c -o layout `pkg-config --cflags --libs gtk+-3.0`

notebook: notebook.c
	gcc -O3 -Wall -g notebook.c -o notebook `pkg-config --cflags --libs gtk+-3.0`


HEADERS = proto.h buffer.h list.h strmap.h strmapbuf.h stack.h
SOURCES = proto.c buffer.c list.c strmap.c strmapbuf.c stack.c test.c


test: $(HEADERS) $(SOURCES)
	gcc -o test $(SOURCES)

doc: $(HEADERS) $(SOURCES)
	echo 'DOXYFILE_ENCODING=UTF-8\nPROJECT_NAME="cbase"\nOUTPUT_LANGUAGE=English\nRECURSIVE=YES\nSEPARATE_MEMBER_PAGES=YES\nREPEAT_BRIEF=NO\nOUTPUT_DIRECTORY=doc\nGENERATE_LATEX=NO\nLAYOUT_FILE=DoxygenLayout.xml' > Doxygen.conf
	echo '<doxygenlayout version="1.0"><navindex><tab type="mainpage" visible="yes" title=""/><tab type="filelist" visible="yes" title="" intro=""/><tab type="globals" visible="yes" title="" intro=""/></navindex></doxygenlayout>' > DoxygenLayout.xml
	doxygen Doxygen.conf
	rm Doxygen.conf DoxygenLayout.xml
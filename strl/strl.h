#ifndef  STRL_H
#define  STRL_H

#ifdef __cplusplus
extern "C" {
#endif

extern void*
strlnewcap ( size_t capacity );

extern void*
strlnew ( void );

extern void
strldest ( void* strl );

extern size_t
strllen ( const void* strl );

extern size_t
strlcap ( const void* strl );

extern unsigned short int
strlerr ( const void* strl );

extern void
strlerrclear ( void* strl );

extern void
strlcapfunc ( void* strl, size_t (*capfunc) ( size_t n, size_t capacity ) );

extern void
strlapnd ( void* strl, const char* str );

extern void
strlapndref ( void* strl, char* str );

extern char**
strlarr ( void* strl );

extern char*
strlrmref ( void* strl, size_t index );

extern void
strlrm ( void* strl, size_t index );

extern void
strlclear ( void* strl );

extern void
strlclearref ( void* strl );

extern void
strlins ( void* strl, size_t index, const char* str );

extern void
strlinsref ( void* strl, size_t index, char* str );

#ifdef __cplusplus
}
#endif

#endif /* STRL_H */
















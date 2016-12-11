#ifndef  SBIF_H
#define  SBIF_H


#include <stdlib.h>


extern void
SBiF_encodeSIZE
(
          void *const out_arr,
         const size_t in_SIZE
);


extern void
SBiF_decodeSIZE
(
        size_t *const out_SIZE,
    const void *const in_arr
);


extern void
SBiF_encodeSIZEArray
(
               void * out_arr,
       const size_t * in_SIZEArr,
         const size_t in_encodeSize
);


extern void
SBiF_decodeSIZEArray
(
             size_t * out_SIZEArr,
         const void * in_arr,
         const size_t in_decodeSize
);


#endif /* SBIF_H */
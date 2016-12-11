
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file sbif.c
 * @author Vitali Baumtrok
 * @date 2015-06-27
 * @brief reference implementation for the SBiF parser
 * 
 * Reference implementation for the SBiF parser.
 */

#include <stdlib.h>


/**
 * @brief Encodes a SIZE number to byte array.
 * 
 * Encodes a SIZE number to byte array.
 * 
 * @param out_arr The array to write the encoded SIZE number to.
 * @param in_SIZE The SIZE number to encode.
 */
void
SBiF_encodeSIZE
(
          void *const out_arr,
         const size_t in_SIZE
)
{
    ((unsigned char *) out_arr)[0] = (unsigned char) (in_SIZE >> 56);
    ((unsigned char *) out_arr)[1] = (unsigned char) (in_SIZE >> 48);
    ((unsigned char *) out_arr)[2] = (unsigned char) (in_SIZE >> 40);
    ((unsigned char *) out_arr)[3] = (unsigned char) (in_SIZE >> 32);
    ((unsigned char *) out_arr)[4] = (unsigned char) (in_SIZE >> 24);
    ((unsigned char *) out_arr)[5] = (unsigned char) (in_SIZE >> 16);
    ((unsigned char *) out_arr)[6] = (unsigned char) (in_SIZE >>  8);
    ((unsigned char *) out_arr)[7] = (unsigned char) (in_SIZE >>  0);
}


/**
 * @brief Decodes a SIZE number from an 8 byte array.
 * 
 * Decodes a SIZE number from an 8 byte array.
 * 
 * @param out_SIZE The resulting SIZE number.
 * @param in_arr The 8 byte array to decode.
 */
void
SBiF_decodeSIZE
(
        size_t *const out_SIZE,
    const void *const in_arr
)
{
    *out_SIZE =
        (size_t) ( ( (size_t) ((unsigned char *) in_arr)[0] ) << 56 ) +
        (size_t) ( ( (size_t) ((unsigned char *) in_arr)[1] ) << 48 ) +
        (size_t) ( ( (size_t) ((unsigned char *) in_arr)[2] ) << 40 ) +
        (size_t) ( ( (size_t) ((unsigned char *) in_arr)[3] ) << 32 ) +
        (size_t) ( ( (size_t) ((unsigned char *) in_arr)[4] ) << 24 ) +
        (size_t) ( ( (size_t) ((unsigned char *) in_arr)[5] ) << 16 ) +
        (size_t) ( ( (size_t) ((unsigned char *) in_arr)[6] ) <<  8 ) +
        (size_t) ( ( (size_t) ((unsigned char *) in_arr)[7] ) <<  0 );
}


/**
 * @brief Encodes a SIZE number array to a byte array.
 * 
 * Encodes a SIZE number array to a byte array.
 * 
 * @param out_arr The array to write the encoded SIZE numbers to.
 * @param in_SIZEArr The SIZE number array to encode.
 * @param in_encodeSize The number of bytes to encode.
 */
void
SBiF_encodeSIZEArray
(
               void * out_arr,
       const size_t * in_SIZEArr,
         const size_t in_encodeSize
)
{
    const size_t encodeSizeModulo = in_encodeSize % sizeof(size_t);
    const size_t *const limit = (const size_t *const) (((size_t) in_SIZEArr) + in_encodeSize - encodeSizeModulo);

    while ( in_SIZEArr < limit )
    {
        const size_t currSIZE = *in_SIZEArr;
        ((unsigned char *) out_arr)[0] = (unsigned char) (currSIZE >> 56);
        ((unsigned char *) out_arr)[1] = (unsigned char) (currSIZE >> 48);
        ((unsigned char *) out_arr)[2] = (unsigned char) (currSIZE >> 40);
        ((unsigned char *) out_arr)[3] = (unsigned char) (currSIZE >> 32);
        ((unsigned char *) out_arr)[4] = (unsigned char) (currSIZE >> 24);
        ((unsigned char *) out_arr)[5] = (unsigned char) (currSIZE >> 16);
        ((unsigned char *) out_arr)[6] = (unsigned char) (currSIZE >>  8);
        ((unsigned char *) out_arr)[7] = (unsigned char) (currSIZE >>  0);
		out_arr = (void*) ( ((size_t) out_arr) + 8 );
        in_SIZEArr++;
    }

    if ( encodeSizeModulo )
    {
        const size_t currSIZE = *in_SIZEArr;

        if ( encodeSizeModulo == 1 )
        {
            ((unsigned char *) out_arr)[0] = (unsigned char) (currSIZE >> 56);
        }
        else if ( encodeSizeModulo == 2 )
        {
            ((unsigned char *) out_arr)[0] = (unsigned char) (currSIZE >> 56);
            ((unsigned char *) out_arr)[1] = (unsigned char) (currSIZE >> 48);
        }
        else if ( encodeSizeModulo == 3 )
        {
            ((unsigned char *) out_arr)[0] = (unsigned char) (currSIZE >> 56);
            ((unsigned char *) out_arr)[1] = (unsigned char) (currSIZE >> 48);
            ((unsigned char *) out_arr)[2] = (unsigned char) (currSIZE >> 40);
        }
        else if ( encodeSizeModulo == 4 )
        {
            ((unsigned char *) out_arr)[0] = (unsigned char) (currSIZE >> 56);
            ((unsigned char *) out_arr)[1] = (unsigned char) (currSIZE >> 48);
            ((unsigned char *) out_arr)[2] = (unsigned char) (currSIZE >> 40);
            ((unsigned char *) out_arr)[3] = (unsigned char) (currSIZE >> 32);
        }
        else if ( encodeSizeModulo == 5 )
        {
            ((unsigned char *) out_arr)[0] = (unsigned char) (currSIZE >> 56);
            ((unsigned char *) out_arr)[1] = (unsigned char) (currSIZE >> 48);
            ((unsigned char *) out_arr)[2] = (unsigned char) (currSIZE >> 40);
            ((unsigned char *) out_arr)[3] = (unsigned char) (currSIZE >> 32);
            ((unsigned char *) out_arr)[4] = (unsigned char) (currSIZE >> 24);
        }
        else if ( encodeSizeModulo == 6 )
        {
            ((unsigned char *) out_arr)[0] = (unsigned char) (currSIZE >> 56);
            ((unsigned char *) out_arr)[1] = (unsigned char) (currSIZE >> 48);
            ((unsigned char *) out_arr)[2] = (unsigned char) (currSIZE >> 40);
            ((unsigned char *) out_arr)[3] = (unsigned char) (currSIZE >> 32);
            ((unsigned char *) out_arr)[4] = (unsigned char) (currSIZE >> 24);
            ((unsigned char *) out_arr)[5] = (unsigned char) (currSIZE >> 16);
        }
        else /* encodeSizeModulo == 7 */
        {
            ((unsigned char *) out_arr)[0] = (unsigned char) (currSIZE >> 56);
            ((unsigned char *) out_arr)[1] = (unsigned char) (currSIZE >> 48);
            ((unsigned char *) out_arr)[2] = (unsigned char) (currSIZE >> 40);
            ((unsigned char *) out_arr)[3] = (unsigned char) (currSIZE >> 32);
            ((unsigned char *) out_arr)[4] = (unsigned char) (currSIZE >> 24);
            ((unsigned char *) out_arr)[5] = (unsigned char) (currSIZE >> 16);
            ((unsigned char *) out_arr)[6] = (unsigned char) (currSIZE >>  8);
        }
    }
}


/**
 * @brief Decodes SIZE numbers from an array.
 * 
 * Decodes SIZE numbers from an array.
 * 
 * @param out_SIZEArr The resulting (decoded) SIZE numbers.
 * @param in_arr The array of SIZE numbers to decode.
 * @param in_decodeSize The number of bytes to decode.
 */
void
SBiF_decodeSIZEArray
(
             size_t * out_SIZEArr,
         const void * in_arr,
         const size_t in_decodeSize
)
{
    const size_t decodeSizeModulo = in_decodeSize % sizeof(size_t);
    const size_t *const limit = (const size_t *const) (((size_t) out_SIZEArr) + in_decodeSize - decodeSizeModulo);

    while ( out_SIZEArr < limit )
    {
        *out_SIZEArr =
            (size_t) ( ( (size_t) ((unsigned char *) in_arr)[0] ) << 56 ) +
            (size_t) ( ( (size_t) ((unsigned char *) in_arr)[1] ) << 48 ) +
            (size_t) ( ( (size_t) ((unsigned char *) in_arr)[2] ) << 40 ) +
            (size_t) ( ( (size_t) ((unsigned char *) in_arr)[3] ) << 32 ) +
            (size_t) ( ( (size_t) ((unsigned char *) in_arr)[4] ) << 24 ) +
            (size_t) ( ( (size_t) ((unsigned char *) in_arr)[5] ) << 16 ) +
            (size_t) ( ( (size_t) ((unsigned char *) in_arr)[6] ) <<  8 ) +
            (size_t) ( ( (size_t) ((unsigned char *) in_arr)[7] ) <<  0 );
		in_arr = (const void*) ( ((size_t) in_arr) + 8 );
        out_SIZEArr++;
    }

    if ( decodeSizeModulo )
    {
        if ( decodeSizeModulo == 1 )
        {
            *out_SIZEArr =
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[0] ) << 56 );
        }
        else if ( decodeSizeModulo == 2 )
        {
            *out_SIZEArr =
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[0] ) << 56 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[1] ) << 48 );
        }
        else if ( decodeSizeModulo == 3 )
        {
            *out_SIZEArr =
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[0] ) << 56 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[1] ) << 48 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[2] ) << 40 );
        }
        else if ( decodeSizeModulo == 4 )
        {
            *out_SIZEArr =
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[0] ) << 56 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[1] ) << 48 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[2] ) << 40 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[3] ) << 32 );
        }
        else if ( decodeSizeModulo == 5 )
        {
            *out_SIZEArr =
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[0] ) << 56 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[1] ) << 48 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[2] ) << 40 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[3] ) << 32 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[4] ) << 24 );
        }
        else if ( decodeSizeModulo == 6 )
        {
            *out_SIZEArr =
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[0] ) << 56 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[1] ) << 48 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[2] ) << 40 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[3] ) << 32 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[4] ) << 24 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[5] ) << 16 );
        }
        else /* decodeSizeModulo == 7 */
        {
            *out_SIZEArr =
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[0] ) << 56 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[1] ) << 48 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[2] ) << 40 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[3] ) << 32 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[4] ) << 24 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[5] ) << 16 ) +
                (size_t) ( ( (size_t) ((unsigned char *) in_arr)[6] ) <<  8 );
        }
    }
}


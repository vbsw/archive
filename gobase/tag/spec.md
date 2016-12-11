# Tag Specification

## Descriptions Layout
The descriptions consitsts of three arrays:
- descriptions data,
- descriptions offsets,
- descriptions lengths.

The *descriptions data* is a consecutive sequence of descriptions. A description is a byte array. Descriptions are sorted incrementaly.

The *descriptions offsets* is an integer array with as many entries as descriptions in *descriptions data*. Each entry points to the beginning of a description.

The *descriptions lengths* is an integer array with as many entries as descriptions in *descriptions data*. Each entry represents the length of a description. The index of a length corresponds to the index of an offset for the same description in *descriptions data*.

## Words Layout
Words have the same layout as descriptions.

## Tags Layout
Tags have the same layout as descriptions.

## Maps Layout
The maps consists of three arrays:
- maps data,
- maps offsets,
- maps lengths.

The *maps data* is a consecutive sequence of maps. A map is an integer array. Integers in a map are sorted incrementaly.

The *maps offsets* is an integer array with as many entries as maps in *maps data*. Each entry points to the beginning of a map.

The *maps lengths* is an integer array with as many entries as maps in *maps data*. Each entry represents the length of a map. The index of a length corresponds to the index of an offset for the same map in *maps data*.

The number of maps is the same as the number of tags. The index of a map corresponds to the index of a tag. Each tag has a map, even if the length of the map is zero. The entries of a map represents the indices of words and descriptions, i.e. they reference them.

## File Encoding
Every integer is encoded as a 32 bit value in a file, i.e. 4 bytes.

The file consists of five parts:
- encoded parts lengths,
- encoded descriptions,
- encoded words,
- encoded tags,
- encoded maps.

The *encoded parts lengths* consists of four integers. First integer represents the size of the *encoded descriptions*, second integer represents the size of the *encoded words*, third represents the size of the *encoded tags* and fourth represents the size of the *encoded maps* in bytes.

The *encoded descriptions* consists of one integer, *descriptions offsets*, *descriptions lengths* and *descriptions data*. The integer represents the number of descriptions.

The *encoded words* consists of one integer, *words offsets*, *words lengths* and *words data*. The integer represents the number of words.

The *encoded tags* consists of one integer, *tags offsets*, *tags lengths* and *tags data*. The integer represents the number of tags.

The *encoded maps* consists of one integer, *maps offsets*, *maps lengths* and *maps data*. The integer represents the number of maps.

## Data Overhead
The total overhead of stored data is 32 bytes, i.e. 8 integers of 32 bits.
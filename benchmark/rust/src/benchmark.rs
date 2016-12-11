/*
* This is free and unencumbered software released into the public domain.
*
* Anyone is free to copy, modify, publish, use, compile, sell, or
* distribute this software, either in source code form or as a compiled
* binary, for any purpose, commercial or non-commercial, and by any
* means.
*
* In jurisdictions that recognize copyright laws, the author or authors
* of this software dedicate any and all copyright interest in the
* software to the public domain. We make this dedication for the benefit
* of the public at large and to the detriment of our heirs and
* successors. We intend this dedication to be an overt act of
* relinquishment in perpetuity of all present and future rights to this
* software under copyright law.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
* EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
* MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
* IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
* OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
* ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
* OTHER DEALINGS IN THE SOFTWARE.
*
* For more information, please refer to <http://unlicense.org>
*/


extern crate time;

use std::fs::File;
use std::io::Read;


const INSERT_REMOVE_TEST_NAME:      &'static str     = "insert/remove";/*
const INSERT_REMOVE_MOD_TEST_NAME:  &'static str     = "insert/remove (mod)";*/
const BINARY_SEARCH_TEST_NAME:      &'static str     = "binary search";
const FILE_NAMES:                  [&'static str; 9] =
[
    "benchmark.parameters.txt",
    "c/benchmark.parameters.txt",
    "c/Release/benchmark.parameters.txt",
    "benchmark/benchmark.parameters.txt",
    "benchmark/c/benchmark.parameters.txt",
    "benchmark/c/Release/benchmark.parameters.txt",
    "../benchmark.parameters.txt",
    "../../benchmark.parameters.txt",
    "../../../benchmark.parameters.txt"
];
const NUMBER_OF_ITERATIONS_STR:     &'static str     = "number of iterations";
const SIZE_OF_ARRAY_STR:            &'static str     = "length of array";
const FILE_NAME:                    &'static str     = FILE_NAMES[0];
const DEFAULT_ITERATION_NUMBER:     i32              = 5;
const DEFAULT_ARRAY_LENGTH:         i32              = 32768 - 2;


fn print_benchmark_header () {
    println! ( "\n" );
    println! ( "/////////////////////////////////////////////////////////////////" );
    println! ( "//                                                             //" );
    println! ( "//                    Rust benchmark                           //" );
    println! ( "//                                                             //" );
    println! ( "/////////////////////////////////////////////////////////////////" );
}

fn read_file () -> Vec<u8> {
    for i in 0..FILE_NAMES.len () {
        let file_name = FILE_NAMES[i];
        let file = File::open ( file_name );

        match file {
            Ok ( mut f ) => {
                 let mut file_bytes: Vec<u8> = Vec::new ();
                 let bytes_read = f.read_to_end ( &mut file_bytes );

                 match bytes_read {
                     Ok ( _ ) => {
                         return file_bytes;
                     }
                     Err ( e ) => {
                         println!("error reading existing file: {:?}", e);
                         return Vec::new ();
                     }
                 }
            }
            Err ( _ ) => {
            }
        }
    }

    return Vec::new ();
}

fn parse_iter_number ( bytes: &Vec<u8> ) -> i32 {
    if bytes.len () > 0 {
        let mut b: u8 = bytes[0];
        let mut result: i32 = -1;

        if ( b > 0x2F ) && ( b < 0x3A ) {
            result = ( b - 0x30 ) as i32;

            for i in 1..bytes.len () {
                b = bytes[i];

                if (b > 0x2F) && (b < 0x3A) {
                    result = result * 10 + ( b - 0x30 ) as i32;
                }
                else {
                    break;
                }
            }
        }
        return result;
    }
    else {
        return -1;
    }
}

fn parse_array_size ( bytes: &Vec<u8> ) -> i32 {
    for i in 0..bytes.len () {

        if bytes[i] == 0x0A {

            if i+1 < bytes.len () {
                let mut b: u8 = bytes[i+1];
                let mut result: i32 = -1;

                if ( b > 0x2F ) && ( b < 0x3A ) {
                    result = ( b - 0x30 ) as i32;

                    for j in i+2..bytes.len () {
                        b = bytes[j];

                        if (b > 0x2F) && (b < 0x3A) {
                            result = result * 10 + ( b - 0x30 ) as i32;
                        }
                        else {
                            break;
                        }
                    }
                }
                return result;
            }
            else {
                return -1;
            }
        }
    }

    return -1;
}

fn parse_parameters ( parameters: &Vec<u8>, parsed_iteration_number: &mut i32, parsed_array_size: &mut i32 ) {
    *parsed_iteration_number = parse_iter_number ( parameters );
    *parsed_array_size = parse_array_size ( parameters );
}

fn set_default_settings ( iter_number: &mut i32, array_length: &mut i32 ) {
    *iter_number = DEFAULT_ITERATION_NUMBER;
    *array_length = DEFAULT_ARRAY_LENGTH;
}

fn assign_settings ( parameters: &Vec<u8>, parsed_iteration_number: &mut i32, parsed_array_size: &mut i32, iter_number: &mut i32, array_length: &mut i32 ) {
    if parameters.is_empty () {
        set_default_settings ( iter_number, array_length );
    }
    else {
        parse_parameters ( parameters, parsed_iteration_number, parsed_array_size );
    }
}

fn correct_settings ( parsed_iteration_number: i32, parsed_array_size: i32, iter_number: &mut i32, array_length: &mut i32 ) {
    if parsed_iteration_number > 0 {
        *iter_number = parsed_iteration_number;
    }
    else {
        *iter_number = DEFAULT_ITERATION_NUMBER;
    }

    if parsed_array_size > 0 {
        *array_length = parsed_array_size;
    }
    else {
        *array_length = DEFAULT_ARRAY_LENGTH;
    }
}

fn print_notice_header () {
    println! ( "//                                                             //" );
    println! ( "//                          Notice                             //" );
    println! ( "//                                                             //" );
}

fn print_notice_some_unparsed_values () {
    println! ( "// Some parameters could not be parsed and are set to default. //" );
}

fn print_notice_some_too_low_values () {
    println! ( "// Some parameters are too low and are set to default.         //" );
}

fn print_notice_footer () {
    println! ( "//                                                             //" );
    println! ( "/////////////////////////////////////////////////////////////////" );
}

fn print_notice_couldnt_read_file () {
    println! ( "// Could not read file '{}{}", FILE_NAME, "'.             //" );
}

fn print_notice ( file_read: bool, parsed_iteration_number: i32, parsed_array_size: i32 ) {
    if file_read {
        if parsed_iteration_number <= 0 || parsed_array_size <= 0 {
            print_notice_header ();

            if ( parsed_iteration_number < 0 ) || ( parsed_array_size < 0 ) {
                print_notice_some_unparsed_values ();
            }

            if ( parsed_iteration_number == 0 ) || ( parsed_array_size == 0 ) {
                print_notice_some_too_low_values ();
            }
            print_notice_footer ();
        }
    }
    else {
        print_notice_header ();
        print_notice_couldnt_read_file ();
        print_notice_footer ();
    }
}

fn print_parameters_header () {
    println! ( "//                                                             //" );
    println! ( "//                        Parameters                           //" );
    println! ( "//                                                             //" );
}

fn print_parameter ( param_name: &str, valid_param: bool, param_value: i32, default_param_value: i32 ) {
    if valid_param {
        println! ( "// {:25}:{:12}                      //", param_name, param_value );
    }
    else {
        println! ( "// {:25}:{:12} (default value)      //", param_name, default_param_value );
    }
}

fn print_parameters_footer () {
    println! ( "//                                                             //" );
    println! ( "/////////////////////////////////////////////////////////////////" );
}

fn print_parameters ( file_read: bool, parsed_iteration_number: i32, parsed_array_size: i32 ) {
    print_parameters_header ();
    print_parameter ( NUMBER_OF_ITERATIONS_STR, file_read && (parsed_iteration_number>0), parsed_iteration_number, DEFAULT_ITERATION_NUMBER );
    print_parameter ( SIZE_OF_ARRAY_STR, file_read && (parsed_array_size>0), parsed_array_size, DEFAULT_ARRAY_LENGTH );
    print_parameters_footer ();
}

fn read_parameters ( iter_number: &mut i32, array_length: &mut i32 ) {
    let file_bytes: Vec<u8> = read_file ();
    let mut parsed_iteration_number: i32 = 0;
    let mut parsed_array_size: i32 = 0;

    assign_settings ( &file_bytes, &mut parsed_iteration_number, &mut parsed_array_size, iter_number, array_length );
    correct_settings ( parsed_iteration_number, parsed_array_size, iter_number, array_length );
    print_notice ( !file_bytes.is_empty (), parsed_iteration_number, parsed_array_size );
    print_parameters ( !file_bytes.is_empty (), parsed_iteration_number, parsed_array_size );
}

fn print_results_header () {
    println! ( "//                                                             //" );
    println! ( "//                         Results                             //" );
    println! ( "//                                                             //" );
}

fn get_time_in_seconds () -> f64 {
    return time::precise_time_s ();
}

fn print_test_time ( test_name: &str, passed_seconds: f64 ) {
    println! ( "// {:25}:{:12.3} seconds              //", test_name, passed_seconds );
}

fn benchmark_insert_remove ( iter_number: i32, array_length: i32 ) {
    let mut arr: Vec<i32> = Vec::new ();
    let start_time: f64 = get_time_in_seconds ();
    let arr_limit: i32 = array_length - 1;

    for _ in 0..iter_number {
        let mut flag: i32 = 0;

        // inserting
        for curr_length in 0..arr_limit {

            if flag != 0 {
                arr.insert ( 1, curr_length );
                flag = 0;
            }
            else {
                arr.insert ( 0, curr_length );
                flag = 1;
            }
        }
        flag = array_length % 2;

        // removing
        for _ in 0..arr_limit {

            if flag != 0 {
                arr.remove ( 1 );
                flag = 0;
            }
            else {
                arr.remove ( 0 );
                flag = 1;
            }
        }
    }
    print_test_time ( INSERT_REMOVE_TEST_NAME, ( get_time_in_seconds () - start_time ) / iter_number as f64 );
}

fn new_word ( length: usize, value: u8 ) -> Vec<u8> {
    let mut chars: Vec<u8> = Vec::with_capacity ( length );
    for _ in 0..chars.capacity () {
        chars.push ( value );
    }
    return chars;
}

fn allocate_words ( words: &mut Vec<Vec<u8>> ) {
    let max_value: usize = 127;
    let alphabet_array_length: usize = words.capacity () + max_value - words.capacity () % max_value;
    let step_length: usize = alphabet_array_length / max_value;
    let mut i: usize = 0;

    while i < words.capacity () {
        let j_limit: usize = i + step_length;
        let number: u8 = ( i / step_length + 1 ) as u8;
        let mut j: usize = i;

        while (j < j_limit) && (j < words.capacity ()) {
            let word_length: usize = (j % max_value) + 1;
            words.push ( new_word ( word_length, number ) );
            j = j + 1;
        }
        i = j;
    }
}

fn allocate_words_to_search ( words: &mut Vec<Vec<u8>> ) {
    let mut i: usize = 0;

    while i < words.capacity () {
        let mut j: i32 = -2;

        while (j < 128) && (i < words.capacity ()) {
            words.push ( new_word ( 10, j as u8 ) );
            j = j + 3;
            i = i + 1;
        }
    }
}

fn benchmark_binary_search ( iter_number: i32, array_length: i32 ) {
    let words_to_search_length: usize = 32766;
    let mut words: Vec<Vec<u8>> = Vec::with_capacity ( array_length as usize );
    let mut words_to_search: Vec<Vec<u8>> = Vec::with_capacity ( words_to_search_length );

    allocate_words ( &mut words );
    allocate_words_to_search ( &mut words_to_search );

    let start_time: f64 = get_time_in_seconds ();
    let a_value: u8 = words_to_search[0][0];
    let b_value: u8 = ( a_value * 2 ) as u8;

    for _ in 0..iter_number {
        for j in 0..words_to_search_length {
            let result = words.binary_search ( &(words_to_search[j]) );

            match result {
                Ok ( _ ) => {
                    words_to_search[0][0] = a_value;
                }
                Err ( _ ) => {
                    words_to_search[0][0] = b_value;
                }
            }
        }
    }
    print_test_time ( BINARY_SEARCH_TEST_NAME, ( get_time_in_seconds () - start_time ) / iter_number as f64 );
}

fn benchmark ( iter_number: i32, array_length: i32 ) {
    benchmark_insert_remove ( iter_number, array_length );/*
    benchmark_insert_remove_mod ( iter_number, array_length );*/
    benchmark_binary_search ( iter_number, array_length );
}

fn print_results_footer () {
    println! ( "//                                                             //" );
    println! ( "/////////////////////////////////////////////////////////////////" );
}

fn main () {
    let mut iter_number: i32 = 0;
    let mut array_length: i32 = 0;

    print_benchmark_header ();
    read_parameters ( &mut iter_number, &mut array_length );
    print_results_header ();
    benchmark ( iter_number, array_length );
    print_results_footer ();
    println! ( "\n" );
}

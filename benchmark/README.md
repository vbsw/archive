# benchmark

## Abstract
This benchmark implements the same functionality for various languages
in different ways. The languages are C, Java, D, Rust, Chapel, Go.

Remember, when you run the benchmarks, you messure the level of code
optimization of your tools/compiler and not of the language itself.

## Copying
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org>

## Clone Repository
[Download](http://git-scm.com/downloads) and install Git.

Clone this repository to folder "benchmark":

	$ git clone https://github.com/vitalibaumtrok/benchmark benchmark

## Compiling

### C
On __Windows__ [download Visual Studio Express] (https://www.visualstudio.com/en-us/products/visual-studio-express-vs.aspx)
(Community Edition, or something comparable)
and install it. Dobble click on file _benchmark.vcxproj_.
The project opens in _Visual Studio_. In the toolbar select the
_Release_ configuration, then select the project and press F7.
The executable _benchmark.exe_ is created in folder _Release_.
Dobble click on it to execute.

On __Linux__ to compile and execute run something like

	$ gcc main.linux.c helper.c -o benchmark
	$ ./benchmark

### Java

[Download Eclipse](https://eclipse.org/downloads/), start it, create new
project, drag-n-drop the file _Benchmark.java_ into the _src_ folder of
your project and click on the run button.

Or compile with

	$ javac Benchmark.java

and run with

	$ java Benchmark

### D

[Download DMD](http://dlang.org/download.html) and install it. To compile
and execute run something like

	$ dmd Benchmark.d -o benchmark
	$ ./benchmark

### Rust

[Download Rust](http://www.rust-lang.org/install.html) and install it.

To compile/run go to directory _"name of the project"/rust/_ (where Cargo.toml is) and run

	$ cargo run --release

Note: When executing the first time, some files will be downloaded.

### Chapel

Not available, yet.

### Go

Not available, yet.

## Results

	              insert/remove   binary search   LOC   MH
	C (vs)            0.045 s         0.003 s
	C (gcc)           0.019 s         0.010 s     689  12.0
	Java (JDT)        0.116 s         0.004 s     410   3.5
	D (DMD)           2.775 s        11.300 s     351   4.2
	Rust (rustc)      0.117 s         0.004 s     416   6.0

The idea was to test how much time you win programming in a modern language
compared to an old one (language C) and how much differs the execution time
of that program, if you "just write" the code down and "just compile" it.
(MH means man-hour, i.e. the time needed to write the code)

__Conclusion__: D seems not to be suitable for beginners. It seems you need to
know the right compiler, the right compile flags, the right algorithms and
so on, before you can write fast code in D. Rust has unusual syntax and memory
management. You need more time to become occustomed to it.

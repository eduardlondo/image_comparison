# Image Search Program using Histogram Intersection Approach

This program is designed to search for repeated images within a folder using the Histogram Intersection approach. It is written in the Go programming language. To compile and run the program, you need to have Go installed on your machine.

## Prerequisites

Go installed on your machine

https://go.dev/doc/install

## Compilation

1. Open a terminal or command prompt.
2. Run the following command inside the src folder to compile the program:

 `go build -o image_comparison.exe .` 

This command will generate an executable file named image_comparison.exe.

## Usage

The compiled executable file accepts two command-line arguments:

1. Image Route: The route to the image file you want to search for repeated instances of.
2. Folder Route: The route to the folder containing the images to be searched.

To execute the program, open a terminal or command prompt, navigate to the directory containing the compiled image_search.exe file, and run the following command:

`image_search.exe <image_route> <folder_route>`

Replace <image_route> with the route to the image file you want to search for repetitions of, and <folder_route> with the route to the folder containing the images to be analized.
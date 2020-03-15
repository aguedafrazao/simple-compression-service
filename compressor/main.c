#include <stdio.h>
#include <stdlib.h>
#include "HuffmanApplication.h"

int main(int argc, char** argv) {
	if(argc != 3) {
		printf("missing arguments: ./app {inputFile} {outputFile}\n");
		return 1;
	}
	char* inputFile = argv[1];
	char* outputFile = argv[2];
    	int result = onCompress(inputFile, outputFile);
	if(result != 0)
		return 1;
	return 0;
}


#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "HuffmanApplication.h"

int main(int argc, char **argv)
{
	if (argc != 4)
	{
		printf("missing arguments: ./app {option C (compress) or D(decompress)} {inputFile} {outputFile}\n");
		return 1;
	}
	char *option = argv[1];
	char *inputFile = argv[2];
	char *outputFile = argv[3];
	int result;
	int compare = strcmp(option, "C");
	if (compare == 0)
		result = onCompress(inputFile, outputFile);
	else
		result = onDecompress(inputFile, outputFile);
	if (result != 0)
		return 55;
	return 0;
}

#include "HuffmanHandler.h"
#include "HuffmanTree.h"
#include "List.h"
#include "utils.h"
#include "PriorityQueue.h"
#include "onCompressUtil.h"
#include "onDecompressUtil.h"
#include "Header.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define DEBUG if(0)

/**********************************************************
            Contract's functions implementation
***********************************************************/
int onCompress(char* inputPathFile, char* outputPathFile) {
    FILE* inputFile = fopen(inputPathFile, "rb"); //opening inputfile
	if(!inputFile) {
		printf("failed to open file %s\n", inputPathFile);
		return 1;
	}	
    int* bytesFrequency = getBytesFrequency(inputFile); //getting bytes frenquency
    fseek(inputFile, 0, SEEK_SET); //because we've gone through the file, so get back to start
    HuffmanTree* tree = buildHuffmanTree(bytesFrequency); //building huffman tree
    byte** matrixPath = buildPaths(tree); //building the matrix that helps to handle bytes
    strcat(outputPathFile, VALID_EXTENSION); //appending .huff to given output name
	FILE* outputFile = fopen(outputPathFile,"wb"); //opening output file 
	int treeSize = getTreeSize(tree); //getting tree size
    Header* header = getHeaderInfo(matrixPath, treeSize, inputFile); //creating the header
    fseek(inputFile, 0, SEEK_SET); //because we've gone through the file, so get back to start
    writeSources(header, tree, matrixPath, outputFile, inputFile); //writes header, tree, and matrix
    fclose(inputFile);
    fclose(outputFile);
	return 0;
}

int onDecompress(char* inputPathFile, char* outputPathFile) {
   	if(!isValidFile(inputPathFile)) {
		printf("given file to decompress is not valid\n");
		return 1;
	}
   	FILE* inputFile = fopen(inputPathFile, "rb");
   	if(!inputFile) {
   		printf("failed to open file %s\n", inputPathFile);
        return 1;
	}
    byte firstByte, secondByte;
    fscanf(inputFile, "%c", &firstByte); //getting first byte
    int trash = getTrash(firstByte);  //getting trash
    fscanf(inputFile, "%c", &secondByte);  //getting second byte
    int treeSize = retrieveTreeSize(firstByte, secondByte);  //get size tree
    byte* treeBytes = huffmanTreeBytes(inputFile, treeSize);
   	HuffmanTree* tree = reassembleHuffmanTree(treeBytes, treeSize); //reassembling huffman tree from its bytes
   	FILE* outputFile = fopen(outputPathFile, "wb");
   	rewriteOriginal(tree, trash, inputFile, outputFile); //creating output file
   	fclose(inputFile);
   	fclose(outputFile);
}

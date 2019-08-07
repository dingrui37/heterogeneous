#include <stdio.h>
#include <string.h>
#include "libarbitrator.h"

int main(int argc, char* argv[]) {
    char* cfgFile = "../config.json";
    int length = strlen(cfgFile);

    GoString value = {
        .p = cfgFile,
        .n = length
    };

    ArbitratorInit(value);

    int param1[] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    int param2[] = {2, 3, 4, 5, 6, 7, 8, 9, 10,11};
    int i = 0;
    for(; i < sizeof(param1) / sizeof(int); i++) {
        ArbitratorAdd(param1[i],param2[i]);
    }
}
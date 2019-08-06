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
    ArbitratorAdd(1,2);
}
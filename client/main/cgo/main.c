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

    //测试用例
    int param1[] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    int param2[] = {2, 3, 4, 5, 6, 7, 8, 9, 10,11};
    int i = 0;
    for(; i < sizeof(param1) / sizeof(int); i++) {
        ArbitratorAdd(param1[i], param2[i]);
    }
}

// dingrui@ubuntu:~/projects/heterogeneous/client/main/cgo$ ./main
// 2019/08/07 10:09:00 New sigleton Arbitrator success
// 2019/08/07 10:09:00 Parse config file:../config.json success, configs:&{0xc000312600 0xc00027d540}
// 2019/08/07 10:09:00 Container started, id: 26b912fc1d8b8f5f57b37c22050a2ab96565d9cad1c73476a401cbbcf27f526b, image: heterogeneous_python:v1.0.0, ServicePort: 50051
// 2019/08/07 10:09:01 Container started, id: 37aa285695caee4f477fb38ba25449d69d4e6b95798075d7cc42f88a66fc0d2c, image: heterogeneous_java:v1.0.0, ServicePort: 50052
// 2019/08/07 10:09:01 Container started, id: 4e357c9125ad90fe3ee9997fd839b5daff9bc6e33e3e0f748fa80193d307f9e8, image: heterogeneous_cpp:v1.0.0, ServicePort: 50053
// 2019/08/07 10:09:02 Container started, id: a5593a57db3d16b7379c192ec1ae3e93700f6409cfbb5e28b2f04819bb580fad, image: heterogeneous_go:v1.0.0, ServicePort: 50054
// 2019/08/07 10:09:02 Connect to 127.0.0.1:50052
// 2019/08/07 10:09:02 Connect to 127.0.0.1:50053
// 2019/08/07 10:09:02 Connect to 127.0.0.1:50054
// 2019/08/07 10:09:02 Connect to 127.0.0.1:50051
// 2019/08/07 10:09:02 Result: 3, Server Type: GOLANG, Server Id: id:"a5593a57db3d"
// 2019/08/07 10:09:02 Result: 3, Server Type: GOLANG, Server Id: id:"4e357c9125ad"
// 2019/08/07 10:09:02 Receive result from channel: 3, GOLANG,id:"a5593a57db3d"
// 2019/08/07 10:09:02 Receive result from channel: 3, GOLANG,id:"4e357c9125ad"
// 2019/08/07 10:09:02 Result: 3, Server Type: PYTHON, Server Id: id:"37aa285695ca"
// 2019/08/07 10:09:02 Result: 3, Server Type: PYTHON, Server Id: id:"26b912fc1d8b"
// 2019/08/07 10:09:02 Receive result from channel: 3, PYTHON,id:"37aa285695ca"
// 2019/08/07 10:09:02 Receive result from channel: 3, PYTHON,id:"26b912fc1d8b"
// 2019/08/07 10:09:02 Result =  3
// 2019/08/07 10:09:02 Connect to 127.0.0.1:50054
// 2019/08/07 10:09:02 Connect to 127.0.0.1:50052
// 2019/08/07 10:09:02 Connect to 127.0.0.1:50053
// 2019/08/07 10:09:02 Connect to 127.0.0.1:50051
// 2019/08/07 10:09:02 Result: 5, Server Type: GOLANG, Server Id: id:"a5593a57db3d"
// 2019/08/07 10:09:02 Receive result from channel: 5, GOLANG,id:"a5593a57db3d"
// 2019/08/07 10:09:02 Result: 0, Server Type: GOLANG, Server Id: id:"4e357c9125ad"
// 2019/08/07 10:09:02 Receive result from channel: 0, GOLANG,id:"4e357c9125ad"
// 2019/08/07 10:09:02 Result: 5, Server Type: PYTHON, Server Id: id:"37aa285695ca"
// 2019/08/07 10:09:02 Receive result from channel: 5, PYTHON,id:"37aa285695ca"
// 2019/08/07 10:09:02 Result: 5, Server Type: PYTHON, Server Id: id:"26b912fc1d8b"
// 2019/08/07 10:09:02 Receive result from channel: 5, PYTHON,id:"26b912fc1d8b"
// maxCount = 3, key = 5, keys = map[0:0xc00027c3c0 5:0xc00027c3a0]
// 2019/08/07 10:09:02 c.ID = 26b912fc1d8b8f5f57b37c22050a2ab96565d9cad1c73476a401cbbcf27f526b success = 2/0
// 2019/08/07 10:09:02 c.ID = 37aa285695caee4f477fb38ba25449d69d4e6b95798075d7cc42f88a66fc0d2c success = 2/0
// 2019/08/07 10:09:02 c.ID = 4e357c9125ad90fe3ee9997fd839b5daff9bc6e33e3e0f748fa80193d307f9e8 failure = 1/1
// 2019/08/07 10:09:02 c.ID = a5593a57db3d16b7379c192ec1ae3e93700f6409cfbb5e28b2f04819bb580fad success = 2/0
// 2019/08/07 10:09:02 Container removed, id: 4e357c9125ad90fe3ee9997fd839b5daff9bc6e33e3e0f748fa80193d307f9e8, workable images:[heterogeneous_go:v1.0.0 heterogeneous_python:v1.0.0 heterogeneous_java:v1.0.0 heterogeneous_cpp:v1.0.0], exception images:[]
// 2019/08/07 10:09:03 Container started, id: fffaa26703170e7c6b11993fa5f9eedea54baf902a97f638adac8ffc4da63755, image: heterogeneous_java:v1.0.0, ServicePort: 50053
// 2019/08/07 10:09:03 Container:4e357c9125ad90fe3ee9997fd839b5daff9bc6e33e3e0f748fa80193d307f9e8 be removed, Image:heterogeneous_cpp:v1.0.0,
//                                                         ServiceAddress:0.0.0.0:50053, SuccCount:1,
//                                                         ContinuousFailureCount:1, TotalFailureCount:1
// 2019/08/07 10:09:03 Result =  5
// 2019/08/07 10:09:03 Connect to 127.0.0.1:50051
// 2019/08/07 10:09:03 Connect to 127.0.0.1:50054
// 2019/08/07 10:09:03 Connect to 127.0.0.1:50053
// 2019/08/07 10:09:03 Connect to 127.0.0.1:50052
// 2019/08/07 10:09:03 Result: 7, Server Type: GOLANG, Server Id: id:"a5593a57db3d"
// 2019/08/07 10:09:03 Receive result from channel: 7, GOLANG,id:"a5593a57db3d"
// 2019/08/07 10:09:03 Result: 7, Server Type: PYTHON, Server Id: id:"26b912fc1d8b"
// 2019/08/07 10:09:03 Receive result from channel: 7, PYTHON,id:"26b912fc1d8b"
// 2019/08/07 10:09:03 Result: 7, Server Type: PYTHON, Server Id: id:"37aa285695ca"
// 2019/08/07 10:09:03 Receive result from channel: 7, PYTHON,id:"37aa285695ca"
// 2019/08/07 10:09:04 Result: 7, Server Type: PYTHON, Server Id: id:"fffaa2670317"
// 2019/08/07 10:09:04 Receive result from channel: 7, PYTHON,id:"fffaa2670317"
// 2019/08/07 10:09:04 Result =  7
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50052
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50053
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50051
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50054
// 2019/08/07 10:09:04 Result: 9, Server Type: GOLANG, Server Id: id:"a5593a57db3d"
// 2019/08/07 10:09:04 Receive result from channel: 9, GOLANG,id:"a5593a57db3d"
// 2019/08/07 10:09:04 Result: 9, Server Type: PYTHON, Server Id: id:"37aa285695ca"
// 2019/08/07 10:09:04 Result: 9, Server Type: PYTHON, Server Id: id:"fffaa2670317"
// 2019/08/07 10:09:04 Receive result from channel: 9, PYTHON,id:"37aa285695ca"
// 2019/08/07 10:09:04 Receive result from channel: 9, PYTHON,id:"fffaa2670317"
// 2019/08/07 10:09:04 Result: 9, Server Type: PYTHON, Server Id: id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Receive result from channel: 9, PYTHON,id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Result =  9
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50052
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50053
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50051
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50054
// 2019/08/07 10:09:04 Result: 11, Server Type: GOLANG, Server Id: id:"a5593a57db3d"
// 2019/08/07 10:09:04 Receive result from channel: 11, GOLANG,id:"a5593a57db3d"
// 2019/08/07 10:09:04 Result: 11, Server Type: PYTHON, Server Id: id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Receive result from channel: 11, PYTHON,id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Result: 11, Server Type: PYTHON, Server Id: id:"fffaa2670317"
// 2019/08/07 10:09:04 Receive result from channel: 11, PYTHON,id:"fffaa2670317"
// 2019/08/07 10:09:04 Result: 11, Server Type: PYTHON, Server Id: id:"37aa285695ca"
// 2019/08/07 10:09:04 Receive result from channel: 11, PYTHON,id:"37aa285695ca"
// 2019/08/07 10:09:04 Result =  11
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50051
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50052
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50053
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50054
// 2019/08/07 10:09:04 Result: 13, Server Type: GOLANG, Server Id: id:"a5593a57db3d"
// 2019/08/07 10:09:04 Receive result from channel: 13, GOLANG,id:"a5593a57db3d"
// 2019/08/07 10:09:04 Result: 13, Server Type: PYTHON, Server Id: id:"37aa285695ca"
// 2019/08/07 10:09:04 Receive result from channel: 13, PYTHON,id:"37aa285695ca"
// 2019/08/07 10:09:04 Result: 13, Server Type: PYTHON, Server Id: id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Receive result from channel: 13, PYTHON,id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Result: 13, Server Type: PYTHON, Server Id: id:"fffaa2670317"
// 2019/08/07 10:09:04 Receive result from channel: 13, PYTHON,id:"fffaa2670317"
// 2019/08/07 10:09:04 Result =  13
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50052
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50053
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50054
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50051
// 2019/08/07 10:09:04 Result: 15, Server Type: GOLANG, Server Id: id:"a5593a57db3d"
// 2019/08/07 10:09:04 Receive result from channel: 15, GOLANG,id:"a5593a57db3d"
// 2019/08/07 10:09:04 Result: 15, Server Type: PYTHON, Server Id: id:"37aa285695ca"
// 2019/08/07 10:09:04 Receive result from channel: 15, PYTHON,id:"37aa285695ca"
// 2019/08/07 10:09:04 Result: 15, Server Type: PYTHON, Server Id: id:"fffaa2670317"
// 2019/08/07 10:09:04 Receive result from channel: 15, PYTHON,id:"fffaa2670317"
// 2019/08/07 10:09:04 Result: 15, Server Type: PYTHON, Server Id: id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Receive result from channel: 15, PYTHON,id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Result =  15
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50052
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50051
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50053
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50054
// 2019/08/07 10:09:04 Result: 17, Server Type: GOLANG, Server Id: id:"a5593a57db3d"
// 2019/08/07 10:09:04 Receive result from channel: 17, GOLANG,id:"a5593a57db3d"
// 2019/08/07 10:09:04 Result: 17, Server Type: PYTHON, Server Id: id:"37aa285695ca"
// 2019/08/07 10:09:04 Receive result from channel: 17, PYTHON,id:"37aa285695ca"
// 2019/08/07 10:09:04 Result: 17, Server Type: PYTHON, Server Id: id:"fffaa2670317"
// 2019/08/07 10:09:04 Result: 17, Server Type: PYTHON, Server Id: id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Receive result from channel: 17, PYTHON,id:"fffaa2670317"
// 2019/08/07 10:09:04 Receive result from channel: 17, PYTHON,id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Result =  17
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50051
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50054
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50053
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50052
// 2019/08/07 10:09:04 Result: 19, Server Type: GOLANG, Server Id: id:"a5593a57db3d"
// 2019/08/07 10:09:04 Receive result from channel: 19, GOLANG,id:"a5593a57db3d"
// 2019/08/07 10:09:04 Result: 19, Server Type: PYTHON, Server Id: id:"fffaa2670317"
// 2019/08/07 10:09:04 Result: 19, Server Type: PYTHON, Server Id: id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Receive result from channel: 19, PYTHON,id:"fffaa2670317"
// 2019/08/07 10:09:04 Receive result from channel: 19, PYTHON,id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Result: 19, Server Type: PYTHON, Server Id: id:"37aa285695ca"
// 2019/08/07 10:09:04 Receive result from channel: 19, PYTHON,id:"37aa285695ca"
// 2019/08/07 10:09:04 Result =  19
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50052
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50051
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50054
// 2019/08/07 10:09:04 Connect to 127.0.0.1:50053
// 2019/08/07 10:09:04 Result: 21, Server Type: GOLANG, Server Id: id:"a5593a57db3d"
// 2019/08/07 10:09:04 Receive result from channel: 21, GOLANG,id:"a5593a57db3d"
// 2019/08/07 10:09:04 Result: 21, Server Type: PYTHON, Server Id: id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Result: 21, Server Type: PYTHON, Server Id: id:"fffaa2670317"
// 2019/08/07 10:09:04 Result: 21, Server Type: PYTHON, Server Id: id:"37aa285695ca"
// 2019/08/07 10:09:04 Receive result from channel: 21, PYTHON,id:"26b912fc1d8b"
// 2019/08/07 10:09:04 Receive result from channel: 21, PYTHON,id:"fffaa2670317"
// 2019/08/07 10:09:04 Receive result from channel: 21, PYTHON,id:"37aa285695ca"
// 2019/08/07 10:09:04 Result =  21

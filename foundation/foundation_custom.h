#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

typedef struct {
    int len;
    const void* data;
} Data;

void* String_New(const char* str);

const char* String_Value(void* ptr);

void* Data_New(void* data, int len);

Data Data_ToBytes(void* ptr);

void* Array_New(void* data, int len);
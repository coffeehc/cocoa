#import <stdlib.h>
#import <stdbool.h>

typedef struct {
    int len;
    const void* data;
} Array;


typedef struct {
    int len;
    const void* key_data;
    const void* value_data;
} Dictionary;
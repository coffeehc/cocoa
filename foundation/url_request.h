#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* C_URLRequest_Alloc();

void* C_NSURLRequest_InitWithURL(void* ptr, void* URL);
void* C_NSURLRequest_Init(void* ptr);
void* C_NSURLRequest_RequestWithURL(void* URL);
void* C_NSURLRequest_ValueForHTTPHeaderField(void* ptr, void* field);
void* C_NSURLRequest_HTTPMethod(void* ptr);
void* C_NSURLRequest_URL(void* ptr);
Array C_NSURLRequest_HTTPBody(void* ptr);
void* C_NSURLRequest_MainDocumentURL(void* ptr);
double C_NSURLRequest_TimeoutInterval(void* ptr);
bool C_NSURLRequest_HTTPShouldHandleCookies(void* ptr);
bool C_NSURLRequest_HTTPShouldUsePipelining(void* ptr);
bool C_NSURLRequest_AllowsCellularAccess(void* ptr);
bool C_NSURLRequest_AllowsConstrainedNetworkAccess(void* ptr);
bool C_NSURLRequest_AllowsExpensiveNetworkAccess(void* ptr);
bool C_NSURLRequest_AssumesHTTP3Capable(void* ptr);

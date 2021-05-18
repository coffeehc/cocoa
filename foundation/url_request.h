#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_URLRequest_Alloc();

void* C_NSURLRequest_InitWithURL(void* ptr, void* URL);
void* C_NSURLRequest_InitWithURL_CachePolicy_TimeoutInterval(void* ptr, void* URL, unsigned int cachePolicy, double timeoutInterval);
void* C_NSURLRequest_Init(void* ptr);
void* C_NSURLRequest_URLRequest_RequestWithURL(void* URL);
void* C_NSURLRequest_URLRequest_RequestWithURL_CachePolicy_TimeoutInterval(void* URL, unsigned int cachePolicy, double timeoutInterval);
void* C_NSURLRequest_ValueForHTTPHeaderField(void* ptr, void* field);
unsigned int C_NSURLRequest_CachePolicy(void* ptr);
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
unsigned int C_NSURLRequest_NetworkServiceType(void* ptr);
bool C_NSURLRequest_URLRequest_SupportsSecureCoding();
bool C_NSURLRequest_AssumesHTTP3Capable(void* ptr);

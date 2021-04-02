#import <stdlib.h>
#import <utils.h>


void* ImageSymbolConfiguration_ConfigurationWithScale(long scale);
void* ImageSymbolConfiguration_ConfigurationWithPointSize(double pointSize, double weight);
void* ImageSymbolConfiguration_ConfigurationWithPointSizeAndScale(double pointSize, double weight, long scale);
void* ImageSymbolConfiguration_ConfigurationWithTextStyle(const char* style);
void* ImageSymbolConfiguration_configurationWithTextStyleAndScale(const char* style, long scale);

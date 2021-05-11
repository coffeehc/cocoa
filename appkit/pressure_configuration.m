#import <Appkit/Appkit.h>
#import "pressure_configuration.h"

void* C_PressureConfiguration_Alloc() {
	return [NSPressureConfiguration alloc];
}

void* C_NSPressureConfiguration_Init(void* ptr) {
	NSPressureConfiguration* nSPressureConfiguration = (NSPressureConfiguration*)ptr;
	NSPressureConfiguration* result = [nSPressureConfiguration init];
	return result;
}

void C_NSPressureConfiguration_Set(void* ptr) {
	NSPressureConfiguration* nSPressureConfiguration = (NSPressureConfiguration*)ptr;
	[nSPressureConfiguration set];
}

#import <Appkit/Appkit.h>
#import "shadow.h"

void* C_Shadow_Alloc() {
	return [NSShadow alloc];
}

void* C_NSShadow_Init(void* ptr) {
	NSShadow* nSShadow = (NSShadow*)ptr;
	NSShadow* result = [nSShadow init];
	return result;
}

void C_NSShadow_Set(void* ptr) {
	NSShadow* nSShadow = (NSShadow*)ptr;
	[nSShadow set];
}

CGSize C_NSShadow_ShadowOffset(void* ptr) {
	NSShadow* nSShadow = (NSShadow*)ptr;
	CGSize result = [nSShadow shadowOffset];
	return result;
}

void C_NSShadow_SetShadowOffset(void* ptr, CGSize value) {
	NSShadow* nSShadow = (NSShadow*)ptr;
	[nSShadow setShadowOffset:value];
}

double C_NSShadow_ShadowBlurRadius(void* ptr) {
	NSShadow* nSShadow = (NSShadow*)ptr;
	double result = [nSShadow shadowBlurRadius];
	return result;
}

void C_NSShadow_SetShadowBlurRadius(void* ptr, double value) {
	NSShadow* nSShadow = (NSShadow*)ptr;
	[nSShadow setShadowBlurRadius:value];
}

void* C_NSShadow_ShadowColor(void* ptr) {
	NSShadow* nSShadow = (NSShadow*)ptr;
	id result = [nSShadow shadowColor];
	return result;
}

void C_NSShadow_SetShadowColor(void* ptr, void* value) {
	NSShadow* nSShadow = (NSShadow*)ptr;
	[nSShadow setShadowColor:(id)value];
}

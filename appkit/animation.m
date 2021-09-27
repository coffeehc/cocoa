#import "animation.h"
#import <AppKit/NSAnimation.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

@interface NSAnimatablePropertyContainerAdaptor : NSObject <NSAnimatablePropertyContainer>
@property (assign) uintptr_t goID;
@end

@implementation NSAnimatablePropertyContainerAdaptor


- (NSObject*)animator {
    void* result_ = animatablePropertyContainer_Animator([self goID]);
    return (NSObject*)result_;
}

- (id)animationForKey:(NSAnimatablePropertyKey)key {
    void* result_ = animatablePropertyContainer_AnimationForKey([self goID], key);
    return (id)result_;
}

- (void)setAnimations:(NSDictionary*)value {
    Dictionary valueArray;
    NSArray * valueKeys = [value allKeys];
    int valueCount = [valueKeys count];
    if (valueCount > 0) {
    	void** valueKeyData = malloc(valueCount * sizeof(void*));
    	void** valueValueData = malloc(valueCount * sizeof(void*));
    	for (int i = 0; i < valueCount; i++) {
    		NSAnimatablePropertyKey kp = [valueKeys objectAtIndex:i];
    		id vp = value[kp];
    		 valueKeyData[i] = kp;
    		 valueValueData[i] = vp;
    	}
    	valueArray.key_data = valueKeyData;
    	valueArray.value_data = valueValueData;
    	valueArray.len = valueCount;
    }
    animatablePropertyContainer_SetAnimations([self goID], valueArray);
}

- (NSDictionary*)animations {
    Dictionary result_ = animatablePropertyContainer_Animations([self goID]);
    NSMutableDictionary* objcResult_ = [NSMutableDictionary dictionaryWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_KeyData = (void**)result_.key_data;
    	void** result_ValueData = (void**)result_.value_data;
    	for (int i = 0; i < result_.len; i++) {
    		void* kp = result_KeyData[i];
    		void* vp = result_ValueData[i];
    		[objcResult_ setObject:(id)vp forKey:(id<NSCopying>)(NSString*)kp];
    	}
    }
    return objcResult_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return AnimatablePropertyContainer_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapAnimatablePropertyContainer(uintptr_t goID) {
    NSAnimatablePropertyContainerAdaptor* adaptor = [[NSAnimatablePropertyContainerAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}

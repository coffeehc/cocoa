#import "sound.h"
#import <AppKit/NSSound.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_Sound_Alloc() {
    return [NSSound alloc];
}

void* C_NSSound_InitWithContentsOfFile_ByReference(void* ptr, void* path, bool byRef) {
    NSSound* nSSound = (NSSound*)ptr;
    NSSound* result_ = [nSSound initWithContentsOfFile:(NSString*)path byReference:byRef];
    return result_;
}

void* C_NSSound_InitWithContentsOfURL_ByReference(void* ptr, void* url, bool byRef) {
    NSSound* nSSound = (NSSound*)ptr;
    NSSound* result_ = [nSSound initWithContentsOfURL:(NSURL*)url byReference:byRef];
    return result_;
}

void* C_NSSound_InitWithData(void* ptr, void* data) {
    NSSound* nSSound = (NSSound*)ptr;
    NSSound* result_ = [nSSound initWithData:(NSData*)data];
    return result_;
}

void* C_NSSound_InitWithPasteboard(void* ptr, void* pasteboard) {
    NSSound* nSSound = (NSSound*)ptr;
    NSSound* result_ = [nSSound initWithPasteboard:(NSPasteboard*)pasteboard];
    return result_;
}

void* C_NSSound_AllocSound() {
    NSSound* result_ = [NSSound alloc];
    return result_;
}

void* C_NSSound_Init(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    NSSound* result_ = [nSSound init];
    return result_;
}

void* C_NSSound_NewSound() {
    NSSound* result_ = [NSSound new];
    return result_;
}

void* C_NSSound_Autorelease(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    NSSound* result_ = [nSSound autorelease];
    return result_;
}

void* C_NSSound_Retain(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    NSSound* result_ = [nSSound retain];
    return result_;
}

bool C_NSSound_Sound_CanInitWithPasteboard(void* pasteboard) {
    BOOL result_ = [NSSound canInitWithPasteboard:(NSPasteboard*)pasteboard];
    return result_;
}

bool C_NSSound_SetName(void* ptr, void* _string) {
    NSSound* nSSound = (NSSound*)ptr;
    BOOL result_ = [nSSound setName:(NSString*)_string];
    return result_;
}

void* C_NSSound_SoundNamed(void* name) {
    NSSound* result_ = [NSSound soundNamed:(NSString*)name];
    return result_;
}

bool C_NSSound_Pause(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    BOOL result_ = [nSSound pause];
    return result_;
}

bool C_NSSound_Play(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    BOOL result_ = [nSSound play];
    return result_;
}

bool C_NSSound_Resume(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    BOOL result_ = [nSSound resume];
    return result_;
}

bool C_NSSound_Stop(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    BOOL result_ = [nSSound stop];
    return result_;
}

void C_NSSound_WriteToPasteboard(void* ptr, void* pasteboard) {
    NSSound* nSSound = (NSSound*)ptr;
    [nSSound writeToPasteboard:(NSPasteboard*)pasteboard];
}

void* C_NSSound_Delegate(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    id result_ = [nSSound delegate];
    return result_;
}

void C_NSSound_SetDelegate(void* ptr, void* value) {
    NSSound* nSSound = (NSSound*)ptr;
    [nSSound setDelegate:(id)value];
}

void* C_NSSound_Name(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    NSSoundName result_ = [nSSound name];
    return result_;
}

float C_NSSound_Volume(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    float result_ = [nSSound volume];
    return result_;
}

void C_NSSound_SetVolume(void* ptr, float value) {
    NSSound* nSSound = (NSSound*)ptr;
    [nSSound setVolume:value];
}

double C_NSSound_CurrentTime(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    NSTimeInterval result_ = [nSSound currentTime];
    return result_;
}

void C_NSSound_SetCurrentTime(void* ptr, double value) {
    NSSound* nSSound = (NSSound*)ptr;
    [nSSound setCurrentTime:value];
}

bool C_NSSound_Loops(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    BOOL result_ = [nSSound loops];
    return result_;
}

void C_NSSound_SetLoops(void* ptr, bool value) {
    NSSound* nSSound = (NSSound*)ptr;
    [nSSound setLoops:value];
}

void* C_NSSound_PlaybackDeviceIdentifier(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    NSSoundPlaybackDeviceIdentifier result_ = [nSSound playbackDeviceIdentifier];
    return result_;
}

void C_NSSound_SetPlaybackDeviceIdentifier(void* ptr, void* value) {
    NSSound* nSSound = (NSSound*)ptr;
    [nSSound setPlaybackDeviceIdentifier:(NSString*)value];
}

Array C_NSSound_SoundUnfilteredTypes() {
    NSArray* result_ = [NSSound soundUnfilteredTypes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

double C_NSSound_Duration(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    NSTimeInterval result_ = [nSSound duration];
    return result_;
}

bool C_NSSound_IsPlaying(void* ptr) {
    NSSound* nSSound = (NSSound*)ptr;
    BOOL result_ = [nSSound isPlaying];
    return result_;
}

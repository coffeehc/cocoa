#import <Appkit/Appkit.h>
#import "sound.h"

void* C_Sound_Alloc() {
	return [NSSound alloc];
}

void* C_NSSound_InitWithContentsOfFile_ByReference(void* ptr, void* path, bool byRef) {
	NSSound* nSSound = (NSSound*)ptr;
	NSSound* result = [nSSound initWithContentsOfFile:(NSString*)path byReference:byRef];
	return result;
}

void* C_NSSound_InitWithContentsOfURL_ByReference(void* ptr, void* url, bool byRef) {
	NSSound* nSSound = (NSSound*)ptr;
	NSSound* result = [nSSound initWithContentsOfURL:(NSURL*)url byReference:byRef];
	return result;
}

void* C_NSSound_InitWithData(void* ptr, Array data) {
	NSSound* nSSound = (NSSound*)ptr;
	NSSound* result = [nSSound initWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
	return result;
}

void* C_NSSound_InitWithPasteboard(void* ptr, void* pasteboard) {
	NSSound* nSSound = (NSSound*)ptr;
	NSSound* result = [nSSound initWithPasteboard:(NSPasteboard*)pasteboard];
	return result;
}

void* C_NSSound_Init(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	NSSound* result = [nSSound init];
	return result;
}

bool C_NSSound_SoundCanInitWithPasteboard(void* pasteboard) {
	bool result = [NSSound canInitWithPasteboard:(NSPasteboard*)pasteboard];
	return result;
}

bool C_NSSound_SetName(void* ptr, void* _string) {
	NSSound* nSSound = (NSSound*)ptr;
	bool result = [nSSound setName:(NSString*)_string];
	return result;
}

void* C_NSSound_SoundNamed(void* name) {
	NSSound* result = [NSSound soundNamed:(NSString*)name];
	return result;
}

bool C_NSSound_Pause(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	bool result = [nSSound pause];
	return result;
}

bool C_NSSound_Play(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	bool result = [nSSound play];
	return result;
}

bool C_NSSound_Resume(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	bool result = [nSSound resume];
	return result;
}

bool C_NSSound_Stop(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	bool result = [nSSound stop];
	return result;
}

void C_NSSound_WriteToPasteboard(void* ptr, void* pasteboard) {
	NSSound* nSSound = (NSSound*)ptr;
	[nSSound writeToPasteboard:(NSPasteboard*)pasteboard];
}

void* C_NSSound_Name(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	NSString* result = [nSSound name];
	return result;
}

float C_NSSound_Volume(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	float result = [nSSound volume];
	return result;
}

void C_NSSound_SetVolume(void* ptr, float value) {
	NSSound* nSSound = (NSSound*)ptr;
	[nSSound setVolume:value];
}

double C_NSSound_CurrentTime(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	double result = [nSSound currentTime];
	return result;
}

void C_NSSound_SetCurrentTime(void* ptr, double value) {
	NSSound* nSSound = (NSSound*)ptr;
	[nSSound setCurrentTime:value];
}

bool C_NSSound_Loops(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	bool result = [nSSound loops];
	return result;
}

void C_NSSound_SetLoops(void* ptr, bool value) {
	NSSound* nSSound = (NSSound*)ptr;
	[nSSound setLoops:value];
}

void* C_NSSound_PlaybackDeviceIdentifier(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	NSString* result = [nSSound playbackDeviceIdentifier];
	return result;
}

void C_NSSound_SetPlaybackDeviceIdentifier(void* ptr, void* value) {
	NSSound* nSSound = (NSSound*)ptr;
	[nSSound setPlaybackDeviceIdentifier:(NSString*)value];
}

double C_NSSound_Duration(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	double result = [nSSound duration];
	return result;
}

bool C_NSSound_IsPlaying(void* ptr) {
	NSSound* nSSound = (NSSound*)ptr;
	bool result = [nSSound isPlaying];
	return result;
}

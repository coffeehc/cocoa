#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Sound_Alloc();

void* C_NSSound_InitWithContentsOfFile_ByReference(void* ptr, void* path, bool byRef);
void* C_NSSound_InitWithContentsOfURL_ByReference(void* ptr, void* url, bool byRef);
void* C_NSSound_InitWithData(void* ptr, void* data);
void* C_NSSound_InitWithPasteboard(void* ptr, void* pasteboard);
void* C_NSSound_AllocSound();
void* C_NSSound_Init(void* ptr);
void* C_NSSound_NewSound();
void* C_NSSound_Autorelease(void* ptr);
void* C_NSSound_Retain(void* ptr);
bool C_NSSound_Sound_CanInitWithPasteboard(void* pasteboard);
bool C_NSSound_SetName(void* ptr, void* _string);
void* C_NSSound_SoundNamed(void* name);
bool C_NSSound_Pause(void* ptr);
bool C_NSSound_Play(void* ptr);
bool C_NSSound_Resume(void* ptr);
bool C_NSSound_Stop(void* ptr);
void C_NSSound_WriteToPasteboard(void* ptr, void* pasteboard);
void* C_NSSound_Delegate(void* ptr);
void C_NSSound_SetDelegate(void* ptr, void* value);
void* C_NSSound_Name(void* ptr);
float C_NSSound_Volume(void* ptr);
void C_NSSound_SetVolume(void* ptr, float value);
double C_NSSound_CurrentTime(void* ptr);
void C_NSSound_SetCurrentTime(void* ptr, double value);
bool C_NSSound_Loops(void* ptr);
void C_NSSound_SetLoops(void* ptr, bool value);
void* C_NSSound_PlaybackDeviceIdentifier(void* ptr);
void C_NSSound_SetPlaybackDeviceIdentifier(void* ptr, void* value);
Array C_NSSound_SoundUnfilteredTypes();
double C_NSSound_Duration(void* ptr);
bool C_NSSound_IsPlaying(void* ptr);

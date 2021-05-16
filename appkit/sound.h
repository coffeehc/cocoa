#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Sound_Alloc();

void* C_NSSound_InitWithContentsOfFile_ByReference(void* ptr, void* path, bool byRef);
void* C_NSSound_InitWithContentsOfURL_ByReference(void* ptr, void* url, bool byRef);
void* C_NSSound_InitWithData(void* ptr, Array data);
void* C_NSSound_InitWithPasteboard(void* ptr, void* pasteboard);
void* C_NSSound_Init(void* ptr);
bool C_NSSound_SoundCanInitWithPasteboard(void* pasteboard);
bool C_NSSound_SetName(void* ptr, void* _string);
void* C_NSSound_SoundNamed(void* name);
bool C_NSSound_Pause(void* ptr);
bool C_NSSound_Play(void* ptr);
bool C_NSSound_Resume(void* ptr);
bool C_NSSound_Stop(void* ptr);
void C_NSSound_WriteToPasteboard(void* ptr, void* pasteboard);
void* C_NSSound_Name(void* ptr);
float C_NSSound_Volume(void* ptr);
void C_NSSound_SetVolume(void* ptr, float value);
double C_NSSound_CurrentTime(void* ptr);
void C_NSSound_SetCurrentTime(void* ptr, double value);
bool C_NSSound_Loops(void* ptr);
void C_NSSound_SetLoops(void* ptr, bool value);
void* C_NSSound_PlaybackDeviceIdentifier(void* ptr);
void C_NSSound_SetPlaybackDeviceIdentifier(void* ptr, void* value);
double C_NSSound_Duration(void* ptr);
bool C_NSSound_IsPlaying(void* ptr);
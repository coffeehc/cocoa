#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_RulerMarker_Alloc();

void* C_NSRulerMarker_InitWithRulerView_MarkerLocation_Image_ImageOrigin(void* ptr, void* ruler, double location, void* image, CGPoint imageOrigin);
void* C_NSRulerMarker_InitWithCoder(void* ptr, void* coder);
void C_NSRulerMarker_DrawRect(void* ptr, CGRect rect);
bool C_NSRulerMarker_TrackMouse_Adding(void* ptr, void* mouseDownEvent, bool isAdding);
void* C_NSRulerMarker_Ruler(void* ptr);
void* C_NSRulerMarker_Image(void* ptr);
void C_NSRulerMarker_SetImage(void* ptr, void* value);
CGPoint C_NSRulerMarker_ImageOrigin(void* ptr);
void C_NSRulerMarker_SetImageOrigin(void* ptr, CGPoint value);
CGRect C_NSRulerMarker_ImageRectInRuler(void* ptr);
double C_NSRulerMarker_ThicknessRequiredInRuler(void* ptr);
bool C_NSRulerMarker_IsMovable(void* ptr);
void C_NSRulerMarker_SetMovable(void* ptr, bool value);
bool C_NSRulerMarker_IsRemovable(void* ptr);
void C_NSRulerMarker_SetRemovable(void* ptr, bool value);
double C_NSRulerMarker_MarkerLocation(void* ptr);
void C_NSRulerMarker_SetMarkerLocation(void* ptr, double value);
bool C_NSRulerMarker_IsDragging(void* ptr);

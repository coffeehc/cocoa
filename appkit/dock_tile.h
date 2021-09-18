#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_DockTile_Alloc();

void* C_NSDockTile_AllocDockTile();
void* C_NSDockTile_Init(void* ptr);
void* C_NSDockTile_NewDockTile();
void* C_NSDockTile_Autorelease(void* ptr);
void* C_NSDockTile_Retain(void* ptr);
void C_NSDockTile_Display(void* ptr);
void* C_NSDockTile_ContentView(void* ptr);
void C_NSDockTile_SetContentView(void* ptr, void* value);
CGSize C_NSDockTile_Size(void* ptr);
void* C_NSDockTile_Owner(void* ptr);
bool C_NSDockTile_ShowsApplicationBadge(void* ptr);
void C_NSDockTile_SetShowsApplicationBadge(void* ptr, bool value);
void* C_NSDockTile_BadgeLabel(void* ptr);
void C_NSDockTile_SetBadgeLabel(void* ptr, void* value);

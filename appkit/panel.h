#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool Panel_IsFloatingPanel(void* ptr);
void Panel_SetFloatingPanel(void* ptr, bool floatingPanel);
bool Panel_BecomesKeyOnlyIfNeeded(void* ptr);
void Panel_SetBecomesKeyOnlyIfNeeded(void* ptr, bool becomesKeyOnlyIfNeeded);
bool Panel_WorksWhenModal(void* ptr);
void Panel_SetWorksWhenModal(void* ptr, bool worksWhenModal);


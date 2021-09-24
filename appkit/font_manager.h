#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_FontManager_Alloc();

void* C_NSFontManager_AllocFontManager();
void* C_NSFontManager_Init(void* ptr);
void* C_NSFontManager_NewFontManager();
void* C_NSFontManager_Autorelease(void* ptr);
void* C_NSFontManager_Retain(void* ptr);
Array C_NSFontManager_AvailableFontNamesWithTraits(void* ptr, unsigned int someTraits);
void C_NSFontManager_SetSelectedFont_IsMultiple(void* ptr, void* fontObj, bool flag);
bool C_NSFontManager_SendAction(void* ptr);
void* C_NSFontManager_LocalizedNameForFamily_Face(void* ptr, void* family, void* faceKey);
void C_NSFontManager_AddFontTrait(void* ptr, void* sender);
void C_NSFontManager_RemoveFontTrait(void* ptr, void* sender);
void C_NSFontManager_ModifyFont(void* ptr, void* sender);
void C_NSFontManager_ModifyFontViaPanel(void* ptr, void* sender);
void C_NSFontManager_OrderFrontStylesPanel(void* ptr, void* sender);
void C_NSFontManager_OrderFrontFontPanel(void* ptr, void* sender);
void* C_NSFontManager_ConvertFont(void* ptr, void* fontObj);
void* C_NSFontManager_ConvertFont_ToFace(void* ptr, void* fontObj, void* typeface);
void* C_NSFontManager_ConvertFont_ToFamily(void* ptr, void* fontObj, void* family);
void* C_NSFontManager_ConvertFont_ToHaveTrait(void* ptr, void* fontObj, unsigned int trait);
void* C_NSFontManager_ConvertFont_ToNotHaveTrait(void* ptr, void* fontObj, unsigned int trait);
void* C_NSFontManager_ConvertFont_ToSize(void* ptr, void* fontObj, double size);
void* C_NSFontManager_ConvertWeight_OfFont(void* ptr, bool upFlag, void* fontObj);
unsigned int C_NSFontManager_ConvertFontTraits(void* ptr, unsigned int traits);
void* C_NSFontManager_FontWithFamily_Traits_Weight_Size(void* ptr, void* family, unsigned int traits, int weight, double size);
unsigned int C_NSFontManager_TraitsOfFont(void* ptr, void* fontObj);
bool C_NSFontManager_FontNamed_HasTraits(void* ptr, void* fName, unsigned int someTraits);
int C_NSFontManager_WeightOfFont(void* ptr, void* fontObj);
void* C_NSFontManager_FontPanel(void* ptr, bool create);
void C_NSFontManager_SetFontMenu(void* ptr, void* newMenu);
void* C_NSFontManager_FontMenu(void* ptr, bool create);
void C_NSFontManager_SetSelectedAttributes_IsMultiple(void* ptr, Dictionary attributes, bool flag);
Dictionary C_NSFontManager_ConvertAttributes(void* ptr, Dictionary attributes);
void* C_NSFontManager_SharedFontManager();
Array C_NSFontManager_AvailableFonts(void* ptr);
Array C_NSFontManager_AvailableFontFamilies(void* ptr);
void* C_NSFontManager_SelectedFont(void* ptr);
bool C_NSFontManager_IsMultiple(void* ptr);
unsigned int C_NSFontManager_CurrentFontAction(void* ptr);
bool C_NSFontManager_IsEnabled(void* ptr);
void C_NSFontManager_SetEnabled(void* ptr, bool value);
void* C_NSFontManager_Action(void* ptr);
void C_NSFontManager_SetAction(void* ptr, void* value);
void* C_NSFontManager_Target(void* ptr);
void C_NSFontManager_SetTarget(void* ptr, void* value);

#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TextStorage_Alloc();

void* C_NSTextStorage_InitWithString(void* ptr, void* str);
void* C_NSTextStorage_InitWithString_Attributes(void* ptr, void* str, Dictionary attrs);
void* C_NSTextStorage_InitWithAttributedString(void* ptr, void* attrStr);
void* C_NSTextStorage_AllocTextStorage();
void* C_NSTextStorage_Init(void* ptr);
void* C_NSTextStorage_NewTextStorage();
void* C_NSTextStorage_Autorelease(void* ptr);
void* C_NSTextStorage_Retain(void* ptr);
void C_NSTextStorage_AddLayoutManager(void* ptr, void* aLayoutManager);
void C_NSTextStorage_RemoveLayoutManager(void* ptr, void* aLayoutManager);
void C_NSTextStorage_Edited_Range_ChangeInLength(void* ptr, unsigned int editedMask, NSRange editedRange, int delta);
void C_NSTextStorage_ProcessEditing(void* ptr);
void C_NSTextStorage_InvalidateAttributesInRange(void* ptr, NSRange _range);
void C_NSTextStorage_EnsureAttributesAreFixedInRange(void* ptr, NSRange _range);
void* C_NSTextStorage_Delegate(void* ptr);
void C_NSTextStorage_SetDelegate(void* ptr, void* value);
Array C_NSTextStorage_LayoutManagers(void* ptr);
bool C_NSTextStorage_FixesAttributesLazily(void* ptr);
unsigned int C_NSTextStorage_EditedMask(void* ptr);
NSRange C_NSTextStorage_EditedRange(void* ptr);
int C_NSTextStorage_ChangeInLength(void* ptr);
Array C_NSTextStorage_AttributeRuns(void* ptr);
void C_NSTextStorage_SetAttributeRuns(void* ptr, Array value);
Array C_NSTextStorage_Paragraphs(void* ptr);
void C_NSTextStorage_SetParagraphs(void* ptr, Array value);
Array C_NSTextStorage_Words(void* ptr);
void C_NSTextStorage_SetWords(void* ptr, Array value);
Array C_NSTextStorage_Characters(void* ptr);
void C_NSTextStorage_SetCharacters(void* ptr, Array value);
void* C_NSTextStorage_Font(void* ptr);
void C_NSTextStorage_SetFont(void* ptr, void* value);
void* C_NSTextStorage_ForegroundColor(void* ptr);
void C_NSTextStorage_SetForegroundColor(void* ptr, void* value);

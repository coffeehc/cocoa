#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_PrintPanel_Alloc();

void* C_NSPrintPanel_Init(void* ptr);
void* C_NSPrintPanel_PrintPanel_();
void* C_NSPrintPanel_DefaultButtonTitle(void* ptr);
void C_NSPrintPanel_SetDefaultButtonTitle(void* ptr, void* defaultButtonTitle);
void C_NSPrintPanel_AddAccessoryController(void* ptr, void* accessoryController);
void C_NSPrintPanel_RemoveAccessoryController(void* ptr, void* accessoryController);
int C_NSPrintPanel_RunModal(void* ptr);
int C_NSPrintPanel_RunModalWithPrintInfo(void* ptr, void* printInfo);
void* C_NSPrintPanel_JobStyleHint(void* ptr);
void C_NSPrintPanel_SetJobStyleHint(void* ptr, void* value);
unsigned int C_NSPrintPanel_Options(void* ptr);
void C_NSPrintPanel_SetOptions(void* ptr, unsigned int value);
void* C_NSPrintPanel_HelpAnchor(void* ptr);
void C_NSPrintPanel_SetHelpAnchor(void* ptr, void* value);
Array C_NSPrintPanel_AccessoryControllers(void* ptr);
void* C_NSPrintPanel_PrintInfo(void* ptr);

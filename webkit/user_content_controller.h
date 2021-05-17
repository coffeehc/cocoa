#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_UserContentController_Alloc();

void* C_WKUserContentController_Init(void* ptr);
void C_WKUserContentController_AddUserScript(void* ptr, void* userScript);
void C_WKUserContentController_RemoveAllUserScripts(void* ptr);
void C_WKUserContentController_AddScriptMessageHandler_Name(void* ptr, void* scriptMessageHandler, void* name);
void C_WKUserContentController_AddScriptMessageHandler_ContentWorld_Name(void* ptr, void* scriptMessageHandler, void* world, void* name);
void C_WKUserContentController_AddScriptMessageHandlerWithReply_ContentWorld_Name(void* ptr, void* scriptMessageHandlerWithReply, void* contentWorld, void* name);
void C_WKUserContentController_RemoveScriptMessageHandlerForName(void* ptr, void* name);
void C_WKUserContentController_RemoveScriptMessageHandlerForName_ContentWorld(void* ptr, void* name, void* contentWorld);
void C_WKUserContentController_RemoveAllScriptMessageHandlersFromContentWorld(void* ptr, void* contentWorld);
void C_WKUserContentController_RemoveAllScriptMessageHandlers(void* ptr);
void C_WKUserContentController_AddContentRuleList(void* ptr, void* contentRuleList);
void C_WKUserContentController_RemoveContentRuleList(void* ptr, void* contentRuleList);
void C_WKUserContentController_RemoveAllContentRuleLists(void* ptr);
Array C_WKUserContentController_UserScripts(void* ptr);

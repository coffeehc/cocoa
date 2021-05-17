#import <WebKit/WebKit.h>
#import "user_content_controller.h"

void* C_UserContentController_Alloc() {
    return [WKUserContentController alloc];
}

void* C_WKUserContentController_Init(void* ptr) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    WKUserContentController* result_ = [wKUserContentController init];
    return result_;
}

void C_WKUserContentController_AddUserScript(void* ptr, void* userScript) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController addUserScript:(WKUserScript*)userScript];
}

void C_WKUserContentController_RemoveAllUserScripts(void* ptr) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController removeAllUserScripts];
}

void C_WKUserContentController_AddScriptMessageHandler_Name(void* ptr, void* scriptMessageHandler, void* name) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController addScriptMessageHandler:(id)scriptMessageHandler name:(NSString*)name];
}

void C_WKUserContentController_AddScriptMessageHandler_ContentWorld_Name(void* ptr, void* scriptMessageHandler, void* world, void* name) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController addScriptMessageHandler:(id)scriptMessageHandler contentWorld:(WKContentWorld*)world name:(NSString*)name];
}

void C_WKUserContentController_AddScriptMessageHandlerWithReply_ContentWorld_Name(void* ptr, void* scriptMessageHandlerWithReply, void* contentWorld, void* name) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController addScriptMessageHandlerWithReply:(id)scriptMessageHandlerWithReply contentWorld:(WKContentWorld*)contentWorld name:(NSString*)name];
}

void C_WKUserContentController_RemoveScriptMessageHandlerForName(void* ptr, void* name) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController removeScriptMessageHandlerForName:(NSString*)name];
}

void C_WKUserContentController_RemoveScriptMessageHandlerForName_ContentWorld(void* ptr, void* name, void* contentWorld) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController removeScriptMessageHandlerForName:(NSString*)name contentWorld:(WKContentWorld*)contentWorld];
}

void C_WKUserContentController_RemoveAllScriptMessageHandlersFromContentWorld(void* ptr, void* contentWorld) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController removeAllScriptMessageHandlersFromContentWorld:(WKContentWorld*)contentWorld];
}

void C_WKUserContentController_RemoveAllScriptMessageHandlers(void* ptr) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController removeAllScriptMessageHandlers];
}

void C_WKUserContentController_AddContentRuleList(void* ptr, void* contentRuleList) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController addContentRuleList:(WKContentRuleList*)contentRuleList];
}

void C_WKUserContentController_RemoveContentRuleList(void* ptr, void* contentRuleList) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController removeContentRuleList:(WKContentRuleList*)contentRuleList];
}

void C_WKUserContentController_RemoveAllContentRuleLists(void* ptr) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    [wKUserContentController removeAllContentRuleLists];
}

Array C_WKUserContentController_UserScripts(void* ptr) {
    WKUserContentController* wKUserContentController = (WKUserContentController*)ptr;
    NSArray* result_ = [wKUserContentController userScripts];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

#import <Appkit/Appkit.h>
#import "text_view_delegate.h"
#import "_cgo_export.h"

@implementation NSTextViewDelegateAdaptor


- (NSUndoManager*)undoManagerForTextView:(NSTextView*)view {
    void* result_ = textViewDelegate_UndoManagerForTextView([self goID], view);
    return (NSUndoManager*)result_;
}

- (NSString*)textView:(NSTextView*)textView willDisplayToolTip:(NSString*)tooltip forCharacterAtIndex:(NSUInteger)characterIndex {
    void* result_ = textViewDelegate_TextView_WillDisplayToolTip_ForCharacterAtIndex([self goID], textView, tooltip, characterIndex);
    return (NSString*)result_;
}

- (NSURL*)textView:(NSTextView*)textView URLForContentsOfTextAttachment:(NSTextAttachment*)textAttachment atIndex:(NSUInteger)charIndex {
    void* result_ = textViewDelegate_TextView_URLForContentsOfTextAttachment_AtIndex([self goID], textView, textAttachment, charIndex);
    return (NSURL*)result_;
}

- (NSRange)textView:(NSTextView*)textView willChangeSelectionFromCharacterRange:(NSRange)oldSelectedCharRange toCharacterRange:(NSRange)newSelectedCharRange {
    NSRange result_ = textViewDelegate_TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange([self goID], textView, oldSelectedCharRange, newSelectedCharRange);
    return result_;
}

- (NSArray*)textView:(NSTextView*)textView willChangeSelectionFromCharacterRanges:(NSArray*)oldSelectedCharRanges toCharacterRanges:(NSArray*)newSelectedCharRanges {
    int oldSelectedCharRangescount = [oldSelectedCharRanges count];
    void** oldSelectedCharRangesData = malloc(oldSelectedCharRangescount * sizeof(void*));
    for (int i = 0; i < oldSelectedCharRangescount; i++) {
    	 void* p = [oldSelectedCharRanges objectAtIndex:i];
    	 oldSelectedCharRangesData[i] = p;
    }
    Array oldSelectedCharRangesArray;
    oldSelectedCharRangesArray.data = oldSelectedCharRangesData;
    oldSelectedCharRangesArray.len = oldSelectedCharRangescount;
    int newSelectedCharRangescount = [newSelectedCharRanges count];
    void** newSelectedCharRangesData = malloc(newSelectedCharRangescount * sizeof(void*));
    for (int i = 0; i < newSelectedCharRangescount; i++) {
    	 void* p = [newSelectedCharRanges objectAtIndex:i];
    	 newSelectedCharRangesData[i] = p;
    }
    Array newSelectedCharRangesArray;
    newSelectedCharRangesArray.data = newSelectedCharRangesData;
    newSelectedCharRangesArray.len = newSelectedCharRangescount;
    Array result_ = textViewDelegate_TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges([self goID], textView, oldSelectedCharRangesArray, newSelectedCharRangesArray);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    void** result_Data = (void**)result_.data;
    for (int i = 0; i < result_.len; i++) {
    	void* p = result_Data[i];
    	[objcResult_ addObject:(NSValue*)(NSValue*)p];
    }
    return objcResult_;
}

- (void)textViewDidChangeSelection:(NSNotification*)notification {
    textViewDelegate_TextViewDidChangeSelection([self goID], notification);
}

- (NSArray*)textView:(NSTextView*)view writablePasteboardTypesForCell:(id)cell atIndex:(NSUInteger)charIndex {
    Array result_ = textViewDelegate_TextView_WritablePasteboardTypesForCell_AtIndex([self goID], view, cell, charIndex);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    void** result_Data = (void**)result_.data;
    for (int i = 0; i < result_.len; i++) {
    	void* p = result_Data[i];
    	[objcResult_ addObject:(NSPasteboardType)(NSString*)p];
    }
    return objcResult_;
}

- (BOOL)textView:(NSTextView*)view writeCell:(id)cell atIndex:(NSUInteger)charIndex toPasteboard:(NSPasteboard*)pboard type:(NSPasteboardType)_type {
    bool result_ = textViewDelegate_TextView_WriteCell_AtIndex_ToPasteboard_Type([self goID], view, cell, charIndex, pboard, _type);
    return result_;
}

- (BOOL)textView:(NSTextView*)textView shouldChangeTextInRange:(NSRange)affectedCharRange replacementString:(NSString*)replacementString {
    bool result_ = textViewDelegate_TextView_ShouldChangeTextInRange_ReplacementString([self goID], textView, affectedCharRange, replacementString);
    return result_;
}

- (BOOL)textView:(NSTextView*)textView shouldChangeTextInRanges:(NSArray*)affectedRanges replacementStrings:(NSArray*)replacementStrings {
    int affectedRangescount = [affectedRanges count];
    void** affectedRangesData = malloc(affectedRangescount * sizeof(void*));
    for (int i = 0; i < affectedRangescount; i++) {
    	 void* p = [affectedRanges objectAtIndex:i];
    	 affectedRangesData[i] = p;
    }
    Array affectedRangesArray;
    affectedRangesArray.data = affectedRangesData;
    affectedRangesArray.len = affectedRangescount;
    int replacementStringscount = [replacementStrings count];
    void** replacementStringsData = malloc(replacementStringscount * sizeof(void*));
    for (int i = 0; i < replacementStringscount; i++) {
    	 void* p = [replacementStrings objectAtIndex:i];
    	 replacementStringsData[i] = p;
    }
    Array replacementStringsArray;
    replacementStringsArray.data = replacementStringsData;
    replacementStringsArray.len = replacementStringscount;
    bool result_ = textViewDelegate_TextView_ShouldChangeTextInRanges_ReplacementStrings([self goID], textView, affectedRangesArray, replacementStringsArray);
    return result_;
}

- (void)textViewDidChangeTypingAttributes:(NSNotification*)notification {
    textViewDelegate_TextViewDidChangeTypingAttributes([self goID], notification);
}

- (void)textView:(NSTextView*)textView clickedOnCell:(id)cell inRect:(NSRect)cellFrame atIndex:(NSUInteger)charIndex {
    textViewDelegate_TextView_ClickedOnCell_InRect_AtIndex([self goID], textView, cell, cellFrame, charIndex);
}

- (void)textView:(NSTextView*)textView doubleClickedOnCell:(id)cell inRect:(NSRect)cellFrame atIndex:(NSUInteger)charIndex {
    textViewDelegate_TextView_DoubleClickedOnCell_InRect_AtIndex([self goID], textView, cell, cellFrame, charIndex);
}

- (BOOL)textView:(NSTextView*)textView clickedOnLink:(id)link atIndex:(NSUInteger)charIndex {
    bool result_ = textViewDelegate_TextView_ClickedOnLink_AtIndex([self goID], textView, link, charIndex);
    return result_;
}

- (NSInteger)textView:(NSTextView*)textView shouldSetSpellingState:(NSInteger)value range:(NSRange)affectedCharRange {
    int result_ = textViewDelegate_TextView_ShouldSetSpellingState_Range([self goID], textView, value, affectedCharRange);
    return result_;
}

- (void)textView:(NSTextView*)view draggedCell:(id)cell inRect:(NSRect)rect event:(NSEvent*)event atIndex:(NSUInteger)charIndex {
    textViewDelegate_TextView_DraggedCell_InRect_Event_AtIndex([self goID], view, cell, rect, event, charIndex);
}

- (NSSharingServicePicker*)textView:(NSTextView*)textView willShowSharingServicePicker:(NSSharingServicePicker*)servicePicker forItems:(NSArray*)items {
    int itemscount = [items count];
    void** itemsData = malloc(itemscount * sizeof(void*));
    for (int i = 0; i < itemscount; i++) {
    	 void* p = [items objectAtIndex:i];
    	 itemsData[i] = p;
    }
    Array itemsArray;
    itemsArray.data = itemsData;
    itemsArray.len = itemscount;
    void* result_ = textViewDelegate_TextView_WillShowSharingServicePicker_ForItems([self goID], textView, servicePicker, itemsArray);
    return (NSSharingServicePicker*)result_;
}

- (BOOL)textView:(NSTextView*)textView doCommandBySelector:(SEL)commandSelector {
    bool result_ = textViewDelegate_TextView_DoCommandBySelector([self goID], textView, commandSelector);
    return result_;
}

- (NSMenu*)textView:(NSTextView*)view menu:(NSMenu*)menu forEvent:(NSEvent*)event atIndex:(NSUInteger)charIndex {
    void* result_ = textViewDelegate_TextView_Menu_ForEvent_AtIndex([self goID], view, menu, event, charIndex);
    return (NSMenu*)result_;
}

- (NSArray*)textView:(NSTextView*)textView candidates:(NSArray*)candidates forSelectedRange:(NSRange)selectedRange {
    int candidatescount = [candidates count];
    void** candidatesData = malloc(candidatescount * sizeof(void*));
    for (int i = 0; i < candidatescount; i++) {
    	 void* p = [candidates objectAtIndex:i];
    	 candidatesData[i] = p;
    }
    Array candidatesArray;
    candidatesArray.data = candidatesData;
    candidatesArray.len = candidatescount;
    Array result_ = textViewDelegate_TextView_Candidates_ForSelectedRange([self goID], textView, candidatesArray, selectedRange);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    void** result_Data = (void**)result_.data;
    for (int i = 0; i < result_.len; i++) {
    	void* p = result_Data[i];
    	[objcResult_ addObject:(NSTextCheckingResult*)(NSTextCheckingResult*)p];
    }
    return objcResult_;
}

- (NSArray*)textView:(NSTextView*)textView candidatesForSelectedRange:(NSRange)selectedRange {
    Array result_ = textViewDelegate_TextView_CandidatesForSelectedRange([self goID], textView, selectedRange);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    void** result_Data = (void**)result_.data;
    for (int i = 0; i < result_.len; i++) {
    	void* p = result_Data[i];
    	[objcResult_ addObject:(NSObject*)(NSObject*)p];
    }
    return objcResult_;
}

- (BOOL)textView:(NSTextView*)textView shouldSelectCandidateAtIndex:(NSUInteger)index {
    bool result_ = textViewDelegate_TextView_ShouldSelectCandidateAtIndex([self goID], textView, index);
    return result_;
}

- (NSArray*)textView:(NSTextView*)textView shouldUpdateTouchBarItemIdentifiers:(NSArray*)identifiers {
    int identifierscount = [identifiers count];
    void** identifiersData = malloc(identifierscount * sizeof(void*));
    for (int i = 0; i < identifierscount; i++) {
    	 void* p = [identifiers objectAtIndex:i];
    	 identifiersData[i] = p;
    }
    Array identifiersArray;
    identifiersArray.data = identifiersData;
    identifiersArray.len = identifierscount;
    Array result_ = textViewDelegate_TextView_ShouldUpdateTouchBarItemIdentifiers([self goID], textView, identifiersArray);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    void** result_Data = (void**)result_.data;
    for (int i = 0; i < result_.len; i++) {
    	void* p = result_Data[i];
    	[objcResult_ addObject:(NSTouchBarItemIdentifier)(NSString*)p];
    }
    return objcResult_;
}

- (void)textDidChange:(NSNotification*)notification {
    textViewDelegate_TextDidChange([self goID], notification);
}

- (BOOL)textShouldBeginEditing:(NSText*)textObject {
    bool result_ = textViewDelegate_TextShouldBeginEditing([self goID], textObject);
    return result_;
}

- (void)textDidBeginEditing:(NSNotification*)notification {
    textViewDelegate_TextDidBeginEditing([self goID], notification);
}

- (BOOL)textShouldEndEditing:(NSText*)textObject {
    bool result_ = textViewDelegate_TextShouldEndEditing([self goID], textObject);
    return result_;
}

- (void)textDidEndEditing:(NSNotification*)notification {
    textViewDelegate_TextDidEndEditing([self goID], notification);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TextViewDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteTextViewDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapTextViewDelegate(uintptr_t goID) {
    NSTextViewDelegateAdaptor* adaptor = [[NSTextViewDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}

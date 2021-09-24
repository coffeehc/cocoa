#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_DiffableDataSourceSnapshot_Alloc();

void* C_NSDiffableDataSourceSnapshot_AllocDiffableDataSourceSnapshot();
void* C_NSDiffableDataSourceSnapshot_Init(void* ptr);
void* C_NSDiffableDataSourceSnapshot_NewDiffableDataSourceSnapshot();
void* C_NSDiffableDataSourceSnapshot_Autorelease(void* ptr);
void* C_NSDiffableDataSourceSnapshot_Retain(void* ptr);
void C_NSDiffableDataSourceSnapshot_AppendSectionsWithIdentifiers(void* ptr, Array sectionIdentifiers);
void C_NSDiffableDataSourceSnapshot_DeleteAllItems(void* ptr);
int C_NSDiffableDataSourceSnapshot_NumberOfItems(void* ptr);
int C_NSDiffableDataSourceSnapshot_NumberOfSections(void* ptr);

#import <Foundation/Foundation.h>
#import "affine_transform.h"

void* C_AffineTransform_Alloc() {
    return [NSAffineTransform alloc];
}

void* C_NSAffineTransform_InitWithTransform(void* ptr, void* transform) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    NSAffineTransform* result_ = [nSAffineTransform initWithTransform:(NSAffineTransform*)transform];
    return result_;
}

void* C_NSAffineTransform_AllocAffineTransform() {
    NSAffineTransform* result_ = [NSAffineTransform alloc];
    return result_;
}

void* C_NSAffineTransform_Autorelease(void* ptr) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    NSAffineTransform* result_ = [nSAffineTransform autorelease];
    return result_;
}

void* C_NSAffineTransform_Retain(void* ptr) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    NSAffineTransform* result_ = [nSAffineTransform retain];
    return result_;
}

void* C_NSAffineTransform_AffineTransform_Transform() {
    NSAffineTransform* result_ = [NSAffineTransform transform];
    return result_;
}

void C_NSAffineTransform_RotateByDegrees(void* ptr, double angle) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    [nSAffineTransform rotateByDegrees:angle];
}

void C_NSAffineTransform_RotateByRadians(void* ptr, double angle) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    [nSAffineTransform rotateByRadians:angle];
}

void C_NSAffineTransform_ScaleBy(void* ptr, double scale) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    [nSAffineTransform scaleBy:scale];
}

void C_NSAffineTransform_ScaleXBy_YBy(void* ptr, double scaleX, double scaleY) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    [nSAffineTransform scaleXBy:scaleX yBy:scaleY];
}

void C_NSAffineTransform_TranslateXBy_YBy(void* ptr, double deltaX, double deltaY) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    [nSAffineTransform translateXBy:deltaX yBy:deltaY];
}

void C_NSAffineTransform_AppendTransform(void* ptr, void* transform) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    [nSAffineTransform appendTransform:(NSAffineTransform*)transform];
}

void C_NSAffineTransform_PrependTransform(void* ptr, void* transform) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    [nSAffineTransform prependTransform:(NSAffineTransform*)transform];
}

void C_NSAffineTransform_Invert(void* ptr) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    [nSAffineTransform invert];
}

CGPoint C_NSAffineTransform_TransformPoint(void* ptr, CGPoint aPoint) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    NSPoint result_ = [nSAffineTransform transformPoint:aPoint];
    return result_;
}

CGSize C_NSAffineTransform_TransformSize(void* ptr, CGSize aSize) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    NSSize result_ = [nSAffineTransform transformSize:aSize];
    return result_;
}

NSAffineTransformStruct C_NSAffineTransform_TransformStruct(void* ptr) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    NSAffineTransformStruct result_ = [nSAffineTransform transformStruct];
    return result_;
}

void C_NSAffineTransform_SetTransformStruct(void* ptr, NSAffineTransformStruct value) {
    NSAffineTransform* nSAffineTransform = (NSAffineTransform*)ptr;
    [nSAffineTransform setTransformStruct:value];
}

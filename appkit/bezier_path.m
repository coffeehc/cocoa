#import <Appkit/Appkit.h>
#import "bezier_path.h"

void* C_BezierPath_Alloc() {
    return [NSBezierPath alloc];
}

void* C_NSBezierPath_Init(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSBezierPath* result_ = [nSBezierPath init];
    return result_;
}

void* C_NSBezierPath_BezierPath_() {
    NSBezierPath* result_ = [NSBezierPath bezierPath];
    return result_;
}

void* C_NSBezierPath_BezierPathWithOvalInRect(CGRect rect) {
    NSBezierPath* result_ = [NSBezierPath bezierPathWithOvalInRect:rect];
    return result_;
}

void* C_NSBezierPath_BezierPathWithRect(CGRect rect) {
    NSBezierPath* result_ = [NSBezierPath bezierPathWithRect:rect];
    return result_;
}

void* C_NSBezierPath_BezierPathWithRoundedRect_XRadius_YRadius(CGRect rect, double xRadius, double yRadius) {
    NSBezierPath* result_ = [NSBezierPath bezierPathWithRoundedRect:rect xRadius:xRadius yRadius:yRadius];
    return result_;
}

void C_NSBezierPath_MoveToPoint(void* ptr, CGPoint point) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath moveToPoint:point];
}

void C_NSBezierPath_LineToPoint(void* ptr, CGPoint point) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath lineToPoint:point];
}

void C_NSBezierPath_CurveToPoint_ControlPoint1_ControlPoint2(void* ptr, CGPoint endPoint, CGPoint controlPoint1, CGPoint controlPoint2) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath curveToPoint:endPoint controlPoint1:controlPoint1 controlPoint2:controlPoint2];
}

void C_NSBezierPath_ClosePath(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath closePath];
}

void C_NSBezierPath_RelativeMoveToPoint(void* ptr, CGPoint point) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath relativeMoveToPoint:point];
}

void C_NSBezierPath_RelativeLineToPoint(void* ptr, CGPoint point) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath relativeLineToPoint:point];
}

void C_NSBezierPath_RelativeCurveToPoint_ControlPoint1_ControlPoint2(void* ptr, CGPoint endPoint, CGPoint controlPoint1, CGPoint controlPoint2) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath relativeCurveToPoint:endPoint controlPoint1:controlPoint1 controlPoint2:controlPoint2];
}

void C_NSBezierPath_AppendBezierPath(void* ptr, void* path) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath appendBezierPath:(NSBezierPath*)path];
}

void C_NSBezierPath_AppendBezierPathWithOvalInRect(void* ptr, CGRect rect) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath appendBezierPathWithOvalInRect:rect];
}

void C_NSBezierPath_AppendBezierPathWithArcFromPoint_ToPoint_Radius(void* ptr, CGPoint point1, CGPoint point2, double radius) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath appendBezierPathWithArcFromPoint:point1 toPoint:point2 radius:radius];
}

void C_NSBezierPath_AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle(void* ptr, CGPoint center, double radius, double startAngle, double endAngle) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath appendBezierPathWithArcWithCenter:center radius:radius startAngle:startAngle endAngle:endAngle];
}

void C_NSBezierPath_AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle_Clockwise(void* ptr, CGPoint center, double radius, double startAngle, double endAngle, bool clockwise) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath appendBezierPathWithArcWithCenter:center radius:radius startAngle:startAngle endAngle:endAngle clockwise:clockwise];
}

void C_NSBezierPath_AppendBezierPathWithRect(void* ptr, CGRect rect) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath appendBezierPathWithRect:rect];
}

void C_NSBezierPath_AppendBezierPathWithRoundedRect_XRadius_YRadius(void* ptr, CGRect rect, double xRadius, double yRadius) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath appendBezierPathWithRoundedRect:rect xRadius:xRadius yRadius:yRadius];
}

void C_NSBezierPath_Stroke(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath stroke];
}

void C_NSBezierPath_Fill(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath fill];
}

void C_NSBezierPath_BezierPath_FillRect(CGRect rect) {
    [NSBezierPath fillRect:rect];
}

void C_NSBezierPath_BezierPath_StrokeRect(CGRect rect) {
    [NSBezierPath strokeRect:rect];
}

void C_NSBezierPath_BezierPath_StrokeLineFromPoint_ToPoint(CGPoint point1, CGPoint point2) {
    [NSBezierPath strokeLineFromPoint:point1 toPoint:point2];
}

void C_NSBezierPath_AddClip(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath addClip];
}

void C_NSBezierPath_SetClip(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath setClip];
}

void C_NSBezierPath_BezierPath_ClipRect(CGRect rect) {
    [NSBezierPath clipRect:rect];
}

bool C_NSBezierPath_ContainsPoint(void* ptr, CGPoint point) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    BOOL result_ = [nSBezierPath containsPoint:point];
    return result_;
}

void C_NSBezierPath_TransformUsingAffineTransform(void* ptr, void* transform) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath transformUsingAffineTransform:(NSAffineTransform*)transform];
}

unsigned int C_NSBezierPath_ElementAtIndex(void* ptr, int index) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSBezierPathElement result_ = [nSBezierPath elementAtIndex:index];
    return result_;
}

void C_NSBezierPath_RemoveAllPoints(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath removeAllPoints];
}

void* C_NSBezierPath_BezierPathByFlatteningPath(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSBezierPath* result_ = [nSBezierPath bezierPathByFlatteningPath];
    return result_;
}

void* C_NSBezierPath_BezierPathByReversingPath(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSBezierPath* result_ = [nSBezierPath bezierPathByReversingPath];
    return result_;
}

unsigned int C_NSBezierPath_WindingRule(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSWindingRule result_ = [nSBezierPath windingRule];
    return result_;
}

void C_NSBezierPath_SetWindingRule(void* ptr, unsigned int value) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath setWindingRule:value];
}

unsigned int C_NSBezierPath_LineCapStyle(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSLineCapStyle result_ = [nSBezierPath lineCapStyle];
    return result_;
}

void C_NSBezierPath_SetLineCapStyle(void* ptr, unsigned int value) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath setLineCapStyle:value];
}

unsigned int C_NSBezierPath_LineJoinStyle(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSLineJoinStyle result_ = [nSBezierPath lineJoinStyle];
    return result_;
}

void C_NSBezierPath_SetLineJoinStyle(void* ptr, unsigned int value) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath setLineJoinStyle:value];
}

double C_NSBezierPath_LineWidth(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    CGFloat result_ = [nSBezierPath lineWidth];
    return result_;
}

void C_NSBezierPath_SetLineWidth(void* ptr, double value) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath setLineWidth:value];
}

double C_NSBezierPath_MiterLimit(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    CGFloat result_ = [nSBezierPath miterLimit];
    return result_;
}

void C_NSBezierPath_SetMiterLimit(void* ptr, double value) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath setMiterLimit:value];
}

double C_NSBezierPath_Flatness(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    CGFloat result_ = [nSBezierPath flatness];
    return result_;
}

void C_NSBezierPath_SetFlatness(void* ptr, double value) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    [nSBezierPath setFlatness:value];
}

unsigned int C_NSBezierPath_BezierPath_DefaultWindingRule() {
    NSWindingRule result_ = [NSBezierPath defaultWindingRule];
    return result_;
}

void C_NSBezierPath_BezierPath_SetDefaultWindingRule(unsigned int value) {
    [NSBezierPath setDefaultWindingRule:value];
}

unsigned int C_NSBezierPath_BezierPath_DefaultLineCapStyle() {
    NSLineCapStyle result_ = [NSBezierPath defaultLineCapStyle];
    return result_;
}

void C_NSBezierPath_BezierPath_SetDefaultLineCapStyle(unsigned int value) {
    [NSBezierPath setDefaultLineCapStyle:value];
}

unsigned int C_NSBezierPath_BezierPath_DefaultLineJoinStyle() {
    NSLineJoinStyle result_ = [NSBezierPath defaultLineJoinStyle];
    return result_;
}

void C_NSBezierPath_BezierPath_SetDefaultLineJoinStyle(unsigned int value) {
    [NSBezierPath setDefaultLineJoinStyle:value];
}

double C_NSBezierPath_BezierPath_DefaultLineWidth() {
    CGFloat result_ = [NSBezierPath defaultLineWidth];
    return result_;
}

void C_NSBezierPath_BezierPath_SetDefaultLineWidth(double value) {
    [NSBezierPath setDefaultLineWidth:value];
}

double C_NSBezierPath_BezierPath_DefaultMiterLimit() {
    CGFloat result_ = [NSBezierPath defaultMiterLimit];
    return result_;
}

void C_NSBezierPath_BezierPath_SetDefaultMiterLimit(double value) {
    [NSBezierPath setDefaultMiterLimit:value];
}

double C_NSBezierPath_BezierPath_DefaultFlatness() {
    CGFloat result_ = [NSBezierPath defaultFlatness];
    return result_;
}

void C_NSBezierPath_BezierPath_SetDefaultFlatness(double value) {
    [NSBezierPath setDefaultFlatness:value];
}

CGRect C_NSBezierPath_Bounds(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSRect result_ = [nSBezierPath bounds];
    return result_;
}

CGRect C_NSBezierPath_ControlPointBounds(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSRect result_ = [nSBezierPath controlPointBounds];
    return result_;
}

CGPoint C_NSBezierPath_CurrentPoint(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSPoint result_ = [nSBezierPath currentPoint];
    return result_;
}

bool C_NSBezierPath_IsEmpty(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    BOOL result_ = [nSBezierPath isEmpty];
    return result_;
}

int C_NSBezierPath_ElementCount(void* ptr) {
    NSBezierPath* nSBezierPath = (NSBezierPath*)ptr;
    NSInteger result_ = [nSBezierPath elementCount];
    return result_;
}

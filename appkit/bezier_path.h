#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_BezierPath_Alloc();

void* C_NSBezierPath_AllocBezierPath();
void* C_NSBezierPath_Init(void* ptr);
void* C_NSBezierPath_NewBezierPath();
void* C_NSBezierPath_Autorelease(void* ptr);
void* C_NSBezierPath_Retain(void* ptr);
void* C_NSBezierPath_BezierPath_();
void* C_NSBezierPath_BezierPathWithOvalInRect(CGRect rect);
void* C_NSBezierPath_BezierPathWithRect(CGRect rect);
void* C_NSBezierPath_BezierPathWithRoundedRect_XRadius_YRadius(CGRect rect, double xRadius, double yRadius);
void C_NSBezierPath_MoveToPoint(void* ptr, CGPoint point);
void C_NSBezierPath_LineToPoint(void* ptr, CGPoint point);
void C_NSBezierPath_CurveToPoint_ControlPoint1_ControlPoint2(void* ptr, CGPoint endPoint, CGPoint controlPoint1, CGPoint controlPoint2);
void C_NSBezierPath_ClosePath(void* ptr);
void C_NSBezierPath_RelativeMoveToPoint(void* ptr, CGPoint point);
void C_NSBezierPath_RelativeLineToPoint(void* ptr, CGPoint point);
void C_NSBezierPath_RelativeCurveToPoint_ControlPoint1_ControlPoint2(void* ptr, CGPoint endPoint, CGPoint controlPoint1, CGPoint controlPoint2);
void C_NSBezierPath_AppendBezierPath(void* ptr, void* path);
void C_NSBezierPath_AppendBezierPathWithOvalInRect(void* ptr, CGRect rect);
void C_NSBezierPath_AppendBezierPathWithArcFromPoint_ToPoint_Radius(void* ptr, CGPoint point1, CGPoint point2, double radius);
void C_NSBezierPath_AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle(void* ptr, CGPoint center, double radius, double startAngle, double endAngle);
void C_NSBezierPath_AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle_Clockwise(void* ptr, CGPoint center, double radius, double startAngle, double endAngle, bool clockwise);
void C_NSBezierPath_AppendBezierPathWithRect(void* ptr, CGRect rect);
void C_NSBezierPath_AppendBezierPathWithRoundedRect_XRadius_YRadius(void* ptr, CGRect rect, double xRadius, double yRadius);
void C_NSBezierPath_Stroke(void* ptr);
void C_NSBezierPath_Fill(void* ptr);
void C_NSBezierPath_BezierPath_FillRect(CGRect rect);
void C_NSBezierPath_BezierPath_StrokeRect(CGRect rect);
void C_NSBezierPath_BezierPath_StrokeLineFromPoint_ToPoint(CGPoint point1, CGPoint point2);
void C_NSBezierPath_AddClip(void* ptr);
void C_NSBezierPath_SetClip(void* ptr);
void C_NSBezierPath_BezierPath_ClipRect(CGRect rect);
bool C_NSBezierPath_ContainsPoint(void* ptr, CGPoint point);
void C_NSBezierPath_TransformUsingAffineTransform(void* ptr, void* transform);
unsigned int C_NSBezierPath_ElementAtIndex(void* ptr, int index);
void C_NSBezierPath_RemoveAllPoints(void* ptr);
void* C_NSBezierPath_BezierPathByFlatteningPath(void* ptr);
void* C_NSBezierPath_BezierPathByReversingPath(void* ptr);
unsigned int C_NSBezierPath_WindingRule(void* ptr);
void C_NSBezierPath_SetWindingRule(void* ptr, unsigned int value);
unsigned int C_NSBezierPath_LineCapStyle(void* ptr);
void C_NSBezierPath_SetLineCapStyle(void* ptr, unsigned int value);
unsigned int C_NSBezierPath_LineJoinStyle(void* ptr);
void C_NSBezierPath_SetLineJoinStyle(void* ptr, unsigned int value);
double C_NSBezierPath_LineWidth(void* ptr);
void C_NSBezierPath_SetLineWidth(void* ptr, double value);
double C_NSBezierPath_MiterLimit(void* ptr);
void C_NSBezierPath_SetMiterLimit(void* ptr, double value);
double C_NSBezierPath_Flatness(void* ptr);
void C_NSBezierPath_SetFlatness(void* ptr, double value);
unsigned int C_NSBezierPath_BezierPath_DefaultWindingRule();
void C_NSBezierPath_BezierPath_SetDefaultWindingRule(unsigned int value);
unsigned int C_NSBezierPath_BezierPath_DefaultLineCapStyle();
void C_NSBezierPath_BezierPath_SetDefaultLineCapStyle(unsigned int value);
unsigned int C_NSBezierPath_BezierPath_DefaultLineJoinStyle();
void C_NSBezierPath_BezierPath_SetDefaultLineJoinStyle(unsigned int value);
double C_NSBezierPath_BezierPath_DefaultLineWidth();
void C_NSBezierPath_BezierPath_SetDefaultLineWidth(double value);
double C_NSBezierPath_BezierPath_DefaultMiterLimit();
void C_NSBezierPath_BezierPath_SetDefaultMiterLimit(double value);
double C_NSBezierPath_BezierPath_DefaultFlatness();
void C_NSBezierPath_BezierPath_SetDefaultFlatness(double value);
CGRect C_NSBezierPath_Bounds(void* ptr);
CGRect C_NSBezierPath_ControlPointBounds(void* ptr);
CGPoint C_NSBezierPath_CurrentPoint(void* ptr);
bool C_NSBezierPath_IsEmpty(void* ptr);
int C_NSBezierPath_ElementCount(void* ptr);

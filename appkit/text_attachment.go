package appkit

// #include "text_attachment.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextAttachment interface {
	objc.Object
	Bounds() foundation.Rect
	SetBounds(value foundation.Rect)
	Contents() []byte
	SetContents(value []byte)
	FileType() string
	SetFileType(value string)
	Image() Image
	SetImage(value Image)
	FileWrapper() foundation.FileWrapper
	SetFileWrapper(value foundation.FileWrapper)
	AttachmentCell() objc.Object
	SetAttachmentCell(value objc.Object)
}

type NSTextAttachment struct {
	objc.NSObject
}

func MakeTextAttachment(ptr unsafe.Pointer) NSTextAttachment {
	return NSTextAttachment{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSTextAttachment) InitWithFileWrapper(fileWrapper foundation.FileWrapper) NSTextAttachment {
	result_ := C.C_NSTextAttachment_InitWithFileWrapper(n.Ptr(), objc.ExtractPtr(fileWrapper))
	return MakeTextAttachment(result_)
}

func (n NSTextAttachment) InitWithData_OfType(contentData []byte, uti string) NSTextAttachment {
	result_ := C.C_NSTextAttachment_InitWithData_OfType(n.Ptr(), foundation.NewData(contentData).Ptr(), foundation.NewString(uti).Ptr())
	return MakeTextAttachment(result_)
}

func AllocTextAttachment() NSTextAttachment {
	result_ := C.C_NSTextAttachment_AllocTextAttachment()
	return MakeTextAttachment(result_)
}

func (n NSTextAttachment) Init() NSTextAttachment {
	result_ := C.C_NSTextAttachment_Init(n.Ptr())
	return MakeTextAttachment(result_)
}

func NewTextAttachment() NSTextAttachment {
	result_ := C.C_NSTextAttachment_NewTextAttachment()
	return MakeTextAttachment(result_)
}

func (n NSTextAttachment) Autorelease() NSTextAttachment {
	result_ := C.C_NSTextAttachment_Autorelease(n.Ptr())
	return MakeTextAttachment(result_)
}

func (n NSTextAttachment) Retain() NSTextAttachment {
	result_ := C.C_NSTextAttachment_Retain(n.Ptr())
	return MakeTextAttachment(result_)
}

func (n NSTextAttachment) Bounds() foundation.Rect {
	result_ := C.C_NSTextAttachment_Bounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSTextAttachment) SetBounds(value foundation.Rect) {
	C.C_NSTextAttachment_SetBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n NSTextAttachment) Contents() []byte {
	result_ := C.C_NSTextAttachment_Contents(n.Ptr())
	return foundation.MakeData(result_).ToBytes()
}

func (n NSTextAttachment) SetContents(value []byte) {
	C.C_NSTextAttachment_SetContents(n.Ptr(), foundation.NewData(value).Ptr())
}

func (n NSTextAttachment) FileType() string {
	result_ := C.C_NSTextAttachment_FileType(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSTextAttachment) SetFileType(value string) {
	C.C_NSTextAttachment_SetFileType(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSTextAttachment) Image() Image {
	result_ := C.C_NSTextAttachment_Image(n.Ptr())
	return MakeImage(result_)
}

func (n NSTextAttachment) SetImage(value Image) {
	C.C_NSTextAttachment_SetImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextAttachment) FileWrapper() foundation.FileWrapper {
	result_ := C.C_NSTextAttachment_FileWrapper(n.Ptr())
	return foundation.MakeFileWrapper(result_)
}

func (n NSTextAttachment) SetFileWrapper(value foundation.FileWrapper) {
	C.C_NSTextAttachment_SetFileWrapper(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextAttachment) AttachmentCell() objc.Object {
	result_ := C.C_NSTextAttachment_AttachmentCell(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTextAttachment) SetAttachmentCell(value objc.Object) {
	C.C_NSTextAttachment_SetAttachmentCell(n.Ptr(), objc.ExtractPtr(value))
}

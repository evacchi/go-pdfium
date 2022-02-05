// Code generated by tool. DO NOT EDIT.
// See the code_generation package.

package multi_threaded

import (
    "errors"
    "io/ioutil"

	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/responses"
)

func (i *pdfiumInstance) FPDFAction_GetDest(request *requests.FPDFAction_GetDest) (*responses.FPDFAction_GetDest, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAction_GetDest(request)
}

func (i *pdfiumInstance) FPDFAction_GetFilePath(request *requests.FPDFAction_GetFilePath) (*responses.FPDFAction_GetFilePath, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAction_GetFilePath(request)
}

func (i *pdfiumInstance) FPDFAction_GetType(request *requests.FPDFAction_GetType) (*responses.FPDFAction_GetType, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAction_GetType(request)
}

func (i *pdfiumInstance) FPDFAction_GetURIPath(request *requests.FPDFAction_GetURIPath) (*responses.FPDFAction_GetURIPath, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAction_GetURIPath(request)
}

func (i *pdfiumInstance) FPDFAttachment_GetFile(request *requests.FPDFAttachment_GetFile) (*responses.FPDFAttachment_GetFile, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAttachment_GetFile(request)
}

func (i *pdfiumInstance) FPDFAttachment_GetName(request *requests.FPDFAttachment_GetName) (*responses.FPDFAttachment_GetName, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAttachment_GetName(request)
}

func (i *pdfiumInstance) FPDFAttachment_GetStringValue(request *requests.FPDFAttachment_GetStringValue) (*responses.FPDFAttachment_GetStringValue, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAttachment_GetStringValue(request)
}

func (i *pdfiumInstance) FPDFAttachment_GetValueType(request *requests.FPDFAttachment_GetValueType) (*responses.FPDFAttachment_GetValueType, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAttachment_GetValueType(request)
}

func (i *pdfiumInstance) FPDFAttachment_HasKey(request *requests.FPDFAttachment_HasKey) (*responses.FPDFAttachment_HasKey, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAttachment_HasKey(request)
}

func (i *pdfiumInstance) FPDFAttachment_SetFile(request *requests.FPDFAttachment_SetFile) (*responses.FPDFAttachment_SetFile, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAttachment_SetFile(request)
}

func (i *pdfiumInstance) FPDFAttachment_SetStringValue(request *requests.FPDFAttachment_SetStringValue) (*responses.FPDFAttachment_SetStringValue, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFAttachment_SetStringValue(request)
}

func (i *pdfiumInstance) FPDFBookmark_Find(request *requests.FPDFBookmark_Find) (*responses.FPDFBookmark_Find, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFBookmark_Find(request)
}

func (i *pdfiumInstance) FPDFBookmark_GetAction(request *requests.FPDFBookmark_GetAction) (*responses.FPDFBookmark_GetAction, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFBookmark_GetAction(request)
}

func (i *pdfiumInstance) FPDFBookmark_GetDest(request *requests.FPDFBookmark_GetDest) (*responses.FPDFBookmark_GetDest, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFBookmark_GetDest(request)
}

func (i *pdfiumInstance) FPDFBookmark_GetFirstChild(request *requests.FPDFBookmark_GetFirstChild) (*responses.FPDFBookmark_GetFirstChild, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFBookmark_GetFirstChild(request)
}

func (i *pdfiumInstance) FPDFBookmark_GetNextSibling(request *requests.FPDFBookmark_GetNextSibling) (*responses.FPDFBookmark_GetNextSibling, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFBookmark_GetNextSibling(request)
}

func (i *pdfiumInstance) FPDFBookmark_GetTitle(request *requests.FPDFBookmark_GetTitle) (*responses.FPDFBookmark_GetTitle, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFBookmark_GetTitle(request)
}

func (i *pdfiumInstance) FPDFCatalog_IsTagged(request *requests.FPDFCatalog_IsTagged) (*responses.FPDFCatalog_IsTagged, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFCatalog_IsTagged(request)
}

func (i *pdfiumInstance) FPDFDest_GetDestPageIndex(request *requests.FPDFDest_GetDestPageIndex) (*responses.FPDFDest_GetDestPageIndex, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDest_GetDestPageIndex(request)
}

func (i *pdfiumInstance) FPDFDest_GetLocationInPage(request *requests.FPDFDest_GetLocationInPage) (*responses.FPDFDest_GetLocationInPage, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDest_GetLocationInPage(request)
}

func (i *pdfiumInstance) FPDFDest_GetView(request *requests.FPDFDest_GetView) (*responses.FPDFDest_GetView, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDest_GetView(request)
}

func (i *pdfiumInstance) FPDFDoc_AddAttachment(request *requests.FPDFDoc_AddAttachment) (*responses.FPDFDoc_AddAttachment, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDoc_AddAttachment(request)
}

func (i *pdfiumInstance) FPDFDoc_CloseJavaScriptAction(request *requests.FPDFDoc_CloseJavaScriptAction) (*responses.FPDFDoc_CloseJavaScriptAction, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDoc_CloseJavaScriptAction(request)
}

func (i *pdfiumInstance) FPDFDoc_DeleteAttachment(request *requests.FPDFDoc_DeleteAttachment) (*responses.FPDFDoc_DeleteAttachment, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDoc_DeleteAttachment(request)
}

func (i *pdfiumInstance) FPDFDoc_GetAttachment(request *requests.FPDFDoc_GetAttachment) (*responses.FPDFDoc_GetAttachment, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDoc_GetAttachment(request)
}

func (i *pdfiumInstance) FPDFDoc_GetAttachmentCount(request *requests.FPDFDoc_GetAttachmentCount) (*responses.FPDFDoc_GetAttachmentCount, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDoc_GetAttachmentCount(request)
}

func (i *pdfiumInstance) FPDFDoc_GetJavaScriptAction(request *requests.FPDFDoc_GetJavaScriptAction) (*responses.FPDFDoc_GetJavaScriptAction, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDoc_GetJavaScriptAction(request)
}

func (i *pdfiumInstance) FPDFDoc_GetJavaScriptActionCount(request *requests.FPDFDoc_GetJavaScriptActionCount) (*responses.FPDFDoc_GetJavaScriptActionCount, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDoc_GetJavaScriptActionCount(request)
}

func (i *pdfiumInstance) FPDFDoc_GetPageMode(request *requests.FPDFDoc_GetPageMode) (*responses.FPDFDoc_GetPageMode, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFDoc_GetPageMode(request)
}

func (i *pdfiumInstance) FPDFJavaScriptAction_GetName(request *requests.FPDFJavaScriptAction_GetName) (*responses.FPDFJavaScriptAction_GetName, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFJavaScriptAction_GetName(request)
}

func (i *pdfiumInstance) FPDFJavaScriptAction_GetScript(request *requests.FPDFJavaScriptAction_GetScript) (*responses.FPDFJavaScriptAction_GetScript, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFJavaScriptAction_GetScript(request)
}

func (i *pdfiumInstance) FPDFLink_CloseWebLinks(request *requests.FPDFLink_CloseWebLinks) (*responses.FPDFLink_CloseWebLinks, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_CloseWebLinks(request)
}

func (i *pdfiumInstance) FPDFLink_CountQuadPoints(request *requests.FPDFLink_CountQuadPoints) (*responses.FPDFLink_CountQuadPoints, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_CountQuadPoints(request)
}

func (i *pdfiumInstance) FPDFLink_CountRects(request *requests.FPDFLink_CountRects) (*responses.FPDFLink_CountRects, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_CountRects(request)
}

func (i *pdfiumInstance) FPDFLink_CountWebLinks(request *requests.FPDFLink_CountWebLinks) (*responses.FPDFLink_CountWebLinks, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_CountWebLinks(request)
}

func (i *pdfiumInstance) FPDFLink_Enumerate(request *requests.FPDFLink_Enumerate) (*responses.FPDFLink_Enumerate, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_Enumerate(request)
}

func (i *pdfiumInstance) FPDFLink_GetAction(request *requests.FPDFLink_GetAction) (*responses.FPDFLink_GetAction, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_GetAction(request)
}

func (i *pdfiumInstance) FPDFLink_GetAnnot(request *requests.FPDFLink_GetAnnot) (*responses.FPDFLink_GetAnnot, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_GetAnnot(request)
}

func (i *pdfiumInstance) FPDFLink_GetAnnotRect(request *requests.FPDFLink_GetAnnotRect) (*responses.FPDFLink_GetAnnotRect, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_GetAnnotRect(request)
}

func (i *pdfiumInstance) FPDFLink_GetDest(request *requests.FPDFLink_GetDest) (*responses.FPDFLink_GetDest, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_GetDest(request)
}

func (i *pdfiumInstance) FPDFLink_GetLinkAtPoint(request *requests.FPDFLink_GetLinkAtPoint) (*responses.FPDFLink_GetLinkAtPoint, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_GetLinkAtPoint(request)
}

func (i *pdfiumInstance) FPDFLink_GetLinkZOrderAtPoint(request *requests.FPDFLink_GetLinkZOrderAtPoint) (*responses.FPDFLink_GetLinkZOrderAtPoint, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_GetLinkZOrderAtPoint(request)
}

func (i *pdfiumInstance) FPDFLink_GetQuadPoints(request *requests.FPDFLink_GetQuadPoints) (*responses.FPDFLink_GetQuadPoints, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_GetQuadPoints(request)
}

func (i *pdfiumInstance) FPDFLink_GetRect(request *requests.FPDFLink_GetRect) (*responses.FPDFLink_GetRect, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_GetRect(request)
}

func (i *pdfiumInstance) FPDFLink_GetTextRange(request *requests.FPDFLink_GetTextRange) (*responses.FPDFLink_GetTextRange, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_GetTextRange(request)
}

func (i *pdfiumInstance) FPDFLink_GetURL(request *requests.FPDFLink_GetURL) (*responses.FPDFLink_GetURL, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_GetURL(request)
}

func (i *pdfiumInstance) FPDFLink_LoadWebLinks(request *requests.FPDFLink_LoadWebLinks) (*responses.FPDFLink_LoadWebLinks, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFLink_LoadWebLinks(request)
}

func (i *pdfiumInstance) FPDFPage_Flatten(request *requests.FPDFPage_Flatten) (*responses.FPDFPage_Flatten, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFPage_Flatten(request)
}

func (i *pdfiumInstance) FPDFPage_GetDecodedThumbnailData(request *requests.FPDFPage_GetDecodedThumbnailData) (*responses.FPDFPage_GetDecodedThumbnailData, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFPage_GetDecodedThumbnailData(request)
}

func (i *pdfiumInstance) FPDFPage_GetRawThumbnailData(request *requests.FPDFPage_GetRawThumbnailData) (*responses.FPDFPage_GetRawThumbnailData, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFPage_GetRawThumbnailData(request)
}

func (i *pdfiumInstance) FPDFPage_GetRotation(request *requests.FPDFPage_GetRotation) (*responses.FPDFPage_GetRotation, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFPage_GetRotation(request)
}

func (i *pdfiumInstance) FPDFPage_GetThumbnailAsBitmap(request *requests.FPDFPage_GetThumbnailAsBitmap) (*responses.FPDFPage_GetThumbnailAsBitmap, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFPage_GetThumbnailAsBitmap(request)
}

func (i *pdfiumInstance) FPDFPage_HasTransparency(request *requests.FPDFPage_HasTransparency) (*responses.FPDFPage_HasTransparency, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFPage_HasTransparency(request)
}

func (i *pdfiumInstance) FPDFPage_SetRotation(request *requests.FPDFPage_SetRotation) (*responses.FPDFPage_SetRotation, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFPage_SetRotation(request)
}

func (i *pdfiumInstance) FPDFSignatureObj_GetByteRange(request *requests.FPDFSignatureObj_GetByteRange) (*responses.FPDFSignatureObj_GetByteRange, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFSignatureObj_GetByteRange(request)
}

func (i *pdfiumInstance) FPDFSignatureObj_GetContents(request *requests.FPDFSignatureObj_GetContents) (*responses.FPDFSignatureObj_GetContents, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFSignatureObj_GetContents(request)
}

func (i *pdfiumInstance) FPDFSignatureObj_GetDocMDPPermission(request *requests.FPDFSignatureObj_GetDocMDPPermission) (*responses.FPDFSignatureObj_GetDocMDPPermission, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFSignatureObj_GetDocMDPPermission(request)
}

func (i *pdfiumInstance) FPDFSignatureObj_GetReason(request *requests.FPDFSignatureObj_GetReason) (*responses.FPDFSignatureObj_GetReason, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFSignatureObj_GetReason(request)
}

func (i *pdfiumInstance) FPDFSignatureObj_GetSubFilter(request *requests.FPDFSignatureObj_GetSubFilter) (*responses.FPDFSignatureObj_GetSubFilter, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFSignatureObj_GetSubFilter(request)
}

func (i *pdfiumInstance) FPDFSignatureObj_GetTime(request *requests.FPDFSignatureObj_GetTime) (*responses.FPDFSignatureObj_GetTime, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFSignatureObj_GetTime(request)
}

func (i *pdfiumInstance) FPDFText_ClosePage(request *requests.FPDFText_ClosePage) (*responses.FPDFText_ClosePage, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_ClosePage(request)
}

func (i *pdfiumInstance) FPDFText_CountChars(request *requests.FPDFText_CountChars) (*responses.FPDFText_CountChars, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_CountChars(request)
}

func (i *pdfiumInstance) FPDFText_CountRects(request *requests.FPDFText_CountRects) (*responses.FPDFText_CountRects, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_CountRects(request)
}

func (i *pdfiumInstance) FPDFText_FindClose(request *requests.FPDFText_FindClose) (*responses.FPDFText_FindClose, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_FindClose(request)
}

func (i *pdfiumInstance) FPDFText_FindNext(request *requests.FPDFText_FindNext) (*responses.FPDFText_FindNext, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_FindNext(request)
}

func (i *pdfiumInstance) FPDFText_FindPrev(request *requests.FPDFText_FindPrev) (*responses.FPDFText_FindPrev, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_FindPrev(request)
}

func (i *pdfiumInstance) FPDFText_FindStart(request *requests.FPDFText_FindStart) (*responses.FPDFText_FindStart, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_FindStart(request)
}

func (i *pdfiumInstance) FPDFText_GetBoundedText(request *requests.FPDFText_GetBoundedText) (*responses.FPDFText_GetBoundedText, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetBoundedText(request)
}

func (i *pdfiumInstance) FPDFText_GetCharAngle(request *requests.FPDFText_GetCharAngle) (*responses.FPDFText_GetCharAngle, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetCharAngle(request)
}

func (i *pdfiumInstance) FPDFText_GetCharBox(request *requests.FPDFText_GetCharBox) (*responses.FPDFText_GetCharBox, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetCharBox(request)
}

func (i *pdfiumInstance) FPDFText_GetCharIndexAtPos(request *requests.FPDFText_GetCharIndexAtPos) (*responses.FPDFText_GetCharIndexAtPos, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetCharIndexAtPos(request)
}

func (i *pdfiumInstance) FPDFText_GetCharIndexFromTextIndex(request *requests.FPDFText_GetCharIndexFromTextIndex) (*responses.FPDFText_GetCharIndexFromTextIndex, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetCharIndexFromTextIndex(request)
}

func (i *pdfiumInstance) FPDFText_GetCharOrigin(request *requests.FPDFText_GetCharOrigin) (*responses.FPDFText_GetCharOrigin, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetCharOrigin(request)
}

func (i *pdfiumInstance) FPDFText_GetFillColor(request *requests.FPDFText_GetFillColor) (*responses.FPDFText_GetFillColor, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetFillColor(request)
}

func (i *pdfiumInstance) FPDFText_GetFontInfo(request *requests.FPDFText_GetFontInfo) (*responses.FPDFText_GetFontInfo, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetFontInfo(request)
}

func (i *pdfiumInstance) FPDFText_GetFontSize(request *requests.FPDFText_GetFontSize) (*responses.FPDFText_GetFontSize, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetFontSize(request)
}

func (i *pdfiumInstance) FPDFText_GetFontWeight(request *requests.FPDFText_GetFontWeight) (*responses.FPDFText_GetFontWeight, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetFontWeight(request)
}

func (i *pdfiumInstance) FPDFText_GetLooseCharBox(request *requests.FPDFText_GetLooseCharBox) (*responses.FPDFText_GetLooseCharBox, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetLooseCharBox(request)
}

func (i *pdfiumInstance) FPDFText_GetMatrix(request *requests.FPDFText_GetMatrix) (*responses.FPDFText_GetMatrix, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetMatrix(request)
}

func (i *pdfiumInstance) FPDFText_GetRect(request *requests.FPDFText_GetRect) (*responses.FPDFText_GetRect, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetRect(request)
}

func (i *pdfiumInstance) FPDFText_GetSchCount(request *requests.FPDFText_GetSchCount) (*responses.FPDFText_GetSchCount, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetSchCount(request)
}

func (i *pdfiumInstance) FPDFText_GetSchResultIndex(request *requests.FPDFText_GetSchResultIndex) (*responses.FPDFText_GetSchResultIndex, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetSchResultIndex(request)
}

func (i *pdfiumInstance) FPDFText_GetStrokeColor(request *requests.FPDFText_GetStrokeColor) (*responses.FPDFText_GetStrokeColor, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetStrokeColor(request)
}

func (i *pdfiumInstance) FPDFText_GetText(request *requests.FPDFText_GetText) (*responses.FPDFText_GetText, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetText(request)
}

func (i *pdfiumInstance) FPDFText_GetTextIndexFromCharIndex(request *requests.FPDFText_GetTextIndexFromCharIndex) (*responses.FPDFText_GetTextIndexFromCharIndex, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetTextIndexFromCharIndex(request)
}

func (i *pdfiumInstance) FPDFText_GetTextRenderMode(request *requests.FPDFText_GetTextRenderMode) (*responses.FPDFText_GetTextRenderMode, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetTextRenderMode(request)
}

func (i *pdfiumInstance) FPDFText_GetUnicode(request *requests.FPDFText_GetUnicode) (*responses.FPDFText_GetUnicode, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_GetUnicode(request)
}

func (i *pdfiumInstance) FPDFText_LoadPage(request *requests.FPDFText_LoadPage) (*responses.FPDFText_LoadPage, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDFText_LoadPage(request)
}

func (i *pdfiumInstance) FPDF_CloseDocument(request *requests.FPDF_CloseDocument) (*responses.FPDF_CloseDocument, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_CloseDocument(request)
}

func (i *pdfiumInstance) FPDF_ClosePage(request *requests.FPDF_ClosePage) (*responses.FPDF_ClosePage, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_ClosePage(request)
}

func (i *pdfiumInstance) FPDF_CloseXObject(request *requests.FPDF_CloseXObject) (*responses.FPDF_CloseXObject, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_CloseXObject(request)
}

func (i *pdfiumInstance) FPDF_CopyViewerPreferences(request *requests.FPDF_CopyViewerPreferences) (*responses.FPDF_CopyViewerPreferences, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_CopyViewerPreferences(request)
}

func (i *pdfiumInstance) FPDF_CreateNewDocument(request *requests.FPDF_CreateNewDocument) (*responses.FPDF_CreateNewDocument, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_CreateNewDocument(request)
}

func (i *pdfiumInstance) FPDF_GetDocPermissions(request *requests.FPDF_GetDocPermissions) (*responses.FPDF_GetDocPermissions, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetDocPermissions(request)
}

func (i *pdfiumInstance) FPDF_GetFileIdentifier(request *requests.FPDF_GetFileIdentifier) (*responses.FPDF_GetFileIdentifier, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetFileIdentifier(request)
}

func (i *pdfiumInstance) FPDF_GetFileVersion(request *requests.FPDF_GetFileVersion) (*responses.FPDF_GetFileVersion, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetFileVersion(request)
}

func (i *pdfiumInstance) FPDF_GetLastError(request *requests.FPDF_GetLastError) (*responses.FPDF_GetLastError, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetLastError(request)
}

func (i *pdfiumInstance) FPDF_GetMetaText(request *requests.FPDF_GetMetaText) (*responses.FPDF_GetMetaText, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetMetaText(request)
}

func (i *pdfiumInstance) FPDF_GetPageAAction(request *requests.FPDF_GetPageAAction) (*responses.FPDF_GetPageAAction, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetPageAAction(request)
}

func (i *pdfiumInstance) FPDF_GetPageCount(request *requests.FPDF_GetPageCount) (*responses.FPDF_GetPageCount, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetPageCount(request)
}

func (i *pdfiumInstance) FPDF_GetPageHeight(request *requests.FPDF_GetPageHeight) (*responses.FPDF_GetPageHeight, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetPageHeight(request)
}

func (i *pdfiumInstance) FPDF_GetPageLabel(request *requests.FPDF_GetPageLabel) (*responses.FPDF_GetPageLabel, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetPageLabel(request)
}

func (i *pdfiumInstance) FPDF_GetPageSizeByIndex(request *requests.FPDF_GetPageSizeByIndex) (*responses.FPDF_GetPageSizeByIndex, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetPageSizeByIndex(request)
}

func (i *pdfiumInstance) FPDF_GetPageWidth(request *requests.FPDF_GetPageWidth) (*responses.FPDF_GetPageWidth, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetPageWidth(request)
}

func (i *pdfiumInstance) FPDF_GetSecurityHandlerRevision(request *requests.FPDF_GetSecurityHandlerRevision) (*responses.FPDF_GetSecurityHandlerRevision, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetSecurityHandlerRevision(request)
}

func (i *pdfiumInstance) FPDF_GetSignatureCount(request *requests.FPDF_GetSignatureCount) (*responses.FPDF_GetSignatureCount, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetSignatureCount(request)
}

func (i *pdfiumInstance) FPDF_GetSignatureObject(request *requests.FPDF_GetSignatureObject) (*responses.FPDF_GetSignatureObject, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_GetSignatureObject(request)
}

func (i *pdfiumInstance) FPDF_ImportNPagesToOne(request *requests.FPDF_ImportNPagesToOne) (*responses.FPDF_ImportNPagesToOne, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_ImportNPagesToOne(request)
}

func (i *pdfiumInstance) FPDF_ImportPages(request *requests.FPDF_ImportPages) (*responses.FPDF_ImportPages, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_ImportPages(request)
}

func (i *pdfiumInstance) FPDF_ImportPagesByIndex(request *requests.FPDF_ImportPagesByIndex) (*responses.FPDF_ImportPagesByIndex, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_ImportPagesByIndex(request)
}

func (i *pdfiumInstance) FPDF_LoadCustomDocument(request *requests.FPDF_LoadCustomDocument) (*responses.FPDF_LoadCustomDocument, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	// Since multi-threaded usage implements gRPC, it can't serialize the reader onto that.
	// To make it support the full interface, we just completely read the file into memory.
	fileData, err := ioutil.ReadAll(request.Reader)
	if err != nil {
		return nil, err
	}

	doc, err := i.OpenDocument(&requests.OpenDocument{
		File:     &fileData,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &responses.FPDF_LoadCustomDocument{Document: doc.Document}, nil
}

func (i *pdfiumInstance) FPDF_LoadDocument(request *requests.FPDF_LoadDocument) (*responses.FPDF_LoadDocument, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_LoadDocument(request)
}

func (i *pdfiumInstance) FPDF_LoadMemDocument(request *requests.FPDF_LoadMemDocument) (*responses.FPDF_LoadMemDocument, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_LoadMemDocument(request)
}

func (i *pdfiumInstance) FPDF_LoadMemDocument64(request *requests.FPDF_LoadMemDocument64) (*responses.FPDF_LoadMemDocument64, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_LoadMemDocument64(request)
}

func (i *pdfiumInstance) FPDF_LoadPage(request *requests.FPDF_LoadPage) (*responses.FPDF_LoadPage, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_LoadPage(request)
}

func (i *pdfiumInstance) FPDF_NewFormObjectFromXObject(request *requests.FPDF_NewFormObjectFromXObject) (*responses.FPDF_NewFormObjectFromXObject, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_NewFormObjectFromXObject(request)
}

func (i *pdfiumInstance) FPDF_NewXObjectFromPage(request *requests.FPDF_NewXObjectFromPage) (*responses.FPDF_NewXObjectFromPage, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_NewXObjectFromPage(request)
}

func (i *pdfiumInstance) FPDF_SaveAsCopy(request *requests.FPDF_SaveAsCopy) (*responses.FPDF_SaveAsCopy, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_SaveAsCopy(request)
}

func (i *pdfiumInstance) FPDF_SaveWithVersion(request *requests.FPDF_SaveWithVersion) (*responses.FPDF_SaveWithVersion, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_SaveWithVersion(request)
}

func (i *pdfiumInstance) FPDF_SetSandBoxPolicy(request *requests.FPDF_SetSandBoxPolicy) (*responses.FPDF_SetSandBoxPolicy, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FPDF_SetSandBoxPolicy(request)
}

func (i *pdfiumInstance) FSDK_SetLocaltimeFunction(request *requests.FSDK_SetLocaltimeFunction) (*responses.FSDK_SetLocaltimeFunction, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FSDK_SetLocaltimeFunction(request)
}

func (i *pdfiumInstance) FSDK_SetTimeFunction(request *requests.FSDK_SetTimeFunction) (*responses.FSDK_SetTimeFunction, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FSDK_SetTimeFunction(request)
}

func (i *pdfiumInstance) FSDK_SetUnSpObjProcessHandler(request *requests.FSDK_SetUnSpObjProcessHandler) (*responses.FSDK_SetUnSpObjProcessHandler, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.FSDK_SetUnSpObjProcessHandler(request)
}

func (i *pdfiumInstance) GetActionInfo(request *requests.GetActionInfo) (*responses.GetActionInfo, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.GetActionInfo(request)
}

func (i *pdfiumInstance) GetAttachments(request *requests.GetAttachments) (*responses.GetAttachments, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.GetAttachments(request)
}

func (i *pdfiumInstance) GetBookmarks(request *requests.GetBookmarks) (*responses.GetBookmarks, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.GetBookmarks(request)
}

func (i *pdfiumInstance) GetDestInfo(request *requests.GetDestInfo) (*responses.GetDestInfo, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.GetDestInfo(request)
}

func (i *pdfiumInstance) GetJavaScriptActions(request *requests.GetJavaScriptActions) (*responses.GetJavaScriptActions, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.GetJavaScriptActions(request)
}

func (i *pdfiumInstance) GetMetaData(request *requests.GetMetaData) (*responses.GetMetaData, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.GetMetaData(request)
}

func (i *pdfiumInstance) GetPageSize(request *requests.GetPageSize) (*responses.GetPageSize, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.GetPageSize(request)
}

func (i *pdfiumInstance) GetPageSizeInPixels(request *requests.GetPageSizeInPixels) (*responses.GetPageSizeInPixels, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.GetPageSizeInPixels(request)
}

func (i *pdfiumInstance) GetPageText(request *requests.GetPageText) (*responses.GetPageText, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.GetPageText(request)
}

func (i *pdfiumInstance) GetPageTextStructured(request *requests.GetPageTextStructured) (*responses.GetPageTextStructured, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.GetPageTextStructured(request)
}

func (i *pdfiumInstance) OpenDocument(request *requests.OpenDocument) (*responses.OpenDocument, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.OpenDocument(request)
}

func (i *pdfiumInstance) RenderPageInDPI(request *requests.RenderPageInDPI) (*responses.RenderPageInDPI, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.RenderPageInDPI(request)
}

func (i *pdfiumInstance) RenderPageInPixels(request *requests.RenderPageInPixels) (*responses.RenderPageInPixels, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.RenderPageInPixels(request)
}

func (i *pdfiumInstance) RenderPagesInDPI(request *requests.RenderPagesInDPI) (*responses.RenderPagesInDPI, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.RenderPagesInDPI(request)
}

func (i *pdfiumInstance) RenderPagesInPixels(request *requests.RenderPagesInPixels) (*responses.RenderPagesInPixels, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.RenderPagesInPixels(request)
}

func (i *pdfiumInstance) RenderToFile(request *requests.RenderToFile) (*responses.RenderToFile, error) {
	if i.closed {
		return nil, errors.New("instance is closed")
	}

	return i.worker.plugin.RenderToFile(request)
}

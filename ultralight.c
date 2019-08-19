#include <Ultralight/CAPI.h>
#include <AppCore/CAPI.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

extern void appUpdateFunction(void*);
extern void viewChangeTitleFunction(void*, ULView, ULString);
extern void viewChangeURLFunction(void*, ULView, ULString);
extern void viewChangeTooltipFunction(void*, ULView, ULString);
extern void viewChangeCursorFunction(void*, ULView, ULCursor);
extern void viewAddConsoleMessageFunction(void*, ULView, ULMessageSource, ULMessageLevel, ULString, unsigned int, unsigned int, ULString);
extern void viewBeginLoadingFunction(void*, ULView);
extern void viewFinishLoadingFunction(void*, ULView);
extern void viewUpdateHistoryFunction(void*, ULView);
extern void viewDOMReadyFunction(void*, ULView);
extern void winCloseFunction(void*);
extern void winResizeFunction(void*, int, int);
extern JSValueRef objFunctionCallback(JSContextRef ctx, JSObjectRef function, JSObjectRef thisObject,
                                      size_t argumentCount, JSValueRef *arguments, JSValueRef* exception);

static inline char * printParams(JSContextRef ctx, JSValueRef *params, size_t count) {
	if (count == 0) {
		return "";
	}
	size_t destLen = 0;
	char delimeter[4] = "\t\t\t";

	for (int i = 0; i < count; i++) {
		if (JSValueIsString(ctx, *params)) {
			JSStringRef strOut = JSValueToStringCopy(ctx, (params)[i], NULL);
			destLen += (JSStringGetLength(strOut)+1);
			if (i != (count-1)) {
				destLen += sizeof(delimeter);
			}
		}
	}

	char output[destLen];	

	for (int i = 0; i < count; i++) {
		if (JSValueIsString(ctx, *params)) {
			JSStringRef strOut = JSValueToStringCopy(ctx, (params)[i], NULL);
			int length = (JSStringGetLength(strOut)+1);
			char value[length];
			JSStringGetUTF8CString(strOut, value, length);
			if (i == 0) {
				sprintf(output, "%s", value);
			} else {
				sprintf(output + strlen(output), "%s", value);
			}
			if (i != (count-1)) {
				sprintf(output + strlen(output), "%s", delimeter);
			}
		} else {
			printf("Not a string, have %s\n", (JSType)JSValueGetType(ctx, *params));
		}
	}
	char * out = malloc(destLen * sizeof(char));
	for (int i = 0; i < destLen; i++) {
		out[i] = output[i];
	}
	return out;
}

static inline void setAppUpdateCallback(ULApp app) {
	ulAppSetUpdateCallback(app, appUpdateFunction, NULL);
}

static inline void setViewChangeTitleCallback(ULView view) {
	ulViewSetChangeTitleCallback(view, viewChangeTitleFunction, NULL);
}

static inline void setViewChangeURLCallback(ULView view) {
	ulViewSetChangeURLCallback(view, viewChangeURLFunction, NULL);
}

static inline void setViewChangeTooltipCallback(ULView view) {
	ulViewSetChangeTooltipCallback(view, viewChangeTooltipFunction, NULL);
}

static inline void setViewChangeCursorCallback(ULView view) {
	ulViewSetChangeCursorCallback(view, viewChangeCursorFunction, NULL);
}

static inline void setViewAddConsoleMessageCallback(ULView view) {
	ulViewSetAddConsoleMessageCallback(view, viewAddConsoleMessageFunction, NULL);
}

static inline void setViewBeginLoadingCallback(ULView view) {
	ulViewSetBeginLoadingCallback(view, viewBeginLoadingFunction, NULL);
}

static inline void setViewFinishLoadingCallback(ULView view) {
	ulViewSetFinishLoadingCallback(view, viewFinishLoadingFunction, NULL);
}

static inline void setViewUpdateHistoryCallback(ULView view) {
	ulViewSetUpdateHistoryCallback(view, viewUpdateHistoryFunction, NULL);
}

static inline void setViewDOMReadyCallback(ULView view) {
	ulViewSetDOMReadyCallback(view, viewDOMReadyFunction, NULL);
}

static inline void setWinCloseCallback(ULWindow win) {
	ulWindowSetCloseCallback(win, winCloseFunction, NULL);
}

static inline void setWinResizeCallback(ULWindow win) {
	ulWindowSetResizeCallback(win, winResizeFunction, NULL);
}

static inline char * strconv(ULString str) {
	if (ulStringGetLength(str) == 0) {
		return "";
	}
	size_t len = ulStringGetLength(str);
	ULChar16 *val = ulStringGetData(str);
	char string[len];
	for (int i = 0; i < len; i++) {	
		if (i == 0) {
			sprintf(string, "%c", (char)*val & 0x00FF);
		} else {
			sprintf(string + strlen(string), "%c", (char)*val & 0x00FF);
		}
		val++;
	}
	char *num = malloc(len);
	for (int i = 0; i < len; i++) {
		num[i] = string[i];
	}
	return num;
}

static inline char * evaluateScript(ULView view, ULString script) {
	const JSStringRef values = JSValueToStringCopy(ulViewGetJSContext(view), ulViewEvaluateScript(view, script), NULL);
	int length = (JSStringGetLength(values)+1);
	char value[length];
	JSStringGetUTF8CString(values, value, sizeof(value));
	// Required to be accessible from Go
	char *out = malloc(length * sizeof(char));
	for (int i = 0; i < length; i++) {
		out[i] = value[i];
	}
	return out;
}

static inline JSObjectRef bindScript(ULView view, char* name) {
	JSContextRef ctx = ulViewGetJSContext(view);
	JSObjectRef func = JSObjectMakeFunctionWithCallback(ctx, NULL, (JSObjectCallAsFunctionCallback)objFunctionCallback);
	JSObjectSetProperty(ctx, JSContextGetGlobalObject(ctx), JSStringCreateWithUTF8CString(name), func, 0, NULL);
	return func;
}
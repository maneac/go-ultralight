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
extern void winResizeFunction(void*, unsigned int, unsigned int);
extern JSValueRef objFunctionCallback(JSContextRef ctx, JSObjectRef function, JSObjectRef thisObject,
                                      size_t argumentCount, JSValueRef *arguments, JSValueRef* exception);
char * printParams(JSContextRef ctx, JSValueRef *params, size_t count);
void setAppUpdateCallback(ULApp app);
void setViewChangeTitleCallback(ULView view);
void setViewChangeURLCallback(ULView view);
void setViewChangeTooltipCallback(ULView view);
void setViewChangeCursorCallback(ULView view);
void setViewAddConsoleMessageCallback(ULView view);
void setViewBeginLoadingCallback(ULView view);
void setViewFinishLoadingCallback(ULView view);
void setViewUpdateHistoryCallback(ULView view);
void setViewDOMReadyCallback(ULView view);
void setWinCloseCallback(ULWindow win);
void setWinResizeCallback(ULWindow win);
char * strconv(ULString str);
char * evaluateScript(ULView view, ULString script);
JSObjectRef bindScript(ULView view, char* name);
// Supported return types for JS Callback functions
JSValueRef makeJSValueString(ULView view, char * str);
JSValueRef makeJSValueBool(ULView view, bool val);
JSValueRef makeJSValueNum(ULView view, double val);
JSValueRef makeJSValueJSON(ULView view, char * val);
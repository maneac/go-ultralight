#include <ultralight.h>

char * printParams(JSContextRef ctx, JSValueRef *params, size_t count) {
	if (count == 0) {
		return "";
	}
	size_t destLen = 0;
	char delimeter[4] = "\t\t\t";
	int i;
	for (i = 0; i < count; i++) {
		if (JSValueIsString(ctx, *params)) {
			JSStringRef strOut = JSValueToStringCopy(ctx, (params)[i], NULL);
			destLen += (JSStringGetLength(strOut)+1);
			if (i != (count-1)) {
				destLen += sizeof(delimeter);
			}
		}
	}

	char output[destLen];	

	for (i = 0; i < count; i++) {
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
			printf("Not a string, have %d\n", (JSType)JSValueGetType(ctx, *params));
		}
	}
	char * out = malloc(destLen * sizeof(char));
	for (i = 0; i < destLen; i++) {
		out[i] = output[i];
	}
	return out;
}

void setAppUpdateCallback(ULApp app) {
	ulAppSetUpdateCallback(app, appUpdateFunction, NULL);
}

void setViewChangeTitleCallback(ULView view) {
	ulViewSetChangeTitleCallback(view, viewChangeTitleFunction, NULL);
}

void setViewChangeURLCallback(ULView view) {
	ulViewSetChangeURLCallback(view, viewChangeURLFunction, NULL);
}

void setViewChangeTooltipCallback(ULView view) {
	ulViewSetChangeTooltipCallback(view, viewChangeTooltipFunction, NULL);
}

void setViewChangeCursorCallback(ULView view) {
	ulViewSetChangeCursorCallback(view, viewChangeCursorFunction, NULL);
}

void setViewAddConsoleMessageCallback(ULView view) {
	ulViewSetAddConsoleMessageCallback(view, viewAddConsoleMessageFunction, NULL);
}

void setViewBeginLoadingCallback(ULView view) {
	ulViewSetBeginLoadingCallback(view, viewBeginLoadingFunction, NULL);
}

void setViewFinishLoadingCallback(ULView view) {
	ulViewSetFinishLoadingCallback(view, viewFinishLoadingFunction, NULL);
}

void setViewUpdateHistoryCallback(ULView view) {
	ulViewSetUpdateHistoryCallback(view, viewUpdateHistoryFunction, NULL);
}

void setViewDOMReadyCallback(ULView view) {
	ulViewSetDOMReadyCallback(view, viewDOMReadyFunction, NULL);
}

void setWinCloseCallback(ULWindow win) {
	ulWindowSetCloseCallback(win, winCloseFunction, NULL);
}

void setWinResizeCallback(ULWindow win) {
	ulWindowSetResizeCallback(win, winResizeFunction, NULL);
}

char * strconv(ULString str) {
	if (ulStringGetLength(str) == 0) {
		return "";
	}
	size_t len = ulStringGetLength(str);
	ULChar16 *val = ulStringGetData(str);
	char string[len];
	int i;
	for (i = 0; i < len; i++) {	
		if (i == 0) {
			sprintf(string, "%c", (char)*val & 0x00FF);
		} else {
			sprintf(string + strlen(string), "%c", (char)*val & 0x00FF);
		}
		val++;
	}
	char *num = malloc(len);
	for (i = 0; i < len; i++) {
		num[i] = string[i];
	}
	return num;
}

char * evaluateScript(ULView view, ULString script) {
	const JSStringRef values = JSValueToStringCopy(ulViewGetJSContext(view), ulViewEvaluateScript(view, script), NULL);
	int length = (JSStringGetLength(values)+1);
	char value[length];
	JSStringGetUTF8CString(values, value, sizeof(value));
	// Required to be accessible from Go
	char *out = malloc(length * sizeof(char));
	int i;
	for (i = 0; i < length; i++) {
		out[i] = value[i];
	}
	return out;
}

JSObjectRef bindScript(ULView view, char* name) {
	JSContextRef ctx = ulViewGetJSContext(view);
	JSObjectRef func = JSObjectMakeFunctionWithCallback(ctx, NULL, (JSObjectCallAsFunctionCallback)objFunctionCallback);
	JSObjectSetProperty(ctx, JSContextGetGlobalObject(ctx), JSStringCreateWithUTF8CString(name), func, 0, NULL);
	return func;
}
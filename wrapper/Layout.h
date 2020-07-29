#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggPushID(char const *id);
extern void iggPushIDInt(int id);
extern void iggPopID(void);

extern void iggSeparator(void);
extern void iggSameLine(float posX, float spacingW);
extern void iggSpacing(void);
extern void iggDummy(IggVec2 const *size);
extern void iggBeginGroup(void);
extern void iggEndGroup(void);
extern void iggIndent(float indent_w);
extern void iggUnindent(float indent_w);

extern void iggCursorPos(IggVec2 *pos);
extern float iggCursorPosX(void);
extern float iggCursorPosY(void);
extern void iggCursorStartPos(IggVec2 *pos);
extern void iggCursorScreenPos(IggVec2 *pos);

extern void iggSetCursorPos(IggVec2 const *localPos);
extern void iggSetCursorScreenPos(IggVec2 const *absPos);
extern void iggAlignTextToFramePadding();
extern float iggGetTextLineHeight(void);
extern float iggGetTextLineHeightWithSpacing(void);
extern float iggGetFrameHeight(void);
extern float iggGetFrameHeightWithSpacing(void);

extern void iggBeginHorizontal(char const *id, IggVec2 const *size, float align);
extern void iggEndHorizontal();
extern void iggBeginVertical(char const *id, IggVec2 const *size, float align);
extern void iggEndVertical();
extern void iggSpring(float weight, float spacing);
extern void iggSuspendLayout();
extern void iggResumeLayout();

#ifdef __cplusplus
}
#endif

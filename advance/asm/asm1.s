#include "textflag.h"

TEXT ·Sout(SB), NOSPLIT, $0
    MOVQ $(0x2000000+4), AX
    MOVQ fd+0(FP), DI
    MOVQ msg_data+8(FP), SI
    MOVQ msg_len+16(FP), DX
    SYSCALL
    MOVQ AX, ret+0(FP)
    RET

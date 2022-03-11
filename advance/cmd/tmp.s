"".main STEXT size=421 args=0x0 locals=0xe0 funcid=0x0
        0x0000 00000 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     TEXT    "".main(SB), ABIInternal, $224-0
        0x0000 00000 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     LEAQ    -96(SP), R12
        0x0005 00005 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     CMPQ    R12, 16(R14)
        0x0009 00009 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     PCDATA  $0, $-2
        0x0009 00009 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     JLS     410
        0x000f 00015 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     PCDATA  $0, $-1
        0x000f 00015 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     SUBQ    $224, SP
        0x0016 00022 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     MOVQ    BP, 216(SP)
        0x001e 00030 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     LEAQ    216(SP), BP
        0x0026 00038 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     FUNCDATA        $0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x0026 00038 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     FUNCDATA        $1, gclocals·b9184dd1ef9d9df7f9e3c00538f23cdd(SB)
        0x0026 00038 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     FUNCDATA        $2, "".main.stkobj(SB)
        0x0026 00038 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:13)     PCDATA  $1, $0
        0x0026 00038 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:13)     CALL    "".test1(SB)
        0x002b 00043 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:13)     MOVQ    AX, "".v1+56(SP)
        0x0030 00048 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:13)     MOVQ    BX, "".v1+64(SP)
        0x0035 00053 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:13)     MOVQ    CX, "".v1+72(SP)
        0x003a 00058 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:13)     MOVQ    DI, "".v1+80(SP)
        0x003f 00063 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:13)     NOP
        0x0040 00064 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:14)     CALL    "".test2(SB)
        0x0045 00069 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:14)     MOVQ    AX, "".v2+24(SP)
        0x004a 00074 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:14)     MOVQ    BX, "".v2+32(SP)
        0x004f 00079 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:14)     MOVQ    CX, "".v2+40(SP)
        0x0054 00084 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:14)     MOVQ    DI, "".v2+48(SP)
        0x0059 00089 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    "".v1+56(SP), AX
        0x005e 00094 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    "".v1+64(SP), CX
        0x0063 00099 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    "".v1+72(SP), DX
        0x0068 00104 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    "".v1+80(SP), BX
        0x006d 00109 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    AX, ""..autotmp_2+120(SP)
        0x0072 00114 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    CX, ""..autotmp_2+128(SP)
        0x007a 00122 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    DX, ""..autotmp_2+136(SP)
        0x0082 00130 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    BX, ""..autotmp_2+144(SP)
        0x008a 00138 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    "".v2+24(SP), AX
        0x008f 00143 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    "".v2+32(SP), CX
        0x0094 00148 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    "".v2+40(SP), DX
        0x0099 00153 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    "".v2+48(SP), BX
        0x009e 00158 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    AX, ""..autotmp_3+88(SP)
        0x00a3 00163 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    CX, ""..autotmp_3+96(SP)
        0x00a8 00168 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    DX, ""..autotmp_3+104(SP)
        0x00ad 00173 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    BX, ""..autotmp_3+112(SP)
        0x00b2 00178 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     LEAQ    ""..autotmp_4+184(SP), AX
        0x00ba 00186 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVUPS  X15, (AX)
        0x00be 00190 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     LEAQ    ""..autotmp_4+200(SP), AX
        0x00c6 00198 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVUPS  X15, (AX)
        0x00ca 00202 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     LEAQ    ""..autotmp_4+184(SP), AX
        0x00d2 00210 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    AX, ""..autotmp_6+152(SP)
        0x00da 00218 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     LEAQ    type."".T(SB), AX
        0x00e1 00225 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     LEAQ    ""..autotmp_2+120(SP), BX
        0x00e6 00230 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     PCDATA  $1, $1
        0x00e6 00230 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     CALL    runtime.convT2Enoptr(SB)
        0x00eb 00235 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    ""..autotmp_6+152(SP), CX
        0x00f3 00243 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     TESTB   AL, (CX)
        0x00f5 00245 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    AX, (CX)
        0x00f8 00248 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     LEAQ    8(CX), DI
        0x00fc 00252 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     PCDATA  $0, $-2
        0x00fc 00252 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     CMPL    runtime.writeBarrier(SB), $0
        0x0103 00259 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     JEQ     263
        0x0105 00261 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     JMP     269
        0x0107 00263 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    BX, 8(CX)
        0x010b 00267 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     JMP     276
        0x010d 00269 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     CALL    runtime.gcWriteBarrierBX(SB)
        0x0112 00274 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     JMP     276
        0x0114 00276 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     PCDATA  $0, $-1
        0x0114 00276 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     LEAQ    type."".T(SB), AX
        0x011b 00283 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     LEAQ    ""..autotmp_3+88(SP), BX
        0x0120 00288 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     CALL    runtime.convT2Enoptr(SB)
        0x0125 00293 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    ""..autotmp_6+152(SP), CX
        0x012d 00301 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     TESTB   AL, (CX)
        0x012f 00303 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    AX, 16(CX)
        0x0133 00307 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     LEAQ    24(CX), DI
        0x0137 00311 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     PCDATA  $0, $-2
        0x0137 00311 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     CMPL    runtime.writeBarrier(SB), $0
        0x013e 00318 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     NOP
        0x0140 00320 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     JEQ     324
        0x0142 00322 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     JMP     330
        0x0144 00324 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    BX, 24(CX)
        0x0148 00328 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     JMP     337
        0x014a 00330 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     CALL    runtime.gcWriteBarrierBX(SB)
        0x014f 00335 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     JMP     337
        0x0151 00337 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     PCDATA  $0, $-1
        0x0151 00337 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    ""..autotmp_6+152(SP), AX
        0x0159 00345 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     TESTB   AL, (AX)
        0x015b 00347 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     JMP     349
        0x015d 00349 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    AX, ""..autotmp_5+160(SP)
        0x0165 00357 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    $2, ""..autotmp_5+168(SP)
        0x0171 00369 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    $2, ""..autotmp_5+176(SP)
        0x017d 00381 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVL    $2, BX
        0x0182 00386 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     MOVQ    BX, CX
        0x0185 00389 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     PCDATA  $1, $0
        0x0185 00389 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:15)     CALL    fmt.Println(SB)
        0x018a 00394 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:16)     MOVQ    216(SP), BP
        0x0192 00402 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:16)     ADDQ    $224, SP
        0x0199 00409 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:16)     RET
        0x019a 00410 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:16)     NOP
        0x019a 00410 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     PCDATA  $1, $-1
        0x019a 00410 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     PCDATA  $0, $-2
        0x019a 00410 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     CALL    runtime.morestack_noctxt(SB)
        0x019f 00415 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     PCDATA  $0, $-1
        0x019f 00415 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     NOP
        0x01a0 00416 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:12)     JMP     0
        0x0000 4c 8d 64 24 a0 4d 3b 66 10 0f 86 8b 01 00 00 48  L.d$.M;f.......H
        0x0010 81 ec e0 00 00 00 48 89 ac 24 d8 00 00 00 48 8d  ......H..$....H.
        0x0020 ac 24 d8 00 00 00 e8 00 00 00 00 48 89 44 24 38  .$.........H.D$8
        0x0030 48 89 5c 24 40 48 89 4c 24 48 48 89 7c 24 50 90  H.\$@H.L$HH.|$P.
        0x0040 e8 00 00 00 00 48 89 44 24 18 48 89 5c 24 20 48  .....H.D$.H.\$ H
        0x0050 89 4c 24 28 48 89 7c 24 30 48 8b 44 24 38 48 8b  .L$(H.|$0H.D$8H.
        0x0060 4c 24 40 48 8b 54 24 48 48 8b 5c 24 50 48 89 44  L$@H.T$HH.\$PH.D
        0x0070 24 78 48 89 8c 24 80 00 00 00 48 89 94 24 88 00  $xH..$....H..$..
        0x0080 00 00 48 89 9c 24 90 00 00 00 48 8b 44 24 18 48  ..H..$....H.D$.H
        0x0090 8b 4c 24 20 48 8b 54 24 28 48 8b 5c 24 30 48 89  .L$ H.T$(H.\$0H.
        0x00a0 44 24 58 48 89 4c 24 60 48 89 54 24 68 48 89 5c  D$XH.L$`H.T$hH.\
        0x00b0 24 70 48 8d 84 24 b8 00 00 00 44 0f 11 38 48 8d  $pH..$....D..8H.
        0x00c0 84 24 c8 00 00 00 44 0f 11 38 48 8d 84 24 b8 00  .$....D..8H..$..
        0x00d0 00 00 48 89 84 24 98 00 00 00 48 8d 05 00 00 00  ..H..$....H.....
        0x00e0 00 48 8d 5c 24 78 e8 00 00 00 00 48 8b 8c 24 98  .H.\$x.....H..$.
        0x00f0 00 00 00 84 01 48 89 01 48 8d 79 08 83 3d 00 00  .....H..H.y..=..
        0x0100 00 00 00 74 02 eb 06 48 89 59 08 eb 07 e8 00 00  ...t...H.Y......
        0x0110 00 00 eb 00 48 8d 05 00 00 00 00 48 8d 5c 24 58  ....H......H.\$X
        0x0120 e8 00 00 00 00 48 8b 8c 24 98 00 00 00 84 01 48  .....H..$......H
        0x0130 89 41 10 48 8d 79 18 83 3d 00 00 00 00 00 66 90  .A.H.y..=.....f.
        0x0140 74 02 eb 06 48 89 59 18 eb 07 e8 00 00 00 00 eb  t...H.Y.........
        0x0150 00 48 8b 84 24 98 00 00 00 84 00 eb 00 48 89 84  .H..$........H..
        0x0160 24 a0 00 00 00 48 c7 84 24 a8 00 00 00 02 00 00  $....H..$.......
        0x0170 00 48 c7 84 24 b0 00 00 00 02 00 00 00 bb 02 00  .H..$...........
        0x0180 00 00 48 89 d9 e8 00 00 00 00 48 8b ac 24 d8 00  ..H.......H..$..
        0x0190 00 00 48 81 c4 e0 00 00 00 c3 e8 00 00 00 00 90  ..H.............
        0x01a0 e9 5b fe ff ff                                   .[...
        rel 3+0 t=24 type."".T+0
        rel 3+0 t=24 type."".T+0
        rel 39+4 t=7 "".test1+0
        rel 65+4 t=7 "".test2+0
        rel 221+4 t=15 type."".T+0
        rel 231+4 t=7 runtime.convT2Enoptr+0
        rel 254+4 t=15 runtime.writeBarrier+-1
        rel 270+4 t=7 runtime.gcWriteBarrierBX+0
        rel 279+4 t=15 type."".T+0
        rel 289+4 t=7 runtime.convT2Enoptr+0
        rel 313+4 t=15 runtime.writeBarrier+-1
        rel 331+4 t=7 runtime.gcWriteBarrierBX+0
        rel 390+4 t=7 fmt.Println+0
        rel 411+4 t=7 runtime.morestack_noctxt+0
"".test1 STEXT nosplit size=100 args=0x0 locals=0x28 funcid=0x0
        0x0000 00000 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:18)     TEXT    "".test1(SB), NOSPLIT|ABIInternal, $40-0
        0x0000 00000 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:18)     SUBQ    $40, SP
        0x0004 00004 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:18)     MOVQ    BP, 32(SP)
        0x0009 00009 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:18)     LEAQ    32(SP), BP
        0x000e 00014 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:18)     FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x000e 00014 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:18)     FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x000e 00014 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:18)     MOVUPS  X15, "".t(SP)
        0x0013 00019 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:18)     MOVUPS  X15, "".t+16(SP)
        0x0019 00025 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:19)     MOVUPS  X15, "".t(SP)
        0x001e 00030 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:19)     MOVUPS  X15, "".t+16(SP)
        0x0024 00036 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:19)     MOVQ    $1, "".t(SP)
        0x002c 00044 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:19)     MOVQ    $2, "".t+8(SP)
        0x0035 00053 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:19)     MOVQ    $3, "".t+16(SP)
        0x003e 00062 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:19)     MOVQ    $4, "".t+24(SP)
        0x0047 00071 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:20)     MOVQ    "".t(SP), AX
        0x004b 00075 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:20)     MOVL    $2, BX
        0x0050 00080 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:20)     MOVL    $3, CX
        0x0055 00085 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:20)     MOVL    $4, DI
        0x005a 00090 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:20)     MOVQ    32(SP), BP
        0x005f 00095 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:20)     ADDQ    $40, SP
        0x0063 00099 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:20)     RET
        0x0000 48 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 44 0f  H..(H.l$ H.l$ D.
        0x0010 11 3c 24 44 0f 11 7c 24 10 44 0f 11 3c 24 44 0f  .<$D..|$.D..<$D.
        0x0020 11 7c 24 10 48 c7 04 24 01 00 00 00 48 c7 44 24  .|$.H..$....H.D$
        0x0030 08 02 00 00 00 48 c7 44 24 10 03 00 00 00 48 c7  .....H.D$.....H.
        0x0040 44 24 18 04 00 00 00 48 8b 04 24 bb 02 00 00 00  D$.....H..$.....
        0x0050 b9 03 00 00 00 bf 04 00 00 00 48 8b 6c 24 20 48  ..........H.l$ H
        0x0060 83 c4 28 c3                                      ..(.
"".test2 STEXT nosplit size=137 args=0x0 locals=0x48 funcid=0x0
        0x0000 00000 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:23)     TEXT    "".test2(SB), NOSPLIT|ABIInternal, $72-0
        0x0000 00000 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:23)     SUBQ    $72, SP
        0x0004 00004 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:23)     MOVQ    BP, 64(SP)
        0x0009 00009 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:23)     LEAQ    64(SP), BP
        0x000e 00014 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:23)     FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x000e 00014 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:23)     FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x000e 00014 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:23)     MOVUPS  X15, "".~r0(SP)
        0x0013 00019 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:23)     MOVUPS  X15, "".~r0+16(SP)
        0x0019 00025 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:24)     MOVUPS  X15, "".v+32(SP)
        0x001f 00031 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:24)     MOVUPS  X15, "".v+48(SP)
        0x0025 00037 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:24)     MOVQ    $1, "".v+32(SP)
        0x002e 00046 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:24)     MOVQ    $2, "".v+40(SP)
        0x0037 00055 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:24)     MOVQ    $3, "".v+48(SP)
        0x0040 00064 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:24)     MOVQ    $4, "".v+56(SP)
        0x0049 00073 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     MOVQ    $1, "".~r0(SP)
        0x0051 00081 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     MOVQ    $2, "".~r0+8(SP)
        0x005a 00090 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     MOVQ    $3, "".~r0+16(SP)
        0x0063 00099 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     MOVQ    $4, "".~r0+24(SP)
        0x006c 00108 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     MOVQ    "".~r0(SP), AX
        0x0070 00112 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     MOVL    $2, BX
        0x0075 00117 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     MOVL    $3, CX
        0x007a 00122 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     MOVL    $4, DI
        0x007f 00127 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     MOVQ    64(SP), BP
        0x0084 00132 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     ADDQ    $72, SP
        0x0088 00136 (/Users/cwb/GolandProjects/GoBaseLearn/advance/cmd/main.go:25)     RET
        0x0000 48 83 ec 48 48 89 6c 24 40 48 8d 6c 24 40 44 0f  H..HH.l$@H.l$@D.
        0x0010 11 3c 24 44 0f 11 7c 24 10 44 0f 11 7c 24 20 44  .<$D..|$.D..|$ D
        0x0020 0f 11 7c 24 30 48 c7 44 24 20 01 00 00 00 48 c7  ..|$0H.D$ ....H.
        0x0030 44 24 28 02 00 00 00 48 c7 44 24 30 03 00 00 00  D$(....H.D$0....
        0x0040 48 c7 44 24 38 04 00 00 00 48 c7 04 24 01 00 00  H.D$8....H..$...
        0x0050 00 48 c7 44 24 08 02 00 00 00 48 c7 44 24 10 03  .H.D$.....H.D$..
        0x0060 00 00 00 48 c7 44 24 18 04 00 00 00 48 8b 04 24  ...H.D$.....H..$
        0x0070 bb 02 00 00 00 b9 03 00 00 00 bf 04 00 00 00 48  ...............H
        0x0080 8b 6c 24 40 48 83 c4 48 c3                       .l$@H..H.

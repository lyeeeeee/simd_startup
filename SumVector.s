
TEXT ·SumAssem(SB), $0
	MOVQ paddr+0(FP), BX
	MOVQ size+8(FP), BP
	MOVQ $0, AX

loop:
	CMPQ     BP, $0
	JLE      return

    MOVQ (BX),DX
    ADDQ DX, AX
    ADDQ $8, BX
    SUBQ $1, BP
    JMP  loop

return:
	MOVQ AX, ret+16(FP)
    RET



TEXT ·SumAssem4Way(SB), $0
	MOVQ paddr+0(FP), BX
	MOVQ size+8(FP), BP
	MOVQ $0, AX

loop:
	CMPQ     BP, $0
	JLE      return

    ADDQ (BX), AX

    ADDQ 8(BX), AX
    ADDQ 16(BX), AX
    ADDQ 24(BX),AX

    ADDQ $32, BX
    SUBQ $4, BP
    JMP  loop

return:
	MOVQ AX, ret+16(FP)
    RET


TEXT ·SumAssemSIMD(SB), $0
	MOVQ acc+0(FP), AX
	MOVQ paddr+8(FP), BX
	MOVQ size+16(FP), BP
	VXORPD Y1, Y1, Y1

loop:
	CMPQ     BP, $0
	JLE      return

	VMOVDQU (BX), Y2
	VPADDQ  Y2, Y1, Y1

    ADDQ $32, BX
    SUBQ $4, BP
    JMP  loop

return:
	VMOVDQU Y1, (AX)
    RET

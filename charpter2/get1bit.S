TEXT main.main(SB) /opt/gopl-practice/charpter2/get1bit.go
  get1bit.go:10		0x49abc0		64488b0c25f8ffffff		MOVQ FS:0xfffffff8, CX			
  get1bit.go:10		0x49abc9		488d842458ffffff		LEAQ 0xffffff58(SP), AX			
  get1bit.go:10		0x49abd1		483b4110			CMPQ 0x10(CX), AX			
  get1bit.go:10		0x49abd5		0f8615030000			JBE 0x49aef0				
  get1bit.go:10		0x49abdb		4881ec28010000			SUBQ $0x128, SP				
  get1bit.go:10		0x49abe2		4889ac2420010000		MOVQ BP, 0x120(SP)			
  get1bit.go:10		0x49abea		488dac2420010000		LEAQ 0x120(SP), BP			
  get1bit.go:11		0x49abf2		488b0d4f530b00			MOVQ os.Args+8(SB), CX			
  get1bit.go:11		0x49abf9		488b1540530b00			MOVQ os.Args(SB), DX			
  get1bit.go:11		0x49ac00		488b1d49530b00			MOVQ os.Args+16(SB), BX			
  get1bit.go:11		0x49ac07		eb00				JMP 0x49ac09				
  get1bit.go:11		0x49ac09		4883f901			CMPQ $0x1, CX				
  get1bit.go:11		0x49ac0d		7305				JAE 0x49ac14				
  get1bit.go:11		0x49ac0f		e9d1020000			JMP 0x49aee5				
  get1bit.go:11		0x49ac14		48ffcb				DECQ BX					
  get1bit.go:11		0x49ac17		4889de				MOVQ BX, SI				
  get1bit.go:11		0x49ac1a		48f7db				NEGQ BX					
  get1bit.go:11		0x49ac1d		48c1fb3f			SARQ $0x3f, BX				
  get1bit.go:11		0x49ac21		4883e310			ANDQ $0x10, BX				
  get1bit.go:11		0x49ac25		4801da				ADDQ BX, DX				
  get1bit.go:11		0x49ac28		48899424f0000000		MOVQ DX, 0xf0(SP)			
  get1bit.go:11		0x49ac30		488d51ff			LEAQ -0x1(CX), DX			
  get1bit.go:11		0x49ac34		48899424f8000000		MOVQ DX, 0xf8(SP)			
  get1bit.go:11		0x49ac3c		4889b42400010000		MOVQ SI, 0x100(SP)			
  get1bit.go:11		0x49ac44		0f57c0				XORPS X0, X0				
  get1bit.go:11		0x49ac47		0f118424b8000000		MOVUPS X0, 0xb8(SP)			
  get1bit.go:11		0x49ac4f		488b8c24f8000000		MOVQ 0xf8(SP), CX			
  get1bit.go:11		0x49ac57		488b9424f0000000		MOVQ 0xf0(SP), DX			
  get1bit.go:11		0x49ac5f		90				NOPL					
  get1bit.go:11		0x49ac60		4885c9				TESTQ CX, CX				
  get1bit.go:11		0x49ac63		7705				JA 0x49ac6a				
  get1bit.go:11		0x49ac65		e971020000			JMP 0x49aedb				
  get1bit.go:11		0x49ac6a		488b02				MOVQ 0(DX), AX				
  get1bit.go:11		0x49ac6d		488b4a08			MOVQ 0x8(DX), CX			
  get1bit.go:11		0x49ac71		4889842498000000		MOVQ AX, 0x98(SP)			
  get1bit.go:11		0x49ac79		48898c24a0000000		MOVQ CX, 0xa0(SP)			
  get1bit.go:11		0x49ac81		48890424			MOVQ AX, 0(SP)				
  get1bit.go:11		0x49ac85		48894c2408			MOVQ CX, 0x8(SP)			
  get1bit.go:11		0x49ac8a		e81111fdff			CALL strconv.Atoi(SB)			
  get1bit.go:11		0x49ac8f		488b442410			MOVQ 0x10(SP), AX			
  get1bit.go:11		0x49ac94		488b4c2418			MOVQ 0x18(SP), CX			
  get1bit.go:11		0x49ac99		488b542420			MOVQ 0x20(SP), DX			
  get1bit.go:11		0x49ac9e		4889442460			MOVQ AX, 0x60(SP)			
  get1bit.go:11		0x49aca3		48898c24b8000000		MOVQ CX, 0xb8(SP)			
  get1bit.go:11		0x49acab		48899424c0000000		MOVQ DX, 0xc0(SP)			
  get1bit.go:11		0x49acb3		488b442460			MOVQ 0x60(SP), AX			
  get1bit.go:11		0x49acb8		4889442458			MOVQ AX, 0x58(SP)			
  get1bit.go:11		0x49acbd		488b8424b8000000		MOVQ 0xb8(SP), AX			
  get1bit.go:11		0x49acc5		488b8c24c0000000		MOVQ 0xc0(SP), CX			
  get1bit.go:11		0x49accd		4889842488000000		MOVQ AX, 0x88(SP)			
  get1bit.go:11		0x49acd5		48898c2490000000		MOVQ CX, 0x90(SP)			
  get1bit.go:12		0x49acdd		4883bc248800000000		CMPQ $0x0, 0x88(SP)			
  get1bit.go:12		0x49ace6		7505				JNE 0x49aced				
  get1bit.go:12		0x49ace8		e9e9010000			JMP 0x49aed6				
  get1bit.go:13		0x49aced		0f57c0				XORPS X0, X0				
  get1bit.go:13		0x49acf0		0f118424a8000000		MOVUPS X0, 0xa8(SP)			
  get1bit.go:13		0x49acf8		488d8424a8000000		LEAQ 0xa8(SP), AX			
  get1bit.go:13		0x49ad00		4889442468			MOVQ AX, 0x68(SP)			
  get1bit.go:13		0x49ad05		488b842488000000		MOVQ 0x88(SP), AX			
  get1bit.go:13		0x49ad0d		488b8c2490000000		MOVQ 0x90(SP), CX			
  get1bit.go:13		0x49ad15		48898424c8000000		MOVQ AX, 0xc8(SP)			
  get1bit.go:13		0x49ad1d		48898c24d0000000		MOVQ CX, 0xd0(SP)			
  get1bit.go:13		0x49ad25		4889842480000000		MOVQ AX, 0x80(SP)			
  get1bit.go:13		0x49ad2d		4883bc248000000000		CMPQ $0x0, 0x80(SP)			
  get1bit.go:13		0x49ad36		7505				JNE 0x49ad3d				
  get1bit.go:13		0x49ad38		e994010000			JMP 0x49aed1				
  get1bit.go:13		0x49ad3d		488b4008			MOVQ 0x8(AX), AX			
  get1bit.go:13		0x49ad41		4889842480000000		MOVQ AX, 0x80(SP)			
  get1bit.go:13		0x49ad49		eb00				JMP 0x49ad4b				
  get1bit.go:13		0x49ad4b		488b442468			MOVQ 0x68(SP), AX			
  get1bit.go:13		0x49ad50		8400				TESTB AL, 0(AX)				
  get1bit.go:13		0x49ad52		488b8c24d0000000		MOVQ 0xd0(SP), CX			
  get1bit.go:13		0x49ad5a		488b942480000000		MOVQ 0x80(SP), DX			
  get1bit.go:13		0x49ad62		488910				MOVQ DX, 0(AX)				
  get1bit.go:13		0x49ad65		488d7808			LEAQ 0x8(AX), DI			
  get1bit.go:13		0x49ad69		833d30270e0000			CMPL $0x0, runtime.writeBarrier(SB)	
  get1bit.go:13		0x49ad70		7405				JE 0x49ad77				
  get1bit.go:13		0x49ad72		e950010000			JMP 0x49aec7				
  get1bit.go:13		0x49ad77		48894808			MOVQ CX, 0x8(AX)			
  get1bit.go:13		0x49ad7b		eb00				JMP 0x49ad7d				
  get1bit.go:13		0x49ad7d		488b442468			MOVQ 0x68(SP), AX			
  get1bit.go:13		0x49ad82		8400				TESTB AL, 0(AX)				
  get1bit.go:13		0x49ad84		eb00				JMP 0x49ad86				
  get1bit.go:13		0x49ad86		48898424d8000000		MOVQ AX, 0xd8(SP)			
  get1bit.go:13		0x49ad8e		48c78424e000000001000000	MOVQ $0x1, 0xe0(SP)			
  get1bit.go:13		0x49ad9a		48c78424e800000001000000	MOVQ $0x1, 0xe8(SP)			
  get1bit.go:13		0x49ada6		488b0d734d0b00			MOVQ os.Stderr(SB), CX			
  get1bit.go:13		0x49adad		488d1534290400			LEAQ go.itab.*os.File,io.Writer(SB), DX	
  get1bit.go:13		0x49adb4		48891424			MOVQ DX, 0(SP)				
  get1bit.go:13		0x49adb8		48894c2408			MOVQ CX, 0x8(SP)			
  get1bit.go:13		0x49adbd		488d0dac3f0200			LEAQ 0x23fac(IP), CX			
  get1bit.go:13		0x49adc4		48894c2410			MOVQ CX, 0x10(SP)			
  get1bit.go:13		0x49adc9		48c74424180c000000		MOVQ $0xc, 0x18(SP)			
  get1bit.go:13		0x49add2		4889442420			MOVQ AX, 0x20(SP)			
  get1bit.go:13		0x49add7		48c744242801000000		MOVQ $0x1, 0x28(SP)			
  get1bit.go:13		0x49ade0		48c744243001000000		MOVQ $0x1, 0x30(SP)			
  get1bit.go:13		0x49ade9		e8527dffff			CALL fmt.Fprintf(SB)			
  get1bit.go:14		0x49adee		48c7042401000000		MOVQ $0x1, 0(SP)			
  get1bit.go:14		0x49adf6		e8e547ffff			CALL os.Exit(SB)			
  get1bit.go:14		0x49adfb		eb00				JMP 0x49adfd				
  get1bit.go:16		0x49adfd		488b442458			MOVQ 0x58(SP), AX			
  get1bit.go:16		0x49ae02		48890424			MOVQ AX, 0(SP)				
  get1bit.go:16		0x49ae06		e815cffcff			CALL charpter2/popcount.PopCount(SB)	
  get1bit.go:16		0x49ae0b		488b442408			MOVQ 0x8(SP), AX			
  get1bit.go:16		0x49ae10		4889442450			MOVQ AX, 0x50(SP)			
  get1bit.go:17		0x49ae15		0f57c0				XORPS X0, X0				
  get1bit.go:17		0x49ae18		0f118424a8000000		MOVUPS X0, 0xa8(SP)			
  get1bit.go:17		0x49ae20		488d8424a8000000		LEAQ 0xa8(SP), AX			
  get1bit.go:17		0x49ae28		4889442478			MOVQ AX, 0x78(SP)			
  get1bit.go:17		0x49ae2d		488b442450			MOVQ 0x50(SP), AX			
  get1bit.go:17		0x49ae32		48890424			MOVQ AX, 0(SP)				
  get1bit.go:17		0x49ae36		e8a5f5f6ff			CALL runtime.convT64(SB)		
  get1bit.go:17		0x49ae3b		488b442408			MOVQ 0x8(SP), AX			
  get1bit.go:17		0x49ae40		4889442470			MOVQ AX, 0x70(SP)			
  get1bit.go:17		0x49ae45		488b4c2478			MOVQ 0x78(SP), CX			
  get1bit.go:17		0x49ae4a		8401				TESTB AL, 0(CX)				
  get1bit.go:17		0x49ae4c		488d15cda20000			LEAQ 0xa2cd(IP), DX			
  get1bit.go:17		0x49ae53		488911				MOVQ DX, 0(CX)				
  get1bit.go:17		0x49ae56		488d7908			LEAQ 0x8(CX), DI			
  get1bit.go:17		0x49ae5a		833d3f260e0000			CMPL $0x0, runtime.writeBarrier(SB)	
  get1bit.go:17		0x49ae61		7402				JE 0x49ae65				
  get1bit.go:17		0x49ae63		eb5a				JMP 0x49aebf				
  get1bit.go:17		0x49ae65		48894108			MOVQ AX, 0x8(CX)			
  get1bit.go:17		0x49ae69		eb00				JMP 0x49ae6b				
  get1bit.go:17		0x49ae6b		488b442478			MOVQ 0x78(SP), AX			
  get1bit.go:17		0x49ae70		8400				TESTB AL, 0(AX)				
  get1bit.go:17		0x49ae72		eb00				JMP 0x49ae74				
  get1bit.go:17		0x49ae74		4889842408010000		MOVQ AX, 0x108(SP)			
  get1bit.go:17		0x49ae7c		48c784241001000001000000	MOVQ $0x1, 0x110(SP)			
  get1bit.go:17		0x49ae88		48c784241801000001000000	MOVQ $0x1, 0x118(SP)			
  get1bit.go:17		0x49ae94		48890424			MOVQ AX, 0(SP)				
  get1bit.go:17		0x49ae98		48c744240801000000		MOVQ $0x1, 0x8(SP)			
  get1bit.go:17		0x49aea1		48c744241001000000		MOVQ $0x1, 0x10(SP)			
  get1bit.go:17		0x49aeaa		e8b17effff			CALL fmt.Println(SB)			
  get1bit.go:18		0x49aeaf		488bac2420010000		MOVQ 0x120(SP), BP			
  get1bit.go:18		0x49aeb7		4881c428010000			ADDQ $0x128, SP				
  get1bit.go:18		0x49aebe		c3				RET					
  get1bit.go:17		0x49aebf		90				NOPL					
  get1bit.go:17		0x49aec0		e8db92fcff			CALL runtime.gcWriteBarrier(SB)		
  get1bit.go:17		0x49aec5		eba4				JMP 0x49ae6b				
  get1bit.go:13		0x49aec7		e8d493fcff			CALL runtime.gcWriteBarrierCX(SB)	
  get1bit.go:13		0x49aecc		e9acfeffff			JMP 0x49ad7d				
  get1bit.go:13		0x49aed1		e975feffff			JMP 0x49ad4b				
  get1bit.go:12		0x49aed6		e922ffffff			JMP 0x49adfd				
  get1bit.go:11		0x49aedb		31c0				XORL AX, AX				
  get1bit.go:11		0x49aedd		0f1f00				NOPL 0(AX)				
  get1bit.go:11		0x49aee0		e85b97fcff			CALL runtime.panicIndex(SB)		
  get1bit.go:11		0x49aee5		b801000000			MOVL $0x1, AX				
  get1bit.go:11		0x49aeea		e81198fcff			CALL runtime.panicSliceB(SB)		
  get1bit.go:11		0x49aeef		90				NOPL					
  get1bit.go:10		0x49aef0		e80b77fcff			CALL runtime.morestack_noctxt(SB)	
  get1bit.go:10		0x49aef5		e9c6fcffff			JMP main.main(SB)			

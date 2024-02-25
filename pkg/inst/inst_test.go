package inst

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

type opcodes map[string]struct {
	Mnemonic string `json:"mnemonic"`
	Operands []struct {
		Name      string `json:"name"`
		Immediate bool   `json:"immediate,omitempty"`
	} `json:"operands"`
	Flags struct {
		Z string `json:"Z"`
		N string `json:"N"`
		H string `json:"H"`
		C string `json:"C"`
	} `json:"flags"`
}

type opcodeTable struct {
	Unprefixed opcodes `json:"unprefixed"`
	CBPrefixed opcodes `json:"cbprefixed"`
}

func TestInitHandlers(t *testing.T) {
	var o opcodeTable

	bs, err := os.ReadFile("../../testdata/opcodes.json")

	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(bs, &o); err != nil {
		t.Fatal(err)
	}

	InitHandlers()

	for code, op := range o.Unprefixed {
		if strings.HasPrefix(op.Mnemonic, "ILLEGAL") {
			continue
		}

		codeInt, err := strconv.ParseInt(code, 0, 32)

		if err != nil {
			t.Fatal(err)
		}

		if x := h.handler(byte(codeInt)); x == nil {
			var operands []string

			for _, operand := range op.Operands {
				if operand.Immediate {
					operands = append(operands, operand.Name)
				} else {
					operands = append(operands, fmt.Sprintf("[%s]", operand.Name))
				}
			}

			t.Errorf("missing handler for unprefixed opcode 0x%02X: %s %s", codeInt, op.Mnemonic, strings.Join(operands, ", "))
		}
	}

	for code, op := range o.CBPrefixed {
		if strings.HasPrefix(op.Mnemonic, "ILLEGAL") {
			continue
		}

		codeInt, err := strconv.ParseInt(code, 0, 32)

		if err != nil {
			t.Fatal(err)
		}

		if x := p.handler(byte(codeInt)); x == nil {
			var operands []string

			for _, operand := range op.Operands {
				if operand.Immediate {
					operands = append(operands, operand.Name)
				} else {
					operands = append(operands, fmt.Sprintf("[%s]", operand.Name))
				}
			}

			t.Errorf("missing handler for prefixed opcode 0xCB 0x%02X: %s %s", codeInt, op.Mnemonic, strings.Join(operands, ", "))
		}
	}
}

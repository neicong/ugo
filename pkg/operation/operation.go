package operation

import (
	"bytes"
	"fmt"
	"io"
)

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func (n *Node) Run() string {
	o := &Operation{}
	return o.GenLLIR(n)

}

type Operation struct {
	ID int
}

func (o *Operation) GenLLIR(node *Node) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "define i32 @main() {\n")
	fmt.Fprintf(&buf, "    ret i32 %s\n", o.genValue(&buf, node))
	fmt.Fprintf(&buf, "}\n")

	return buf.String()
}

func (o *Operation) genId() string {
	id := fmt.Sprintf("%%t%d", o.ID)
	o.ID++
	return id
}

func (o *Operation) genValue(w io.Writer, node *Node) string {
	if node == nil {
		return ""
	}
	id := o.genId()
	switch node.Val {
	case "+":
		fmt.Fprintf(w, "\t%s = add i32 %s, %s\n",
			id, o.genValue(w, node.Left), o.genValue(w, node.Right),
		)
	case "-":
		fmt.Fprintf(w, "\t%s = sub  i32 %s, %s\n",
			id, o.genValue(w, node.Left), o.genValue(w, node.Right),
		)
	case "*":
		fmt.Fprintf(w, "\t%s = mul   i32 %s, %s\n",
			id, o.genValue(w, node.Left), o.genValue(w, node.Right),
		)
	case "/":
		fmt.Fprintf(w, "\t%s = sdiv  i32 %s, %s\n",
			id, o.genValue(w, node.Left), o.genValue(w, node.Right),
		)
	default:
		fmt.Fprintf(w, "\t%[1]s = add i32 0, %[2]s; %[1]s = %[2]s\n",
			id, node.Val,
		)
	}
	return id
}

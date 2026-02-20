package bases

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Test() {
	fmt.Println("Prube exitosa")
}
func GetOps(t string, lt []string, tipo int) string {

	if tipo == 0 {
		fmt.Println(t)
	}
	if tipo == 1 {
		fmt.Print(t)
	}

	for _, x := range lt {
		fmt.Println(x)
	}
	reader := bufio.NewReader(os.Stdin)
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

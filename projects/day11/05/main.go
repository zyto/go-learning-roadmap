package main

//исправил импорт:
import (
	"fmt"
	"strings"
)

//было так:
/*import (
	"fmt"
	"strconv"
	"strings"
)*/

func main() {
	fmt.Println(strings.TrimSpace("   www.ru"))
}

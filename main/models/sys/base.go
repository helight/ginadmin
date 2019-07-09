package sys

import (
	"fmt"

	"github.com/it234/goapp/main/models/basemodel"
)

func TableName(name string) string {
	return fmt.Sprintf("%s%s%s", basemodel.GetTablePrefix(), "sys_", name)
}

package cloud

import (
	"fmt"
)

func (a AWS) Link(fileName string) string {
	return fmt.Sprintf("https://storage.yandexcloud.net/lms-user/%s", fileName)
}

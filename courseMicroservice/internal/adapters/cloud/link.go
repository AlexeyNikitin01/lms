package cloud

import (
	"fmt"
)

func (a AWS) Link(fileName string) string {
	return fmt.Sprintf("https://storage.yandexcloud.net/%s/%s", a.bucket, fileName)
}

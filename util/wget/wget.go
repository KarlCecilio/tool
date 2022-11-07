package wget

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func WgetFile(downloadFile string, saveDir string) (string, error) {
	var b bytes.Buffer
	if _, err := os.Stat(saveDir); os.IsNotExist(err) { //路径不存在，创建
		err = os.MkdirAll(saveDir, 0700)
		if err != nil {
			return "", err
		}
	}
	cmd := exec.Command("/bin/bash", "-c", "wget "+downloadFile+" -P "+saveDir+" -o "+saveDir+strconv.FormatInt(time.Now().Unix(), 10)) //不加第一个第二个参数会报错
	cmd.Stdout = &b
	cmd.Stderr = &b
	err := cmd.Run()
	if err != nil {
		return "", errors.New(fmt.Sprintf(err.Error() + ":" + b.String()))
	} else {
		return string(b.Bytes()), nil
	}
}

package karl_file

import (
	"bufio"
	"os"
)

func WriteBufio(path, str string) error {
	fileObj, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 132)
	if err != nil {
		return err
	}
	defer fileObj.Close()
	write := bufio.NewWriter(fileObj)
	// 此时写在了缓存中
	write.Write([]byte(str))
	// 需要从缓存中写入到文件
	err1 := write.Flush()
	if err1 != nil {
		return err1
	}
	return nil
}

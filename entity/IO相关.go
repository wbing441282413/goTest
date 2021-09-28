package entity

import "io/ioutil"

func ReadFileByFilePath(fp string) ([]byte, error) { //小文件一次性读出
	filePath := fp
	data, err := ioutil.ReadFile(filePath)
	if err == nil {
		return data, nil
	}
	return nil, err
}

//TODO 待续

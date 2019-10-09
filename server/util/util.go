package util

import(
	"server/config"
	"bytes"
	"os/exec"
	"fmt"
)

const BadInputError = "Bad input string"
const BadResponseError = "Bad response string"

type UtilService interface{
	Exec(input string) (string, error)
}

func NewUtilService() UtilService{
	return &Util{
		execPath: config.GetString(config.UtilPath),
		errorString: "error",
		utilExecThrottling: make(chan struct{},
			config.GetInt32(config.UtilWorkersMaxCount)),
	}
}

type Util struct {
	execPath string
	errorString string
	utilExecThrottling chan struct{}
}

func (u *Util) Exec(input string) (string, error){
	if input == u.errorString {
		return "", fmt.Errorf(BadInputError)
	}

	u.utilExecThrottling <- struct{}{}
	result, err := u.runUtilWithInput(input)
	<- u.utilExecThrottling

	if err != nil {
		return "", err
	}

	if result == u.errorString {
		return "", fmt.Errorf(BadResponseError)
	}

	return result, nil
}

func (u *Util) runUtilWithInput(input string) (string, error){
	var param = "-reverse=" + input
	cmd := exec.Command(u.execPath, param)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

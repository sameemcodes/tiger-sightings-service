package utils

import "tigerhall-kittens/cmd/constants"

func HandleError(msg string, err error) {
	if err != nil {
		constants.Logger.Errorf(msg+" \n", err)
	}
}

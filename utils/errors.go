package utils

import (
	"errors"
)

var (
	UnableToParseTimeErr = errors.New("unable to parse time")
	PastTimeErr          = errors.New("set a future time")
	ChDirErr             = errors.New("couldn't change directory to home directory")
	HomeDirErr           = errors.New("couldn't get users home directory")
	CreatingDirErr       = errors.New("error creating directory")
	OpenFileErr          = errors.New("error opening or creating file")
	WritingFileErr       = errors.New("error writing in file")
	DeletBackSpaces      = errors.New("delete unnecessary backspace in noty-data/alarms-data in user directory")
)

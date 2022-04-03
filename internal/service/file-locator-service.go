package service

import (
	"errors"
	"github.com/leddzip/back-find-cli/internal/domain"
	"path"
	"path/filepath"
	"strings"
)

// FindFileBetween try to find a suitable file that should exist inside the
func FindFileBetween(fileName string, fromFolder string, backToFolder string) (domain.FilePresence, error) {

	cleanAbsFrom, cleanAbsBackTo, err := validateFolderRange(fromFolder, backToFolder)
	if err != nil {
		return domain.FileAbsent(), err
	}

	folder := domain.Folder{Path: cleanAbsFrom}
	backToParentFolder := domain.Folder{Path: cleanAbsBackTo}
	backToParentFolder.NavigateToParent()
	for f := &folder; f.Path != backToParentFolder.Path; f.NavigateToParent() {
		isFileInFolder, err := f.ContainsFile(fileName)
		if err != nil {
			return domain.FileAbsent(), err
		}
		if isFileInFolder {
			return domain.FilePresent(path.Join(f.Path, fileName)), nil
		}
	}

	// edge case. When the backTo folder is the root, the previous loop does not check if
	// the filename is present in the folder. So we have to manage this edge case independently
	if cleanAbsBackTo == "/" {
		root := domain.Folder{Path: "/"}
		isFileInFolder, err := root.ContainsFile(fileName)
		if err != nil {
			return domain.FileAbsent(), err
		}
		if isFileInFolder {
			return domain.FilePresent(path.Join(root.Path, fileName)), nil
		}
	}

	return domain.FileAbsent(), nil
}

// validateFolderRange transform the two folder (from=origin and backTo=destination) into their
// cleaned and absolute variant. It also checks that thw 'from' folder is a child (regardless of the
// level) of the 'backTo' folder.
func validateFolderRange(from string, backTo string) (string, string, error) {
	cleanAbsFrom, err := filepath.Abs(from)
	if err != nil {
		return from, backTo, errors.New("unable to convert the 'from' into its absolute form")
	}

	cleanAbsBackTo, err := filepath.Abs(backTo)
	if err != nil {
		return from, backTo, errors.New("unable to convert the 'backTo' into its absolute from")
	}

	if strings.Index(cleanAbsFrom, cleanAbsBackTo) != 0 {
		return from, backTo, errors.New("the 'from' folder is not a subdirectory of the 'backTo' folder")
	}

	return cleanAbsFrom, cleanAbsBackTo, nil
}

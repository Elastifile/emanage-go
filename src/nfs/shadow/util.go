package shadow

import (
	"sort"

	"github.com/go-errors/errors"

	"gopkg.in/fatih/set.v0"

	"nfs"
)

var VerifyDirectories = true // FIXME: Global state, may be problematic

func shadowReadDir(
	opRefr func() (<-chan *nfs.DirEntry, error),
	opElfs func() (<-chan *nfs.DirEntry, error),
) (<-chan *nfs.DirEntry, error) {

	makeSet := func(entriesIn <-chan *nfs.DirEntry, entriesOut chan<- *nfs.DirEntry, err error) (set.Interface, error) {
		if err != nil {
			return nil, err
		}

		s := set.NewNonTS()
		numEntries := 0
		for entry := range entriesIn {
			if err := entry.Err; err != nil {
				return nil, err
			}

			numEntries++
			s.Add(entry.Name)

			if entriesOut != nil {
				entriesOut <- entry
			}
		}

		if entriesOut != nil {
			close(entriesOut)
		}

		if s.Size() != numEntries {
			return s, errors.Errorf("Multiple entries with identical names were found")
		}

		return s, nil
	}

	entriesRefr, errRefr := opRefr()
	entriesElfs, errElfs := opElfs()
	entries := make(chan *nfs.DirEntry)

	var shadowErr error

	go func() {
		setRefr, refrErr := makeSet(entriesRefr, nil, errRefr)
		setElfs, elfsErr := makeSet(entriesElfs, entries, errElfs)

		err := checkShadow("ReadDir", refrErr, elfsErr)
		if err != nil {
			shadowErr = err
			return
		}

		if VerifyDirectories {
			if !setElfs.IsEqual(setRefr) {
				err := &ShadowDirectoryMismatchError{
					OnlyInElfs: makeSorted(set.Difference(setElfs, setRefr)),
					OnlyInRefr: makeSorted(set.Difference(setRefr, setElfs)),
				}
				if err != nil {
					shadowErr = err
					return
				}
			}
		}
	}()

	return entries, shadowErr
}

func makeSorted(s set.Interface) []string {
	slice := set.StringSlice(s)
	sort.Strings(slice)
	return slice
}

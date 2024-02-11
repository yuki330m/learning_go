package main

import "fmt"

func notUsedDeferError(val1 int, val2 string) (string, error) {
	val3, err := doThing1(val1)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	result, err := doThing3(val3, val4)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	return result, nil
}

func useDeferError(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
	}()
	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}
	result, err := doThing3(val3, val4)
	if err != nil {
		return "", err
	}
	return result, nil
}

func doThing1(v any) (string, error) {
	return "", nil
}
func doThing2(v any) (string, error) {
	return "", nil
}
func doThing3(v any, v2 any) (string, error) {
	return "", nil
}

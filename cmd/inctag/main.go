package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const op = "increment tag"

	args := os.Args

	if len(args) < 3 {
		panic(op + ": not enough args")
	}

	tag, err := updateTag(args[1], args[2])
	if err != nil {
		panic(err)
	}

	fmt.Println(tag)
}

func updateTag(tag, part string) (string, error) {
	const op = "update tag"

	parts := strings.Split(tag, "/")
	if len(parts) == 0 {
		return "", fmt.Errorf("%s: empty tag", op)
	}

	last := parts[len(parts)-1]
	if len(last) == 0 {
		return "", fmt.Errorf("%s: empty tag", op)
	}

	withV := false
	if strings.HasPrefix(last, "v") {
		last = last[1:]
		withV = true
	}

	semVerParts := strings.Split(last, ".")
	if len(semVerParts) != 3 {
		return "", fmt.Errorf("%s: tag version must be of semver format 1.2.3", op)
	}

	var versions []int
	for _, v := range semVerParts {
		i, err := strconv.Atoi(v)
		if err != nil {
			return "", fmt.Errorf("%s: tag version must be of semver format 1.2.3", op)
		}
		versions = append(versions, i)
	}

	switch part {
	case "minor":
		versions[1] = versions[1] + 1
	case "patch":
		versions[2] = versions[2] + 1
	}

	newSemVer := fmt.Sprintf("%d.%d.%d", versions[0], versions[1], versions[2])
	if withV {
		newSemVer = "v" + newSemVer
	}

	return strings.Join(append(parts[:len(parts)-1], newSemVer), "/"), nil
}

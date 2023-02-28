package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Data struct {
	ID            string
	IDLike        string
	Name          string
	PrettyName    string
	Version       string
	VersionID     string
	KernelVersion string
	Arch          string
}

const (
	EtcOsRelease string = "/etc/os-release"
	DebianID            = "debian"
	FedoraID            = "fedora"
	UbuntuID            = "ubuntu"
	RhelID              = "rhel"
	CentosID            = "centos"
)

func OSInfo() (*Data, error) {
	content, err := ExecCMDWithResult("cat", []string{EtcOsRelease}, 5)
	if err != nil {
		return nil, fmt.Errorf("Get os info failed ")
	}
	if len(content) == 0 {
		return nil, fmt.Errorf("Failed to get os info data ")
	}
	data := Parse(string(content))
	log.Infof("OS info: %v", data)
	return data, nil
}

// Parse is to parse a os release file content.
func Parse(content string) (data *Data) {
	data = new(Data)
	lines, err := parseString(content)
	if err != nil {
		return
	}

	info := make(map[string]string)
	for _, v := range lines {
		key, value, err := parseLine(v)
		if err == nil {
			info[key] = value
		}
	}
	data.ID = info["ID"]
	data.IDLike = info["ID_LIKE"]
	data.Name = info["NAME"]
	data.PrettyName = info["PRETTY_NAME"]
	data.Version = info["VERSION"]
	data.VersionID = info["VERSION_ID"]
	return
}

func parseString(content string) (lines []string, err error) {
	in := bytes.NewBufferString(content)
	reader := bufio.NewReader(in)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseLine(line string) (string, string, error) {
	// skip empty lines
	if line == "" {
		return "", "", errors.New("Skipping: zero-length ")
	}

	// skip comments
	if line[0] == '#' {
		return "", "", errors.New("Skipping: comment ")
	}

	// try to split string at the first '='
	splitString := strings.SplitN(line, "=", 2)
	if len(splitString) != 2 {
		return "", "", errors.New("Can not extract key=value ")
	}

	// trim white space from key and value
	key := splitString[0]
	key = strings.Trim(key, " ")
	value := splitString[1]
	value = strings.Trim(value, " ")

	// Handle double quotes
	if strings.ContainsAny(value, `"`) {
		first := value[0:1]
		last := value[len(value)-1:]

		if first == last && strings.ContainsAny(first, `"'`) {
			value = strings.TrimPrefix(value, `'`)
			value = strings.TrimPrefix(value, `"`)
			value = strings.TrimSuffix(value, `'`)
			value = strings.TrimSuffix(value, `"`)
		}
	}

	// expand anything else that could be escaped
	value = strings.Replace(value, `\"`, `"`, -1)
	value = strings.Replace(value, `\$`, `$`, -1)
	value = strings.Replace(value, `\\`, `\`, -1)
	value = strings.Replace(value, "\\`", "`", -1)
	value = strings.TrimRight(value, "\r\n")
	value = strings.TrimLeft(value, "\"")
	value = strings.TrimRight(value, "\"")
	return key, value, nil
}

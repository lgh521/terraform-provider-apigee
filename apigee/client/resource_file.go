package client

import (
	"strconv"
	"strings"
)

const (
	OrganizationResourceFilePath       = "organizations/%s/resourcefiles"
	OrganizationResourceFilePathOfType = OrganizationResourceFilePath + "/%s"
	OrganizationResourceFilePathGet    = OrganizationResourceFilePathOfType + "/%s"
	EnvironmentResourceFilePath        = "organizations/%s/environments/%s/resourcefiles"
	EnvironmentResourceFilePathOfType  = EnvironmentResourceFilePath + "/%s"
	EnvironmentResourceFilePathGet     = EnvironmentResourceFilePathOfType + "/%s"
	ProxyResourceFilePath              = "organizations/%s/apis/%s/revisions/%d/resourcefiles"
	ProxyResourceFilePathOfType        = ProxyResourceFilePath + "/%s"
	ProxyResourceFilePathGet           = ProxyResourceFilePathOfType + "/%s"
)

type ResourceFile struct {
	Type string `json:"type"`
	Name string `json:"name"`
	//Only used for Environment context
	EnvironmentName string `json:"-"`
	//Only used for Proxy context
	ProxyName string `json:"-"`
	//Only used for Proxy context
	Revision int `json:"-"`
}
type ResourceFilesOfType struct {
	Files []ResourceFile `json:"resourceFile"`
}

func (c *ResourceFile) OrganizationResourceFileEncodeId() string {
	return c.Type + IdSeparator + c.Name
}

func (c *ResourceFile) EnvironmentResourceFileEncodeId() string {
	return c.EnvironmentName + IdSeparator + c.Type + IdSeparator + c.Name
}

func (c *ResourceFile) ProxyResourceFileEncodeId() string {
	return c.ProxyName + IdSeparator + strconv.Itoa(c.Revision) + IdSeparator + c.Type + IdSeparator + c.Name
}

func OrganizationResourceFileDecodeId(s string) (string, string) {
	tokens := strings.Split(s, IdSeparator)
	return tokens[0], tokens[1]
}

func EnvironmentResourceFileDecodeId(s string) (string, string, string) {
	tokens := strings.Split(s, IdSeparator)
	return tokens[0], tokens[1], tokens[2]
}

func ProxyResourceFileDecodeId(s string) (string, int, string, string) {
	tokens := strings.Split(s, IdSeparator)
	revision, _ := strconv.Atoi(tokens[1])
	return tokens[0], revision, tokens[2], tokens[3]
}

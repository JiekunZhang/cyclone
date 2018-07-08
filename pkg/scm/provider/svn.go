/*
Copyright 2017 caicloud authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"fmt"
	"strings"

	"github.com/caicloud/cyclone/pkg/api"
	"github.com/caicloud/cyclone/pkg/executil"
	"github.com/caicloud/cyclone/pkg/scm"
	log "github.com/golang/glog"
)

// SVN represents the SCM provider of SVN.
type SVN struct{}

func init() {
	if err := scm.RegisterProvider(api.SVN, new(SVN)); err != nil {
		log.Errorln(err)
	}
}

func (s *SVN) spilitToken(token string) (string, string, error) {
	userPwd := strings.Split(token, api.SVNUsernPwdSep)
	if len(userPwd) != 2 {
		err := fmt.Errorf("split token fail as the length of userPwd equals %v", len(userPwd))
		return "", "", err
	}

	return userPwd[0], userPwd[1], nil
}

func (s *SVN) GetToken(scm *api.SCMConfig) (string, error) {
	return scm.Username + api.SVNUsernPwdSep + scm.Password, nil
}

func (s *SVN) ListRepos(scm *api.SCMConfig) ([]api.Repository, error) {
	return nil, nil
}

func (s *SVN) ListBranches(scm *api.SCMConfig, repo string) ([]string, error) {
	return nil, nil
}

func (s *SVN) ListTags(scm *api.SCMConfig, repo string) ([]string, error) {
	return nil, nil
}

func (s *SVN) CheckToken(scm *api.SCMConfig) bool {
	//username, password, err := s.spilitToken(scm.Token)
	//if err != nil {
	//	return false
	//}
	//fmt.Println(username)
	//fmt.Println(password)
	//
	//url := scm.Server
	//args := []string{"list", "--username", username, "--password", password,
	//	"--non-interactive", "--trust-server-cert-failures", "unknown-ca", "--no-auth-cache", url}
	//_, err = executil.RunInDir("./", "svn", args...)
	//if err != nil {
	//	log.Errorf("Error when list repos as : %v", err)
	//	return false
	//}
	return true
}

func (s *SVN) NewTagFromLatest(scm *api.SCMConfig, tagName, description, commitID, url string) error {
	username, password, err := s.spilitToken(scm.Token)
	if err != nil {
		return err
	}

	if !strings.Contains(url, "/trunk") {
		return fmt.Errorf("not standard SVN dirs, cannot create tag")
	}

	tagURL := strings.Split(url, "/trunk")[0] + "/tags/" + tagName + "/"
	log.Infof("trunk[%s] tag[%s]", url, tagURL)
	args := []string{"copy", url, tagURL, "-m", "Cyclone auto tag " + tagName,
		"--username", username, "--password", password,
		"--non-interactive", "--trust-server-cert-failures", "unknown-ca", "--no-auth-cache"}

	output, err := executil.RunInDir("./", "svn", args...)
	log.Infof("Command output: %+v", string(output))
	if err == nil {
		log.Infof("Successfully svn create tag.")
	}
	return err
}

func (s *SVN) CreateWebHook(scm *api.SCMConfig, repoURL string, webHook *scm.WebHook) error {
	return nil
}
func (s *SVN) DeleteWebHook(scm *api.SCMConfig, repoURL string, webHookUrl string) error {
	return nil
}

// TODO: oauth by SVN
func (s *SVN) GetAuthCodeURL(state string, scmType api.SCMType) (string, error) {
	return "", fmt.Errorf("Not implemented")
}
func (s *SVN) Authcallback(code string, state string) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

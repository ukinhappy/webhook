package hook

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/ukinhappy/webhook/logger"
	"os/exec"
	"sync"
)

type webhook struct {
	Name      string
	Url       string
	Branch    []string
	Event     []string
	ShellPath string
}

func (w *webhook) Verify(branch, event string) bool {
	var bt, et bool
	for _, b := range w.Branch {
		if b == branch {
			bt = true
		}
	}
	// TODO TEST
	bt = true
	for _, e := range w.Event {
		if e == event {
			et = true
		}
	}
	return bt && et
}
func (w *webhook) Do() error {
	cmd := exec.Command("/bin/bash", "-c", w.ShellPath)
	body, err := cmd.Output()
	logger.Infof("bash do %s", body)
	return err
}

var lock sync.RWMutex
var allWebHook []*webhook
var webHookMap map[string]*webhook

func LoadAllWebHook() {
	lock.Lock()
	defer lock.Unlock()
	webHookMap = make(map[string]*webhook, 0)
	if err := viper.UnmarshalKey("projects", &allWebHook); err != nil {
		logger.Panicf("UnmarshalKey projects %v", err)
	}
	for _, w := range allWebHook {
		webHookMap[w.Name] = w
	}
}

func Do(name, branch, event string) error {
	lock.RLock()
	defer lock.RUnlock()
	if w, ok := webHookMap[name]; ok && w.Verify(branch, event) {
		return w.Do()
	}

	return fmt.Errorf("%s项目不支持")
}

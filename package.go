package debpack

import (
	"bytes"
	"fmt"
	"github.com/xor-gate/debpkg"
	"log"
	"os/exec"
	"text/template"
)

type Package struct {
	Name            string
	Version         string
	Description     string
	MaintainerName  string
	MaintainerEmail string
	User            string

	ExecFileName      string
	ConfigFileName    string
	DefaultFileName   string
	LogFileName       string
	InitFileName      string
	LogRotateFileName string

	ConfigFile    string
	DefaultFile   string
	InitFile      string
	LogRotateFile string

	ConfFileNames []string
	deb           *debpkg.DebPkg
}

func NewPackage(name string, version string) *Package {
	self := Package{Name: name, Version: version}
	self.deb = debpkg.New()

	return &self
}

func (self *Package) SetMeta() {

	self.User = self.Name
	self.deb.SetName(self.Name)
	self.deb.SetVersion(self.Version)
	self.deb.SetArchitecture("amd64")
	self.deb.SetSection("misc")
	self.deb.SetPriority(debpkg.PriorityOptional)
	self.deb.SetShortDescription(self.Description)
	self.deb.SetDepends("lsb-base")
	self.deb.SetMaintainer(self.MaintainerName)
	self.deb.SetMaintainerEmail(self.MaintainerEmail)
}

func (self *Package) Build(entrypoint string) {
	flags := fmt.Sprintf("-X main.Version=%s", self.Version)
	args := []string{"build", "-ldflags", flags, entrypoint}
	if out, err := exec.Command("go", args...).CombinedOutput(); err != nil {
		log.Fatal("Exec error", err)
	} else {
		log.Println("Exec done", string(out))
	}

}

func (self *Package) SetDefaultFilenames() {
	self.ExecFileName = fmt.Sprintf("/usr/sbin/%s", self.Name)
	self.ConfigFileName = fmt.Sprintf("/etc/%s/%s.conf", self.Name, self.Name)
	self.DefaultFileName = fmt.Sprintf("/etc/default/%s", self.Name)
	self.LogFileName = fmt.Sprintf("/var/log/%s/%s.log", self.Name, self.Name)
	self.LogRotateFileName = fmt.Sprintf("/etc/logrotate.d/%s", self.Name)
	self.InitFileName = fmt.Sprintf("/etc/init.d/%s", self.Name)
	self.ConfFileNames = []string{self.ConfigFileName, self.InitFileName, self.LogRotateFileName, self.DefaultFileName}
}

func (self *Package) SetDefaultFiles() {
	self.ConfigFile = "# Empty config file"
	self.DefaultFile = "# Empty default file"

	self.InitFile = self.MustApplyTemplate(Sysvinit)
	self.LogRotateFile = self.MustApplyTemplate(Logrotate)
}

func (self *Package) AddControls() {
	controls := map[string]string{
		"preinst":  self.MustApplyTemplate(Preinst),
		"postinst": self.MustApplyTemplate(Postinst),
		"prerm":    self.MustApplyTemplate(Prerm),
		"postrm":   self.MustApplyTemplate(Postrm),
	}

	for k, v := range controls {
		if err := self.deb.AddControlExtraString(k, v); err != nil {
			log.Println(err)
		}
	}

}

func (self *Package) AddFiles() {
	//add files
	if err := self.deb.AddFileString(self.InitFile, self.InitFileName); err != nil {
		log.Println(err)
	}

	if err := self.deb.AddFileString(self.LogRotateFile, self.LogRotateFileName); err != nil {
		log.Println(err)
	}

	if err := self.deb.AddFileString(self.ConfigFile, self.ConfigFileName); err != nil {
		log.Println(err)
	}
	if err := self.deb.AddFileString(self.DefaultFile, self.DefaultFileName); err != nil {
		log.Println(err)
	}

	if err := self.deb.AddFile(self.Name, self.ExecFileName); err != nil {
		log.Println(err)
	}

}

func (self *Package) MarkConfig() {
	//mark as config
	for _, conf := range self.ConfFileNames {
		if err := self.deb.MarkConfigFile(conf); err != nil {
			log.Println(err)
		}
	}
}
func (self *Package) Pack() {
	//create deb
	if err := self.deb.Write(""); err != nil {
		log.Fatalf("Error writing outputfile: %v", err)
	}
}

func (self *Package) ApplyTemplate(tpl string) (string, error) {
	if tpl, err := template.New("").Parse(tpl); err != nil {
		return "", err
	} else {
		var buf bytes.Buffer
		if err := tpl.Execute(&buf, self); err != nil {
			return "", err
		} else {
			return buf.String(), nil
		}
	}
}

func (self *Package) MustApplyTemplate(tpl string) string {
	if res, err := self.ApplyTemplate(tpl); err != nil {
		panic(err)
	} else {
		return res
	}
}

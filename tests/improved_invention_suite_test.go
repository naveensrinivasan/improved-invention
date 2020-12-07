package tests

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestImprovedInvention(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ImprovedInvention Suite")
}

var config Config

func init() {
	configor.Load(&config, "config.yml")
}

var timeout = config.Standard.TimeoutInSeconds
var interval = config.Standard.PollIntervalInMilliseconds

var _ = BeforeSuite(func() {
	if config.Db.Initialize {
		c := exec.Command("docker", "run", "-p", "3306:3306", "--name", "some-mysql", "-e",
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", config.Db.Password), "-d", "mysql:latest", "MYSQL_DATABASE=default")
		err := c.Run()
		Expect(err).ShouldNot(HaveOccurred())
	}
})

var _ = AfterSuite(func() {
	if config.Db.Initialize {
		c := exec.Command("docker", "stop", "some-mysql")
		err := c.Run()
		Expect(err).ShouldNot(HaveOccurred())

		c = exec.Command("docker", "rm", "some-mysql")
		err = c.Run()
		Expect(err).ShouldNot(HaveOccurred())
	}
})

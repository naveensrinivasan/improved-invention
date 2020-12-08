package tests

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"testing"

	"github.com/jinzhu/configor"

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

// randStringRunes is used for generating random string.
func randStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz") // for generating random names for tests.
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// getEnvDefault returns the value of the given environment variable or a
// default value if the given environment variable is not set.
func getEnvDefault(variable string, defaultVal string) string {
	envVar, exists := os.LookupEnv(variable)
	if !exists {
		return defaultVal
	}
	return envVar
}

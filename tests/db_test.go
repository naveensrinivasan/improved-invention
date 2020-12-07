package tests

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DATABASE:Checking for Database connectivity ", func() {
	Context("when running within the DMZ", func() {
		It("Should be able to connect to the mysql server", func() {
			Eventually(func() bool {
				_, err := sql.Open("mysql",
					fmt.Sprintf("%s:%s@/%s", config.Db.User, config.Db.Password, config.Db.Name))
				return err == nil
			}, timeout, interval).Should(BeTrue())
		})
	})
})

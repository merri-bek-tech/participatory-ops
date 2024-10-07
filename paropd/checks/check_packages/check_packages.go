package check_packages

import (
	"log"
	"os/exec"
	"paropd/checks/checks_helpers"
	"strings"
	"time"
)

type CheckPackages struct {
}

// Verify that CheckPackages implements Check
var _ checks_helpers.Check = &CheckPackages{}

func (c *CheckPackages) Measure() {
	log.Println("CheckPackages measure")

	result := countUpgradeablePackages()

	log.Println("CheckPackages measure done", result)
}

func (c *CheckPackages) MeasureFrequency() time.Duration {
	return 10 * time.Second
}

// PRIVATE

func countUpgradeablePackages() int {
	out := shellCmdOutput("apt list --upgradable")
	return strings.Count(out, "\n")
}

func shellCmdOutput(shellCmd string) string {
	out, err := exec.Command("sh", "-c", shellCmd).Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(out))
}

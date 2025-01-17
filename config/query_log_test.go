package config

import (
	"time"

	"github.com/creasty/defaults"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("QueryLogConfig", func() {
	var cfg QueryLogConfig

	suiteBeforeEach()

	BeforeEach(func() {
		cfg = QueryLogConfig{
			Target:           "/dev/null",
			Type:             QueryLogTypeCsvClient,
			LogRetentionDays: 0,
			CreationAttempts: 1,
			CreationCooldown: Duration(time.Millisecond),
		}
	})

	Describe("IsEnabled", func() {
		It("should be true by default", func() {
			cfg := QueryLogConfig{}
			Expect(defaults.Set(&cfg)).Should(Succeed())

			Expect(cfg.IsEnabled()).Should(BeTrue())
		})

		When("enabled", func() {
			It("should be true", func() {
				Expect(cfg.IsEnabled()).Should(BeTrue())
			})
		})

		When("disabled", func() {
			It("should be false", func() {
				cfg := QueryLogConfig{
					Type: QueryLogTypeNone,
				}

				Expect(cfg.IsEnabled()).Should(BeFalse())
			})
		})
	})

	Describe("LogConfig", func() {
		It("should log configuration", func() {
			cfg.LogConfig(logger)

			Expect(hook.Calls).ShouldNot(BeEmpty())
			Expect(hook.Messages).Should(ContainElement(ContainSubstring("logRetentionDays:")))
		})
	})

	Describe("SetDefaults", func() {
		It("should log configuration", func() {
			cfg := QueryLogConfig{}
			Expect(cfg.Fields).Should(BeEmpty())

			Expect(defaults.Set(&cfg)).Should(Succeed())

			Expect(cfg.Fields).ShouldNot(BeEmpty())
		})
	})
})

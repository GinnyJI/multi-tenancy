package e2e

import (
	. "github.com/onsi/ginkgo"
	. "sigs.k8s.io/multi-tenancy/incubator/hnc/pkg/testutils"
)

var _ = Describe("HNS set-config", func() {
	It("Should use '--force' flag to change from 'Ignore' to 'Propagate'", func() {
		MustRun("kubectl hns config set-type --apiVersion v1 --kind Secret Ignore")
		MustNotRun("kubectl hns config set-type --apiVersion v1 --kind Secret Propagate")
		MustRun("kubectl hns config set-type --apiVersion v1 --kind Secret Propagate --force")
		// check that we don't need '--force' flag when changing it back
		MustRun("kubectl hns config set-type --apiVersion v1 --kind Secret Ignore")
	})
})

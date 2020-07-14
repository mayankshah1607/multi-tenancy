package test

import (
	"os"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Delete-anchor-crd", func() {

	hncRecoverPath := os.Getenv("HNC_REPAIR")

	const (
		nsParent = "delete-crd-parent"
		nsChild = "delete-crd-child"
	)

	AfterEach(func() {
		cleanupNamespaces(nsParent, nsChild)
		err := tryRun("kubectl apply -f", hncRecoverPath)
		if err != nil {
			GinkgoT().Log("-----------------------------WARNING------------------------------")
			GinkgoT().Logf("WARNING: COULDN'T REPAIR HNC: %v", err)
			GinkgoT().Log("ANY TEST AFTER THIS COULD FAIL BECAUSE WE COULDN'T REPAIR HNC HERE")
			GinkgoT().Log("------------------------------------------------------------------")
			GinkgoT().FailNow()
		}
	})

	It("should create parent and deletable child, and delete the CRD", func() {
		// we don't want to destroy the HNC without being able to repair it, so skip this test if recovery path not set
		if hncRecoverPath == ""{
			Skip("Environment variable HNC_REPAIR not set. Skipping reocovering HNC.")
		}
		// set up
		mustRun("kubectl create ns", nsParent)
		mustRun("kubectl hns create", nsChild, "-n", nsParent)
		// test
		mustRun("kubectl delete customresourcedefinition/subnamespaceanchors.hnc.x-k8s.io")
		// verify
		mustRun("kubectl get ns", nsChild)
	})
})

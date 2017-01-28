package linter

import (
	"io/ioutil"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linter", func() {
	Describe("LintFiles", func() {
		It("should lint the files in the given directory", func() {
			dir, err := ioutil.TempDir("", "test-dir")
			Expect(err).NotTo(HaveOccurred())

			file, err := ioutil.TempFile(dir, "tempfile with spaces")
			err = LintFiles(dir)
			Expect(err).ToNot(HaveOccurred())

			Expect(file.Name()).ShouldNot(BeAnExistingFile())
			Expect(strings.Replace(file.Name(), " ", "-", -1)).Should(BeAnExistingFile())
			err = os.RemoveAll(dir)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should rename the files with spaces to without spaces", func() {
			_, err := os.Create("banana file here")
			Expect(err).ToNot(HaveOccurred())
			LintFiles(".")
			Expect("banana file here").ShouldNot(BeAnExistingFile())
			Expect("banana-file-here").Should(BeAnExistingFile())
			err = os.Remove("banana-file-here")
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when the rename fails", func() {
			It("should return a useful error", func() {
				dir, err := ioutil.TempDir("", "test-dir")
				Expect(err).NotTo(HaveOccurred())

				_, err = ioutil.TempFile(dir, "tempfile with spaces")
				err = os.RemoveAll(dir)
				Expect(err).ToNot(HaveOccurred())

				err = LintFiles(dir)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("listFiles", func() {
		It("should list all the files in the current directory", func() {
			filenames, err := listFiles(".")
			Expect(err).ToNot(HaveOccurred())
			Expect(filenames).To(ContainElement("linter.go"))
			Expect(filenames).To(ContainElement("linter_suite_test.go"))
			Expect(filenames).To(ContainElement("linter_test.go"))
		})
	})

	Describe("lint", func() {
		Context("when the filename does not have spaces", func() {
			It("should leave it as it is", func() {
				filename := "bananaFile"
				lintedFilename := lint(filename)
				Expect(lintedFilename).Should(Equal(filename))
			})
		})

		Context("when the filename does have spaces", func() {
			It("should remove all the spaces", func() {
				filename := "bananaFile and an applefile"
				lintedFilename := lint(filename)
				Expect(strings.ContainsAny(lintedFilename, " ")).To(BeFalse())
			})
		})
	})
})

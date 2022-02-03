package shared_tests

import (
	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
	"io/ioutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func RunfpdfPpoTests(pdfiumContainer pdfium.Pdfium, testsPath string, prefix string) {
	Describe("fpdf_ppo", func() {
		Context("no document", func() {
			When("is opened", func() {
				It("returns an error when calling FPDF_ImportPages", func() {
					FPDF_ImportPages, err := pdfiumContainer.FPDF_ImportPages(&requests.FPDF_ImportPages{})
					Expect(err).To(MatchError("document not given"))
					Expect(FPDF_ImportPages).To(BeNil())
				})

				It("returns an error when calling FPDF_CopyViewerPreferences", func() {
					FPDF_CopyViewerPreferences, err := pdfiumContainer.FPDF_CopyViewerPreferences(&requests.FPDF_CopyViewerPreferences{})
					Expect(err).To(MatchError("document not given"))
					Expect(FPDF_CopyViewerPreferences).To(BeNil())
				})

				It("returns an error when calling FPDF_ImportPagesByIndex", func() {
					FPDF_ImportPagesByIndex, err := pdfiumContainer.FPDF_ImportPagesByIndex(&requests.FPDF_ImportPagesByIndex{})
					Expect(err).To(MatchError("document not given"))
					Expect(FPDF_ImportPagesByIndex).To(BeNil())
				})

				It("returns an error when calling FPDF_ImportNPagesToOne", func() {
					FPDF_ImportNPagesToOne, err := pdfiumContainer.FPDF_ImportNPagesToOne(&requests.FPDF_ImportNPagesToOne{})
					Expect(err).To(MatchError("document not given"))
					Expect(FPDF_ImportNPagesToOne).To(BeNil())
				})

				It("returns an error when calling FPDF_NewXObjectFromPage", func() {
					FPDF_NewXObjectFromPage, err := pdfiumContainer.FPDF_NewXObjectFromPage(&requests.FPDF_NewXObjectFromPage{})
					Expect(err).To(MatchError("document not given"))
					Expect(FPDF_NewXObjectFromPage).To(BeNil())
				})

				It("returns an error when calling FPDF_CloseXObject", func() {
					FPDF_CloseXObject, err := pdfiumContainer.FPDF_CloseXObject(&requests.FPDF_CloseXObject{})
					Expect(err).To(MatchError("xObject not given"))
					Expect(FPDF_CloseXObject).To(BeNil())
				})

				It("returns an error when calling FPDF_NewFormObjectFromXObject", func() {
					FPDF_NewFormObjectFromXObject, err := pdfiumContainer.FPDF_NewFormObjectFromXObject(&requests.FPDF_NewFormObjectFromXObject{})
					Expect(err).To(MatchError("xObject not given"))
					Expect(FPDF_NewFormObjectFromXObject).To(BeNil())
				})
			})
		})

		Context("a normal PDF file", func() {
			var doc references.FPDF_DOCUMENT

			BeforeEach(func() {
				pdfData, err := ioutil.ReadFile(testsPath + "/testdata/test.pdf")
				Expect(err).To(BeNil())
				if err != nil {
					return
				}

				newDoc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData)
				if err != nil {
					return
				}

				doc = *newDoc
			})

			AfterEach(func() {
				err := pdfiumContainer.FPDF_CloseDocument(doc)
				Expect(err).To(BeNil())
			})

			When("is opened", func() {
				When("no destination document is given", func() {
					It("returns an error when calling FPDF_ImportPages", func() {
						FPDF_ImportPages, err := pdfiumContainer.FPDF_ImportPages(&requests.FPDF_ImportPages{
							Source: doc,
						})
						Expect(err).To(MatchError("document not given"))
						Expect(FPDF_ImportPages).To(BeNil())
					})

					It("returns an error when calling FPDF_CopyViewerPreferences", func() {
						FPDF_CopyViewerPreferences, err := pdfiumContainer.FPDF_CopyViewerPreferences(&requests.FPDF_CopyViewerPreferences{
							Source: doc,
						})
						Expect(err).To(MatchError("document not given"))
						Expect(FPDF_CopyViewerPreferences).To(BeNil())
					})

					It("returns an error when calling FPDF_ImportPagesByIndex", func() {
						FPDF_ImportPagesByIndex, err := pdfiumContainer.FPDF_ImportPagesByIndex(&requests.FPDF_ImportPagesByIndex{
							Source: doc,
						})
						Expect(err).To(MatchError("document not given"))
						Expect(FPDF_ImportPagesByIndex).To(BeNil())
					})

					It("returns an error when calling FPDF_NewXObjectFromPage", func() {
						FPDF_NewXObjectFromPage, err := pdfiumContainer.FPDF_NewXObjectFromPage(&requests.FPDF_NewXObjectFromPage{
							Source: doc,
						})
						Expect(err).To(MatchError("document not given"))
						Expect(FPDF_NewXObjectFromPage).To(BeNil())
					})
				})

				When("no page information is given to FPDF_ImportNPagesToOne", func() {
					It("returns an error when calling FPDF_ImportNPagesToOne", func() {
						FPDF_ImportNPagesToOne, err := pdfiumContainer.FPDF_ImportNPagesToOne(&requests.FPDF_ImportNPagesToOne{
							Source: doc,
						})
						Expect(err).To(MatchError("import of pages failed"))
						Expect(FPDF_ImportNPagesToOne).To(BeNil())
					})
				})

				When("page information is given to FPDF_ImportNPagesToOne", func() {
					It("returns a new document when calling FPDF_ImportNPagesToOne", func() {
						FPDF_ImportNPagesToOne, err := pdfiumContainer.FPDF_ImportNPagesToOne(&requests.FPDF_ImportNPagesToOne{
							Source:          doc,
							OutputWidth:     2000,
							OutputHeight:    2000,
							NumPagesOnXAxis: 2,
							NumPagesOnYAxis: 2,
						})
						Expect(err).To(BeNil())
						Expect(FPDF_ImportNPagesToOne).To(Not(BeNil()))
						Expect(FPDF_ImportNPagesToOne.Document).To(Not(BeNil()))
					})
				})

				Context("a second PDF file is opened", func() {
					var doc2 references.FPDF_DOCUMENT

					BeforeEach(func() {
						pdfData, err := ioutil.ReadFile(testsPath + "/testdata/viewer_ref.pdf")
						Expect(err).To(BeNil())
						if err != nil {
							return
						}

						newDoc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData)
						if err != nil {
							return
						}

						doc2 = *newDoc
					})

					AfterEach(func() {
						err := pdfiumContainer.FPDF_CloseDocument(doc2)
						Expect(err).To(BeNil())
					})

					When("is opened", func() {
						It("returns no error when FPDF_ImportPages is called", func() {
							FPDF_ImportPages, err := pdfiumContainer.FPDF_ImportPages(&requests.FPDF_ImportPages{
								Source:      doc2,
								Destination: doc,
							})
							Expect(err).To(BeNil())
							Expect(FPDF_ImportPages).To(Not(BeNil()))
						})

						It("returns no error when FPDF_ImportPages is called with a valid pagerange", func() {
							pageRange := "1"
							FPDF_ImportPages, err := pdfiumContainer.FPDF_ImportPages(&requests.FPDF_ImportPages{
								Source:      doc2,
								Destination: doc,
								PageRange:   &pageRange,
							})
							Expect(err).To(BeNil())
							Expect(FPDF_ImportPages).To(Not(BeNil()))
						})

						It("returns no error when FPDF_ImportPages is called with an invalid pagerange", func() {
							pageRange := "32"
							FPDF_ImportPages, err := pdfiumContainer.FPDF_ImportPages(&requests.FPDF_ImportPages{
								Source:      doc2,
								Destination: doc,
								PageRange:   &pageRange,
							})
							Expect(err).To(MatchError("import of pages failed"))
							Expect(FPDF_ImportPages).To(BeNil())
						})

						It("returns an error when calling FPDF_CopyViewerPreferences with a source document that has no viewer preferences", func() {
							FPDF_CopyViewerPreferences, err := pdfiumContainer.FPDF_CopyViewerPreferences(&requests.FPDF_CopyViewerPreferences{
								Source:      doc,
								Destination: doc2,
							})
							Expect(err).To(MatchError("copy of viewer preferences failed"))
							Expect(FPDF_CopyViewerPreferences).To(BeNil())
						})

						It("returns no error when calling FPDF_CopyViewerPreferences with a source document that has viewer preferences", func() {
							FPDF_CopyViewerPreferences, err := pdfiumContainer.FPDF_CopyViewerPreferences(&requests.FPDF_CopyViewerPreferences{
								Source:      doc2,
								Destination: doc,
							})
							Expect(err).To(BeNil())
							Expect(FPDF_CopyViewerPreferences).To(Not(BeNil()))
						})

						It("returns no error when calling FPDF_ImportPagesByIndex", func() {
							FPDF_ImportPagesByIndex, err := pdfiumContainer.FPDF_ImportPagesByIndex(&requests.FPDF_ImportPagesByIndex{
								Source:      doc2,
								Destination: doc,
							})
							Expect(err).To(BeNil())
							Expect(FPDF_ImportPagesByIndex).To(Not(BeNil()))
						})

						It("returns no error when calling FPDF_ImportPagesByIndex with indices", func() {
							FPDF_ImportPagesByIndex, err := pdfiumContainer.FPDF_ImportPagesByIndex(&requests.FPDF_ImportPagesByIndex{
								Source:      doc2,
								Destination: doc,
								PageIndices: []int{0},
							})
							Expect(err).To(BeNil())
							Expect(FPDF_ImportPagesByIndex).To(Not(BeNil()))
						})

						It("returns no error when calling FPDF_ImportPagesByIndex with invalid indices", func() {
							FPDF_ImportPagesByIndex, err := pdfiumContainer.FPDF_ImportPagesByIndex(&requests.FPDF_ImportPagesByIndex{
								Source:      doc2,
								Destination: doc,
								PageIndices: []int{32},
							})
							Expect(err).To(MatchError("import of pages failed"))
							Expect(FPDF_ImportPagesByIndex).To(BeNil())
						})

						It("returns an xobject when calling FPDF_NewXObjectFromPage", func() {
							FPDF_NewXObjectFromPage, err := pdfiumContainer.FPDF_NewXObjectFromPage(&requests.FPDF_NewXObjectFromPage{
								Source:          doc2,
								Destination:     doc,
								SourcePageIndex: 0,
							})
							Expect(err).To(BeNil())
							Expect(FPDF_NewXObjectFromPage).To(Not(BeNil()))
							Expect(FPDF_NewXObjectFromPage.XObject).To(Not(BeNil()))
						})

						It("closes an xobject when calling FPDF_NewXObjectFromPage", func() {
							FPDF_NewXObjectFromPage, err := pdfiumContainer.FPDF_NewXObjectFromPage(&requests.FPDF_NewXObjectFromPage{
								Source:          doc2,
								Destination:     doc,
								SourcePageIndex: 0,
							})
							Expect(err).To(BeNil())
							Expect(FPDF_NewXObjectFromPage).To(Not(BeNil()))
							Expect(FPDF_NewXObjectFromPage.XObject).To(Not(BeNil()))

							FPDF_CloseXObject, err := pdfiumContainer.FPDF_CloseXObject(&requests.FPDF_CloseXObject{
								XObject: FPDF_NewXObjectFromPage.XObject,
							})
							Expect(err).To(BeNil())
							Expect(FPDF_CloseXObject).To(Not(BeNil()))
						})

						It("returns an page object when calling FPDF_NewXObjectFromPage", func() {
							FPDF_NewXObjectFromPage, err := pdfiumContainer.FPDF_NewXObjectFromPage(&requests.FPDF_NewXObjectFromPage{
								Source:          doc2,
								Destination:     doc,
								SourcePageIndex: 0,
							})
							Expect(err).To(BeNil())
							Expect(FPDF_NewXObjectFromPage).To(Not(BeNil()))
							Expect(FPDF_NewXObjectFromPage.XObject).To(Not(BeNil()))

							FPDF_NewFormObjectFromXObject, err := pdfiumContainer.FPDF_NewFormObjectFromXObject(&requests.FPDF_NewFormObjectFromXObject{
								XObject: FPDF_NewXObjectFromPage.XObject,
							})
							Expect(err).To(BeNil())
							Expect(FPDF_NewFormObjectFromXObject).To(Not(BeNil()))
							Expect(FPDF_NewFormObjectFromXObject.PageObject).To(Not(BeNil()))
						})
					})
				})
			})
		})
	})
}

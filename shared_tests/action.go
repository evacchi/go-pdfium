package shared_tests

import (
	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/enums"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
	"io/ioutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

func RunActionTests(pdfiumContainer pdfium.Pdfium, testsPath string, prefix string) {
	Describe("action", func() {
		Context("no document", func() {
			When("is opened", func() {
				It("returns an error when calling GetActionInfo", func() {
					getActionInfo, err := pdfiumContainer.GetActionInfo(&requests.GetActionInfo{})
					Expect(err).To(MatchError("document not given"))
					Expect(getActionInfo).To(BeNil())
				})

				It("returns an error when calling GetDestInfo", func() {
					getDestInfo, err := pdfiumContainer.GetDestInfo(&requests.GetDestInfo{})
					Expect(err).To(MatchError("document not given"))
					Expect(getDestInfo).To(BeNil())
				})
			})
		})

		Context("a normal PDF file", func() {
			var doc references.FPDF_DOCUMENT

			BeforeEach(func() {
				pdfData, err := ioutil.ReadFile(testsPath + "/testdata/test.pdf")
				Expect(err).To(BeNil())

				newDoc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData)
				Expect(err).To(BeNil())

				doc = *newDoc
			})

			AfterEach(func() {
				err := pdfiumContainer.FPDF_CloseDocument(doc)
				Expect(err).To(BeNil())
			})

			When("GetActionInfo is called without an action", func() {
				It("returns an error", func() {
					actionInfo, err := pdfiumContainer.GetActionInfo(&requests.GetActionInfo{
						Document: doc,
					})
					Expect(err).To(MatchError("action not given"))
					Expect(actionInfo).To(BeNil())
				})
			})

			When("GetDestInfo is called without a dest", func() {
				It("returns an error", func() {
					destInfo, err := pdfiumContainer.GetDestInfo(&requests.GetDestInfo{
						Document: doc,
					})
					Expect(err).To(MatchError("dest not given"))
					Expect(destInfo).To(BeNil())
				})
			})
		})

		Context("a PDF file with a launch action", func() {
			var doc references.FPDF_DOCUMENT
			var link references.FPDF_LINK
			BeforeEach(func() {
				pdfData, err := ioutil.ReadFile(testsPath + "/testdata/launch_action.pdf")
				Expect(err).To(BeNil())

				newDoc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData)
				Expect(err).To(BeNil())

				doc = *newDoc

				pageLink, err := pdfiumContainer.FPDFLink_GetLinkAtPoint(&requests.FPDFLink_GetLinkAtPoint{
					Page: requests.Page{
						ByIndex: &requests.PageByIndex{
							Document: doc,
							Index:    0,
						},
					},
					X: 100,
					Y: 100,
				})
				Expect(err).To(BeNil())
				Expect(pageLink).To(Not(BeNil()))
				Expect(pageLink.Link).To(Not(BeNil()))
				link = *pageLink.Link
			})

			AfterEach(func() {
				err := pdfiumContainer.FPDF_CloseDocument(doc)
				Expect(err).To(BeNil())
			})

			Context("A launch action is loaded", func() {
				var action references.FPDF_ACTION

				BeforeEach(func() {
					actionResp, err := pdfiumContainer.FPDFLink_GetAction(&requests.FPDFLink_GetAction{
						Link: link,
					})
					Expect(err).To(BeNil())
					Expect(actionResp).To(Not(BeNil()))
					Expect(actionResp.Action).To(Not(BeNil()))
					action = *actionResp.Action
				})

				When("GetActionInfo is called", func() {
					It("returns the correct info", func() {
						actionInfo, err := pdfiumContainer.GetActionInfo(&requests.GetActionInfo{
							Document: doc,
							Action:   action,
						})
						Expect(err).To(BeNil())
						Expect(actionInfo).To(Not(BeNil()))
						Expect(actionInfo.ActionInfo).To(Not(BeNil()))
						Expect(actionInfo.ActionInfo).To(MatchAllFields(Fields{
							"Type":      Equal(enums.FPDF_ACTION_ACTION_LAUNCH),
							"Reference": Not(BeNil()),
							"DestInfo":  BeNil(),
							"FilePath":  PointTo(Equal("test.pdf")),
							"URIPath":   BeNil(),
						}))
					})
				})
			})
		})

		Context("a PDF file with a uri action", func() {
			var doc references.FPDF_DOCUMENT
			var link references.FPDF_LINK
			BeforeEach(func() {
				pdfData, err := ioutil.ReadFile(testsPath + "/testdata/uri_action.pdf")
				Expect(err).To(BeNil())

				newDoc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData)
				Expect(err).To(BeNil())

				doc = *newDoc

				pageLink, err := pdfiumContainer.FPDFLink_GetLinkAtPoint(&requests.FPDFLink_GetLinkAtPoint{
					Page: requests.Page{
						ByIndex: &requests.PageByIndex{
							Document: doc,
							Index:    0,
						},
					},
					X: 100,
					Y: 100,
				})
				Expect(err).To(BeNil())
				Expect(pageLink).To(Not(BeNil()))
				Expect(pageLink.Link).To(Not(BeNil()))
				link = *pageLink.Link
			})

			AfterEach(func() {
				err := pdfiumContainer.FPDF_CloseDocument(doc)
				Expect(err).To(BeNil())
			})

			Context("A uri action is loaded", func() {
				var action references.FPDF_ACTION

				BeforeEach(func() {
					actionResp, err := pdfiumContainer.FPDFLink_GetAction(&requests.FPDFLink_GetAction{
						Link: link,
					})
					Expect(err).To(BeNil())
					Expect(actionResp).To(Not(BeNil()))
					Expect(actionResp.Action).To(Not(BeNil()))
					action = *actionResp.Action
				})

				When("GetActionInfo is called", func() {
					It("returns the correct info", func() {
						actionInfo, err := pdfiumContainer.GetActionInfo(&requests.GetActionInfo{
							Document: doc,
							Action:   action,
						})
						Expect(err).To(BeNil())
						Expect(actionInfo).To(Not(BeNil()))
						Expect(actionInfo.ActionInfo).To(Not(BeNil()))
						Expect(actionInfo.ActionInfo).To(MatchAllFields(Fields{
							"Type":      Equal(enums.FPDF_ACTION_ACTION_URI),
							"Reference": Not(BeNil()),
							"DestInfo":  BeNil(),
							"FilePath":  BeNil(),
							"URIPath":   PointTo(Equal("https://example.com/page.html")),
						}))
					})
				})
			})
		})

		Context("a PDF file with a goto action", func() {
			var doc references.FPDF_DOCUMENT
			var link references.FPDF_LINK
			BeforeEach(func() {
				pdfData, err := ioutil.ReadFile(testsPath + "/testdata/goto_action.pdf")
				Expect(err).To(BeNil())

				newDoc, err := pdfiumContainer.NewDocumentFromBytes(&pdfData)
				Expect(err).To(BeNil())

				doc = *newDoc

				pageLink, err := pdfiumContainer.FPDFLink_GetLinkAtPoint(&requests.FPDFLink_GetLinkAtPoint{
					Page: requests.Page{
						ByIndex: &requests.PageByIndex{
							Document: doc,
							Index:    0,
						},
					},
					X: 100,
					Y: 100,
				})
				Expect(err).To(BeNil())
				Expect(pageLink).To(Not(BeNil()))
				Expect(pageLink.Link).To(Not(BeNil()))
				link = *pageLink.Link
			})

			AfterEach(func() {
				err := pdfiumContainer.FPDF_CloseDocument(doc)
				Expect(err).To(BeNil())
			})

			Context("A goto action is loaded", func() {
				var action references.FPDF_ACTION

				BeforeEach(func() {
					actionResp, err := pdfiumContainer.FPDFLink_GetAction(&requests.FPDFLink_GetAction{
						Link: link,
					})
					Expect(err).To(BeNil())
					Expect(actionResp).To(Not(BeNil()))
					Expect(actionResp.Action).To(Not(BeNil()))
					action = *actionResp.Action
				})

				When("GetActionInfo is called", func() {
					It("returns the correct info", func() {
						actionInfo, err := pdfiumContainer.GetActionInfo(&requests.GetActionInfo{
							Document: doc,
							Action:   action,
						})
						Expect(err).To(BeNil())
						Expect(actionInfo).To(Not(BeNil()))
						Expect(actionInfo.ActionInfo).To(Not(BeNil()))
						Expect(actionInfo.ActionInfo).To(MatchAllFields(Fields{
							"Type":      Equal(enums.FPDF_ACTION_ACTION_GOTO),
							"Reference": Not(BeNil()),
							"DestInfo": PointTo(MatchAllFields(Fields{
								"PageIndex": Equal(1),
								"Reference": Not(BeNil()),
							})),
							"FilePath": BeNil(),
							"URIPath":  BeNil(),
						}))
					})
				})

				Context("A dest action is loaded", func() {
					var dest references.FPDF_DEST

					BeforeEach(func() {
						actionDest, err := pdfiumContainer.FPDFAction_GetDest(&requests.FPDFAction_GetDest{
							Document: doc,
							Action:   action,
						})
						Expect(err).To(BeNil())
						Expect(actionDest).To(Not(BeNil()))
						Expect(actionDest.Dest).To(Not(BeNil()))
						dest = *actionDest.Dest
					})

					When("GetDestInfo is called", func() {
						It("returns the correct dest info", func() {
							destInfo, err := pdfiumContainer.GetDestInfo(&requests.GetDestInfo{
								Document: doc,
								Dest:     dest,
							})
							Expect(err).To(BeNil())
							Expect(destInfo).To(Not(BeNil()))
							Expect(destInfo.DestInfo).To(Not(BeNil()))
							Expect(destInfo.DestInfo).To(MatchAllFields(Fields{
								"PageIndex": Equal(1),
								"Reference": Not(BeNil()),
							}))
						})
					})
				})
			})
		})
	})
}

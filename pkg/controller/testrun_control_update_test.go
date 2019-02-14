package controller

import (
	"testing"

	"github.com/gardener/test-infra/pkg/testmachinery/testdefinition"

	argov1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	tmv1beta1 "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testmachinery Controller Suite")
}

var (
	workflowTmpl argov1.Workflow
	testrunTmpl  tmv1beta1.Testrun
)

var _ = Describe("Testmachinery controller update", func() {

	BeforeSuite(func() {
		testrunTmpl = tmv1beta1.Testrun{
			Status: tmv1beta1.TestrunStatus{
				Steps: [][]*tmv1beta1.TestflowStepStatus{
					[]*tmv1beta1.TestflowStepStatus{
						&tmv1beta1.TestflowStepStatus{
							Phase: tmv1beta1.PhaseStatusInit,
							TestDefinition: tmv1beta1.TestflowStepStatusTestDefinition{
								Name: "testdef1",
								Position: map[string]string{
									testdefinition.AnnotationPosition: "0/0",
									testdefinition.AnnotationFlow:     "flow",
								},
							},
						},
					},
				},
			},
		}
		workflowTmpl = argov1.Workflow{
			Spec: argov1.WorkflowSpec{
				Templates: []argov1.Template{
					argov1.Template{
						Name: "template1",
						Metadata: argov1.Metadata{
							Annotations: map[string]string{
								testdefinition.AnnotationPosition:    "0/0",
								testdefinition.AnnotationFlow:        "flow",
								testdefinition.AnnotationTestDefName: "testdef1",
								"testannotation":                     "anything",
							},
						},
					},
				},
			},
			Status: argov1.WorkflowStatus{
				Nodes: map[string]argov1.NodeStatus{
					"node1": argov1.NodeStatus{
						TemplateName: "template1",
						Phase:        argov1.NodeRunning,
					},
				},
			},
		}
	})

	Context("Update status", func() {
		It("should update the status of 1 step and 1 template", func() {
			tr := testrunTmpl
			wf := workflowTmpl
			updateStepsStatus(&tr, &wf)
			Expect(tr.Status.Steps).To(Equal([][]*tmv1beta1.TestflowStepStatus{
				[]*tmv1beta1.TestflowStepStatus{
					&tmv1beta1.TestflowStepStatus{
						Phase: argov1.NodeRunning,
						TestDefinition: tmv1beta1.TestflowStepStatusTestDefinition{
							Name: "testdef1",
							Position: map[string]string{
								testdefinition.AnnotationPosition: "0/0",
								testdefinition.AnnotationFlow:     "flow",
							},
						},
					},
				},
			}))
		})

		// It("should update the status of multiple steps and templates", func() {
		// 	tr := testrunTmpl
		// 	tr.Status.Steps = [][]*tmv1beta1.TestflowStepStatus{
		// 		[]*tmv1beta1.TestflowStepStatus{
		// 			&tmv1beta1.TestflowStepStatus{
		// 				Phase: tmv1beta1.PhaseStatusInit,
		// 				TestDefinition: tmv1beta1.TestflowStepStatusTestDefinition{
		// 					Name: "testdef1",
		// 					Position: map[string]string{
		// 						testdefinition.AnnotationPosition: "0/0",
		// 						testdefinition.AnnotationFlow:     "flow",
		// 					},
		// 				},
		// 			},
		// 		},
		// 		[]*tmv1beta1.TestflowStepStatus{
		// 			&tmv1beta1.TestflowStepStatus{
		// 				Phase: tmv1beta1.PhaseStatusInit,
		// 				TestDefinition: tmv1beta1.TestflowStepStatusTestDefinition{
		// 					Name: "testdef2",
		// 					Position: map[string]string{
		// 						testdefinition.AnnotationPosition: "1/0",
		// 						testdefinition.AnnotationFlow:     "flow",
		// 					},
		// 				},
		// 			},
		// 			&tmv1beta1.TestflowStepStatus{
		// 				Phase: tmv1beta1.PhaseStatusInit,
		// 				TestDefinition: tmv1beta1.TestflowStepStatusTestDefinition{
		// 					Name: "testdef3",
		// 					Position: map[string]string{
		// 						testdefinition.AnnotationPosition: "1/1",
		// 						testdefinition.AnnotationFlow:     "flow",
		// 					},
		// 				},
		// 			},
		// 		},
		// 	}
		// 	wf := workflowTmpl
		// 	updateStepsStatus(&tr, &wf)
		// 	Expect(tr.Status.Steps).To(Equal([][]*tmv1beta1.TestflowStepStatus{
		// 		[]*tmv1beta1.TestflowStepStatus{
		// 			&tmv1beta1.TestflowStepStatus{
		// 				Phase: argov1.NodeRunning,
		// 				TestDefinition: tmv1beta1.TestflowStepStatusTestDefinition{
		// 					Name: "testdef1",
		// 					Position: map[string]string{
		// 						testdefinition.AnnotationPosition: "0/0",
		// 						testdefinition.AnnotationFlow:     "flow",
		// 					},
		// 				},
		// 			},
		// 		},
		// 	}))
		// })

	})
})

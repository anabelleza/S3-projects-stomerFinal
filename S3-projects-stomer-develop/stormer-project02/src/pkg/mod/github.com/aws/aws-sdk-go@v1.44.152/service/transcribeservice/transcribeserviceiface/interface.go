// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package transcribeserviceiface provides an interface to enable mocking the Amazon Transcribe Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package transcribeserviceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
)

// TranscribeServiceAPI provides an interface to enable mocking the
// transcribeservice.TranscribeService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Transcribe Service.
//	func myFunc(svc transcribeserviceiface.TranscribeServiceAPI) bool {
//	    // Make svc.CreateCallAnalyticsCategory request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := transcribeservice.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockTranscribeServiceClient struct {
//	    transcribeserviceiface.TranscribeServiceAPI
//	}
//	func (m *mockTranscribeServiceClient) CreateCallAnalyticsCategory(input *transcribeservice.CreateCallAnalyticsCategoryInput) (*transcribeservice.CreateCallAnalyticsCategoryOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockTranscribeServiceClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type TranscribeServiceAPI interface {
	CreateCallAnalyticsCategory(*transcribeservice.CreateCallAnalyticsCategoryInput) (*transcribeservice.CreateCallAnalyticsCategoryOutput, error)
	CreateCallAnalyticsCategoryWithContext(aws.Context, *transcribeservice.CreateCallAnalyticsCategoryInput, ...request.Option) (*transcribeservice.CreateCallAnalyticsCategoryOutput, error)
	CreateCallAnalyticsCategoryRequest(*transcribeservice.CreateCallAnalyticsCategoryInput) (*request.Request, *transcribeservice.CreateCallAnalyticsCategoryOutput)

	CreateLanguageModel(*transcribeservice.CreateLanguageModelInput) (*transcribeservice.CreateLanguageModelOutput, error)
	CreateLanguageModelWithContext(aws.Context, *transcribeservice.CreateLanguageModelInput, ...request.Option) (*transcribeservice.CreateLanguageModelOutput, error)
	CreateLanguageModelRequest(*transcribeservice.CreateLanguageModelInput) (*request.Request, *transcribeservice.CreateLanguageModelOutput)

	CreateMedicalVocabulary(*transcribeservice.CreateMedicalVocabularyInput) (*transcribeservice.CreateMedicalVocabularyOutput, error)
	CreateMedicalVocabularyWithContext(aws.Context, *transcribeservice.CreateMedicalVocabularyInput, ...request.Option) (*transcribeservice.CreateMedicalVocabularyOutput, error)
	CreateMedicalVocabularyRequest(*transcribeservice.CreateMedicalVocabularyInput) (*request.Request, *transcribeservice.CreateMedicalVocabularyOutput)

	CreateVocabulary(*transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error)
	CreateVocabularyWithContext(aws.Context, *transcribeservice.CreateVocabularyInput, ...request.Option) (*transcribeservice.CreateVocabularyOutput, error)
	CreateVocabularyRequest(*transcribeservice.CreateVocabularyInput) (*request.Request, *transcribeservice.CreateVocabularyOutput)

	CreateVocabularyFilter(*transcribeservice.CreateVocabularyFilterInput) (*transcribeservice.CreateVocabularyFilterOutput, error)
	CreateVocabularyFilterWithContext(aws.Context, *transcribeservice.CreateVocabularyFilterInput, ...request.Option) (*transcribeservice.CreateVocabularyFilterOutput, error)
	CreateVocabularyFilterRequest(*transcribeservice.CreateVocabularyFilterInput) (*request.Request, *transcribeservice.CreateVocabularyFilterOutput)

	DeleteCallAnalyticsCategory(*transcribeservice.DeleteCallAnalyticsCategoryInput) (*transcribeservice.DeleteCallAnalyticsCategoryOutput, error)
	DeleteCallAnalyticsCategoryWithContext(aws.Context, *transcribeservice.DeleteCallAnalyticsCategoryInput, ...request.Option) (*transcribeservice.DeleteCallAnalyticsCategoryOutput, error)
	DeleteCallAnalyticsCategoryRequest(*transcribeservice.DeleteCallAnalyticsCategoryInput) (*request.Request, *transcribeservice.DeleteCallAnalyticsCategoryOutput)

	DeleteCallAnalyticsJob(*transcribeservice.DeleteCallAnalyticsJobInput) (*transcribeservice.DeleteCallAnalyticsJobOutput, error)
	DeleteCallAnalyticsJobWithContext(aws.Context, *transcribeservice.DeleteCallAnalyticsJobInput, ...request.Option) (*transcribeservice.DeleteCallAnalyticsJobOutput, error)
	DeleteCallAnalyticsJobRequest(*transcribeservice.DeleteCallAnalyticsJobInput) (*request.Request, *transcribeservice.DeleteCallAnalyticsJobOutput)

	DeleteLanguageModel(*transcribeservice.DeleteLanguageModelInput) (*transcribeservice.DeleteLanguageModelOutput, error)
	DeleteLanguageModelWithContext(aws.Context, *transcribeservice.DeleteLanguageModelInput, ...request.Option) (*transcribeservice.DeleteLanguageModelOutput, error)
	DeleteLanguageModelRequest(*transcribeservice.DeleteLanguageModelInput) (*request.Request, *transcribeservice.DeleteLanguageModelOutput)

	DeleteMedicalTranscriptionJob(*transcribeservice.DeleteMedicalTranscriptionJobInput) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error)
	DeleteMedicalTranscriptionJobWithContext(aws.Context, *transcribeservice.DeleteMedicalTranscriptionJobInput, ...request.Option) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error)
	DeleteMedicalTranscriptionJobRequest(*transcribeservice.DeleteMedicalTranscriptionJobInput) (*request.Request, *transcribeservice.DeleteMedicalTranscriptionJobOutput)

	DeleteMedicalVocabulary(*transcribeservice.DeleteMedicalVocabularyInput) (*transcribeservice.DeleteMedicalVocabularyOutput, error)
	DeleteMedicalVocabularyWithContext(aws.Context, *transcribeservice.DeleteMedicalVocabularyInput, ...request.Option) (*transcribeservice.DeleteMedicalVocabularyOutput, error)
	DeleteMedicalVocabularyRequest(*transcribeservice.DeleteMedicalVocabularyInput) (*request.Request, *transcribeservice.DeleteMedicalVocabularyOutput)

	DeleteTranscriptionJob(*transcribeservice.DeleteTranscriptionJobInput) (*transcribeservice.DeleteTranscriptionJobOutput, error)
	DeleteTranscriptionJobWithContext(aws.Context, *transcribeservice.DeleteTranscriptionJobInput, ...request.Option) (*transcribeservice.DeleteTranscriptionJobOutput, error)
	DeleteTranscriptionJobRequest(*transcribeservice.DeleteTranscriptionJobInput) (*request.Request, *transcribeservice.DeleteTranscriptionJobOutput)

	DeleteVocabulary(*transcribeservice.DeleteVocabularyInput) (*transcribeservice.DeleteVocabularyOutput, error)
	DeleteVocabularyWithContext(aws.Context, *transcribeservice.DeleteVocabularyInput, ...request.Option) (*transcribeservice.DeleteVocabularyOutput, error)
	DeleteVocabularyRequest(*transcribeservice.DeleteVocabularyInput) (*request.Request, *transcribeservice.DeleteVocabularyOutput)

	DeleteVocabularyFilter(*transcribeservice.DeleteVocabularyFilterInput) (*transcribeservice.DeleteVocabularyFilterOutput, error)
	DeleteVocabularyFilterWithContext(aws.Context, *transcribeservice.DeleteVocabularyFilterInput, ...request.Option) (*transcribeservice.DeleteVocabularyFilterOutput, error)
	DeleteVocabularyFilterRequest(*transcribeservice.DeleteVocabularyFilterInput) (*request.Request, *transcribeservice.DeleteVocabularyFilterOutput)

	DescribeLanguageModel(*transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error)
	DescribeLanguageModelWithContext(aws.Context, *transcribeservice.DescribeLanguageModelInput, ...request.Option) (*transcribeservice.DescribeLanguageModelOutput, error)
	DescribeLanguageModelRequest(*transcribeservice.DescribeLanguageModelInput) (*request.Request, *transcribeservice.DescribeLanguageModelOutput)

	GetCallAnalyticsCategory(*transcribeservice.GetCallAnalyticsCategoryInput) (*transcribeservice.GetCallAnalyticsCategoryOutput, error)
	GetCallAnalyticsCategoryWithContext(aws.Context, *transcribeservice.GetCallAnalyticsCategoryInput, ...request.Option) (*transcribeservice.GetCallAnalyticsCategoryOutput, error)
	GetCallAnalyticsCategoryRequest(*transcribeservice.GetCallAnalyticsCategoryInput) (*request.Request, *transcribeservice.GetCallAnalyticsCategoryOutput)

	GetCallAnalyticsJob(*transcribeservice.GetCallAnalyticsJobInput) (*transcribeservice.GetCallAnalyticsJobOutput, error)
	GetCallAnalyticsJobWithContext(aws.Context, *transcribeservice.GetCallAnalyticsJobInput, ...request.Option) (*transcribeservice.GetCallAnalyticsJobOutput, error)
	GetCallAnalyticsJobRequest(*transcribeservice.GetCallAnalyticsJobInput) (*request.Request, *transcribeservice.GetCallAnalyticsJobOutput)

	GetMedicalTranscriptionJob(*transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error)
	GetMedicalTranscriptionJobWithContext(aws.Context, *transcribeservice.GetMedicalTranscriptionJobInput, ...request.Option) (*transcribeservice.GetMedicalTranscriptionJobOutput, error)
	GetMedicalTranscriptionJobRequest(*transcribeservice.GetMedicalTranscriptionJobInput) (*request.Request, *transcribeservice.GetMedicalTranscriptionJobOutput)

	GetMedicalVocabulary(*transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error)
	GetMedicalVocabularyWithContext(aws.Context, *transcribeservice.GetMedicalVocabularyInput, ...request.Option) (*transcribeservice.GetMedicalVocabularyOutput, error)
	GetMedicalVocabularyRequest(*transcribeservice.GetMedicalVocabularyInput) (*request.Request, *transcribeservice.GetMedicalVocabularyOutput)

	GetTranscriptionJob(*transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error)
	GetTranscriptionJobWithContext(aws.Context, *transcribeservice.GetTranscriptionJobInput, ...request.Option) (*transcribeservice.GetTranscriptionJobOutput, error)
	GetTranscriptionJobRequest(*transcribeservice.GetTranscriptionJobInput) (*request.Request, *transcribeservice.GetTranscriptionJobOutput)

	GetVocabulary(*transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error)
	GetVocabularyWithContext(aws.Context, *transcribeservice.GetVocabularyInput, ...request.Option) (*transcribeservice.GetVocabularyOutput, error)
	GetVocabularyRequest(*transcribeservice.GetVocabularyInput) (*request.Request, *transcribeservice.GetVocabularyOutput)

	GetVocabularyFilter(*transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error)
	GetVocabularyFilterWithContext(aws.Context, *transcribeservice.GetVocabularyFilterInput, ...request.Option) (*transcribeservice.GetVocabularyFilterOutput, error)
	GetVocabularyFilterRequest(*transcribeservice.GetVocabularyFilterInput) (*request.Request, *transcribeservice.GetVocabularyFilterOutput)

	ListCallAnalyticsCategories(*transcribeservice.ListCallAnalyticsCategoriesInput) (*transcribeservice.ListCallAnalyticsCategoriesOutput, error)
	ListCallAnalyticsCategoriesWithContext(aws.Context, *transcribeservice.ListCallAnalyticsCategoriesInput, ...request.Option) (*transcribeservice.ListCallAnalyticsCategoriesOutput, error)
	ListCallAnalyticsCategoriesRequest(*transcribeservice.ListCallAnalyticsCategoriesInput) (*request.Request, *transcribeservice.ListCallAnalyticsCategoriesOutput)

	ListCallAnalyticsCategoriesPages(*transcribeservice.ListCallAnalyticsCategoriesInput, func(*transcribeservice.ListCallAnalyticsCategoriesOutput, bool) bool) error
	ListCallAnalyticsCategoriesPagesWithContext(aws.Context, *transcribeservice.ListCallAnalyticsCategoriesInput, func(*transcribeservice.ListCallAnalyticsCategoriesOutput, bool) bool, ...request.Option) error

	ListCallAnalyticsJobs(*transcribeservice.ListCallAnalyticsJobsInput) (*transcribeservice.ListCallAnalyticsJobsOutput, error)
	ListCallAnalyticsJobsWithContext(aws.Context, *transcribeservice.ListCallAnalyticsJobsInput, ...request.Option) (*transcribeservice.ListCallAnalyticsJobsOutput, error)
	ListCallAnalyticsJobsRequest(*transcribeservice.ListCallAnalyticsJobsInput) (*request.Request, *transcribeservice.ListCallAnalyticsJobsOutput)

	ListCallAnalyticsJobsPages(*transcribeservice.ListCallAnalyticsJobsInput, func(*transcribeservice.ListCallAnalyticsJobsOutput, bool) bool) error
	ListCallAnalyticsJobsPagesWithContext(aws.Context, *transcribeservice.ListCallAnalyticsJobsInput, func(*transcribeservice.ListCallAnalyticsJobsOutput, bool) bool, ...request.Option) error

	ListLanguageModels(*transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error)
	ListLanguageModelsWithContext(aws.Context, *transcribeservice.ListLanguageModelsInput, ...request.Option) (*transcribeservice.ListLanguageModelsOutput, error)
	ListLanguageModelsRequest(*transcribeservice.ListLanguageModelsInput) (*request.Request, *transcribeservice.ListLanguageModelsOutput)

	ListLanguageModelsPages(*transcribeservice.ListLanguageModelsInput, func(*transcribeservice.ListLanguageModelsOutput, bool) bool) error
	ListLanguageModelsPagesWithContext(aws.Context, *transcribeservice.ListLanguageModelsInput, func(*transcribeservice.ListLanguageModelsOutput, bool) bool, ...request.Option) error

	ListMedicalTranscriptionJobs(*transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error)
	ListMedicalTranscriptionJobsWithContext(aws.Context, *transcribeservice.ListMedicalTranscriptionJobsInput, ...request.Option) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error)
	ListMedicalTranscriptionJobsRequest(*transcribeservice.ListMedicalTranscriptionJobsInput) (*request.Request, *transcribeservice.ListMedicalTranscriptionJobsOutput)

	ListMedicalTranscriptionJobsPages(*transcribeservice.ListMedicalTranscriptionJobsInput, func(*transcribeservice.ListMedicalTranscriptionJobsOutput, bool) bool) error
	ListMedicalTranscriptionJobsPagesWithContext(aws.Context, *transcribeservice.ListMedicalTranscriptionJobsInput, func(*transcribeservice.ListMedicalTranscriptionJobsOutput, bool) bool, ...request.Option) error

	ListMedicalVocabularies(*transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error)
	ListMedicalVocabulariesWithContext(aws.Context, *transcribeservice.ListMedicalVocabulariesInput, ...request.Option) (*transcribeservice.ListMedicalVocabulariesOutput, error)
	ListMedicalVocabulariesRequest(*transcribeservice.ListMedicalVocabulariesInput) (*request.Request, *transcribeservice.ListMedicalVocabulariesOutput)

	ListMedicalVocabulariesPages(*transcribeservice.ListMedicalVocabulariesInput, func(*transcribeservice.ListMedicalVocabulariesOutput, bool) bool) error
	ListMedicalVocabulariesPagesWithContext(aws.Context, *transcribeservice.ListMedicalVocabulariesInput, func(*transcribeservice.ListMedicalVocabulariesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*transcribeservice.ListTagsForResourceInput) (*transcribeservice.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *transcribeservice.ListTagsForResourceInput, ...request.Option) (*transcribeservice.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*transcribeservice.ListTagsForResourceInput) (*request.Request, *transcribeservice.ListTagsForResourceOutput)

	ListTranscriptionJobs(*transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error)
	ListTranscriptionJobsWithContext(aws.Context, *transcribeservice.ListTranscriptionJobsInput, ...request.Option) (*transcribeservice.ListTranscriptionJobsOutput, error)
	ListTranscriptionJobsRequest(*transcribeservice.ListTranscriptionJobsInput) (*request.Request, *transcribeservice.ListTranscriptionJobsOutput)

	ListTranscriptionJobsPages(*transcribeservice.ListTranscriptionJobsInput, func(*transcribeservice.ListTranscriptionJobsOutput, bool) bool) error
	ListTranscriptionJobsPagesWithContext(aws.Context, *transcribeservice.ListTranscriptionJobsInput, func(*transcribeservice.ListTranscriptionJobsOutput, bool) bool, ...request.Option) error

	ListVocabularies(*transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error)
	ListVocabulariesWithContext(aws.Context, *transcribeservice.ListVocabulariesInput, ...request.Option) (*transcribeservice.ListVocabulariesOutput, error)
	ListVocabulariesRequest(*transcribeservice.ListVocabulariesInput) (*request.Request, *transcribeservice.ListVocabulariesOutput)

	ListVocabulariesPages(*transcribeservice.ListVocabulariesInput, func(*transcribeservice.ListVocabulariesOutput, bool) bool) error
	ListVocabulariesPagesWithContext(aws.Context, *transcribeservice.ListVocabulariesInput, func(*transcribeservice.ListVocabulariesOutput, bool) bool, ...request.Option) error

	ListVocabularyFilters(*transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error)
	ListVocabularyFiltersWithContext(aws.Context, *transcribeservice.ListVocabularyFiltersInput, ...request.Option) (*transcribeservice.ListVocabularyFiltersOutput, error)
	ListVocabularyFiltersRequest(*transcribeservice.ListVocabularyFiltersInput) (*request.Request, *transcribeservice.ListVocabularyFiltersOutput)

	ListVocabularyFiltersPages(*transcribeservice.ListVocabularyFiltersInput, func(*transcribeservice.ListVocabularyFiltersOutput, bool) bool) error
	ListVocabularyFiltersPagesWithContext(aws.Context, *transcribeservice.ListVocabularyFiltersInput, func(*transcribeservice.ListVocabularyFiltersOutput, bool) bool, ...request.Option) error

	StartCallAnalyticsJob(*transcribeservice.StartCallAnalyticsJobInput) (*transcribeservice.StartCallAnalyticsJobOutput, error)
	StartCallAnalyticsJobWithContext(aws.Context, *transcribeservice.StartCallAnalyticsJobInput, ...request.Option) (*transcribeservice.StartCallAnalyticsJobOutput, error)
	StartCallAnalyticsJobRequest(*transcribeservice.StartCallAnalyticsJobInput) (*request.Request, *transcribeservice.StartCallAnalyticsJobOutput)

	StartMedicalTranscriptionJob(*transcribeservice.StartMedicalTranscriptionJobInput) (*transcribeservice.StartMedicalTranscriptionJobOutput, error)
	StartMedicalTranscriptionJobWithContext(aws.Context, *transcribeservice.StartMedicalTranscriptionJobInput, ...request.Option) (*transcribeservice.StartMedicalTranscriptionJobOutput, error)
	StartMedicalTranscriptionJobRequest(*transcribeservice.StartMedicalTranscriptionJobInput) (*request.Request, *transcribeservice.StartMedicalTranscriptionJobOutput)

	StartTranscriptionJob(*transcribeservice.StartTranscriptionJobInput) (*transcribeservice.StartTranscriptionJobOutput, error)
	StartTranscriptionJobWithContext(aws.Context, *transcribeservice.StartTranscriptionJobInput, ...request.Option) (*transcribeservice.StartTranscriptionJobOutput, error)
	StartTranscriptionJobRequest(*transcribeservice.StartTranscriptionJobInput) (*request.Request, *transcribeservice.StartTranscriptionJobOutput)

	TagResource(*transcribeservice.TagResourceInput) (*transcribeservice.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *transcribeservice.TagResourceInput, ...request.Option) (*transcribeservice.TagResourceOutput, error)
	TagResourceRequest(*transcribeservice.TagResourceInput) (*request.Request, *transcribeservice.TagResourceOutput)

	UntagResource(*transcribeservice.UntagResourceInput) (*transcribeservice.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *transcribeservice.UntagResourceInput, ...request.Option) (*transcribeservice.UntagResourceOutput, error)
	UntagResourceRequest(*transcribeservice.UntagResourceInput) (*request.Request, *transcribeservice.UntagResourceOutput)

	UpdateCallAnalyticsCategory(*transcribeservice.UpdateCallAnalyticsCategoryInput) (*transcribeservice.UpdateCallAnalyticsCategoryOutput, error)
	UpdateCallAnalyticsCategoryWithContext(aws.Context, *transcribeservice.UpdateCallAnalyticsCategoryInput, ...request.Option) (*transcribeservice.UpdateCallAnalyticsCategoryOutput, error)
	UpdateCallAnalyticsCategoryRequest(*transcribeservice.UpdateCallAnalyticsCategoryInput) (*request.Request, *transcribeservice.UpdateCallAnalyticsCategoryOutput)

	UpdateMedicalVocabulary(*transcribeservice.UpdateMedicalVocabularyInput) (*transcribeservice.UpdateMedicalVocabularyOutput, error)
	UpdateMedicalVocabularyWithContext(aws.Context, *transcribeservice.UpdateMedicalVocabularyInput, ...request.Option) (*transcribeservice.UpdateMedicalVocabularyOutput, error)
	UpdateMedicalVocabularyRequest(*transcribeservice.UpdateMedicalVocabularyInput) (*request.Request, *transcribeservice.UpdateMedicalVocabularyOutput)

	UpdateVocabulary(*transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error)
	UpdateVocabularyWithContext(aws.Context, *transcribeservice.UpdateVocabularyInput, ...request.Option) (*transcribeservice.UpdateVocabularyOutput, error)
	UpdateVocabularyRequest(*transcribeservice.UpdateVocabularyInput) (*request.Request, *transcribeservice.UpdateVocabularyOutput)

	UpdateVocabularyFilter(*transcribeservice.UpdateVocabularyFilterInput) (*transcribeservice.UpdateVocabularyFilterOutput, error)
	UpdateVocabularyFilterWithContext(aws.Context, *transcribeservice.UpdateVocabularyFilterInput, ...request.Option) (*transcribeservice.UpdateVocabularyFilterOutput, error)
	UpdateVocabularyFilterRequest(*transcribeservice.UpdateVocabularyFilterInput) (*request.Request, *transcribeservice.UpdateVocabularyFilterOutput)
}

var _ TranscribeServiceAPI = (*transcribeservice.TranscribeService)(nil)

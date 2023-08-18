// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elasticsearchservice provides the client and types for making API
// requests to Amazon Elasticsearch Service.
//
// Use the Amazon Elasticsearch Configuration API to create, configure, and
// manage Elasticsearch domains.
//
// For sample code that uses the Configuration API, see the Amazon Elasticsearch
// Service Developer Guide (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-configuration-samples.html).
// The guide also contains sample code for sending signed HTTP requests to the
// Elasticsearch APIs (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-request-signing.html).
//
// The endpoint for configuration service requests is region-specific: es.region.amazonaws.com.
// For example, es.us-east-1.amazonaws.com. For a current list of supported
// regions and endpoints, see Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticsearch-service-regions).
//
// See elasticsearchservice package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/elasticsearchservice/
//
// # Using the Client
//
// To contact Amazon Elasticsearch Service with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Elasticsearch Service client ElasticsearchService for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/elasticsearchservice/#New
package elasticsearchservice

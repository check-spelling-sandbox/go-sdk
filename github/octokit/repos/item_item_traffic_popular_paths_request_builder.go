package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemTrafficPopularPathsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\traffic\popular\paths
type ItemItemTrafficPopularPathsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemTrafficPopularPathsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemTrafficPopularPathsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemTrafficPopularPathsRequestBuilderInternal instantiates a new PathsRequestBuilder and sets the default values.
func NewItemItemTrafficPopularPathsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficPopularPathsRequestBuilder) {
    m := &ItemItemTrafficPopularPathsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/traffic/popular/paths", pathParameters),
    }
    return m
}
// NewItemItemTrafficPopularPathsRequestBuilder instantiates a new PathsRequestBuilder and sets the default values.
func NewItemItemTrafficPopularPathsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficPopularPathsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTrafficPopularPathsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the top 10 popular contents over the last 14 days.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/traffic#get-top-referral-paths
func (m *ItemItemTrafficPopularPathsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemTrafficPopularPathsRequestBuilderGetRequestConfiguration)([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentTrafficable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateContentTrafficFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentTrafficable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentTrafficable)
        }
    }
    return val, nil
}
// ToGetRequestInformation get the top 10 popular contents over the last 14 days.
func (m *ItemItemTrafficPopularPathsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemTrafficPopularPathsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemTrafficPopularPathsRequestBuilder) WithUrl(rawUrl string)(*ItemItemTrafficPopularPathsRequestBuilder) {
    return NewItemItemTrafficPopularPathsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
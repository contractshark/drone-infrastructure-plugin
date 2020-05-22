package main

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v28/github"
)

// MockGithubRepositoryClient is a mock of GithubRepositoryClient interface
type MockGithubRepositoryClient struct {
	ctrl     *gomock.Controller
	recorder *MockGithubRepositoryClientMockRecorder
}

// MockGithubRepositoryClientMockRecorder is the mock recorder for MockGithubRepositoryClient
type MockGithubRepositoryClientMockRecorder struct {
	mock *MockGithubRepositoryClient
}

// NewMockGithubRepositoryClient creates a new mock instance
func NewMockGithubRepositoryClient(ctrl *gomock.Controller) *MockGithubRepositoryClient {
	mock := &MockGithubRepositoryClient{ctrl: ctrl}
	mock.recorder = &MockGithubRepositoryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGithubRepositoryClient) EXPECT() *MockGithubRepositoryClientMockRecorder {
	return m.recorder
}

// GetCommit mocks base method
func (m *MockGithubRepositoryClient) GetCommit(ctx context.Context, owner, repo, sha string) (*github.RepositoryCommit, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommit", ctx, owner, repo, sha)
	ret0, _ := ret[0].(*github.RepositoryCommit)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCommit indicates an expected call of GetCommit
func (mr *MockGithubRepositoryClientMockRecorder) GetCommit(ctx, owner, repo, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommit", reflect.TypeOf((*MockGithubRepositoryClient)(nil).GetCommit), ctx, owner, repo, sha)
}

// CompareCommits mocks base method
func (m *MockGithubRepositoryClient) CompareCommits(ctx context.Context, owner, repo, base, head string) (*github.CommitsComparison, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareCommits", ctx, owner, repo, base, head)
	ret0, _ := ret[0].(*github.CommitsComparison)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CompareCommits indicates an expected call of CompareCommits
func (mr *MockGithubRepositoryClientMockRecorder) CompareCommits(ctx, owner, repo, base, head interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareCommits", reflect.TypeOf((*MockGithubRepositoryClient)(nil).CompareCommits), ctx, owner, repo, base, head)
}

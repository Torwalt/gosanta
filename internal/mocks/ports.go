// Code generated by MockGen. DO NOT EDIT.
// Source: ./ports.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	awards "gosanta/internal"
	events "gosanta/pkg"

	gomock "github.com/golang/mock/gomock"
)

// MockAwardReadRepository is a mock of AwardReadRepository interface.
type MockAwardReadRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAwardReadRepositoryMockRecorder
}

// MockAwardReadRepositoryMockRecorder is the mock recorder for MockAwardReadRepository.
type MockAwardReadRepositoryMockRecorder struct {
	mock *MockAwardReadRepository
}

// NewMockAwardReadRepository creates a new mock instance.
func NewMockAwardReadRepository(ctrl *gomock.Controller) *MockAwardReadRepository {
	mock := &MockAwardReadRepository{ctrl: ctrl}
	mock.recorder = &MockAwardReadRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwardReadRepository) EXPECT() *MockAwardReadRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockAwardReadRepository) Get(id int64) (*awards.PhishingAward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*awards.PhishingAward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAwardReadRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAwardReadRepository)(nil).Get), id)
}

// GetByEmailRef mocks base method.
func (m *MockAwardReadRepository) GetByEmailRef(id awards.UserId, ref string) (*awards.PhishingAward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmailRef", id, ref)
	ret0, _ := ret[0].(*awards.PhishingAward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmailRef indicates an expected call of GetByEmailRef.
func (mr *MockAwardReadRepositoryMockRecorder) GetByEmailRef(id, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmailRef", reflect.TypeOf((*MockAwardReadRepository)(nil).GetByEmailRef), id, ref)
}

// GetUserAwards mocks base method.
func (m *MockAwardReadRepository) GetUserAwards(id awards.UserId) ([]awards.PhishingAward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAwards", id)
	ret0, _ := ret[0].([]awards.PhishingAward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAwards indicates an expected call of GetUserAwards.
func (mr *MockAwardReadRepositoryMockRecorder) GetUserAwards(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAwards", reflect.TypeOf((*MockAwardReadRepository)(nil).GetUserAwards), id)
}

// MockAwardRepository is a mock of AwardRepository interface.
type MockAwardRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAwardRepositoryMockRecorder
}

// MockAwardRepositoryMockRecorder is the mock recorder for MockAwardRepository.
type MockAwardRepositoryMockRecorder struct {
	mock *MockAwardRepository
}

// NewMockAwardRepository creates a new mock instance.
func NewMockAwardRepository(ctrl *gomock.Controller) *MockAwardRepository {
	mock := &MockAwardRepository{ctrl: ctrl}
	mock.recorder = &MockAwardRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwardRepository) EXPECT() *MockAwardRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockAwardRepository) Add(award *awards.PhishingAward) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", award)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockAwardRepositoryMockRecorder) Add(award interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAwardRepository)(nil).Add), award)
}

// Delete mocks base method.
func (m *MockAwardRepository) Delete(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAwardRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAwardRepository)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockAwardRepository) Get(id int64) (*awards.PhishingAward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*awards.PhishingAward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAwardRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAwardRepository)(nil).Get), id)
}

// GetByEmailRef mocks base method.
func (m *MockAwardRepository) GetByEmailRef(id awards.UserId, ref string) (*awards.PhishingAward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmailRef", id, ref)
	ret0, _ := ret[0].(*awards.PhishingAward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmailRef indicates an expected call of GetByEmailRef.
func (mr *MockAwardRepositoryMockRecorder) GetByEmailRef(id, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmailRef", reflect.TypeOf((*MockAwardRepository)(nil).GetByEmailRef), id, ref)
}

// GetUserAwards mocks base method.
func (m *MockAwardRepository) GetUserAwards(id awards.UserId) ([]awards.PhishingAward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAwards", id)
	ret0, _ := ret[0].([]awards.PhishingAward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAwards indicates an expected call of GetUserAwards.
func (mr *MockAwardRepositoryMockRecorder) GetUserAwards(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAwards", reflect.TypeOf((*MockAwardRepository)(nil).GetUserAwards), id)
}

// UpdateExisting mocks base method.
func (m *MockAwardRepository) UpdateExisting(existing, award *awards.PhishingAward) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExisting", existing, award)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExisting indicates an expected call of UpdateExisting.
func (mr *MockAwardRepositoryMockRecorder) UpdateExisting(existing, award interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExisting", reflect.TypeOf((*MockAwardRepository)(nil).UpdateExisting), existing, award)
}

// MockUserReadRepository is a mock of UserReadRepository interface.
type MockUserReadRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserReadRepositoryMockRecorder
}

// MockUserReadRepositoryMockRecorder is the mock recorder for MockUserReadRepository.
type MockUserReadRepositoryMockRecorder struct {
	mock *MockUserReadRepository
}

// NewMockUserReadRepository creates a new mock instance.
func NewMockUserReadRepository(ctrl *gomock.Controller) *MockUserReadRepository {
	mock := &MockUserReadRepository{ctrl: ctrl}
	mock.recorder = &MockUserReadRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserReadRepository) EXPECT() *MockUserReadRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUserReadRepository) Get(uId awards.UserId) (*awards.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", uId)
	ret0, _ := ret[0].(*awards.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserReadRepositoryMockRecorder) Get(uId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserReadRepository)(nil).Get), uId)
}

// GetCompanyUsers mocks base method.
func (m *MockUserReadRepository) GetCompanyUsers(cId awards.CompanyId) ([]awards.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyUsers", cId)
	ret0, _ := ret[0].([]awards.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyUsers indicates an expected call of GetCompanyUsers.
func (mr *MockUserReadRepositoryMockRecorder) GetCompanyUsers(cId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyUsers", reflect.TypeOf((*MockUserReadRepository)(nil).GetCompanyUsers), cId)
}

// MockAwardReadingService is a mock of AwardReadingService interface.
type MockAwardReadingService struct {
	ctrl     *gomock.Controller
	recorder *MockAwardReadingServiceMockRecorder
}

// MockAwardReadingServiceMockRecorder is the mock recorder for MockAwardReadingService.
type MockAwardReadingServiceMockRecorder struct {
	mock *MockAwardReadingService
}

// NewMockAwardReadingService creates a new mock instance.
func NewMockAwardReadingService(ctrl *gomock.Controller) *MockAwardReadingService {
	mock := &MockAwardReadingService{ctrl: ctrl}
	mock.recorder = &MockAwardReadingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwardReadingService) EXPECT() *MockAwardReadingServiceMockRecorder {
	return m.recorder
}

// CalcLeaderboard mocks base method.
func (m *MockAwardReadingService) CalcLeaderboard(uId awards.UserId) (*awards.Leaderboard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcLeaderboard", uId)
	ret0, _ := ret[0].(*awards.Leaderboard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalcLeaderboard indicates an expected call of CalcLeaderboard.
func (mr *MockAwardReadingServiceMockRecorder) CalcLeaderboard(uId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcLeaderboard", reflect.TypeOf((*MockAwardReadingService)(nil).CalcLeaderboard), uId)
}

// Get mocks base method.
func (m *MockAwardReadingService) Get(id int64) (*awards.PhishingAward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*awards.PhishingAward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAwardReadingServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAwardReadingService)(nil).Get), id)
}

// GetUserAwards mocks base method.
func (m *MockAwardReadingService) GetUserAwards(uId awards.UserId) ([]awards.PhishingAward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAwards", uId)
	ret0, _ := ret[0].([]awards.PhishingAward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAwards indicates an expected call of GetUserAwards.
func (mr *MockAwardReadingServiceMockRecorder) GetUserAwards(uId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAwards", reflect.TypeOf((*MockAwardReadingService)(nil).GetUserAwards), uId)
}

// MockEventReadRepository is a mock of EventReadRepository interface.
type MockEventReadRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEventReadRepositoryMockRecorder
}

// MockEventReadRepositoryMockRecorder is the mock recorder for MockEventReadRepository.
type MockEventReadRepositoryMockRecorder struct {
	mock *MockEventReadRepository
}

// NewMockEventReadRepository creates a new mock instance.
func NewMockEventReadRepository(ctrl *gomock.Controller) *MockEventReadRepository {
	mock := &MockEventReadRepository{ctrl: ctrl}
	mock.recorder = &MockEventReadRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventReadRepository) EXPECT() *MockEventReadRepositoryMockRecorder {
	return m.recorder
}

// ClickedExists mocks base method.
func (m *MockEventReadRepository) ClickedExists(uID awards.UserId, emailRef string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClickedExists", uID, emailRef)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClickedExists indicates an expected call of ClickedExists.
func (mr *MockEventReadRepositoryMockRecorder) ClickedExists(uID, emailRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClickedExists", reflect.TypeOf((*MockEventReadRepository)(nil).ClickedExists), uID, emailRef)
}

// GetUnprocessed mocks base method.
func (m *MockEventReadRepository) GetUnprocessed() ([]awards.UserPhishingEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnprocessed")
	ret0, _ := ret[0].([]awards.UserPhishingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnprocessed indicates an expected call of GetUnprocessed.
func (mr *MockEventReadRepositoryMockRecorder) GetUnprocessed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnprocessed", reflect.TypeOf((*MockEventReadRepository)(nil).GetUnprocessed))
}

// MockEventRepositoryProcessor is a mock of EventRepositoryProcessor interface.
type MockEventRepositoryProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepositoryProcessorMockRecorder
}

// MockEventRepositoryProcessorMockRecorder is the mock recorder for MockEventRepositoryProcessor.
type MockEventRepositoryProcessorMockRecorder struct {
	mock *MockEventRepositoryProcessor
}

// NewMockEventRepositoryProcessor creates a new mock instance.
func NewMockEventRepositoryProcessor(ctrl *gomock.Controller) *MockEventRepositoryProcessor {
	mock := &MockEventRepositoryProcessor{ctrl: ctrl}
	mock.recorder = &MockEventRepositoryProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventRepositoryProcessor) EXPECT() *MockEventRepositoryProcessorMockRecorder {
	return m.recorder
}

// ClickedExists mocks base method.
func (m *MockEventRepositoryProcessor) ClickedExists(uID awards.UserId, emailRef string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClickedExists", uID, emailRef)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClickedExists indicates an expected call of ClickedExists.
func (mr *MockEventRepositoryProcessorMockRecorder) ClickedExists(uID, emailRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClickedExists", reflect.TypeOf((*MockEventRepositoryProcessor)(nil).ClickedExists), uID, emailRef)
}

// GetUnprocessed mocks base method.
func (m *MockEventRepositoryProcessor) GetUnprocessed() ([]awards.UserPhishingEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnprocessed")
	ret0, _ := ret[0].([]awards.UserPhishingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnprocessed indicates an expected call of GetUnprocessed.
func (mr *MockEventRepositoryProcessorMockRecorder) GetUnprocessed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnprocessed", reflect.TypeOf((*MockEventRepositoryProcessor)(nil).GetUnprocessed))
}

// MarkAsProcessed mocks base method.
func (m *MockEventRepositoryProcessor) MarkAsProcessed(event *awards.UserPhishingEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkAsProcessed", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkAsProcessed indicates an expected call of MarkAsProcessed.
func (mr *MockEventRepositoryProcessorMockRecorder) MarkAsProcessed(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAsProcessed", reflect.TypeOf((*MockEventRepositoryProcessor)(nil).MarkAsProcessed), event)
}

// MockEventRepository is a mock of EventRepository interface.
type MockEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepositoryMockRecorder
}

// MockEventRepositoryMockRecorder is the mock recorder for MockEventRepository.
type MockEventRepositoryMockRecorder struct {
	mock *MockEventRepository
}

// NewMockEventRepository creates a new mock instance.
func NewMockEventRepository(ctrl *gomock.Controller) *MockEventRepository {
	mock := &MockEventRepository{ctrl: ctrl}
	mock.recorder = &MockEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventRepository) EXPECT() *MockEventRepositoryMockRecorder {
	return m.recorder
}

// ClickedExists mocks base method.
func (m *MockEventRepository) ClickedExists(uID awards.UserId, emailRef string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClickedExists", uID, emailRef)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClickedExists indicates an expected call of ClickedExists.
func (mr *MockEventRepositoryMockRecorder) ClickedExists(uID, emailRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClickedExists", reflect.TypeOf((*MockEventRepository)(nil).ClickedExists), uID, emailRef)
}

// GetUnprocessed mocks base method.
func (m *MockEventRepository) GetUnprocessed() ([]awards.UserPhishingEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnprocessed")
	ret0, _ := ret[0].([]awards.UserPhishingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnprocessed indicates an expected call of GetUnprocessed.
func (mr *MockEventRepositoryMockRecorder) GetUnprocessed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnprocessed", reflect.TypeOf((*MockEventRepository)(nil).GetUnprocessed))
}

// Write mocks base method.
func (m *MockEventRepository) Write(upe awards.UserPhishingEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", upe)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockEventRepositoryMockRecorder) Write(upe interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockEventRepository)(nil).Write), upe)
}

// MockEventQueue is a mock of EventQueue interface.
type MockEventQueue struct {
	ctrl     *gomock.Controller
	recorder *MockEventQueueMockRecorder
}

// MockEventQueueMockRecorder is the mock recorder for MockEventQueue.
type MockEventQueueMockRecorder struct {
	mock *MockEventQueue
}

// NewMockEventQueue creates a new mock instance.
func NewMockEventQueue(ctrl *gomock.Controller) *MockEventQueue {
	mock := &MockEventQueue{ctrl: ctrl}
	mock.recorder = &MockEventQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventQueue) EXPECT() *MockEventQueueMockRecorder {
	return m.recorder
}

// DeleteMessage mocks base method.
func (m *MockEventQueue) DeleteMessage(eventID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", eventID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockEventQueueMockRecorder) DeleteMessage(eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockEventQueue)(nil).DeleteMessage), eventID)
}

// GetNextMessages mocks base method.
func (m *MockEventQueue) GetNextMessages() ([]events.PhishingEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextMessages")
	ret0, _ := ret[0].([]events.PhishingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextMessages indicates an expected call of GetNextMessages.
func (mr *MockEventQueueMockRecorder) GetNextMessages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextMessages", reflect.TypeOf((*MockEventQueue)(nil).GetNextMessages))
}

// MockMailSender is a mock of MailSender interface.
type MockMailSender struct {
	ctrl     *gomock.Controller
	recorder *MockMailSenderMockRecorder
}

// MockMailSenderMockRecorder is the mock recorder for MockMailSender.
type MockMailSenderMockRecorder struct {
	mock *MockMailSender
}

// NewMockMailSender creates a new mock instance.
func NewMockMailSender(ctrl *gomock.Controller) *MockMailSender {
	mock := &MockMailSender{ctrl: ctrl}
	mock.recorder = &MockMailSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailSender) EXPECT() *MockMailSenderMockRecorder {
	return m.recorder
}

// SendToUser mocks base method.
func (m *MockMailSender) SendToUser(arg0 awards.UserAwardEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendToUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendToUser indicates an expected call of SendToUser.
func (mr *MockMailSenderMockRecorder) SendToUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendToUser", reflect.TypeOf((*MockMailSender)(nil).SendToUser), arg0)
}

// MockEventPublisher is a mock of EventPublisher interface.
type MockEventPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockEventPublisherMockRecorder
}

// MockEventPublisherMockRecorder is the mock recorder for MockEventPublisher.
type MockEventPublisherMockRecorder struct {
	mock *MockEventPublisher
}

// NewMockEventPublisher creates a new mock instance.
func NewMockEventPublisher(ctrl *gomock.Controller) *MockEventPublisher {
	mock := &MockEventPublisher{ctrl: ctrl}
	mock.recorder = &MockEventPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventPublisher) EXPECT() *MockEventPublisherMockRecorder {
	return m.recorder
}

// PublishEvent mocks base method.
func (m *MockEventPublisher) PublishEvent(arg0 awards.UserAwardEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishEvent indicates an expected call of PublishEvent.
func (mr *MockEventPublisherMockRecorder) PublishEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishEvent", reflect.TypeOf((*MockEventPublisher)(nil).PublishEvent), arg0)
}

// MockEventLogReader is a mock of EventLogReader interface.
type MockEventLogReader struct {
	ctrl     *gomock.Controller
	recorder *MockEventLogReaderMockRecorder
}

// MockEventLogReaderMockRecorder is the mock recorder for MockEventLogReader.
type MockEventLogReaderMockRecorder struct {
	mock *MockEventLogReader
}

// NewMockEventLogReader creates a new mock instance.
func NewMockEventLogReader(ctrl *gomock.Controller) *MockEventLogReader {
	mock := &MockEventLogReader{ctrl: ctrl}
	mock.recorder = &MockEventLogReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventLogReader) EXPECT() *MockEventLogReaderMockRecorder {
	return m.recorder
}

// GetUnprocessedEvents mocks base method.
func (m *MockEventLogReader) GetUnprocessedEvents() ([]awards.UserPhishingEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnprocessedEvents")
	ret0, _ := ret[0].([]awards.UserPhishingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnprocessedEvents indicates an expected call of GetUnprocessedEvents.
func (mr *MockEventLogReaderMockRecorder) GetUnprocessedEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnprocessedEvents", reflect.TypeOf((*MockEventLogReader)(nil).GetUnprocessedEvents))
}

// LogNewEvents mocks base method.
func (m *MockEventLogReader) LogNewEvents() ([]awards.UserPhishingEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogNewEvents")
	ret0, _ := ret[0].([]awards.UserPhishingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogNewEvents indicates an expected call of LogNewEvents.
func (mr *MockEventLogReaderMockRecorder) LogNewEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogNewEvents", reflect.TypeOf((*MockEventLogReader)(nil).LogNewEvents))
}

// MockAwardAssigner is a mock of AwardAssigner interface.
type MockAwardAssigner struct {
	ctrl     *gomock.Controller
	recorder *MockAwardAssignerMockRecorder
}

// MockAwardAssignerMockRecorder is the mock recorder for MockAwardAssigner.
type MockAwardAssignerMockRecorder struct {
	mock *MockAwardAssigner
}

// NewMockAwardAssigner creates a new mock instance.
func NewMockAwardAssigner(ctrl *gomock.Controller) *MockAwardAssigner {
	mock := &MockAwardAssigner{ctrl: ctrl}
	mock.recorder = &MockAwardAssignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwardAssigner) EXPECT() *MockAwardAssignerMockRecorder {
	return m.recorder
}

// AssignAward mocks base method.
func (m *MockAwardAssigner) AssignAward(event awards.UserPhishingEvent) (awards.UserAwardEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignAward", event)
	ret0, _ := ret[0].(awards.UserAwardEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignAward indicates an expected call of AssignAward.
func (mr *MockAwardAssignerMockRecorder) AssignAward(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignAward", reflect.TypeOf((*MockAwardAssigner)(nil).AssignAward), event)
}

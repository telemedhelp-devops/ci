package serverMethodsLib

import (
	"fmt"
	"sync"

	"gopkg.in/olahol/melody.v1"
)

type Log interface {
	Lock()
	Unlock()
	AddSession(session *melody.Session)
	RemoveSession(session *melody.Session)
	PushLine(line string)
	GetHistory() []string
	CloseSessions()
}

type log struct {
	mutex    *sync.Mutex
	sessions []*melody.Session
	History  []string
}

func newLog() *log {
	return &log{
		mutex: &sync.Mutex{},
	}
}

func (log *log) Lock() {
	log.mutex.Lock()
}
func (log *log) Unlock() {
	log.mutex.Unlock()
}

func (log *log) GetHistory() []string {
	return log.History
}

func (log *log) AddSession(session *melody.Session) {
	log.Lock()
	defer log.Unlock()

	log.sessions = append(log.sessions, session)

	for _, line := range log.GetHistory() {
		err := session.Write([]byte(line))
		if err != nil {
			log.RemoveSession(session)
			break
		}
	}
}

func (log *log) RemoveSession(session *melody.Session) {
	log.Lock()
	defer log.Unlock()
	for idx, s := range log.sessions {
		if s != session {
			continue
		}
		s.Close()
		newSessions := log.sessions[:idx]
		if idx < len(log.sessions)-1 {
			newSessions = append(newSessions, log.sessions[idx+1:]...)
		}
		log.sessions = newSessions
		break
	}
}

func (log *log) PushLine(line string) {
	log.Lock()
	defer log.Unlock()
	log.History = append(log.History, line)
	for idx, session := range log.sessions {
		err := session.Write([]byte(line))
		if err != nil {
			session.Close()
			newSessions := log.sessions[:idx-1]
			if idx < len(log.sessions)-1 {
				newSessions = append(newSessions, log.sessions[idx:]...)
			}
			log.sessions = newSessions
		}
	}
}
func (log *log) CloseSessions() {
	log.Lock()
	defer log.Unlock()
	for _, session := range log.sessions {
		session.Close()
	}
	log.sessions = []*melody.Session{}
}

type Logs interface {
	Get(int) Log
	Remove(int)
	PushLine(int, string)
	Lock()
	Unlock()
}

type logs struct {
	logs  map[int]Log
	mutex *sync.Mutex
}

func NewLogs() *logs {
	return &logs{
		logs:  map[int]Log{},
		mutex: &sync.Mutex{},
	}
}

func (logs *logs) Get(pipelineId int) Log {
	logs.Lock()
	defer logs.Unlock()

	if logs.logs[pipelineId] == nil {
		logs.logs[pipelineId] = newLog()
	}

	return logs.logs[pipelineId]
}

func (logs *logs) Remove(pipelineId int) {
	logs.Lock()
	defer logs.Unlock()

	log := logs.logs[pipelineId]
	logs.logs[pipelineId] = nil

	log.CloseSessions()
}

func (logs *logs) PushLine(pipelineId int, line string) {
	logs.Lock()
	defer logs.Unlock()
	log := logs.logs[pipelineId]
	if log == nil {
		panic(fmt.Errorf("logs.logs[%v] == nil", pipelineId))
	}
	log.PushLine(line)
}

func (logs *logs) Lock() {
	logs.mutex.Lock()
}

func (logs *logs) Unlock() {
	logs.mutex.Unlock()
}

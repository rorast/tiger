package main

import (
	"fmt"
	"log"
	"net/http"
	"tiger/deployserver/logic"
	"time"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	"github.com/lonng/nano/session"
	"strings"
)

type (
	Room struct {
		group *nano.Group
	}

	// RoomManager represents a component that contains a bundle of room
	RoomManager struct {
		component.Base
		timer *nano.Timer
		rooms map[int]*Room
	}

	// UserMessage represents a message that user sent
	UserMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	// NewUser message will be received when new user join room
	NewUser struct {
		Content string `json:"content"`
	}

	// AllMembers contains all members uid
	AllMembers struct {
		Members []int64 `json:"members"`
	}

	// JoinResponse represents the result of joining room
	JoinResponse struct {
		Code   int    `json:"code"`
		Result string `json:"result"`
	}

	stats struct {
		component.Base
		timer         *nano.Timer
		outboundBytes int
		inboundBytes  int
	}

	// 傳給前端的結果及下次轉盤的數字
	ResultResponse struct {
		Code       int64  `json:"code"`
		Scatter    int8
		Result     [3][5]logic.Symbol  `json:"result"`
		NextSymbol [5][25]logic.Symbol `json:"symbol"`
	}
)


func (stats *stats) outbound(s *session.Session, msg nano.Message) error {
	stats.outboundBytes += len(msg.Data)
	return nil
}

func (stats *stats) inbound(s *session.Session, msg nano.Message) error {
	stats.inboundBytes += len(msg.Data)
	return nil
}

func (stats *stats) AfterInit() {
	stats.timer = nano.NewTimer(time.Minute, func() {
		println("OutboundBytes", stats.outboundBytes)
		println("InboundBytes", stats.outboundBytes)
	})
}

const (
	testRoomID = 1
	roomIDKey  = "ROOM_ID"
)

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: map[int]*Room{},
	}
}

// AfterInit component lifetime callback
func (mgr *RoomManager) AfterInit() {
	session.Lifetime.OnClosed(func(s *session.Session) {
		if !s.HasKey(roomIDKey) {
			return
		}
		room := s.Value(roomIDKey).(*Room)
		room.group.Leave(s)
	})
	mgr.timer = nano.NewTimer(time.Minute, func() {
		for roomId, room := range mgr.rooms {
			println(fmt.Sprintf("UserCount: RoomID=%d, Time=%s, Count=%d",
				roomId, time.Now().String(), room.group.Count()))
		}
	})
}

// Spin
func (mgr *RoomManager) Spin(s *session.Session, msg []byte) error {

	fmt.Println(len(msg))

	room, found := mgr.rooms[testRoomID]
	if !found {
		room = &Room{
			group: nano.NewGroup(fmt.Sprintf("room-%d", testRoomID)),
		}
		mgr.rooms[testRoomID] = room
	}

	fakeUID := s.ID()
	s.Bind(fakeUID)
	s.Set(roomIDKey, room)
	s.Push("onSpin", &AllMembers{Members: room.group.Members()})
	room.group.Broadcast("onSpin", &NewUser{Content: fmt.Sprintf("New user: %d", s.ID())})
	// new user join group
	room.group.Add(s) // add session to group
	//return s.Response(&JoinResponse{Result: "success"})

	m := logic.New()
	_, _, err := m.Spin(88, 20)
	if err != nil {
		//Fatal(err)
		fmt.Println(err)
	}

	rrr := ResultResponse{}
	rrr.Code = logic.WinTotal
	rrr.Scatter = logic.Scatter

	//rrr.Code = make([]int, 5)
	//rrr.Result = make([]int,15)

	//rrr.NextSymbol = make([]int,125)
	//seed := time.Now().UnixNano()
	//r := rand.New(rand.NewSource(seed))
	//先隨機計算結果給前端 0 :=沒中，1 := 中彩金，2 :=freeSpin, 3 := bigWin, 4 := jackPort
	//var resultCode [5]int
	//for i :=0 ; i < len(resultCode) ; i ++ {
	//	rrr.Code[i] = r.Intn(5)+1
	//}

	// 15個滾輪，13個symbol 1-13
	//for i:=0 ; i < len(rrr.Result) ; i++ {
	//	rrr.Result[i] = r.Intn(13)+1
	//}

	rrr.Result = logic.Symbols
	// reel為25個symbol， 有 5 個 reel 帶
	//for i:=0 ; i < len(rrr.NextSymbol) ; i++ {
	//	rrr.NextSymbol[i] = r.Intn(13)+1
	//}
	rrr.NextSymbol = logic.ReelStrips


	return s.Response(&rrr)
}

//func GetNumber(num int) json {
//	seed := time.Now().UnixNano()
//	r := rand.New(rand.NewSource(seed))
//	var resultCode [5]int
//	for i :=0 ; i < len(resultCode) ; i ++ {
//		resultCode[i] = r.Intn(num)+1
//	}
//	lang, err := json.Serializer{resultCode}
//	if err == nil {
//		return lang
//	}
//
//	//array 到 json str
//	//arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
//	//lang, err := json.Marshal(arr)
//	//if err == nil {
//	//	fmt.Println("================array 到 json str==")
//	//	fmt.Println(string(lang))
//	//}
//}


// Join room
func (mgr *RoomManager) Join(s *session.Session, msg []byte) error {

	// NOTE: join test room only in demo
	room, found := mgr.rooms[testRoomID]
	if !found {
		room = &Room{
			group: nano.NewGroup(fmt.Sprintf("room-%d", testRoomID)),
		}
		mgr.rooms[testRoomID] = room
	}

	fakeUID := s.ID() //just use s.ID as uid !!!
	s.Bind(fakeUID)   // binding session uids.Set(roomIDKey, room)
	s.Set(roomIDKey, room)
	s.Push("onMembers", &AllMembers{Members: room.group.Members()})
	// notify others
	room.group.Broadcast("onNewUser", &NewUser{Content: fmt.Sprintf("New user: %d", s.ID())})
	// new user join group
	room.group.Add(s) // add session to group
	//return s.Response(&JoinResponse{Result: "success"})
	return s.Response(&JoinResponse{Result: "{1,2,3,4,5,6,7,8,9,10}" +
		"{2,4,6,8,10,12,14,16,18,20}" +
		"{32,34,36,38,30,32,34,36,38,23}"})
}

// Message sync last message to all members
func (mgr *RoomManager) Message(s *session.Session, msg *UserMessage) error {
	if !s.HasKey(roomIDKey) {
		return fmt.Errorf("not join room yet")
	}
	room := s.Value(roomIDKey).(*Room)
	return room.group.Broadcast("onMessage", msg)
}

// 開獎 vatility L1 ~ L5
//func (rrr *ResultResponse) GetPrize(rr *ResultResponse) string {
//	seed := time.Now().UnixNano()
//	r := rand.New(rand.NewSource(seed))
//	// 先隨機計算結果給前端 0 :=沒中，1 := 中彩金，2 :=freeSpin, 3 := bigWin, 4 := jackPort
//	var resultCode [5]int
//	for i :=0 ; i < len(resultCode) ; i ++ {
//		resultCode[i] = r.Intn(5)+1
//	}
//
//	var prize [5][3]int
//	// 15個滾輪，13個symbol 1-13
//	for i:=0 ; i < 5 ; i++ {
//		for j:=0 ; j < 3 ; j++ {
//			prize[i][j] = r.Intn(13)+1
//		}
//		//prize[i] = r.Intn(33)+1
//	}
//	rr.Result = json(prize)
//	// reel為25個symbol， 有 5 個 reel 帶
//	var nextPrize [5][25]int
//	for i:=0 ; i < 5 ; i++ {
//		for j:=0 ; j < 25 ; j++ {
//			nextPrize[i][j] = r.Intn(13)+1
//		}
//		//prize[i] = r.Intn(33)+1
//	}
//
//	return fmt.Sprintf("code = : %v",resultCode,"result = : %v", prize,"symbols = : %v", nextPrize)
//}


func main() {
	// override default serializer
	nano.SetSerializer(json.NewSerializer())

	// rewrite component and handler name
	room := NewRoomManager()
	nano.Register(room,
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower),
	)

	// traffic stats
	pipeline := nano.NewPipeline()
	var stats = &stats{}
	pipeline.Outbound().PushBack(stats.outbound)
	pipeline.Inbound().PushBack(stats.inbound)

	nano.EnableDebug()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	nano.SetWSPath("/")

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))


	nano.SetCheckOriginFunc(func(_ *http.Request) bool { return true })
	nano.ListenWS(":3250", nano.WithPipeline(pipeline))



}
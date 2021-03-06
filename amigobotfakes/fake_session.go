// Code generated by counterfeiter. DO NOT EDIT.
package amigobotfakes

import (
	sync "sync"

	discordgo "github.com/bwmarrin/discordgo"
	amigobot "github.com/ryanmiville/amigobot"
)

type FakeSession struct {
	ChannelMessageDeleteStub        func(string, string) error
	channelMessageDeleteMutex       sync.RWMutex
	channelMessageDeleteArgsForCall []struct {
		arg1 string
		arg2 string
	}
	channelMessageDeleteReturns struct {
		result1 error
	}
	channelMessageDeleteReturnsOnCall map[int]struct {
		result1 error
	}
	ChannelMessagePinStub        func(string, string) error
	channelMessagePinMutex       sync.RWMutex
	channelMessagePinArgsForCall []struct {
		arg1 string
		arg2 string
	}
	channelMessagePinReturns struct {
		result1 error
	}
	channelMessagePinReturnsOnCall map[int]struct {
		result1 error
	}
	ChannelMessageSendStub        func(string, string) (*discordgo.Message, error)
	channelMessageSendMutex       sync.RWMutex
	channelMessageSendArgsForCall []struct {
		arg1 string
		arg2 string
	}
	channelMessageSendReturns struct {
		result1 *discordgo.Message
		result2 error
	}
	channelMessageSendReturnsOnCall map[int]struct {
		result1 *discordgo.Message
		result2 error
	}
	ChannelMessageSendEmbedStub        func(string, *discordgo.MessageEmbed) (*discordgo.Message, error)
	channelMessageSendEmbedMutex       sync.RWMutex
	channelMessageSendEmbedArgsForCall []struct {
		arg1 string
		arg2 *discordgo.MessageEmbed
	}
	channelMessageSendEmbedReturns struct {
		result1 *discordgo.Message
		result2 error
	}
	channelMessageSendEmbedReturnsOnCall map[int]struct {
		result1 *discordgo.Message
		result2 error
	}
	MessageReactionAddStub        func(string, string, string) error
	messageReactionAddMutex       sync.RWMutex
	messageReactionAddArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	messageReactionAddReturns struct {
		result1 error
	}
	messageReactionAddReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSession) ChannelMessageDelete(arg1 string, arg2 string) error {
	fake.channelMessageDeleteMutex.Lock()
	ret, specificReturn := fake.channelMessageDeleteReturnsOnCall[len(fake.channelMessageDeleteArgsForCall)]
	fake.channelMessageDeleteArgsForCall = append(fake.channelMessageDeleteArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ChannelMessageDelete", []interface{}{arg1, arg2})
	fake.channelMessageDeleteMutex.Unlock()
	if fake.ChannelMessageDeleteStub != nil {
		return fake.ChannelMessageDeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.channelMessageDeleteReturns
	return fakeReturns.result1
}

func (fake *FakeSession) ChannelMessageDeleteCallCount() int {
	fake.channelMessageDeleteMutex.RLock()
	defer fake.channelMessageDeleteMutex.RUnlock()
	return len(fake.channelMessageDeleteArgsForCall)
}

func (fake *FakeSession) ChannelMessageDeleteCalls(stub func(string, string) error) {
	fake.channelMessageDeleteMutex.Lock()
	defer fake.channelMessageDeleteMutex.Unlock()
	fake.ChannelMessageDeleteStub = stub
}

func (fake *FakeSession) ChannelMessageDeleteArgsForCall(i int) (string, string) {
	fake.channelMessageDeleteMutex.RLock()
	defer fake.channelMessageDeleteMutex.RUnlock()
	argsForCall := fake.channelMessageDeleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSession) ChannelMessageDeleteReturns(result1 error) {
	fake.channelMessageDeleteMutex.Lock()
	defer fake.channelMessageDeleteMutex.Unlock()
	fake.ChannelMessageDeleteStub = nil
	fake.channelMessageDeleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSession) ChannelMessageDeleteReturnsOnCall(i int, result1 error) {
	fake.channelMessageDeleteMutex.Lock()
	defer fake.channelMessageDeleteMutex.Unlock()
	fake.ChannelMessageDeleteStub = nil
	if fake.channelMessageDeleteReturnsOnCall == nil {
		fake.channelMessageDeleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.channelMessageDeleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSession) ChannelMessagePin(arg1 string, arg2 string) error {
	fake.channelMessagePinMutex.Lock()
	ret, specificReturn := fake.channelMessagePinReturnsOnCall[len(fake.channelMessagePinArgsForCall)]
	fake.channelMessagePinArgsForCall = append(fake.channelMessagePinArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ChannelMessagePin", []interface{}{arg1, arg2})
	fake.channelMessagePinMutex.Unlock()
	if fake.ChannelMessagePinStub != nil {
		return fake.ChannelMessagePinStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.channelMessagePinReturns
	return fakeReturns.result1
}

func (fake *FakeSession) ChannelMessagePinCallCount() int {
	fake.channelMessagePinMutex.RLock()
	defer fake.channelMessagePinMutex.RUnlock()
	return len(fake.channelMessagePinArgsForCall)
}

func (fake *FakeSession) ChannelMessagePinCalls(stub func(string, string) error) {
	fake.channelMessagePinMutex.Lock()
	defer fake.channelMessagePinMutex.Unlock()
	fake.ChannelMessagePinStub = stub
}

func (fake *FakeSession) ChannelMessagePinArgsForCall(i int) (string, string) {
	fake.channelMessagePinMutex.RLock()
	defer fake.channelMessagePinMutex.RUnlock()
	argsForCall := fake.channelMessagePinArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSession) ChannelMessagePinReturns(result1 error) {
	fake.channelMessagePinMutex.Lock()
	defer fake.channelMessagePinMutex.Unlock()
	fake.ChannelMessagePinStub = nil
	fake.channelMessagePinReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSession) ChannelMessagePinReturnsOnCall(i int, result1 error) {
	fake.channelMessagePinMutex.Lock()
	defer fake.channelMessagePinMutex.Unlock()
	fake.ChannelMessagePinStub = nil
	if fake.channelMessagePinReturnsOnCall == nil {
		fake.channelMessagePinReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.channelMessagePinReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSession) ChannelMessageSend(arg1 string, arg2 string) (*discordgo.Message, error) {
	fake.channelMessageSendMutex.Lock()
	ret, specificReturn := fake.channelMessageSendReturnsOnCall[len(fake.channelMessageSendArgsForCall)]
	fake.channelMessageSendArgsForCall = append(fake.channelMessageSendArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ChannelMessageSend", []interface{}{arg1, arg2})
	fake.channelMessageSendMutex.Unlock()
	if fake.ChannelMessageSendStub != nil {
		return fake.ChannelMessageSendStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.channelMessageSendReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSession) ChannelMessageSendCallCount() int {
	fake.channelMessageSendMutex.RLock()
	defer fake.channelMessageSendMutex.RUnlock()
	return len(fake.channelMessageSendArgsForCall)
}

func (fake *FakeSession) ChannelMessageSendCalls(stub func(string, string) (*discordgo.Message, error)) {
	fake.channelMessageSendMutex.Lock()
	defer fake.channelMessageSendMutex.Unlock()
	fake.ChannelMessageSendStub = stub
}

func (fake *FakeSession) ChannelMessageSendArgsForCall(i int) (string, string) {
	fake.channelMessageSendMutex.RLock()
	defer fake.channelMessageSendMutex.RUnlock()
	argsForCall := fake.channelMessageSendArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSession) ChannelMessageSendReturns(result1 *discordgo.Message, result2 error) {
	fake.channelMessageSendMutex.Lock()
	defer fake.channelMessageSendMutex.Unlock()
	fake.ChannelMessageSendStub = nil
	fake.channelMessageSendReturns = struct {
		result1 *discordgo.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) ChannelMessageSendReturnsOnCall(i int, result1 *discordgo.Message, result2 error) {
	fake.channelMessageSendMutex.Lock()
	defer fake.channelMessageSendMutex.Unlock()
	fake.ChannelMessageSendStub = nil
	if fake.channelMessageSendReturnsOnCall == nil {
		fake.channelMessageSendReturnsOnCall = make(map[int]struct {
			result1 *discordgo.Message
			result2 error
		})
	}
	fake.channelMessageSendReturnsOnCall[i] = struct {
		result1 *discordgo.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) ChannelMessageSendEmbed(arg1 string, arg2 *discordgo.MessageEmbed) (*discordgo.Message, error) {
	fake.channelMessageSendEmbedMutex.Lock()
	ret, specificReturn := fake.channelMessageSendEmbedReturnsOnCall[len(fake.channelMessageSendEmbedArgsForCall)]
	fake.channelMessageSendEmbedArgsForCall = append(fake.channelMessageSendEmbedArgsForCall, struct {
		arg1 string
		arg2 *discordgo.MessageEmbed
	}{arg1, arg2})
	fake.recordInvocation("ChannelMessageSendEmbed", []interface{}{arg1, arg2})
	fake.channelMessageSendEmbedMutex.Unlock()
	if fake.ChannelMessageSendEmbedStub != nil {
		return fake.ChannelMessageSendEmbedStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.channelMessageSendEmbedReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSession) ChannelMessageSendEmbedCallCount() int {
	fake.channelMessageSendEmbedMutex.RLock()
	defer fake.channelMessageSendEmbedMutex.RUnlock()
	return len(fake.channelMessageSendEmbedArgsForCall)
}

func (fake *FakeSession) ChannelMessageSendEmbedCalls(stub func(string, *discordgo.MessageEmbed) (*discordgo.Message, error)) {
	fake.channelMessageSendEmbedMutex.Lock()
	defer fake.channelMessageSendEmbedMutex.Unlock()
	fake.ChannelMessageSendEmbedStub = stub
}

func (fake *FakeSession) ChannelMessageSendEmbedArgsForCall(i int) (string, *discordgo.MessageEmbed) {
	fake.channelMessageSendEmbedMutex.RLock()
	defer fake.channelMessageSendEmbedMutex.RUnlock()
	argsForCall := fake.channelMessageSendEmbedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSession) ChannelMessageSendEmbedReturns(result1 *discordgo.Message, result2 error) {
	fake.channelMessageSendEmbedMutex.Lock()
	defer fake.channelMessageSendEmbedMutex.Unlock()
	fake.ChannelMessageSendEmbedStub = nil
	fake.channelMessageSendEmbedReturns = struct {
		result1 *discordgo.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) ChannelMessageSendEmbedReturnsOnCall(i int, result1 *discordgo.Message, result2 error) {
	fake.channelMessageSendEmbedMutex.Lock()
	defer fake.channelMessageSendEmbedMutex.Unlock()
	fake.ChannelMessageSendEmbedStub = nil
	if fake.channelMessageSendEmbedReturnsOnCall == nil {
		fake.channelMessageSendEmbedReturnsOnCall = make(map[int]struct {
			result1 *discordgo.Message
			result2 error
		})
	}
	fake.channelMessageSendEmbedReturnsOnCall[i] = struct {
		result1 *discordgo.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) MessageReactionAdd(arg1 string, arg2 string, arg3 string) error {
	fake.messageReactionAddMutex.Lock()
	ret, specificReturn := fake.messageReactionAddReturnsOnCall[len(fake.messageReactionAddArgsForCall)]
	fake.messageReactionAddArgsForCall = append(fake.messageReactionAddArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("MessageReactionAdd", []interface{}{arg1, arg2, arg3})
	fake.messageReactionAddMutex.Unlock()
	if fake.MessageReactionAddStub != nil {
		return fake.MessageReactionAddStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.messageReactionAddReturns
	return fakeReturns.result1
}

func (fake *FakeSession) MessageReactionAddCallCount() int {
	fake.messageReactionAddMutex.RLock()
	defer fake.messageReactionAddMutex.RUnlock()
	return len(fake.messageReactionAddArgsForCall)
}

func (fake *FakeSession) MessageReactionAddCalls(stub func(string, string, string) error) {
	fake.messageReactionAddMutex.Lock()
	defer fake.messageReactionAddMutex.Unlock()
	fake.MessageReactionAddStub = stub
}

func (fake *FakeSession) MessageReactionAddArgsForCall(i int) (string, string, string) {
	fake.messageReactionAddMutex.RLock()
	defer fake.messageReactionAddMutex.RUnlock()
	argsForCall := fake.messageReactionAddArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSession) MessageReactionAddReturns(result1 error) {
	fake.messageReactionAddMutex.Lock()
	defer fake.messageReactionAddMutex.Unlock()
	fake.MessageReactionAddStub = nil
	fake.messageReactionAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSession) MessageReactionAddReturnsOnCall(i int, result1 error) {
	fake.messageReactionAddMutex.Lock()
	defer fake.messageReactionAddMutex.Unlock()
	fake.MessageReactionAddStub = nil
	if fake.messageReactionAddReturnsOnCall == nil {
		fake.messageReactionAddReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.messageReactionAddReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSession) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.channelMessageDeleteMutex.RLock()
	defer fake.channelMessageDeleteMutex.RUnlock()
	fake.channelMessagePinMutex.RLock()
	defer fake.channelMessagePinMutex.RUnlock()
	fake.channelMessageSendMutex.RLock()
	defer fake.channelMessageSendMutex.RUnlock()
	fake.channelMessageSendEmbedMutex.RLock()
	defer fake.channelMessageSendEmbedMutex.RUnlock()
	fake.messageReactionAddMutex.RLock()
	defer fake.messageReactionAddMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSession) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ amigobot.Session = new(FakeSession)

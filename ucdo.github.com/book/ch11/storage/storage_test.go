package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	// 这里是为了保存原来的状态。免得重写了之后影响之后的处理
	saved := notifyUser
	defer func() { notifyUser = saved }()

	var notifiedUser, notifiedMsg string
	// 这里相当于重写了 storage里面的notifyUser。等于mocking桩
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	const user = "userxxx1@xx.com"

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser is not called")
	}

	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s",
			notifiedUser, user)
	}

	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}

func TestNonUse(t *testing.T) {
	s, sep := "a,b,c", ","
	words := strings.Split(s, sep)
	if got, want := len(words), 5; got != want {
		t.Fatalf("split(%s %s),return %d words ,want %d",
			s, sep, got, want)
	}
}

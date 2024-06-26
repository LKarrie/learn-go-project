package mail

import (
	"testing"

	"github.com/LKarrie/learn-go-project/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)
	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>hello world</h1>
	`
	to := []string{"lkarrie616@gmail.com"}
	attachFiles := []string{"../README.md"}

	sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}

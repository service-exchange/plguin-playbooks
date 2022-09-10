package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPlaybookDetailsURL(t *testing.T) {
	require.Equal(t,
		"http://service-exchange.com/playbooks/playbooks/playbookTestId",
		getPlaybookDetailsURL("http://service-exchange.com", "playbookTestId"),
	)
}

func TestGetPlaybooksNewURL(t *testing.T) {
	require.Equal(t,
		"http://service-exchange.com/playbooks/playbooks/new",
		getPlaybooksNewURL("http://service-exchange.com"),
	)
}

func TestGetPlaybooksURL(t *testing.T) {
	require.Equal(t,
		"http://service-exchange.com/playbooks/playbooks",
		getPlaybooksURL("http://service-exchange.com"),
	)
}

func TestGetPlaybookDetailsRelativeURL(t *testing.T) {
	require.Equal(t,
		"/playbooks/playbooks/testPlaybookId",
		GetPlaybookDetailsRelativeURL("testPlaybookId"),
	)
}

func TestGetRunDetailsRelativeURL(t *testing.T) {
	require.Equal(t,
		"/playbooks/runs/testPlaybookRunId",
		GetRunDetailsRelativeURL("testPlaybookRunId"),
	)
}

func TestGetRunDetailsURL(t *testing.T) {
	require.Equal(t,
		"http://service-exchange.com/playbooks/runs/testPlaybookRunId",
		getRunDetailsURL("http://service-exchange.com", "testPlaybookRunId"),
	)
}

func TestGetRunRetrospectiveURL(t *testing.T) {
	require.Equal(t,
		"http://service-exchange.com/playbooks/runs/testPlaybookRunId/retrospective",
		getRunRetrospectiveURL("http://service-exchange.com", "testPlaybookRunId"),
	)
}

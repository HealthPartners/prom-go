package prometheus

// AlertGroup represents a prometheus webhook alert https://prometheus.io/docs/alerting/configuration/#%3Cwebhook_config%3E
type AlertGroup struct {
	Version           string `json:"version"`
	GroupKey          string `json:"groupKey"`
	Status            string `json:"status"`
	Receiver          string `json:"receiver"`
	ExternalURL       string `json:"externalURL"`
	CommonAnnotations struct {
		Description      string `json:"description"`
		ShortDescription string `json:"summary"`
		RunBook          string `json:"runbook"`
	} `json:"commonAnnotations"`
	Alerts []struct {
		Status      string            `json:"status"`
		Labels      map[string]string `json:"labels"`
		Annotations struct {
			Description      string `json:"description"`
			ShortDescription string `json:"summary"`
			RunBook          string `json:"runbook"`
		} `json:"annotations"`
	} `json:"alerts"`
}
